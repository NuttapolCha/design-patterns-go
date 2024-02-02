package main

import (
	"fmt"
)

type Ducker interface {
	Display()
	Swim()
	Quack()
	Fly()
}

type Duck struct{}

func (d *Duck) Display() {
	fmt.Println("I am a duck")
}

func (d Duck) Swim() {
	fmt.Println("I am swimming")
}

func (d Duck) Quack() {
	fmt.Println("quack quack!")
}

func (d Duck) Fly() {
	fmt.Println("I am flying")
}

type MallardDuck struct {
	Duck
}

func (d MallardDuck) Display() {
	fmt.Println("I am a mullard duck")
}

type RedheadDuck struct {
	Duck
	isTried bool
}

func (d RedheadDuck) Display() {
	fmt.Println("I am a redhead duck")
}

func (d *RedheadDuck) SetTried() {
	d.isTried = true
}

func (d *RedheadDuck) SetIsNotTried() {
	d.isTried = false
}

func (d RedheadDuck) Fly() {
	if d.isTried {
		fmt.Println("I cannot fly")
		return
	}
	fmt.Println("I am flying")
}

type RubberDuck struct {
	Duck
}

func (d RubberDuck) Display() {
	fmt.Println("I am a rubber duck")
}

func (d RubberDuck) Fly() {
	fmt.Println("I cannot fly")
}

type DecoyDuck struct {
	Duck
}

func (d DecoyDuck) Display() {
	fmt.Println("I am a decoy duck")
}

func (d DecoyDuck) Fly() {
	fmt.Println("I cannot fly")
}

func (d DecoyDuck) Quack() {
	fmt.Println("I cannot quack")
}

type MagicDuck struct {
	Duck
	haveFly   bool
	isEndGame bool
}

func (d MagicDuck) Display() {
	fmt.Println("I am a magic duck")
}

func (d *MagicDuck) Fly() {
	d.haveFly = true
	if d.isEndGame {
		fmt.Println("I am flying with rocket power!")
		return
	}
	fmt.Println("I am flying")
}

func (d MagicDuck) Quack() {
	if !d.haveFly {
		fmt.Println("I cannot quack")
		return
	}
	if d.isEndGame {
		fmt.Println("quack quack! very loud!")
		return
	}
	fmt.Println("quack quack!")
}

func (d *MagicDuck) SetEndGame() {
	d.isEndGame = true
}

func main() {
	duck := &Duck{}
	mullardDuck := &MallardDuck{}
	redheadDuck := &RedheadDuck{}
	rubberDuck := &RubberDuck{}
	decoyDuck := &DecoyDuck{}
	magicDuck := &MagicDuck{}
	ducks := []Ducker{duck, mullardDuck, redheadDuck, rubberDuck, decoyDuck, magicDuck}

	fmt.Printf("\n===== Beginning of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		d.Fly()
		fmt.Println()
	}

	// you can add transition logic here
	redheadDuck.SetTried()

	fmt.Printf("\n===== Middle of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		d.Fly()
		fmt.Println()
	}

	// you can add transition logic here
	redheadDuck.SetIsNotTried()
	magicDuck.SetEndGame()

	fmt.Printf("\n===== Ending of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		d.Fly()
		fmt.Println()
	}
}
