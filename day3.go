package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
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

func filter(input []string, requiredBit uint8, pos int) []string {
	outputPosition := 0

	for curPos, value := range input {
		if value[pos] == requiredBit {
			input[outputPosition] = value
			outputPosition++
		}
		if outputPosition == 0 && curPos == len(input)-1 {
			input[outputPosition] = value
			outputPosition++
		}
	}
	return input[:outputPosition]
}

func countZeroesAndOnes(input []string, pos int) (onesCount, zeroesCount int) {
	for _, x := range input {
		if x[pos] == '1' {
			onesCount++
		} else {
			zeroesCount++
		}
	}
	return
}

func getBigger(input []string, pos int) uint8 {
	o, z := countZeroesAndOnes(input, pos)
	if z <= o {
		return '1'
	} else {
		return '0'
	}
}

func getSmaller(input []string, pos int) uint8 {
	o, z := countZeroesAndOnes(input, pos)
	if z <= o {
		return '0'
	} else {
		return '1'
	}
}

func day3main() {
	//file, err := os.Open("test_input")
	file, err := os.Open("day3_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var oxygenGeneratorRatingOutput []string
	var co2ScrubberRatingOutput []string
	var count = 0
	for scanner.Scan() {
		count += 1
		oxygenGeneratorRatingOutput = append(oxygenGeneratorRatingOutput, scanner.Text())
		co2ScrubberRatingOutput = append(co2ScrubberRatingOutput, scanner.Text())
	}

	for pos := 0; pos < len(oxygenGeneratorRatingOutput[0]); pos++ {
		if len(oxygenGeneratorRatingOutput) > 1 {
			oxygenGeneratorRatingOutput = filter(oxygenGeneratorRatingOutput, getBigger(oxygenGeneratorRatingOutput, pos), pos)
		}

		if len(co2ScrubberRatingOutput) > 1 {
			co2ScrubberRatingOutput = filter(co2ScrubberRatingOutput, getSmaller(co2ScrubberRatingOutput, pos), pos)
		}
	}

	log.Printf("\nOxygen :   %v \nCO2 : %v", co2ScrubberRatingOutput, oxygenGeneratorRatingOutput)
	o, _ := strconv.ParseInt(co2ScrubberRatingOutput[0], 2, 32)
	co2, _ := strconv.ParseInt(oxygenGeneratorRatingOutput[0], 2, 32)
	log.Printf("Result %d", o*co2)
}
