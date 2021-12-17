package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type stack []string

func (s *stack) push(what string) {
	*s = append(*s, what)
}

func (s *stack) pop() (res string) {
	if len(*s) > 0 {
		res = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
	}
	return
}

func (s *stack) clear() {
	*s = (*s)[:0]
}

func part1() {
	file, err := os.Open("day10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	resultScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		unclosed.clear()
		for n, i := range strings.Split(scanner.Text(), "") {
			score := resultScore
			/*
				): 3 points.
				]: 57 points.
				}: 1197 points.
				>: 25137 points*/
			switch i {
			case ")":
				if unclosed.pop() != "(" {
					resultScore += 3
				}
			case "]":
				if unclosed.pop() != "[" {
					resultScore += 57
				}
			case "}":
				if unclosed.pop() != "{" {
					resultScore += 1197
				}
			case ">":
				if unclosed.pop() != "<" {
					resultScore += 25137
				}
			default:
				unclosed.push(i)
			}
			if resultScore != score {
				log.Printf("Failed at : %d, %s", n, i)
				break
			}
		}
	}
	log.Printf("Result : %d", resultScore)
}

var unclosed = stack(make([]string, 0, 100))
var scores = make([]int, 0, 100)

func part_two() {
	//file, err := os.Open("test_input.txt")
	file, err := os.Open("day10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		unclosed.clear()
		fail := false
		for _, i := range strings.Split(scanner.Text(), "") {
			switch i {
			case ")":
				fail = unclosed.pop() != "("
			case "]":
				fail = unclosed.pop() != "["
			case "}":
				fail = unclosed.pop() != "{"
			case ">":
				fail = unclosed.pop() != "<"
			default:
				unclosed.push(i)
			}
			if fail {
				break
			}
		}
		if !fail {
			score := 0
			for u := unclosed.pop(); u != ""; u = unclosed.pop() {
				/*
					): 1 point.
					]: 2 points.
					}: 3 points.
					>: 4 points.*/
				switch u {
				case "(":
					score = score*5 + 1
				case "[":
					score = score*5 + 2
				case "{":
					score = score*5 + 3
				case "<":
					score = score*5 + 4
				}
			}
			scores = append(scores, score)
			sort.Ints(scores)
		}
	}
	log.Printf("Result : %d", scores[len(scores)/2])
}
