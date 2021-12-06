package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const days = 256
const growCycle = 8
const matureCycle = 6

func day6main() {
	var lanternFishes [growCycle + 1]uint64
	file, err := os.Open("day6_input.txt")
	//file, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var input = strings.Split(scanner.Text(), ",")
		for _, i := range input {
			f, _ := strconv.Atoi(i)
			lanternFishes[f]++
		}
	}

	for i := 0; i < days; i++ {
		splitters := lanternFishes[0]
		for j := 1; j <= growCycle; j++ {
			lanternFishes[j-1] = lanternFishes[j]
		}
		lanternFishes[growCycle] = splitters
		lanternFishes[matureCycle] += splitters
	}

	result := uint64(0)

	for _, dayCount := range lanternFishes {
		result += dayCount
	}

	log.Printf("result : %d", result)

}
