package main

import (
	"fmt"
)

// -1: nil, 0: dead, 1: alive
func getSingle(x int, y int, space [][] Cell, h int, w int) int {
	if x >= 0 && x < w && y >= 0 && y < h {
		if space[y][x].IsAlive() {
			return 1
		} else {
			return 0
		}
	} else {
		return -1
	}
}

type _pos struct {
	x int
	y int
}

func getAround(x int, y int, space [][] Cell) (int, int) {
	dead, alive, h, w := 0, 0, len(space), len(space[0]) 
	poses := [8] _pos { 
		_pos{x-1, y-1}, 
		_pos{x, y-1}, 
		_pos{x+1, y-1}, 
		_pos{x-1, y}, 
		_pos{x+1, y}, 
		_pos{x-1, y+1}, 
		_pos{x, y+1}, 
		_pos{x+1, y+1}, 
	}
	
	for _, pos := range poses {
		switch status := getSingle(pos.x, pos.y, space, h, w); status {
		case 0:
			dead++
		case 1:
			alive++
		}
	}

	return alive, dead
}

func rule1 (space [][] Cell) {
	for y, s := range space {
		for x, c := range s {
			if c.IsAlive() {
				fmt.Println()
				alive, dead := getAround(x, y, space)
				fmt.Printf("(%d, %d) %d %d\n", x, y, alive, dead)
			}
		}
	}
	
}

func rule2 (space [][] Cell) {
	
}

func rule3 (space [][] Cell) {
	
}

func rule4 (space [][] Cell) {
	
}

func main() {
	w := World{}
	w.Init(10,10,10)
	w.AddRule(rule1)
	w.Start()
}

