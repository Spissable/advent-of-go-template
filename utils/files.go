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
		panic(fmt.Sprintf("File not found for path: %s", err.Error()))
	}

	return string(content)
}
