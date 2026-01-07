package partial_interface

type onlyGetMock struct {
	Store
}

func (g onlyGetMock) Get(_ int) int {
	return 42
}
