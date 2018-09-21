package rule

import (
	"github.com/Zanets/Conway-Game-of-Life/world"
)

// rules from wiki
func Wiki (space [][] world.Cell) {
	for y, s := range space {
		for x, c := range s {
			alive, _ := getAroundLive(x, y, space)
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
