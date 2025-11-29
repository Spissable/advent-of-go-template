package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func ReadInput() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller information")
	}

	filePath := fmt.Sprintf("%s/input.txt", filepath.Dir(file))
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic("Input file not found. Make sure to run `make init` to complete the one-time setup.")
	}

	return string(content)
}
