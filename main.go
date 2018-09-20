package main

import (
	"flag"
	"gameoflife/rule"
	"gameoflife/world"
)

func main() {

	h := flag.Int("h", 10, "Height of world.")
	w := flag.Int("w", 10, "Width of world.")
	alive := flag.Int("alive", 50, "Alive cell at start.")
	isSeed := flag.Bool("rand", false, "Start by random seed.")
	flag.Parse()

	world := world.World{}
	world.IsUsingSeed(*isSeed)
	world.Init(*h, *w, *alive)
	world.AddRule(rule.Wiki)
	world.Start()
}

