package common_dna

import "fmt"

type Animal struct {
	name string
}

func (a Animal) Speak() {
	fmt.Println("UGH!", a.name) // никто на самом деле не будет вызывать этот метод, но он нужен.
}

type Cat struct {
	Animal
}

func (c Cat) Speak() {
	fmt.Println("Meow!", "My name is", c.name) // на самом деле вызывать мы хотим это
}

type Dog struct {
	Animal
}

func (d Dog) Speak() {
	fmt.Println("Woff!", "My name is", d.name) // на самом деле вызывать мы хотим это
}
