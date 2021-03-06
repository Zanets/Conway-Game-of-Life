package world

import (
	"math/rand"
	"fmt"
	"time"
)

type World struct {
	h int
	w int
	isUsingSeed bool
	size int
	space [][] Cell
	rules [] func(space [][] Cell)
}

func (this *World) IsUsingSeed(s bool) {
	this.isUsingSeed = s
}

func (this *World) Init(h int, w int, aliveCell int) {
	this.h = h
	this.w = w
	this.size = h * w
	this.space = make( [][]Cell, h)

	if this.isUsingSeed {
		rand.Seed(time.Now().UTC().UnixNano())
	}

	s := make([]int, this.size)
	for i := 0; i < this.size; i++ {
		s[i] = i
	} 

	for i := 0; i < h; i++ {
		this.space[i] = make( []Cell, w )
	} 

	for i := 0; i < aliveCell; i++ {
		r := rand.Intn(this.size - i)
		x := s[r] % w
		y := s[r] / w

		// fmt.Printf("(%d) Set %d,%d\n", s[r], y, x)
		this.space[y][x].Alive()
		this.space[y][x].Demand()
		s = append(s[:r], s[r+1:]...)
	}
}

func (this *World) Show() {
	fmt.Println()
	fmt.Printf("  ")
	for i := 0; i < this.w; i++ {
		fmt.Printf("- ")
	}
	for _, i := range this.space {
		fmt.Println()
		fmt.Printf("| ")
		for _, j := range i {
			if j.IsAlive() {
				fmt.Printf("%c ", j.Face())
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("| ")
	}
	fmt.Printf("\n  ")
	for i := 0; i < this.w; i++ {
		fmt.Printf("- ")
	}
	fmt.Println()
	time.Sleep(1 * time.Second)
}

func (this *World) AddRule( rule func(space [][] Cell) ) {
	this.rules = append(this.rules, rule)
}

func (this *World) Start() {
	this.Show()
	
	for true {
		for _, rule := range this.rules {
			rule(this.space)
		}

		for y, i := range this.space {
			for x, _ := range i {
				this.space[y][x].Demand()
			}
		}
		this.Show()
	}
}
