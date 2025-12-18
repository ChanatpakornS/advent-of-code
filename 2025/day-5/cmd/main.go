package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type pair struct {
	first  int
	second int
}

func check(e error, info string) {
	if e != nil {
		panic(fmt.Sprintf("%s: %v", info, e))
	}
}

func parseInt(input string) int {
	number, err := strconv.Atoi(input)
	check(err, "Fail to parse integer")
	return number
}

func validateItemValue(item int, interval []pair) bool {
	for _, pair := range interval {
		if item >= pair.first && item <= pair.second {
			return true
		}
	}
	return false
}

func main() {
	dir := filepath.Dir("main.go")
	filePath := filepath.Join(dir, "input.txt")
	data, err := os.ReadFile(filePath)
	check(err, "Fail to read file")

	trimData := strings.Split(string(bytes.TrimSpace(data)), "\n\n")
	rangeFresh := strings.SplitSeq(trimData[0], "\n")
	items := strings.SplitSeq(trimData[1], "\n")

	freshCount := 0
	interval := []pair{}

	// Create the map of interval range
	for line := range rangeFresh {
		freshInterval := strings.Split(line, "-")
		interval = append(interval, pair{
			first:  parseInt(freshInterval[0]),
			second: parseInt(freshInterval[1]),
		})
	}

	// Check if the value is within the interval
	for item := range items {
		if validateItemValue(parseInt(item), interval) {
			freshCount++
		}
	}

	fmt.Println(freshCount)
}
