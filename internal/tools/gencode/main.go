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
		common_chars := make([]string, 0, 2)
		for _, c := range chars {
			char := []rune(c)[0]
			if char < 0x4E00 || char > 0x9FFF {
				continue
			}
			common_chars = append(common_chars, string(char))
		}
		if len(common_chars) == 0 {
			continue
		}
		sort.Strings(common_chars)
		out = append(out, append([]string{code}, common_chars...))
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
	cjwriter.WriteAll(out)
	cjwriter.Flush()
	if err := cjwriter.Error(); err != nil {
		log.Fatalf("Failed to flush CSV writer: %v", err)
	}
}
