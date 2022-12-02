package main

import AH "./adventhelper"

import (
	"strings"
)

type RPS int
const (
	ROCK RPS = iota
	PAPER
	SCISSORS
)

func score(round string)  (int, int) {
	pair := strings.Split(round, " ")
	player1, player2 := int(AH.FirstRune(pair[0])) - 65, int(AH.FirstRune(pair[1])) - 88

	result1 := (player2 - player1 + 3) % 3
	final1 := 0
	if (result1 == 0) {
		final1 = 3
	} else if (result1 == 1) {
		final1 = 6
	}

	final2 := 0
	if (player2 == 1) {
		final2 = 3
	} else if (player2 == 2) {
		final2 = 6
	}

	temp := player1
	if (player2 == 0) {
		temp += -1
	} else if (player2 == 2) {
		temp += 1
	}

	play2 := (temp + 3) % 3

	return 1 + player2 + final1, 1 + play2 + final2
}

func main() {
	rounds, _ := AH.ReadStrFile("../input/input02.txt")
	part1, part2 := 0, 0
	for _, round := range rounds {
		p1, p2 := score(round)
		part1 += p1
		part2 += p2
	}

	AH.PrintSoln(2, part1, part2)

	return
}
