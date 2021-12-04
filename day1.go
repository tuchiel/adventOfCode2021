package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var processingBuffer = make([]int, 4, 4)

func shiftIn(x int) {
	processingBuffer = processingBuffer[1:]
	processingBuffer = append(processingBuffer, x)
}

func main_day1() {
	file, err := os.Open("/mnt/c/Users/matus.blaho/Downloads/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ups := 0
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	processingBuffer[1], _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	processingBuffer[2], _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	processingBuffer[3], _ = strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		x := scanner.Text()
		log.Println(x)
		i, _ := strconv.Atoi(x)
		shiftIn(i)
		if (processingBuffer[0] + processingBuffer[1] + processingBuffer[2]) < (processingBuffer[1] + processingBuffer[2] + processingBuffer[3]) {
			ups += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Result : %d", ups)
}
