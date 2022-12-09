package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type Move struct {
	d string
	n int
}

func parseInput(s string) Move {
	ps := strings.Split(s, " ")
	n, _ := strconv.Atoi(string(ps[1]))
	return Move{ps[0], n}
}

func updateT(posT Pos, posH Pos) Pos {
	dx := posH.x - posT.x
	dy := posH.y - posT.y

	if (AH.AbsInt(dx) > 1) || (AH.AbsInt(dy) > 1) {
		return Pos{posT.x + AH.Sign(dx), posT.y + AH.Sign(dy)}
	} else {
		return posT
	}
}

func pullChain(chain []Pos, histTail map[Pos]int, m Move) ([]Pos) {
		for i := 0; i < m.n; i++ {
		switch m.d {
			case "U":
				chain[0].y -= 1
			case "D":
				chain[0].y += 1
			case "L":
				chain[0].x -= 1
			case "R":
				chain[0].x += 1
		}

		for i := 1; i < len(chain); i++ {
			chain[i] = updateT(chain[i], chain[i-1])
		}

		histTail[chain[len(chain) - 1]] += 1
	}

	return chain
}

func main() {
	lines, _ := AH.ReadStrFile("../input/input09.txt")
	moves := []Move{}
	for _, l := range lines {
		moves = append(moves, parseInput(l))
	}
	hist1 := make(map[Pos]int)
	hist2 := make(map[Pos]int)
	chain1 := []Pos{Pos{0,0}, Pos{0,0}}
	chain2 := []Pos{Pos{0,0}, Pos{0,0}, Pos{0,0}, Pos{0,0}, Pos{0,0}, Pos{0,0},
	               Pos{0,0}, Pos{0,0}, Pos{0,0}, Pos{0,0}}

	for _, m := range moves {
		chain1 = pullChain(chain1, hist1, m)
		chain2 = pullChain(chain2, hist2, m)
	}

	AH.PrintSoln(9, len(hist1), len(hist2))

	return
}
