package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const stepLimit = 10

type element struct {
	value string
	next  *element
}

/*func expandAndSumarize(first, second string, table map[string]string, summary map[string]int, step int) {
	middle := table[first+second]

	summary[middle] = summary[middle] + 1
	if step == stepLimit {
		return
	}
	log.Printf("Occureneces %v", summary)
	expandAndSumarize(first, middle, table, summary, step + 1)
	expandAndSumarize(middle, second, table, summary, step + 1)
}*/

func day_14main() {
	file, err := os.Open("day14_input.txt")
	//file, err := os.Open("test_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var expansionTable = make(map[string]string)
	var occurrences = make(map[string]int)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	var input = strings.Split(scanner.Text(), "")
	for i := range input {
		if i < len(input)-1 {
			occurrences[input[i]+input[i+1]] = occurrences[input[i]+input[i+1]] + 1
		}
	}

	scanner.Scan()
	for scanner.Scan() {
		inputLine := strings.Split(scanner.Text(), " ")
		expansionTable[inputLine[0]] = inputLine[2]
	}

	for stepsCnt := 10; stepsCnt > 0; stepsCnt-- {
		log.Printf("%v", occurrences)
		newOccurences := make(map[string]int)
		for k, v := range occurrences {
			letters := strings.Split(k, "")
			middle := expansionTable[k]
			newOccurences[letters[0]+middle] = newOccurences[letters[0]+middle] + v
			newOccurences[middle+letters[1]] = newOccurences[middle+letters[1]] + v
		}
		occurrences = newOccurences
	}
	log.Printf("%v", occurrences)
	letterOccurences := make(map[string]int)
	letterOccurences[input[0]] = 1
	letterOccurences[input[len(input)-1]] = 1
	for k, v := range occurrences {
		letters := strings.Split(k, "")
		letterOccurences[letters[0]] = letterOccurences[letters[0]] + v
		letterOccurences[letters[1]] = letterOccurences[letters[1]] + v
	}

	var min, max = input[0], input[0]
	for k, v := range letterOccurences {
		if v > letterOccurences[max] {
			max = k
		}
		if v < letterOccurences[min] {
			min = k
		}
	}

	log.Printf("Result : %s :%d , %s, %d %d", max, letterOccurences[max], min, letterOccurences[min], (letterOccurences[max])/2-letterOccurences[min]/2)

}
