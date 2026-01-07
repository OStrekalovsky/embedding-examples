package observer

type Ball struct {
	// Чтобы тип стал "наблюдаем" достаточно встроить в него DefaultSubject.
	*DefaultSubject
	position *Position
}

func (b *Ball) GetPosition() *Position {
	return b.position
}

func (b *Ball) SetPosition(position *Position) {
	b.position = position
	b.Notify()
}

func NewFootBall() *Ball {
	return &Ball{DefaultSubject: NewDefaultSubject()}
}

type Position struct {
	x, y, z int
}

func NewPosition(x, y, z int) *Position {
	return &Position{x, y, z}
}
