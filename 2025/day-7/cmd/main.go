package main

import (
	"path/filepath"
	"strings"

	"aoc2025/internal/pkg/input_reader"
)

func validateSplitted(col int, visitedColumns map[int]struct{}) bool {
	if _, ok := visitedColumns[col]; ok {
		return true
	}
	return false
}

func main() {
	dir := filepath.Dir("main.go")
	fileName := "test-input.txt"
	input := input_reader.ReadInput(dir, fileName)
	floors := strings.SplitSeq(input, "\n")

	numberofSplitted := 0
	visitedColumns := make(map[int]struct{})

	for line := range floors {
		bufferLine := []rune{}
		for col, value := range line {
			switch {
			case value == rune('S'):
				visitedColumns[col] = struct{}{}

				// Assume that there's only one starting point and will not have splitter in this line anymore
				break
			case value == rune('^'):
				if validateSplitted(col, visitedColumns) {
					continue
				}
				bufferLine = append(bufferLine, value)
				numberofSplitted += validSplitAmount(col, bufferLine)
				visitedColumns[col] = struct{}{}
			}
		}
	}
}
