package observer

import (
	"testing"
)

/*
	Встраивание даёт возможность легко реализовать нужный интерфейс: Игроки и Судья подписываются на изменение положение мяча.
*/

func Test(t *testing.T) {
	ball := NewFootBall()
	player1 := NewPlayer("player1", ball)
	player2 := NewPlayer("player2", ball)
	player3 := NewPlayer("player3", ball)
	judge := NewJudge(ball)

	ball.Attach(player1)
	ball.Attach(player2)
	ball.Attach(player3)
	ball.Attach(judge)

	aPosition := NewPosition(1, 2, 3)
	ball.SetPosition(aPosition)
	ball.Detach(player2)
	ball.SetPosition(NewPosition(2, 3, 4))
}
