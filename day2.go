package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var depth = 0
var distance = 0
var aim = 0

func move(dir string, d int) {
	switch dir {
	case "up":
		depth -= d
	case "down":
		depth += d
	case "forward":
		distance += d
	}
}

func move2(dir string, d int) {
	switch dir {
	case "up":
		aim -= d
	case "down":
		aim += d
	case "forward":
		distance += d
		depth += d * aim
	}
}

func main_day2() {
	file, err := os.Open("/mnt/c/Users/matus.blaho/Downloads/input_day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := scanner.Text()
		res := strings.Split(x, " ")
		dist, _ := strconv.Atoi(res[1])
		move2(res[0], dist)
	}
	log.Printf("Result : %d, %d, %d", depth, distance, depth*distance)
}
