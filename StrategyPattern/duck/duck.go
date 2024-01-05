package duck

import "fmt"

type Duck struct{}

func NewDuck() *Duck {
	return &Duck{}
}

func (d Duck) Display() {
	fmt.Println("I am a duck")
}

func (d Duck) Swim() {
	fmt.Println("I am swimming")
}

func (d Duck) PerformFly() {
	fmt.Println("I am flying")
}

func (d Duck) PerformQuack() {
	fmt.Println("I am quacking")
}

type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	baseDuck := NewDuck()
	return &MallardDuck{Duck: *baseDuck}
}

func (d MallardDuck) Display() {
	fmt.Println("I am a mallard duck")
}

type RedheadDuck struct {
	Duck
}

func NewRedheadDuck() *RedheadDuck {
	baseDuck := NewDuck()
	return &RedheadDuck{Duck: *baseDuck}
}

func (d RedheadDuck) Display() {
	fmt.Println("I am a redhead duck")
}

type RubberDuck struct {
	Duck
}

func NewRubberDuck() *RubberDuck {
	baseDuck := NewDuck()
	return &RubberDuck{Duck: *baseDuck}
}

func (d RubberDuck) Display() {
	fmt.Println("I am a rubber duck")
}
