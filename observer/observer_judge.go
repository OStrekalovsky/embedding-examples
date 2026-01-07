package observer

import "fmt"

type Judge struct {
	ball         *Ball
	lastPosition *Position
}

func (j *Judge) Update() {
	j.lastPosition = j.ball.GetPosition()
	fmt.Println("Judge", "noticed that ball has moved to", j.lastPosition)
}

func NewJudge(ball *Ball) *Judge {
	this := new(Judge)
	this.ball = ball
	return this
}
