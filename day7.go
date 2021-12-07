package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convert(input []string) (output []int) {
	output = make([]int, 0, len(input))
	for _, str := range input {
		var val, _ = strconv.Atoi(str)
		output = append(output, val)
	}
	return
}

func median(n ...int) int {
	mNumber := len(n) / 2

	if len(n)%2 == 1 {
		return n[mNumber]
	}

	return (n[mNumber-1] + n[mNumber]) / 2
}

var counts = make(map[int]int)

func getDistanceFrom(x int) (dist int) {
	for k, v := range counts {
		if k != x {
			dist += int(math.Abs(float64(x)-float64(k))) * v
		}
	}
	return dist
}

func getDistanceFrom2(x int) (dist int) {
	for k, v := range counts {
		if k != x {
			price := 0
			for i := 1; i <= int(math.Abs(float64(x)-float64(k))); i++ {
				price += i
			}
			dist += price * v
		}
	}
	return dist
}

func day7_main() {
	//file, err := os.Open("test_input")
	file, err := os.Open("day7_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := convert(strings.Split(scanner.Text(), ","))
	sort.Ints(input)
	med := median(input...)

	for _, x := range input {
		if v, e := counts[x]; e {
			counts[x] = v + 1
		} else {
			counts[x] = 1
		}

	}
	/*
		first part
		dist := getDistanceFrom(med)
	*/

	dist := getDistanceFrom2(med)

	for i := med + 1; i < input[len(input)-1]; i++ {
		newDist := getDistanceFrom2(i)
		if newDist < dist {
			dist = newDist
			log.Printf("New dist : %d at %d", dist, i)
		} else {
			break
		}
	}

	for i := med - 1; i < input[0]; i-- {
		newDist := getDistanceFrom2(i)
		if newDist < dist {
			dist = newDist
			log.Printf("New dist : %d at %d", dist, i)
		} else {
			break
		}
	}

	log.Printf("Result %d", dist)
}
