package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
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

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			char := matches[1]
			freq, _ := strconv.Atoi(matches[2])

			cjwriter.Write([]string{char, strconv.Itoa(freq)})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %v", err)
	}

	cjwriter.Flush()
	if err := cjwriter.Error(); err != nil {
		log.Fatalf("Failed to flush CSV writer: %v", err)
	}
}
