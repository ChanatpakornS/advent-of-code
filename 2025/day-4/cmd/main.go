package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var adjPath = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, -1}, {-1, -1}, {1, 1}, {-1, 1}}

func check(e error, info string) {
	if e != nil {
		panic(fmt.Sprintf("%s: %v", info, e))
	}
}

func validatePaperRoll(view [][]string, r, c int) bool {
	adjPaperRoll := 0
	for _, path := range adjPath {
		newR, newC := r+path[0], c+path[1]
		if newR >= 0 && newR < len(view) && newC >= 0 && newC < len(view[0]) && view[newR][newC] == "@" {
			adjPaperRoll++
		}
	}

	return adjPaperRoll < 4 && view[r][c] == "@"
}

func main() {
	dir := filepath.Dir("main.go")
	filePath := filepath.Join(dir, "input.txt")
	data, err := os.ReadFile(filePath)
	check(err, "Fail to read file")

	trimData := strings.Split(string(bytes.TrimSpace(data)), "\n")

	amountValidPaperRoll := 0

	row_amount := len(trimData)
	col_amount := len(trimData[0])
	overview := make([][]string, row_amount)
	for i := range overview {
		overview[i] = make([]string, col_amount)
	}

	for row, line := range trimData {
		for col, text := range line {
			overview[row][col] = string(text)
		}
	}

	for row := range overview {
		for col := range overview[row] {
			if validatePaperRoll(overview, row, col) {
				amountValidPaperRoll++
			}
		}
	}

	fmt.Println(amountValidPaperRoll)

}
