package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	scanner := bufio.NewScanner(file)
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read file.")
	}

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Failed to create file.")
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to write JSON data.")
	}

	file.Close()
	return nil
}
