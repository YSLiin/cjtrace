package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func main() {
	inputPath := flag.String("in", "", "Path to BIAU.TXT (UTF-8).")
	outputPath := flag.String("out", "", "Path to output CSV.")
	flag.Parse()

	reader, err := os.Open(*inputPath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer reader.Close()

	// NOTE:
	// ║     1  │ 的 │白│08║  32739 │   32739│  1.651 ║
	// ║     2  │ 不 │一│04║  24362 │   57101│  2.879 ║
	re := regexp.MustCompile(`║\s*\d+\s*│\s*([^\s│]+)\s*│[^║]*║\s*(\d+)`)
	scanner := bufio.NewScanner(reader)

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

	// NOTE: Lines that don't match are headers and borders; a matching row must parse cleanly.
	rows := 0
	for lineNo := 1; scanner.Scan(); lineNo++ {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) != 3 {
			continue
		}
		char, freq := matches[1], matches[2]
		if utf8.RuneCountInString(char) != 1 {
			log.Fatalf("Line %d: want one char, got %q", lineNo, char)
		}
		if _, err := strconv.Atoi(freq); err != nil {
			log.Fatalf("Line %d: invalid frequency %q: %v", lineNo, freq, err)
		}
		if err := cjwriter.Write([]string{char, freq}); err != nil {
			log.Fatalf("Failed to write CSV: %v", err)
		}
		rows++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %v", err)
	}
	if rows == 0 {
		log.Fatalf("No frequency rows found in %s", *inputPath)
	}

	cjwriter.Flush()
	if err := cjwriter.Error(); err != nil {
		log.Fatalf("Failed to flush CSV writer: %v", err)
	}
}
