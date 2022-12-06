package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

type Blocks []string
type Move struct {
	n, f, t int
}

func parseInput(ss []string) (bs Blocks) {
	bs = Blocks{"","","","","","","","",""}

	for _, s := range ss {
		for j := 1; j < len(s); j += 4 {
			r := AH.RuneAt(s, j)
			if (r != ' ') {
				bs[ j / 4 ] += string(r)
			}
		}
	}

	return
}

func parseMoves(ss []string) (mvs []Move) {
	mvs = []Move{}

	for _, s := range ss {
		ps := strings.Split(s, " ")
		n, _ := strconv.Atoi(ps[1])
		f, _ := strconv.Atoi(ps[3])
		t, _ := strconv.Atoi(ps[5])

		mvs = append(mvs, Move{n:n, f:(f - 1), t:(t - 1)})
	}

	return
}

func applyMove(v Blocks, m Move, b bool) {
	from_new := v[m.f][m.n:len(v[m.f])]
	moveable := v[m.f][0:m.n]

	if b {
		moveable = AH.ReverseString(moveable)
	}
	to_new := moveable + v[m.t]

	v[m.f] = from_new
	v[m.t] = to_new

	return
}

func main() {
	lines, _ := AH.ReadStrFile("../input/input05.txt")
	bricks1 := parseInput(lines[:8])
	bricks2 := parseInput(lines[:8])

	moves := parseMoves(lines[9:])

	for _, mv := range moves {
		applyMove(bricks1, mv, true)
		applyMove(bricks2, mv, false)
	}

	part1 := ""
	for _, b := range bricks1 {
		part1 += string(AH.FirstRune(b))
	}

	part2 := ""
	for _, b := range bricks2 {
		part2 += string(AH.FirstRune(b))
	}

	AH.PrintSoln(5, part1, part2)

	return
}
