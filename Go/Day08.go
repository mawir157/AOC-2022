package main

import AH "./adventhelper"

import (
	"strconv"
)

type Pos struct {
	x, y int
}

func parseInput(ss []string) (map[Pos]int, int, int) {
	m := make(map[Pos]int)

	for i, s := range ss {
		for j, r := range s {
			val, _ := strconv.Atoi(string(r))
			ps := Pos {x:i, y:j}
			m[ps] = val
		}
	}

	return m, len(ss), len(ss[0])
}

func countVisible(forest map[Pos]int, dx int, dy int) int {
	vis := make(map[Pos]bool)

	//l-r
	for r := 0; r < dy; r++ {
		maxHeight := -1
		for c := 0; c < dx; c++ {
			ps := Pos{r, c}
			if forest[ps] > maxHeight {
				vis[ps] = true
				maxHeight = forest[ps]
			}
		}
	}

	// r-l
	for r := 0; r < dy; r++ {
		maxHeight := -1
		for c := dx - 1; c >= 0; c-- {
			ps := Pos{r, c}
			if forest[ps] > maxHeight {
				vis[ps] = true
				maxHeight = forest[ps]
			}
		}
	}

	// t-b
	for c := 0; c < dx; c++ {
		maxHeight := -1
		for r := 0; r < dy; r++ {
			ps := Pos{r, c}
			if forest[ps] > maxHeight {
				vis[ps] = true
				maxHeight = forest[ps]
			}
		}
	}

	// b-t
	for c := 0; c < dx; c++ {
		maxHeight := -1
		for r := dy - 1; r >= 0; r-- {
			ps := Pos{r, c}
			if forest[ps] > maxHeight {
				vis[ps] = true
				maxHeight = forest[ps]
			}
		}
	}

	return len(vis)
}

func scenicScore(forest map[Pos]int, xMax int, yMax int)  int {
	scenic := make(map[Pos]int)

	for k, height := range forest {
		x := k.x
		y := k.y

		// look left
		leftView := 0
		if (x == 0) {
			// do nothing, no trees to the left
		} else {
			for dx := k.x - 1; dx >= 0; dx-- {
				ps := Pos{dx, y}
				leftView++
				if forest[ps] < height {
				} else {
					break
				}
			} 
		}
		// look right
		rightView := 0
		if (x == xMax - 1) {
			// do nothing, no trees to the left
		} else {
			for dx := k.x + 1; dx < xMax; dx++ {
				ps := Pos{dx, y}
				rightView++
				if forest[ps] < height {
				} else {
					break
				}
			} 
		}
		// look up
		upView := 0
		if (y == 0) {
			// do nothing, no trees to the left
		} else {
			for dy := k.y - 1; dy >= 0; dy-- {
				ps := Pos{x, dy}
				upView++
				if forest[ps] < height {
				} else {
					break
				}
			} 
		}
		// look down
		downView := 0
		if (x == yMax - 1) {
			// do nothing, no trees to the left
		} else {
			for dy := k.y + 1; dy < yMax; dy++ {
				ps := Pos{x, dy}
				downView++
				if forest[ps] < height {
				} else {
					break
				}
			} 
		}

		scenic[k] = leftView * rightView * upView * downView
	}

	// find max scenic-ness
	mostScenic := 0
	for _, v := range scenic {
		if v > mostScenic {
			mostScenic = v
		}
	}

	return mostScenic
}

func main() {
	lines, _ := AH.ReadStrFile("../input/input08.txt")
	forest, dx, dy := parseInput(lines)

	AH.PrintSoln(8, countVisible(forest, dx, dy), scenicScore(forest, dx, dy))

	return
}
