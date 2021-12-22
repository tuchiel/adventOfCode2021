package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const mapSize = 10

type octopus struct {
	v          int
	neighbours []*octopus
}
type octopusMatrix [mapSize][mapSize]*octopus

func (o *octopus) increase() {
	if o.v++; o.v == 10 {
		for _, n := range o.neighbours {
			if n != nil {
				n.increase()
			}
		}
	}
}

func (om *octopusMatrix) get(row, col int) *octopus {
	if row >= 0 && row < mapSize && col >= 0 && col < mapSize {
		if (*om)[row][col] == nil {
			om.add(row, col, 0)
		}
		return (*om)[row][col]
	}
	return nil
}

func (om *octopusMatrix) print() {
	var str = ""
	for _, octoline := range *om {
		str += "\n"
		for _, o := range octoline {
			str += fmt.Sprint(o.v)
		}

	}
	log.Println(str)
}

func (om *octopusMatrix) add(row, col int, val int) {
	if (*om)[row][col] == nil {
		(*om)[row][col] = &octopus{v: val}
		(*om)[row][col].neighbours = make([]*octopus, 0, 8)
		(*om)[row][col].neighbours = append((*om)[row][col].neighbours, om.get(row, col-1))
		(*om)[row][col].neighbours = append((*om)[row][col].neighbours, om.get(row, col+1))
		(*om)[row][col].neighbours = append((*om)[row][col].neighbours, om.get(row-1, col-1))
		(*om)[row][col].neighbours = append((*om)[row][col].neighbours, om.get(row-1, col))
		(*om)[row][col].neighbours = append((*om)[row][col].neighbours, om.get(row-1, col+1))
		(*om)[row][col].neighbours = append((*om)[row][col].neighbours, om.get(row+1, col-1))
		(*om)[row][col].neighbours = append((*om)[row][col].neighbours, om.get(row+1, col))
		(*om)[row][col].neighbours = append((*om)[row][col].neighbours, om.get(row+1, col+1))
	} else {
		(*om)[row][col].v = val
	}

}

func (om *octopusMatrix) collectFlashes() (flashesCount int) {
	for _, octoline := range *om {
		for _, o := range octoline {
			o.increase()
		}
	}

	for _, octoLine := range *om {
		for i, _ := range octoLine {
			if octoLine[i].v > 9 {
				octoLine[i].v = 0
				flashesCount++
			}
		}
	}
	return
}

func day11_main() {
	//file, err := os.Open("test_input.txt")

	file, err := os.Open("day11_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var energyMap octopusMatrix
	var stepCount = 100
	scanner := bufio.NewScanner(file)
	var r = 0
	for scanner.Scan() {
		for c, val := range convert2(scanner.Text()) {
			energyMap.add(r, c, val)
		}
		r++
	}
	var result = 0
	for i := 0; ; i++ {
		mid_result := energyMap.collectFlashes()
		if i < stepCount {
			result += mid_result
		}

		if mid_result == 100 {
			log.Printf("First sync %d", i+1)
			break
		}
		if i == 194 {
			energyMap.print()
		}

	}
	log.Printf("Number of flashes %d", result)

}
