package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error, info string) {
	if e != nil {
		panic(fmt.Sprintf("%s: %v", info, e))
	}
}

func parseInt(input string) int {
	intValue, err := strconv.Atoi(input)
	check(err, "Fail to parse int")
	return intValue
}

func invalidNumber(number int) bool {
	stringNumber := strconv.Itoa(number)
	if len(stringNumber)%2 != 0 || len(stringNumber) < 2 {
		return false
	}
	return stringNumber[:len(stringNumber)/2] == stringNumber[len(stringNumber)/2:]
}

func main() {
	dir := filepath.Dir("main.go")
	filePath := filepath.Join(dir, "input.txt")
	data, err := os.ReadFile(filePath)
	check(err, "Fail to read file")

	trimData := strings.ReplaceAll(string(bytes.TrimSpace(data)), "\n", "")
	rangeIDs := strings.SplitSeq(string(trimData), ",")

	collector := 0

	for rangeID := range rangeIDs {
		interval := strings.Split(string(rangeID), "-")
		start := parseInt(interval[0])
		stop := parseInt(interval[1])

		for i := start; i <= stop; i++ {
			if invalidNumber(i) {
				collector += i
			}
		}
	}

	fmt.Println(collector)
}
