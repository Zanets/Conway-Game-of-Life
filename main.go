package main

import (
	"flag"
	"rule"
	"world"
)

func main() {

	h := flag.Int("h", 10, "Height of world.")
	w := flag.Int("w", 10, "Width of world.")
	alive := flag.Int("alive", 50, "Alive cell at start.")
	flag.Parse()

	world := world.World{}
	world.Init(*h, *w, *alive)
	world.AddRule(rule.Wiki)
	world.Start()
}

