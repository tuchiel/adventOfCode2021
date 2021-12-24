package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const origamiXSize = 1311
const origamiYSize = 1311

type origami [][]int

func foldSum(a, b int) (res int) {
	if res = a + b; res > 1 {
		res = 1
	}
	return
}
func (o *origami) fold(isX bool, axis int) {
	if isX {
		for i, _ := range *o {
			for s, h := axis-1, axis+1; s >= 0; s, h = s-1, h+1 {
				(*o)[i][s] = foldSum((*o)[i][s], (*o)[i][h])
			}
			(*o)[i] = (*o)[i][:axis]
		}
	} else {
		for s, h := axis-1, axis+1; s >= 0; s, h = s-1, h+1 {
			for i, _ := range (*o)[s] {
				(*o)[s][i] = foldSum((*o)[s][i], (*o)[h][i])
			}
		}
		*o = (*o)[:axis]
	}
}

func (o *origami) dotCount() (result int) {
	for _, row := range *o {
		for _, v := range row {
			result += v
		}
	}
	return
}

func (o *origami) print() {
	result := ""
	for _, row := range *o {
		result += "\n"
		for _, v := range row {
			if v == 1 {
				result += "#"
			} else {
				result += "."
			}
		}
	}
	log.Println(result)
	return
}

func myAtoi(s string) (res int) {
	res, _ = strconv.Atoi(s)
	return
}

func day13_main() {
	//file, err := os.Open("test_input.txt")

	file, err := os.Open("day13_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var o origami = make([][]int, origamiYSize)
	for i := 0; i < origamiYSize; i++ {
		o[i] = make([]int, origamiXSize)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputLine := scanner.Text()
		switch {
		case strings.HasPrefix(inputLine, "fold"):
			isX, val := func() (bool, int) { f := strings.Split(inputLine, "="); return f[0][len(f[0])-1] == 'x', myAtoi(f[1]) }()
			o.fold(isX, val)
			log.Printf("Dot count :%d", o.dotCount())
			//o.print()
			break
		case inputLine == "":
			log.Printf("Dot count :%d", o.dotCount())
			//o.print()
			break
		default:
			x, y := func() (int, int) { dot := strings.Split(inputLine, ","); return myAtoi(dot[0]), myAtoi(dot[1]) }()
			o[y][x] = 1
			break
		}
	}
	o.print()
}
