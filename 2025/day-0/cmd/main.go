package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	check(err, "Fail to read file")

	trimData := strings.ReplaceAll(string(bytes.TrimSpace(data)), "\n", "")
	input := strings.SplitSeq(trimData, ",")
	fmt.Println(input)
}
