package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func (p Pos) drop(dx int, dy int) Pos {
	p.x += dx
	p.y += dy
	return p
}

func makeWalls(ss []string, space map[Pos]int) {
	var x_prev, y_prev int
	for _, s := range ss {
		parts := strings.Split(s, " -> ")
		for i, p := range parts {
			indices := strings.Split(p, ",")
			x, _ := strconv.Atoi(indices[0])
			y, _ := strconv.Atoi(indices[1])
			if i != 0 {
				if x != x_prev {
					for ix := x_prev; ix != x; ix += AH.Sign(x - x_prev) {
						space[Pos{ix, y}] = 1
					}
				}
				if y != y_prev {
					for iy := y_prev; iy != y; iy += AH.Sign(y - y_prev) {
						space[Pos{x, iy}] = 1
					}
				}
				space[Pos{x, y}] = 1
			}
			x_prev = x
			y_prev = y
		}
	}

	return
}

func findFloor(space map[Pos]int) int{
	max_y := 0
	for p, _ := range space {
		if p.y > max_y {
			max_y = p.y
		}
	}

	return max_y
}
 
func dropSand(source Pos, space map[Pos]int, floor int) bool {
	carry_on := true
	point := source
	for carry_on {
		
		// scan below
		probe := point.drop(0, 1)
		if floor != 0 {
			if probe.y == floor {
  			space[point] = 2
  			return true				
			}
		}

		if _, blocked := space[probe]; blocked { // the cell directly below is blocked
    	probe = point.drop(-1, 1) // try down and to the left
    	if _, blocked := space[probe]; blocked { // this cell is also blocked
    		probe = point.drop(1, 1) // try down and to the right
    		if _, blocked := space[probe]; blocked { // this cell is also blocked
    			// all the cell below are blocked so our journey ends here
    			space[point] = 2
					if point == source { // if we at the source no more sand can fall
						return false
					} else {
    				return true
					}
    		} else {
    			point = probe // go round again?
    		}
    	} else {
				point = probe // go round again?
    	}
		} else {
			point = probe // go round again?
		}

		if point.y > 500 { // fallen off the end of the world!
			carry_on = false
		}
	}
	return carry_on
}

func keepDroppingSand(p Pos, space map[Pos]int, floor int) (count int){
	ok := true
	count = 0
	for ok {
		ok = dropSand(p, space, floor)
		if ok {
			count++
		}
	}

	return
}

func main() {
	lines, _ := AH.ReadStrFile("../input/input14.txt")
	space := make(map[Pos]int)
	makeWalls(lines, space)

	part1 := keepDroppingSand(Pos{500, 0}, space, 0)

	floor := findFloor(space)
	part2 := keepDroppingSand(Pos{500, 0}, space, floor + 2)

	AH.PrintSoln(14, part1, part1 + part2 + 1)

	return
}
