//https://adventofcode.com/2024/day/2
//Find the distances between the numbers of two lists ordered from smallest to largest.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isValidValue(a int, b int, increasing bool) bool {
	var distance = b - a
	var absDistance = absInt(distance)

	if absDistance < 1 || absDistance > 3 {
		return false
	}

	if increasing && distance < 0 {
		return false
	}

	if !increasing && distance > 0 {
		return false
	}

	return true
}

func isValidSlice(intSlice []string) bool {
	//fmt.Println(intSlice)
	var increasing bool
	var previousValue int
	var isSafe bool = true

	for i, v := range intSlice {
		currentValue, err := strconv.Atoi(v)
		check(err)

		if i == 0 {
			previousValue = currentValue
			continue
		}

		if i == 1 {
			increasing = previousValue < currentValue
		}

		if !isValidValue(previousValue, currentValue, increasing) {
			isSafe = false
			break
		}
		previousValue = currentValue
	}

	return isSafe
}

func processFile(f *os.File) int {

	scanner := bufio.NewScanner(f)
	var sum = 0
	var originalSlice []string

	for scanner.Scan() {
		line := scanner.Text()
		originalSlice = strings.Split(line, " ")
		if len(originalSlice) == 1 {
			sum += 1
			continue
		}

		for i := range originalSlice {
			tempSlice := make([]string, len(originalSlice))
			copy(tempSlice, originalSlice)
			tempSlice = append(tempSlice[:i], tempSlice[i+1:]...)
			isSafe := isValidSlice(tempSlice)
			if isSafe {
				sum += 1
				break
			}
			//fmt.Println(line)
		}
	}
	err := scanner.Err()
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
