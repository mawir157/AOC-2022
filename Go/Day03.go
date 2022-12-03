package main

import AH "./adventhelper"

func stringIntersection(s1 string, s2 string) (intersection string) {
	set := make(map[rune]bool)

	for _, c := range s1 {
		set[c] = true
	}

	intersection = ""
	for _, c := range s2 {
		if _, ok := set[c]; ok {
			intersection = intersection + string(c)
		}
	}

	return
}

func score(bags []string) (total int) {
	total = 0

	for _, bag := range bags {
		half := len(bag) / 2
		lhs := bag[:half]
		rhs := bag[half:]

		intersection := stringIntersection(lhs, rhs)
		score := int(AH.FirstRune(intersection))
		if (score <= 90) {
			score -= 38
		} else {
			score -= 96
		}

		total += score
	}

	return
}

func assignBadges(bags []string) (total int) {
	total = 0
	for i := 0; i < len(bags); i = i + 3 {
		i1 := stringIntersection(bags[i], bags[i+1])
		i2 := stringIntersection(bags[i+2], i1)

		score := int(AH.FirstRune(i2))
		if (score <= 90) {
			score -= 38
		} else {
			score -= 96
		}
		total += score
	}

	return
}

func main() {
	bags, _ := AH.ReadStrFile("../input/input03.txt")

	AH.PrintSoln(3, score(bags), assignBadges(bags))

	return
}
