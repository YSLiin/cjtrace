package cjdict

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"log"
	"slices"
	"strconv"
	"strings"
	"sync"
)

//go:generate go run ../internal/tools/gencode -in $CJ5_JSON -out cjcode.csv
//go:embed cjcode.csv
var rawCodes []byte

//go:generate go run ../internal/tools/genfreq -in $BIAU1_TXT -out freq.csv
//go:embed freq.csv
var rawFreqs []byte

var (
	freqsOnce sync.Once
	cjFreqs   map[rune]int

	fullOnce    sync.Once
	cjFullTable map[rune]CjCode

	poolOnce sync.Once
	cjTable  map[rune]CjCode
)

type CjCode struct {
	Codes    []string
	Bitmasks []uint32
}

func GetFreqs() map[rune]int {
	freqsOnce.Do(func() {
		cjreader := csv.NewReader(bytes.NewReader(rawFreqs))
		records, err := cjreader.ReadAll()
		if err != nil {
			log.Fatalf("Failed to parse freq.csv: %v", err)
		}

		cjFreqs = make(map[rune]int, len(records))
		for _, record := range records {
			char := []rune(record[0])[0]
			freq, err := strconv.Atoi(record[1])
			if err != nil {
				log.Fatalf("Failed to parse freq.csv: %v", err)
			}
			cjFreqs[char] = freq
		}
	})
	return cjFreqs
}

func GetChars() map[rune]CjCode {
	poolOnce.Do(func() {
		freqs := GetFreqs()
		cjTable = make(map[rune]CjCode, len(freqs))
		for char, cjcode := range GetAllChars() {
			if _, ok := freqs[char]; ok {
				cjTable[char] = cjcode
			}
		}
	})
	return cjTable
}

func GetAllChars() map[rune]CjCode {
	fullOnce.Do(func() {
		cjreader := csv.NewReader(bytes.NewReader(rawCodes))
		cjreader.FieldsPerRecord = -1
		records, err := cjreader.ReadAll()
		if err != nil {
			log.Fatalf("Failed to parse cjcode.csv: %v", err)
		}

		// NOTE: Each record starts with a code followed by the chars, e.g., "abjj,暈,暉".
		cjFullTable = make(map[rune]CjCode)
		for _, record := range records {
			code := record[0]
			for _, c := range record[1:] {
				char := []rune(c)[0]
				cjchar := cjFullTable[char]
				cjchar.Codes = append(cjchar.Codes, code)
				cjchar.Bitmasks = append(cjchar.Bitmasks, getBitmask(code))
				cjFullTable[char] = cjchar
			}
		}
	})
	return cjFullTable
}

func GetCandidates(letters string) []rune {
	bitmask := getBitmask(letters)
	return getCandidates(bitmask, GetChars())
}

func GetAllCandidates(letters string) []rune {
	bitmask := getBitmask(letters)
	return getCandidates(bitmask, GetAllChars())
}

func getCandidates(bitmask uint32, pool map[rune]CjCode) []rune {
	chars := make([]rune, 0, 16)
	for char, cjcode := range pool {
		if slices.ContainsFunc(cjcode.Bitmasks, func(mask uint32) bool { return mask&^bitmask == 0 }) {
			chars = append(chars, char)
		}
	}
	slices.Sort(chars)
	return chars
}

func getBitmask(keys string) (bitmask uint32) {
	for _, c := range strings.ToLower(keys) {
		// NOTE: unicodes.IsLetter() admits é/中; unicodes.IsLower() admits é.
		if 'a' <= c && c <= 'z' {
			bitmask |= 1 << (c - 'a')
		}
	}
	return bitmask
}
