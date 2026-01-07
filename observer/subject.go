package observer

// обобщенные компоненты для реализации паттерна Observer

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

type Observer interface {
	Update()
}

type DefaultSubject struct {
	observers []Observer
}

func NewDefaultSubject() *DefaultSubject {
	return &DefaultSubject{
		observers: []Observer{},
	}
}

func (s *DefaultSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}
func (s *DefaultSubject) Detach(observer Observer) {
	for i := 0; i < len(s.observers); i++ {
		currentObserver := s.observers[i]
		if currentObserver == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *DefaultSubject) Notify() {
	for i := 0; i < len(s.observers); i++ {
		observer := s.observers[i]
		observer.Update()
	}
}
