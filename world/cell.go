package world

type Cell struct {
	isAlive bool
	isAliveThen bool
}	

func (this *Cell) Dead() {
	this.isAliveThen = false
}

func (this *Cell) Alive() {
	this.isAliveThen = true
}

func (this *Cell) Face() int32 {
	return '*'
}

func (this *Cell) IsAlive() bool {
	return this.isAlive
}

func (this *Cell) Demand() {
	this.isAlive = this.isAliveThen
}
