package observer

import "fmt"

type Player struct {
	name         string
	lastPosition *Position
	ball         *Ball
}

func (p *Player) Update() {
	p.lastPosition = p.ball.GetPosition()
	fmt.Println(p.name, "noticed that ball has moved to", p.lastPosition)
}

func NewPlayer(name string, ball *Ball) *Player {
	this := new(Player)
	this.name = name
	this.ball = ball
	return this
}
