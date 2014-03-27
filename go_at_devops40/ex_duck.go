package main
import "fmt"

type Duck interface {
	Walk()
	Swim()
	Quack()
}

func WalkSwimQuack(duck Duck) { duck.Walk(); duck.Swim(); duck.Quack() }

type SomeBird struct{}
func (b *SomeBird) Walk() { fmt.Println("I'm walking...") }
func (b *SomeBird) Swim() { fmt.Println("I'm swimming...") }
func (b *SomeBird) Quack() { fmt.Println("I'm quacking...") }

func main() {
	bird := new(SomeBird)
	WalkSwimQuack(bird)
}