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

func processFile(f *os.File) int {

	scanner := bufio.NewScanner(f)
	var sum = 0
	var temp []string
	var err error

	for scanner.Scan() {
		var previousValue int
		var currentValue int
		var increasing bool
		var isSafe bool = true

		line := scanner.Text()
		temp = strings.Split(line, " ")
		if len(temp) == 1 {
			sum += 1
			continue
		}

		check(err)

		for i, v := range temp {
			currentValue, err = strconv.Atoi(v)
			check(err)

			if i == 0 {
				previousValue = currentValue
				continue
			}

			if i == 1 {
				if previousValue < currentValue {
					increasing = true
				} else {
					increasing = false
				}
			}

			if !isValidValue(previousValue, currentValue, increasing) {
				isSafe = false
				break
			}

			previousValue = currentValue
		}

		if isSafe {
			sum += 1
		}
		fmt.Println(isSafe)
		fmt.Println(line)
	}

	err = scanner.Err()
	check(err)

	return sum
}

func main() {

	dir, err := os.Getwd()
	check(err)

	fileName := "input.txt"
	fileName = "testInput.txt"
	file := filepath.Join(dir, fileName)

	f, err := os.Open(file)
	check(err)

	result := processFile(f)
	f.Close()

	fmt.Println(result)
}
