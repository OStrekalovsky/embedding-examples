package inside_out

import "math/rand"

type Chess struct {
	*GameTemplate
}

func (g *Chess) SetPlayers(players int) {
	println("SetPlayers for Chess")
}

func (g *Chess) EndOfGame() bool {
	return rand.Int()%2 == 0
}

func (g *Chess) DoTurn() {
	println("DoTurn")
}
