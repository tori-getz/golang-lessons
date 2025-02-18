package files

import (
	"errors"
	"os"
	"strings"
)

func Read(fileName string) ([]string, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, errors.New("READ_FILE_ERROR")
	}

	urlSlice := strings.Split(string(file), "\n")

	return urlSlice, nil
}
