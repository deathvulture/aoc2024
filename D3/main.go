//https://adventofcode.com/2024/day/1
//Find the distances between the numbers of two lists ordered from smallest to largest.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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

func ProcessFile(f *os.File) int {

	scanner := bufio.NewScanner(f)
	var sum = 0
	var list []int
	var list2 []int
	var temp []string
	var tempVal int
	var tempVal2 int
	var err error

	for scanner.Scan() {
		line := scanner.Text()
		temp = strings.Split(line, "   ")
		tempVal, err = strconv.Atoi(temp[0])
		check(err)

		tempVal2, err = strconv.Atoi(temp[1])
		check(err)

		list = append(list, tempVal)
		list2 = append(list2, tempVal2)

		//fmt.Println(line)
	}

	sort.Ints(list)
	sort.Ints(list2)

	for i := 0; i < len(list); i++ {
		sum += absInt(list[i] - list2[i])
	}
	err = scanner.Err()
	check(err)

	return sum
}

func Main() {

	dir, err := os.Getwd()
	check(err)

	fileName := "input.txt"
	//fileName = "testInput.txt"
	file := filepath.Join(dir, fileName)

	f, err := os.Open(file)
	check(err)

	result := ProcessFile(f)
	f.Close()

	fmt.Println(result)
}
