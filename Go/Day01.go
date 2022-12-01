package main

import AH "./adventhelper"

import (
	"sort"
	"strconv"
	"strings"
)

func calorieCount(cals string, sep string) (total int) {
	total = 0
	cs := strings.Split(cals, sep)
	for _, c := range cs {
		n, _ := strconv.Atoi(c)
		total += n
	}
	return
}

func main() {
	elfs, _ := AH.ParseLineGroups("../input/input01.txt", "|")

	calories := []int{}
	for _, elf := range elfs {
		cals := calorieCount(elf, "|")
		calories = append(calories, cals)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	AH.PrintSoln(1, calories[0], calories[0] + calories[1] + calories[2])

	return
}
