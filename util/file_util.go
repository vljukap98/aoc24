package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileAsText(path string) (string, error) {
	var fileName string = path
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Cannot read file %s", fileName)
		return "", err
	}
	fileContent := string(file)

	return fileContent, err
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadMat(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elem := strings.Split(scanner.Text(), "")
		lines = append(lines, elem)
	}
	return lines, scanner.Err()
}
