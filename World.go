package main

import (
	"math/rand"
	"fmt"
	"time"
)

type World struct {
	h int
	w int
	size int
	space [][] Cell
	rules [] func(space [][] Cell)
}

func (this *World) Init(h int, w int, aliveCell int) {
	this.h = h
	this.w = w
	this.size = h * w
	this.space = make( [][]Cell, h)

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
	fmt.Printf(" ")
	for x, _ := range this.space {
		fmt.Printf("%d ", x)
	}
	for y, i := range this.space {
		fmt.Println()
		fmt.Printf("%d ", y)
		for _, j := range i {
			if j.IsAlive() {
				fmt.Printf("* ")
			} else {
				fmt.Printf("  ")
			}
		}
	}
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

		for _, i := range this.space {
			for _, j := range i {
				j.Demand()
			}
		}
		this.Show()
	}
}
