package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const typeOceanMapSize = 1000

type oceanMap [typeOceanMapSize][typeOceanMapSize]int

func determineDirection(start, end int) int {
	if start > end {
		return -1
	}

	if start < end {
		return 1
	}

	return 0
}

func determineNextCoordinate(start, end, dir int) int {
	if start == end {
		return start
	} else {
		return start + dir
	}
}

func (om *oceanMap) DrawLine(startX, startY, endX, endY int) {

	var (
		xDirection          = determineDirection(startX, endX)
		yDirection          = determineDirection(startY, endY)
		markXCoordinate     = startX
		markYCoordinate     = startY
		nextMarkXCoordinate = -1
		nextMarkYCoordinate = -1
	)

	for {
		//log.Printf("Marking %d, %d", markXCoordinate, markYCoordinate)
		(*om)[markXCoordinate][markYCoordinate] = (*om)[markXCoordinate][markYCoordinate] + 1
		nextMarkXCoordinate = determineNextCoordinate(markXCoordinate, endX, xDirection)
		nextMarkYCoordinate = determineNextCoordinate(markYCoordinate, endY, yDirection)
		//log.Printf("Next %d, %d", nextMarkXCoordinate, nextMarkYCoordinate)
		if nextMarkXCoordinate == markXCoordinate && nextMarkYCoordinate == markYCoordinate {
			break
		} else {
			markXCoordinate = nextMarkXCoordinate
			markYCoordinate = nextMarkYCoordinate
		}
	}
}

func (om *oceanMap) GetValuesAbove(value int) (result int) {

	for _, row := range *om {
		for _, number := range row {
			if number > value {
				result++
			}
		}
	}
	return
}

func (om *oceanMap) toString() string {
	var s []string
	for i := range *om {
		var line []string
		for _, n := range (*om)[i] {
			line = append(line, strconv.Itoa(n))
		}
		s = append(s, strings.Join(line, ","))
	}
	return strings.Join(s, "\n")
}

func readInput(line string) (startX, startY, endX, endY int) {
	var startEnd = strings.Split(line, "->")
	var start = strings.Split(strings.ReplaceAll(startEnd[0], " ", ""), ",")
	var end = strings.Split(strings.ReplaceAll(startEnd[1], " ", ""), ",")
	startX, _ = strconv.Atoi(start[0])
	startY, _ = strconv.Atoi(start[1])
	endX, _ = strconv.Atoi(end[0])
	endY, _ = strconv.Atoi(end[1])
	return
}

func filterNonDiagonals(startX, startY, endX, endY int) bool {
	return (startX == endX) || (startY == endY)
}

func day5_main() {
	file, err := os.Open("day5_input.txt")
	//file, err := os.Open("test_input")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	om := &oceanMap{}
	for scanner.Scan() {
		var sx, sy, ex, ey = readInput(scanner.Text())
		//if filterNonDiagonals(sx, sy, ex, ey) { first part
		om.DrawLine(sx, sy, ex, ey)
		//}
	}
	//log.Println(om.toString())
	log.Printf("Result %d", om.GetValuesAbove(1))
}
