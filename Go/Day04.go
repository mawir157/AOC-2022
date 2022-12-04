package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

func check(s string) (bool, bool) {
	bounds := strings.Split(s, ",")
	lhs, rhs := strings.Split(bounds[0], "-"), strings.Split(bounds[1], "-")
	l1, _ := strconv.Atoi(lhs[0])
	l2, _ := strconv.Atoi(lhs[1])
	r1, _ := strconv.Atoi(rhs[0])
	r2, _ := strconv.Atoi(rhs[1])

	part1 := ((l1 <= r1) && (r2 <= l2)) || ((r1 <= l1) && (l2 <= r2))
	part2 := ((l2 >= r1) && (r2 >= l1))

	return part1, part2
}

func main() {
	space, _ := AH.ReadStrFile("../input/input04.txt")

	part1, part2 := 0, 0
	for _, s := range space {
		p1, p2 := check(s)
		if (p1) {
			part1 += 1
		}
		if (p2) {
			part2 += 1
		}
	}

	AH.PrintSoln(4, part1, part2)

	return
}
