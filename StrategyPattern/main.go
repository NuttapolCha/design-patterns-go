package main

import (
	"fmt"
)

type Ducker interface {
	Display()
	Swim()
	Quack()
	Fly()
	SetIsMidGame(b bool)
	SetIsEndingGame(b bool)
}

type Duck struct {
	isMidGame bool
	isEndGame bool
}

func (d Duck) Display() {
	fmt.Println("I am a duck")
}

func (d Duck) Swim() {
	fmt.Println("I am swimming")
}

func (d Duck) Quack() {
	fmt.Println("I am quacking")
}

func (d Duck) Fly() {
	fmt.Println("I am flying")
}

func (d *Duck) SetIsMidGame(b bool) {
	d.isMidGame = b
}

func (d *Duck) SetIsEndingGame(b bool) {
	d.isEndGame = b
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

func (d RedheadDuck) Fly() {
	if d.Duck.isMidGame {
		fmt.Println("I cannot fly")
	} else {
		d.Duck.Fly()
	}
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

func (d DecoyDuck) Quack() {
	fmt.Println("I cannot quack")
}

func (d DecoyDuck) Fly() {
	fmt.Println("I cannot fly")
}

type MagicDuck struct {
	Duck
	hasFly bool
}

func (d MagicDuck) Display() {
	fmt.Println("I am a magic duck")
}

func (d MagicDuck) Quack() {
	if d.hasFly {
		if d.isEndGame {
			fmt.Println("I am quacking very loud")
		} else {
			d.Duck.Quack()
		}
	} else {
		fmt.Println("I cannot quack")
	}
}

func (d *MagicDuck) Fly() {
	if d.isEndGame {
		fmt.Println("I am flying with rocket power!")
	} else {
		d.Duck.Fly()
	}
	d.hasFly = true
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

	redheadDuck.SetIsMidGame(true)

	fmt.Printf("\n===== Middle of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		d.Fly()
		fmt.Println()
	}

	magicDuck.SetIsEndingGame(true)

	fmt.Printf("\n===== Ending of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		d.Fly()
		fmt.Println()
	}
}
