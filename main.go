package main


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

// rules from wiki
func default_rule (space [][] Cell) {
	for y, s := range space {
		for x, c := range s {
			alive, _ := getAround(x, y, space)
			if c.IsAlive() {
				// fmt.Printf("(%d, %d) %d %d", x, y, alive, dead)
				if alive < 2 {
					space[y][x].Dead()
				} else if alive == 2 || alive == 3 {
					// do nothing
				} else if alive > 3 {
					space[y][x].Dead()
				} 
			} else {
				if alive == 3 {
					space[y][x].Alive()
				}
			}
		}
	}
	
}

func main() {
	w := World{}
	w.Init(10,10,50)
	w.AddRule(default_rule)
	w.Start()
}

