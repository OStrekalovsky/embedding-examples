package inside_out

import "math/rand"

type TicTacToe struct {
	*GameTemplate
}

func (g *TicTacToe) SetPlayers(players int) {
	println("SetPlayers")
}

func (g *TicTacToe) EndOfGame() bool {
	return rand.Int()%2 == 0
}

func (g *TicTacToe) DoTurn() {
	println("DoTurn")
}

// переопределим реализацию
func (g *TicTacToe) PrintWinner() {
	println("Tic tac toe PrintWinner")
}
