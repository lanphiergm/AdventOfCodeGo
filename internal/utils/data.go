package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Parses the integers from a comma-separated list of values
func ParseInts(csv string) []int {
	strs := strings.Split(csv, ",")
	values := make([]int, len(strs))
	for i, n := range strs {
		values[i] = Atoi(n)
	}
	return values
}

// Reads the comma-separated values from the file at the specified path into
// a slice of ints
func ReadCsv(path string) []int {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return ParseInts(string(contents))
}

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
		lines = append(lines, Atoi(scanner.Text()))
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
