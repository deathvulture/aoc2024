//https://adventofcode.com/2024/day/2
//Find the distances between the numbers of two lists ordered from smallest to largest.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func findInstances(s string) [][]string {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(s, -1)
	return matches
}

func processFile(f *os.File) int {

	scanner := bufio.NewScanner(f)
	var sum = 0
	//var temp []string
	var err error

	for scanner.Scan() {
		line := scanner.Text()
		matches := findInstances(line)
		for _, match := range matches {
			a, err := strconv.Atoi(match[1])
			b, err := strconv.Atoi(match[2])
			check(err)
			sum += a * b
		}

	}

	err = scanner.Err()
	check(err)

	return sum
}

func main() {

	dir, err := os.Getwd()
	check(err)

	fileName := "input.txt"
	//fileName = "testInput.txt"
	file := filepath.Join(dir, fileName)

	f, err := os.Open(file)
	check(err)

	result := processFile(f)
	f.Close()

	fmt.Println(result)
}
