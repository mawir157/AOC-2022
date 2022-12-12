package main

import AH "./adventhelper"

type Pos struct {
	x, y int
}

func makeGeography(ss []string) (geo [][]int, start Pos, end Pos) {
	geo = [][]int{}
	for i, s := range ss {
		row := []int{}
		for j, r := range s {
			if r == 'S' {
				row = append(row, 0)
				start = Pos{j,i}
			} else if r == 'E' {
				row = append(row, 26) // READ THE INSTRUCTIONS YOU MORON!
				end = Pos{j,i}
			} else {
				row = append(row, int(r) - 96)
			}
		}
		geo = append(geo, row)
	}

	return
}
 
func getNbhrs(loc Pos, gphy [][]int, step int) []Pos {
	target := gphy[loc.y][loc.x] + step
	nbrs := []Pos{}
	// Up
	if (loc.x > 0) {
		// fmt.Println("Up", gphy[loc.y][loc.x - 1])
		if (gphy[loc.y][loc.x - 1] <= target) {
			nbrs = append(nbrs, Pos{loc.x - 1, loc.y})
		}
	}
	// Down
	if (loc.x < len(gphy[0]) - 1) {
		// fmt.Println("Down", gphy[loc.y][loc.x + 1])
		if (gphy[loc.y][loc.x + 1] <= target) {
			nbrs = append(nbrs, Pos{loc.x + 1, loc.y})		
		}
	}
	// Left
	if (loc.y > 0) {
		// fmt.Println("Left", gphy[loc.y - 1][loc.x])
		if (gphy[loc.y - 1][loc.x] <= target) {
			nbrs = append(nbrs, Pos{loc.x, loc.y - 1})
		}
	}
	// Right
	if (loc.y < len(gphy) - 1) {
		// fmt.Println("Right", gphy[loc.y + 1][loc.x])
		if (gphy[loc.y + 1][loc.x] <= target) {
			nbrs = append(nbrs, Pos{loc.x, loc.y + 1})
		}
	}

	return nbrs
}

func minAlt(flagged map[Pos]int) (pMin Pos) {
	min := -1

	for k, v := range flagged {
		if (min == -1) || (v < min) {
			min = v
			pMin = k
		}
	}
	delete(flagged, pMin)

	return
}

func GraphTraverse (gphy [][]int, source Pos, target Pos) (dist map[Pos]int) {
	Q := make(map[Pos]bool)
	dist = make(map[Pos]int)
	marked := make(map[Pos]int)

	// set all vertices to infinity
	for y := 0; y < len(gphy); y++ {
		for x := 0; x < len(gphy[0]); x++ {
			p := Pos{x, y}
			Q[p] = false
			dist[p] = 1000000
		}
	}
	Q[source] = true
	dist[source] = 0
	marked[source] = 0

	for len(marked) > 0 {
		u := minAlt(marked)
		delete(Q, u)


		distU := dist[u]

		moves := getNbhrs(u, gphy, 1)

		for _, n := range moves {
			if _, ok := Q[n] ; ok {
				alt := distU + 1
				if alt < dist[n] {
					dist[n] = alt
					marked[n] = alt
				}
			}
		}
	}

	return
}

func main() {
	lines, _ := AH.ReadStrFile("../input/input12.txt")
	geography, s, e := makeGeography(lines)

	d := GraphTraverse(geography, s, e)

	part2 := 100000
	for i, v := range geography {
		for j, _ :=	range v {
			if geography[i][j] == 1 {
				p := Pos{j, i}
				d2 := GraphTraverse(geography, p, e)
				if d2[e] < part2 {
					part2 = d2[e]
				}
			}
		}
	}

	AH.PrintSoln(12, d[e], part2)

	return
}
