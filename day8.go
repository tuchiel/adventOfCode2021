package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

var occureneces = [10]int{}
var translationTable = make(map[string]int)
var reverseTranslationTable = make(map[int]string)

func containsAll(container, what string) bool {
	for _, l := range strings.Split(what, "") {
		if !strings.Contains(container, l) {
			return false
		}
	}
	return true
}

func sameCount(a, b string) int {
	same := 0
	for _, l := range strings.Split(a, "") {
		if strings.Contains(b, l) {
			same++
			strings.ReplaceAll(b, l, "")
		}
	}
	return same
}

func day8_main() {
	//file, err := os.Open("test_input.txt")
	file, err := os.Open("day8_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result2 := 0
	for scanner.Scan() {
		translationTable = make(map[string]int)
		reverseTranslationTable = make(map[int]string)
		sequence := strings.Split(strings.TrimSuffix(strings.Split(scanner.Text(), "|")[0], " "), " ")
		input := strings.Split(strings.TrimPrefix(strings.Split(scanner.Text(), "|")[1], " "), " ")
		for i, _ := range sequence {
			s := strings.Split(sequence[i], "")
			sort.Strings(s)
			x := strings.Join(s, "")
			sequence[i] = ""
			switch len(x) {
			case 2:
				translationTable[x] = 1
				reverseTranslationTable[1] = x
			case 3:
				translationTable[x] = 7
				reverseTranslationTable[7] = x
			case 4:
				translationTable[x] = 4
				reverseTranslationTable[4] = x
			case 7:
				translationTable[x] = 8
				reverseTranslationTable[8] = x
			default:
				sequence[i] = x
			}
		}
		for j := range sequence {
			if sequence[j] != "" {
				if len(sequence[j]) == 6 && containsAll(sequence[j], reverseTranslationTable[4]) {
					translationTable[sequence[j]] = 9
					reverseTranslationTable[9] = sequence[j]
					sequence[j] = ""
				}
				if len(sequence[j]) == 5 && containsAll(sequence[j], reverseTranslationTable[7]) {
					translationTable[sequence[j]] = 3
					reverseTranslationTable[3] = sequence[j]
					sequence[j] = ""
				}
			}
		}
		for i := range sequence {
			if sequence[i] != "" {
				if len(sequence[i]) == 5 {
					if sameCount(sequence[i], reverseTranslationTable[9]) == 4 {
						translationTable[sequence[i]] = 2
						reverseTranslationTable[2] = sequence[i]
						sequence[i] = ""
					} else {
						translationTable[sequence[i]] = 5
						reverseTranslationTable[5] = sequence[i]
						sequence[i] = ""
					}
				}
			}
		}
		for i := range sequence {
			if sequence[i] != "" {
				if len(sequence[i]) == 6 {
					if containsAll(sequence[i], reverseTranslationTable[5]) {
						translationTable[sequence[i]] = 6
						reverseTranslationTable[6] = sequence[i]
						sequence[i] = ""
					} else {
						translationTable[sequence[i]] = 0
						reverseTranslationTable[0] = sequence[i]
						sequence[i] = ""
					}
				}
			}
		}
		log.Printf("Table : %v", translationTable)
		lineValue := 0
		for _, in := range input {
			s := strings.Split(in, "")
			sort.Strings(s)
			lineValue = lineValue * 10
			lineValue += translationTable[strings.Join(s, "")]
			switch len(in) {
			case 2:
				occureneces[1]++
			case 3:
				occureneces[7]++
			case 4:
				occureneces[4]++
			case 7:
				occureneces[8]++
			}
		}
		result2 += lineValue
		log.Printf("Line : %d", lineValue)
	}
	log.Printf("Result part 1 : %d", occureneces[1]+occureneces[7]+occureneces[4]+occureneces[8])
	log.Printf("Result part 1 : %d", result2)
}
