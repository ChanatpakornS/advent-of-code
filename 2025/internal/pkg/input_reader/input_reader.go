package input_reader

import (
	"bytes"
	"os"
	"path/filepath"

	"aoc2025/internal/pkg/errs"
)

func ReadInput(dir, fileName string) string {
	filePath := filepath.Join(dir, fileName)
	data, err := os.ReadFile(filePath)
	errs.Error(err, "Fail to read file")

	trimData := string(bytes.TrimSpace(data))
	return trimData
}
