package common_dna

import "testing"

type Speaker interface {
	Speak()
}

// группируем типы по общему поведению.
func callSpeakers(speakers ...Speaker) {
	for _, animal := range speakers {
		animal.Speak()
	}
}

func TestCommonBehavior(t *testing.T) {
	cat := Cat{Animal{"Tom"}}
	dog := Dog{Animal{"Spike"}}
	cat.Speak()
	dog.Speak()
	callSpeakers(cat, dog)
}
