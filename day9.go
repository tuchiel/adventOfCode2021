package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var line []int
var previousLine []int = nil
var nextLine []int = nil

func convert2(input string) (output []int) {
	output = make([]int, 0, len(input))
	for _, str := range strings.Split(input, "") {
		var val, _ = strconv.Atoi(str)
		output = append(output, val)
	}
	return
}

func day9_main() {
	//file, err := os.Open("test_input.txt")
	file, err := os.Open("day9_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line = convert2(scanner.Text())
	scanner.Scan()
	nextLine = convert2(scanner.Text())
	riskScore := 0
	for line != nil {
		for i := range line {
			var (
				previous, next, above, under = 9, 9, 9, 9
			)
			if previousLine != nil {
				above = previousLine[i]
			}
			if nextLine != nil {
				under = nextLine[i]
			}
			if i < len(line)-1 {
				next = line[i+1]
			}
			if i > 0 {
				previous = line[i-1]
			}
			if (line[i] < previous) && (line[i] < next) && (line[i] < above) && (line[i] < under) {
				log.Printf("Low point %d", line[i])
				riskScore += line[i] + 1
			}
		}
		previousLine = line
		line = nextLine
		if scanner.Scan() {
			nextLine = convert2(scanner.Text())
		} else {
			nextLine = nil
		}
	}
	log.Printf("Part 1 result %d", riskScore)
}
