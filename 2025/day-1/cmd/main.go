package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

func check(e error, info string) {
	if e != nil {
		panic(fmt.Sprintf("%s: %v", info, e))
	}
}

func main() {
	dir := filepath.Dir("main.go")
	filePath := filepath.Join(dir, "input.txt")
	data, err := os.ReadFile(filePath)
	check(err, "Failed to read file")

	// Dial start at 50
	dial := 50
	password := 0

	// Right is POSITIVE
	// Left is NEGATIVE
	trimInput := bytes.TrimSpace(data)
	actionList := strings.SplitSeq(string(trimInput), "\n")

	for action := range actionList {
		splitIndex := strings.IndexFunc(action, unicode.IsDigit)
		if splitIndex == -1 {
			check(errors.New("Invalid action format"), "Failed to split the action")
		}

		prefixAction := action[:splitIndex]
		valueAction := action[splitIndex:]
		value, err := strconv.Atoi(valueAction)
		check(err, "Failed to parse the action value")

		switch prefixAction {
		case "R":
			dial += (value % 100)
		case "L":
			dial -= (value % 100)
		default:
			check(errors.New("Invalid action prefix"), "Failed to parse the action prefix")
		}

		dial %= 100
		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			password++
		}
	}

	fmt.Println(password)
}
