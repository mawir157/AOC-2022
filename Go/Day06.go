package main

import AH "./adventhelper"

func getMarker(s string, n int) int {
	runes := []rune(s)

	for i := 0; i < len(runes) - n; i++ {
		p := runes[i:i+n]

		hash := make(map[rune]bool)

		for _, r := range p {
			hash[r] = true
		}

		if len(hash) == n {
			return i + n
		}
	}

	return 0
}

func main() {
	space, _ := AH.ReadStrFile("../input/input06.txt")

	AH.PrintSoln(6, getMarker(space[0], 4), getMarker(space[0], 14))

	return
}
