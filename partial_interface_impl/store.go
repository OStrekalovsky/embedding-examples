package partial_interface

type Store interface {
	Add(id int)
	Del(id int)
	Get(id int) int
}
