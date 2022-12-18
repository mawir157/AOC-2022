package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

// global variables, cos I'm a lazy bastard living in a suit
var Rates []int // the flow rate of the non-zero valves (and AA)
var DistMat [][]int // the distance between non-zero valves (and AA)
var DP []int // the best score for each state in the DP recursion

type Valve struct {
	flow int
	nbhrs []string
}

func buildGraph(ss []string) (valves map[string]Valve) {
	valves = make(map[string]Valve)
	for _, s := range ss {
		label := AH.Take(AH.Drop(s, 6), 2)
		ps := strings.Split(s, "; ")
		flow_str := AH.Drop(ps[0], 23)
		flow, _ := strconv.Atoi(flow_str)

		nbrs_str := ""
		if AH.RuneAt(ps[1], 6) == 's' {
			nbrs_str = AH.Drop(ps[1], 23)
		} else {
			nbrs_str = AH.Drop(ps[1], 22)
		}
		nbhrs := strings.Split(nbrs_str, ", ")

		valve := Valve{flow:flow, nbhrs:nbhrs}
		valves[label] = valve
	}

	return
}

func minAlt(flagged map[string]int) (pMin string) {
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

func distPair(valves map[string]Valve, s string) (dist map[string]int) {
	dist = make(map[string]int)
	marked := make(map[string]int)

	// set all vertices to infinity
	for k, _ := range valves {
		dist[k] = 1000000
	}
	dist[s] = 0
	marked[s] = 0

	for len(marked) > 0 {
		u := minAlt(marked)
		distU := dist[u]
		nbhrs := valves[u].nbhrs

		for _, n := range nbhrs {
			alt := distU + 1
			if alt < dist[n] {
				dist[n] = alt
				marked[n] = alt
			}
		}
	}

	return
}

func buildDistMatrix(valves map[string]Valve) {
	lut := make(map[string]int)
	revLut := []string{}

	revLut = append(revLut, "AA")
	lut["AA"] = 0
	Rates = append(Rates, 0)

	index := 1
	for k, v := range valves {
		if v.flow > 0 {
			Rates = append(Rates, v.flow)
			lut[k] = index
			revLut = append(revLut, k)
			index++
		}
	}

	for _, s := range revLut {
		temp := []int{}
		q := distPair(valves, s)
		for _, e := range revLut {
			temp = append(temp, q[e])
		}
		DistMat = append(DistMat, temp)		
	}

	n := (1 << len(revLut)) * len(revLut) * 31 * 2
	DP = make([]int, n)
	for i:= 0; i < len(DP); i++{
	    DP[i] = -1
	}

	return
}

func dynamic(pos int, open int, time int, elephant int) int {
	if time <= 0 {
		if elephant > 0 { // set the elephant going...
			return dynamic(0, open, 26, elephant - 1)
		} else {
			return 0;
		}
	}

	// convert to state index
  key := (open * len(Rates) * 31 * 2) + (pos * 31 * 2) + (time * 2) + elephant;
  if DP[key] >= 0 { // have we already calculated this state...
    return DP[key]; // so just return it
  }

  ans := 0
  if ((open & (1 << pos) == 0)) { // we haven't opened this valve yet...
  	newOpen := open | (1 << pos) // so open it and recur
  	ans = AH.MaxPair(ans, (time - 1) * Rates[pos] + dynamic(pos, newOpen, time-1, elephant))
  }

  for i, t := range DistMat[pos] { // move to another valve
  	if i != pos { // stay where we are? no!, t =  time taken to get there
	  		ans = AH.MaxPair(ans, dynamic(i, open, time - t, elephant))
	  }
  }

  DP[key] = ans
  return ans;
} 


func main() {
	lines, _ := AH.ReadStrFile("../input/input16.txt")
	// parse the input
	valves := buildGraph(lines)
	// use Dijkstra to build a distance matrix between non-zero valves
	buildDistMatrix(valves)
	// use DP to get the answers
	part1 := dynamic(0,0, 30, 0)
	part2 := dynamic(0,0, 26, 1)
	AH.PrintSoln(16, part1, part2)

	return
}
