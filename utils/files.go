package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func ReadInput(t *testing.T) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("Failed to get caller information. Can't determine which day to load.")
	}

	filePath := fmt.Sprintf("%s/input.txt", filepath.Dir(file))
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal("Input file not found. Make sure to run `make init` to complete the one-time setup.")
	}

	return string(content)
}
