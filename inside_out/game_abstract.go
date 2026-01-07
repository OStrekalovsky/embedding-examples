package inside_out

/*
	Как обходить невозможность вызвать методы "внешней структуры" из "внутренней".
	Применение Client-Specified Self pattern.
*/

type Game interface {
	// методы будет абстрактыми, т.к. его не реализует "родитель".
	SetPlayers(players int)
	EndOfGame() bool
	DoTurn()
	// для этих двух методов сделаем "дефолтную реализацию"
	InitGame()
	PrintWinner()
}

// GameTemplate делается "абстрактным" и не реализует Game.
type GameTemplate struct{}

func (this *GameTemplate) InitGame() {
	println("Common: Game init")
}

func (this *GameTemplate) PrintWinner() {
	println("Common: printing game winner")
}

// ядро алгоритма.
func PlayGame(game Game, players int) {
	game.SetPlayers(players)
	game.InitGame()
	for !game.EndOfGame() {
		game.DoTurn()
	}
	game.PrintWinner()
}
