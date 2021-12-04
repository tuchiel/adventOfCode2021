package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

var zeroes []int
var ones []int

func compute(inputNum string) {
	if zeroes == nil {
		zeroes = make([]int, len(inputNum))
		ones = make([]int, len(inputNum))
	}
	for pos, char := range inputNum {
		if char == '1' {
			ones[pos] = ones[pos] + 1
		} else {
			zeroes[pos] = zeroes[pos] + 1
		}
	}
}

func powTwo(x int) int {
	return int(math.Pow(float64(2), float64(x)))
}

func partOne() {
	file, err := os.Open("/mnt/c/Users/matus.blaho/Downloads/input_day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		compute(scanner.Text())
	}
	gammaRate := 0
	epsilonRate := 0

	gammaRateStr := ""
	epsilonRateStr := ""

	log.Printf("Ones   %v", ones)
	log.Printf("Zeroes %v", zeroes)

	for pos, num := range ones {
		if zeroes[pos] < num {
			gammaRate += powTwo(len(ones) - pos - 1)
			gammaRateStr += "1"
			epsilonRateStr += "0"
		} else {
			epsilonRate += powTwo(len(ones) - pos - 1)
			gammaRateStr += "0"
			epsilonRateStr += "1"
		}
	}
	log.Printf("\nGamma :   %s \nEpsilon : %s", gammaRateStr, epsilonRateStr)
	log.Printf("\nGamma :   %d \nEpsilon : %d", gammaRate, epsilonRate)
	log.Printf("Result %d", epsilonRate*gammaRate)
}

func filter(inFile *bufio.Scanner, outFile *io.Writer, requiredBit uint8, pos int) (result string) {
	x := ""
	count := 0
	for inFile.Scan() {
		x = inFile.Text()
		if x[pos] == requiredBit {
			count += 1
			fmt.Fprintln(*outFile, x)
		}
	}
	if count == 0 {
		return x
	}
	return
}

func main() {
	file, err := os.Open("/mnt/c/Users/matus.blaho/Downloads/input_day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	b := bufio.NewWriter(file)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	oxygenGeneratorRating := ""
	co2ScrubberRating := ""

}
