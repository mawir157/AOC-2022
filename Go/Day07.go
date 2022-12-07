package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

type Node struct {
    name string
    size int
    file bool
    parent *Node
    children []*Node
}


func (n *Node) dirSize() int {
		if n.file {
			return n.size
		} else {
			subsum := 0
			for _, c := range n.children {
				subsum += c.dirSize()
			}
    	return subsum
		} 
}

func changeDir(s string, n *Node) (*Node) {
	if s == ".." {
		return n.parent
	}

	for _, c := range n.children {
		if c.name == s {

			return c
		}
	}

	return n
}

func parseCLI(ss []string, n *Node) {
	ptr := n
	for _, s := range ss {
		if s == "$ ls" { // list everything in directory
			// nothing ? 
		} else if (s[0:4] == "$ cd") { // change pointer
			data := strings.Split(s, " ")
			ptr = changeDir(data[2], ptr)
		} else { // file / dir
			data := strings.Split(s, " ")
			if data[0] == "dir" { // is directory
				n_new := Node {name:data[1], size:0, file:false, parent:ptr, children:[]*Node{}}
				ptr.children = append(ptr.children, &n_new)
			} else { // is file
				fsize, _:= strconv.Atoi(data[0])
				n_new := Node {name:data[1], size:fsize, file:true, parent:ptr, children:[]*Node{}}
				ptr.children = append(ptr.children, &n_new)				
			}
		}
	}
}

func (n *Node)part1(size int) (int) {
	counter := 0
	if n.file {
		return 0
	} else {
		if n.dirSize() < size {
			counter += n.dirSize()
		}
	}

	for _, c := range n.children {
		counter += c.part1(size)
	}

	return counter
}

func (n *Node)part2(m map[string]int) {
	if n.file {
		return
	} else {
		m[n.name] = n.dirSize()
	}

	for _, c := range n.children {
		c.part2(m)
	}

	return
}

func main() {
	cli_instructions, _ := AH.ReadStrFile("../input/input07.txt")
	root := Node {name:"root", size:0, file:false, parent:nil, children:[]*Node{}}

	parseCLI(cli_instructions[1:], &root)

	toDelete := root.dirSize() - 40000000

	dirSizes := make(map[string]int)
	root.part2(dirSizes)

	part2 := root.dirSize()
	for _, v := range dirSizes {
		if (v > toDelete) && (v < part2) {
			part2 = v
		}
	}

	AH.PrintSoln(7, root.part1(100000), part2)

	return
}
