//https://adventofcode.com/2024/day/1
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

func processFile(f *os.File) int {

	scanner := bufio.NewScanner(f)
	var sum = 0
	var temp []string
	var tempVal int
	var tempVal2 int
	var err error
	var leftMap (map[int]int) = make(map[int]int)
	var rightMap (map[int]int) = make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		temp = strings.Split(line, "   ")
		tempVal, err = strconv.Atoi(temp[0])
		check(err)

		tempVal2, err = strconv.Atoi(temp[1])
		check(err)

		//fmt.Println(line)
		if leftMap[tempVal] == 0 {
			leftMap[tempVal] = 1
		} else {
			leftMap[tempVal] += 1
		}

		if rightMap[tempVal2] == 0 {
			rightMap[tempVal2] = 1
		} else {
			rightMap[tempVal2] += 1
		}
	}

	for k, v := range leftMap {
		sum += k * v * rightMap[k]
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
