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

func parseInt(input rune) int {
	intValue, err := strconv.Atoi(string(input))
	check(err, "Fail to parse int")
	return intValue
}

func main() {
	dir := filepath.Dir("main.go")
	filePath := filepath.Join(dir, "input.txt")
	data, err := os.ReadFile(filePath)
	check(err, "Fail to read file")

	input := strings.SplitSeq(string(bytes.TrimSpace(data)), "\n")
	totalJoltage := 0

	for batteryBank := range input {
		prefix := -1
		suffix := -1
		// assume that the input will not less than 2 characters
		for i := 0; i < len(batteryBank)-1; i++ {
			value := parseInt(rune(batteryBank[i]))
			if value > prefix {
				prefix = value
				suffix = parseInt(rune(batteryBank[i+1]))
			} else if value > suffix {
				suffix = value
			}
		}
		lastDigit := parseInt(rune(batteryBank[len(batteryBank)-1]))
		if lastDigit > suffix {
			suffix = lastDigit
		}

		joltage := prefix*10 + suffix
		totalJoltage += joltage
		// fmt.Println(joltage)
	}

	fmt.Println(totalJoltage)
}
