package main

func getSingle(x int, y int) bool {

}

func getAround(x int, y int, space [][] Cell) (int, int) {
	dead, alive, _x, _y, h, w := 0, 0, 0, 0, len(space), len(space[0]) 
	
	_x, _y = i-1, j-1
	if _x >= 0 && _x < w && _y >= 0 _y < h && _y > 0 && space[_y][_x].isAlive() {
		alive++
	}

	return alive, dead
}

func rule1 (space [][] Cell) {
	h := len(space)
	for i, s := range space {
		for j, c := range s {
			if c.isAlive() {
				_x, _y  := 0, 0
				alive := 0

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
	w.Init(25,50,22)
	w.Start()
}

