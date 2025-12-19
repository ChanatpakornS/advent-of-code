package parse

import (
	"aoc2025/internal/pkg/errs"
	"strconv"
)

func ParseRuneInt(r rune) int {
	number, err := strconv.Atoi(string(r))
	errs.Error(err, "Failed to parse rune type to int")
	return number
}

func ParseStringToInt(s string) int {
	number, err := strconv.Atoi(s)
	errs.Error(err, "Failed to parse string type to int")
	return number
}
