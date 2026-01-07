package common_dna

import (
	"testing"
)

// Попытка сгруппировать животных по общему встроенному типу
func callAnimals(animals ...Animal) {
	for _, animal := range animals {
		animal.Speak()
	}
}

func TestCommonDNA(t *testing.T) {
	cat := Cat{Animal{"Tom"}}
	dog := Dog{Animal{"Spike"}}
	cat.Speak()
	dog.Speak()
	// Ошибка компиляции ниже
	// callAnimals(cat, dog)
}
