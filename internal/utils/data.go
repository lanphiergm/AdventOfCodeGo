package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Reads the file at the specified path into a slice of ints
func ReadInts(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		lines = append(lines, int(val))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// Reads the file at the specified path into a slice of strings
func ReadStrings(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
