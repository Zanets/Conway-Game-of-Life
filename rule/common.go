package rule

import (
	"world"
)

type _pos struct {
	x int
	y int
}

// -1: nil, 0: dead, 1: alive
func getSingleLive(x int, y int, space [][] world.Cell, h int, w int) int {
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

func getAroundLive(x int, y int, space [][] world.Cell) (int, int) {
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
		switch status := getSingleLive(pos.x, pos.y, space, h, w); status {
		case 0:
			dead++
		case 1:
			alive++
		}
	}

	return alive, dead
}
