package main

import (
	"fmt"
)

type Ducker interface {
	Display()
	Swim()
	Quack()
}

type Duck struct{}

func (d Duck) Display() {
	fmt.Println("I am a duck")
}

func (d Duck) Swim() {
	fmt.Println("I am swimming")
}

func (d Duck) Quack() {
	fmt.Println("I am quacking")
}

type MallardDuck struct {
	Duck
}

func (d MallardDuck) Display() {
	fmt.Println("I am a mullard duck")
}

type RedheadDuck struct {
	Duck
}

func (d RedheadDuck) Display() {
	fmt.Println("I am a redhead duck")
}

type RubberDuck struct {
	Duck
}

func (d RubberDuck) Display() {
	fmt.Println("I am a rubber duck")
}

func main() {
	duck := &Duck{}
	mullardDuck := &MallardDuck{}
	redheadDuck := &RedheadDuck{}
	rubberDuck := &RubberDuck{}
	ducks := []Ducker{duck, mullardDuck, redheadDuck, rubberDuck}

	fmt.Printf("\n===== Beginning of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		fmt.Println()
	}
}
