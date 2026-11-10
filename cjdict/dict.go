package cjdict

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:generate go run ../internal/tools/gencode -in $CJ5_JSON -out cjcode.csv
//go:embed cjcode.csv
var RawCodes []byte

//go:generate go run ../internal/tools/genfreq -in $BIAU1_TXT -out freq.csv
//go:embed freq.csv
var RawFreqs []byte

var cjFullTable = make(map[rune]CjCode)
var cjFreqs = make(map[rune]int)

type CjCode struct {
	Codes    []string
	Bitmasks []uint32
}

func GetFreqs() map[rune]int {
	if len(cjFreqs) > 0 {
		return cjFreqs
	}

	cjreader := csv.NewReader(bytes.NewReader(RawFreqs))
	records, err := cjreader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse freq.csv: %v", err)
	}

	for _, record := range records {
		char := []rune(record[0])[0]
		freq, _ := strconv.Atoi(record[1])
		cjFreqs[char] = freq
	}
	return cjFreqs
}

func GetChars() map[rune]CjCode {
	cjTable := make(map[rune]CjCode)
	for char, cjcode := range GetAllChars() {
		if _, ok := GetFreqs()[char]; ok {
			cjTable[char] = cjcode
		}
	}
	return cjTable
}

func GetAllChars() map[rune]CjCode {
	if len(cjFullTable) > 0 {
		return cjFullTable
	}

	cjreader := csv.NewReader(bytes.NewReader(RawCodes))
	cjreader.FieldsPerRecord = -1
	records, err := cjreader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse cjcode.csv: %v", err)
	}

	// NOTE: Each record start with a code followed by the chars, e.g., "abjj,暈,暉".
	for _, record := range records {
		code := record[0]
		for _, c := range record[1:] {
			char := []rune(c)[0]
			cjchar, ok := cjFullTable[char]
			if ok {
				cjchar.Codes = append(cjchar.Codes, code)
				cjchar.Bitmasks = append(cjchar.Bitmasks, getBitmask(code))
				cjFullTable[char] = cjchar
			} else {
				cjFullTable[char] = CjCode{
					Codes:    []string{code},
					Bitmasks: []uint32{getBitmask(code)},
				}
			}
		}
	}
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
