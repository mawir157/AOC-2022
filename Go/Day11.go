package main

import AH "./adventhelper"

import (
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items []int
	operation string
	op_value int
	test int
	if_true int
	if_false int
	counter int
}

func parseInput(ss []string, sep string) []Monkey {
	monkeys := []Monkey{}

	for _, s := range ss {
		lines := strings.Split(s, sep)
		// parse items
		strItems := strings.Split(lines[1][18:], ", ")
		tempItems := []int{}
		for _, item := range strItems {
			n, _ := strconv.Atoi(item)
			tempItems = append(tempItems, n)
		}
		// parse operation
		tempOp := lines[2][23:24]
		tempOpValue, err := strconv.Atoi(lines[2][25:])
		if err != nil {
			tempOpValue = 0
		}
		// Parse Divisor
		tempDiv, _ := strconv.Atoi(lines[3][21:])
		// Parse IfTrue
		tempTrue, _ := strconv.Atoi(lines[4][29:])
		// Parse IfFalse
		tempFalse, _ := strconv.Atoi(lines[5][30:])

		monkeys = append(monkeys, Monkey{
			tempItems, tempOp, tempOpValue, tempDiv, tempTrue, tempFalse, 0})
	}

	return monkeys
}

func (m * Monkey) monkeySee(worry bool) (int, int) {
	m.counter += 1

	value := m.items[0]
	m.items = m.items[1:]

	if m.operation == "+" {
		if m.op_value == 0 {
			value += value
		} else {
			value += m.op_value
		}
	} else {
		if m.op_value == 0 {
			value *= value
		} else {
			value *= m.op_value
		}		
	}

	if worry {
		value /= 3
	}

	if value % m.test == 0 {
		return value, m.if_true
	} else {
		return value, m.if_false
	}
}

func monkeyDo(ms []Monkey, worry bool) {
	reducer := 1
	for _, m := range ms {
		reducer *= m.test
	} 

	for index := 0; index < len(ms); index++ {
		for _, _ = range ms[index].items {
			val, to := ms[index].monkeySee(worry)
			ms[to].items = append(ms[to].items, val % reducer)
		}
	}

	return
}

func monkeyBusiness(ms []Monkey) int {
	business := []int{}
	for _, m := range ms {
		business = append(business, m.counter)
	}

	sort.Slice(business, func(i, j int) bool { return business[i] > business[j] })
	return (business[0] * business[1])
}

func main() {
	lines, _ := AH.ParseLineGroups("../input/input11.txt", "|")
	ms1 := parseInput(lines, "|")
	ms2 := parseInput(lines, "|")

	for i := 0; i < 20; i++ {
		monkeyDo(ms1, true)
	}
	for i := 0; i < 10000; i++ {
		monkeyDo(ms2, false)
	}

	AH.PrintSoln(11, monkeyBusiness(ms1), monkeyBusiness(ms2))

	return
}
