package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"sort"
)

type CJ5 struct {
	CharDefs map[string][]string `json:"chardefs"`
}

func main() {
	inputPath := flag.String("in", "", "Path to cj5.json.")
	outputPath := flag.String("out", "", "Path to output csv file.")
	flag.Parse()

	file, err := os.Open(*inputPath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	cjbytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	file.Close()

	cj5 := CJ5{}
	if err := json.Unmarshal(cjbytes, &cj5); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	cjcodes := cj5.CharDefs
	out := make([][]string, 0)
	for code, chars := range cjcodes {
		commonChars := make([]string, 0, 2)
		for _, c := range chars {
			if !isValidCode(code) || c == "" {
				log.Printf("Skipping due to invalid code %q or empty char under code %q.", code, code)
				continue
			}

			// NOTE: Only CJK Unified Ideographs.
			char := []rune(c)[0]
			if char < 0x4E00 || char > 0x9FFF {
				continue
			}
			commonChars = append(commonChars, string(char))
		}
		if len(commonChars) > 0 {
			sort.Strings(commonChars)
			out = append(out, append([]string{code}, commonChars...))
		}
	}
	if len(out) == 0 {
		log.Fatalf("No codes produced from %s", *inputPath)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i][0] < out[j][0]
	})

	var cjwriter *csv.Writer
	if *outputPath == "" {
		cjwriter = csv.NewWriter(os.Stdout)
	} else {
		file, err := os.OpenFile(*outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to open output file: %v", err)
		}
		defer file.Close()
		cjwriter = csv.NewWriter(file)
	}
	if err := cjwriter.WriteAll(out); err != nil {
		log.Fatalf("Failed to write CSV: %v", err)
	}
}

func isValidCode(code string) bool {
	if len(code) < 1 || len(code) > 5 {
		return false
	}
	for _, c := range code {
		if c < 'a' || c > 'y' {
			return false
		}
	}
	return true
}
