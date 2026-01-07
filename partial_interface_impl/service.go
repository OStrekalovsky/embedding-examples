package partial_interface

type Service struct {
	store Store
}

func (s *Service) AddSome(id int) {
	s.store.Add(id)
}
func (s *Service) DelSome(id int) {
	s.store.Del(id)
}
func (s *Service) GetSome(id int) int {
	return s.store.Get(id)
}
