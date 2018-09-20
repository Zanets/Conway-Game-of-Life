package main

type Cell struct {
	isAlive bool
	isAlive2 bool
}	

func (this *Cell) Dead() {
	this.isAlive2 = false
}

func (this *Cell) Alive() {
	this.isAlive2 = true
}

func (this *Cell) Face() int32 {
	return '*'
}

func (this *Cell) IsAlive() bool {
	return this.isAlive
}

func (this *Cell) Demand() {
	this.isAlive = this.isAlive2
}
