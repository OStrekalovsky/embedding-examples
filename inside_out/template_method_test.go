package inside_out

import (
	"testing"
)

func TestGame(t *testing.T) {
	chess := &Chess{}
	PlayGame(chess, 2)

	println()

	ttt := &TicTacToe{}
	PlayGame(ttt, 2)
}
