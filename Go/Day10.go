package main

import AH "./adventhelper"

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	d string
	n int
}

func parseInput(s string) Instruction {
	ps := strings.Split(s, " ")
	n := 0
	if ps[0] == "addx" {
		n, _ = strconv.Atoi(string(ps[1]))
	}
	return Instruction{ps[0], n}
}

func runInstruction(is []Instruction) []int {
	reg := 1
	values := []int{reg}
	for _, ins := range is {
		values = append(values, reg)
		if ins.d == "addx" {
			reg += ins.n
			values = append(values, reg)
		}
	}

	return values
}

func part1(ns []int, is []int) (int) {
	total := 0
	for _, i := range is {
		total += (i * ns[i - 1])
	}

	return total
}

func renderCRT(ns []int) string {
	crt := ""
	for pixel, n := range ns {
		if ((n - 1) <= (pixel % 40)) && ((pixel % 40) <= (n + 1)) {
			crt = crt + "#"
		} else {
			crt = crt + " "
		}
	}

	return crt
}

func main() {
	lines, _ := AH.ReadStrFile("../input/input10.txt")
	is := []Instruction{}
	for _, l := range lines {
		is = append(is, parseInput(l))
	}

	out := runInstruction(is)
	idx := []int{20, 60, 100, 140, 180, 220}
	crt := renderCRT(out)

	AH.PrintSoln(10, part1(out, idx), 0)

	for i := 0; i < 6; i++ {
		fmt.Println(crt[40*i:40*(i+1)])
	}

	return
}
