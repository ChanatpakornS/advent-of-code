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

func parseInt(s string) int {
	number, err := strconv.Atoi(s)
	check(err, "Fail to parse integer")
	return number
}

func main() {
	dir := filepath.Dir("main.go")
	filePath := filepath.Join(dir, "input.txt")
	data, err := os.ReadFile(filePath)
	check(err, "Fail to read file")

	trimData := strings.Split(string(bytes.TrimSpace(data)), "\n")
	operations := strings.Fields(trimData[len(trimData)-1])

	totalFactor := len(strings.Fields(trimData[0]))
	numberBuffer := make([][]int, totalFactor)
	totalOutput := 0

	// Put the number to the summation bucket
	for index, line := range trimData {
		if index == len(trimData)-1 {
			break
		}
		members := strings.Fields(line)
		for bucket, number := range members {
			numberBuffer[bucket] = append(numberBuffer[bucket], parseInt(number))
		}
	}

	// Calculate the total output
	for idx, numbers := range numberBuffer {
		switch operations[idx] {
		case "+":
			buffer := 0
			for _, value := range numbers {
				buffer += value
			}
			totalOutput += buffer
		case "*":
			buffer := 1
			for _, value := range numbers {
				buffer *= value
			}
			totalOutput += buffer
		default:
			check(fmt.Errorf("Invalid operation: %s", operations[idx]), "Invalid operation")
		}
	}

	fmt.Println(totalOutput)
}
