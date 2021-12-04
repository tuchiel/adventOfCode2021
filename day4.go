package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	boardSize = 5
)

type bingoNumber struct {
	value  int
	marked bool
}

func (n *bingoNumber) load(sn string) {
	n.value, _ = strconv.Atoi(sn)
}

func (n *bingoNumber) mark(picked int) bool {
	if !n.marked {
		n.marked = picked == n.value
	}
	return n.marked
}

type numbersCollection struct {
	numbers [boardSize]*bingoNumber
}

func (c *numbersCollection) load(numbers string) {
	i := 0
	for _, n := range strings.Split(numbers, " ") {
		if n != "" {
			c.numbers[i] = &bingoNumber{marked: false}
			c.numbers[i].load(n)
			i++
		}

	}
}

func (r *numbersCollection) isFullyMarkedBy(picked int) (result bool) {
	result = true
	for _, n := range r.numbers {
		if !n.mark(picked) {
			result = false
		}
	}
	return
}

func (r *numbersCollection) isFullyMarked() bool {
	for _, n := range r.numbers {
		if !n.marked {
			return false
		}
	}
	return true
}

func (c *numbersCollection) unmarkedSum() (res int) {
	res = 0
	for _, n := range c.numbers {
		if !n.marked {
			log.Printf("Adding %d", n.value)
			res += n.value
		}
	}
	return
}

type board struct {
	rows        [boardSize]numbersCollection
	columns     [boardSize]numbersCollection
	fullyMarked bool
}

func newBoard() (b *board) {
	b = new(board)
	b.fullyMarked = false
	for i := 0; i < boardSize; i++ {
		b.rows[i] = numbersCollection{}
		b.columns[i] = numbersCollection{}
	}
	return
}

func (b *board) load(rows []string) {
	for i, r := range rows {
		b.rows[i].load(r)
		for j := 0; j < boardSize; j++ {
			b.columns[j].numbers[i] = b.rows[i].numbers[j]
		}
	}
}

func (b *board) isFullyMarkedBy(picked int) bool {
	if b.fullyMarked {
		return true
	}
	for _, r := range b.rows {
		if r.isFullyMarkedBy(picked) {
			b.fullyMarked = true
		}
	}
	for _, c := range b.columns {
		if c.isFullyMarked() {
			b.fullyMarked = true
		}
	}
	return b.fullyMarked
}

func (b *board) unmarkedSum() (res int) {
	res = 0
	for _, r := range b.rows {
		res += r.unmarkedSum()
	}
	return
}

func day4_main() {
	file, err := os.Open("day4_input.txt")
	//file, err := os.Open("test_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numbers := scanner.Text()
	var rowNum int
	var boards []*board
	var boardData = make([]string, boardSize)
	var loadedBoard *board
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			if rowNum == boardSize {
				rowNum = 0
				loadedBoard.load(boardData)
				boards = append(boards, loadedBoard)
			}
			loadedBoard = newBoard()
		} else {
			boardData[rowNum] = row
			rowNum++
		}
	}
	loadedBoard.load(boardData)
	boards = append(boards, loadedBoard)
	log.Printf("Boards count %d", len(boards))

	var markedBoards = 0

	for _, n := range strings.Split(numbers, ",") {
		number, _ := strconv.Atoi(n)
		for i, b := range boards {
			if !b.fullyMarked {
				if b.isFullyMarkedBy(number) {
					if markedBoards == 0 {
						log.Printf("Winning board : %d", i+1)
						log.Printf("Winning sum %d", b.unmarkedSum())
						log.Printf("Result : %d", number*b.unmarkedSum())
					}
					markedBoards++
					if markedBoards == len(boards) {
						log.Printf("Loosing board : %d", i+1)
						log.Printf("Loosing sum %d", b.unmarkedSum())
						log.Printf("Result : %d", number*b.unmarkedSum())
						return
					}
				}
			}
		}
	}

}
