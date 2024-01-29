package main

import (
	"fmt"
)

type Ducker interface {
	Display()
	Swim()
	SetQuackBehavior(q QuackBehavior)
	Quack()
	SetFlyBehavior(f FlyBehavior)
	Fly()
}

/*

	Quack Behaviors

*/

type QuackBehavior interface {
	PerformQuack()
}

type canQuack struct{}

func (q canQuack) PerformQuack() {
	fmt.Println("I am quacking")
}

type cannotQuack struct{}

func (q cannotQuack) PerformQuack() {
	fmt.Println("I cannot quack")
}

type veryLoudQuack struct{}

func (q veryLoudQuack) PerformQuack() {
	fmt.Println("I am quacking very loud!")
}

/*

	Fly Behaviors

*/

type FlyBehavior interface {
	PerformFly()
}

type canFly struct{}

func (f canFly) PerformFly() {
	fmt.Println("I am flying")
}

type cannotFly struct{}

func (f cannotFly) PerformFly() {
	fmt.Println("I cannot fly")
}

type rocketPowerFly struct{}

func (f rocketPowerFly) PerformFly() {
	fmt.Println("I am flying with rocket power!")
}

/*

	Concrete Ducks

*/

type Duck struct {
	quackBehavier QuackBehavior
	flyBehavior   FlyBehavior
}

func NewDuck(q QuackBehavior, f FlyBehavior) *Duck {
	return &Duck{
		quackBehavier: q,
		flyBehavior:   f,
	}
}

func (d Duck) Display() {
	fmt.Println("I am a duck")
}

func (d Duck) Swim() {
	fmt.Println("I am swimming")
}

func (d *Duck) SetQuackBehavior(q QuackBehavior) {
	d.quackBehavier = q
}

func (d Duck) Quack() {
	d.quackBehavier.PerformQuack()
}

func (d *Duck) SetFlyBehavior(f FlyBehavior) {
	d.flyBehavior = f
}

func (d Duck) Fly() {
	d.flyBehavior.PerformFly()
}

type MallardDuck struct {
	Duck
}

func NewMallardDuck(q QuackBehavior, f FlyBehavior) *MallardDuck {
	duck := NewDuck(q, f)
	return &MallardDuck{
		Duck: *duck,
	}
}

func (d MallardDuck) Display() {
	fmt.Println("I am a mullard duck")
}

type RedheadDuck struct {
	Duck
}

func NewRedheadDuck(q QuackBehavior, f FlyBehavior) *RedheadDuck {
	duck := NewDuck(q, f)
	return &RedheadDuck{
		Duck: *duck,
	}
}

func (d RedheadDuck) Display() {
	fmt.Println("I am a redhead duck")
}

type RubberDuck struct {
	Duck
}

func NewRubberDuck(q QuackBehavior, f FlyBehavior) *RubberDuck {
	duck := NewDuck(q, f)
	return &RubberDuck{
		Duck: *duck,
	}
}

func (d RubberDuck) Display() {
	fmt.Println("I am a rubber duck")
}

type DecoyDuck struct {
	Duck
}

func NewDecoyDuck(q QuackBehavior, f FlyBehavior) *DecoyDuck {
	duck := NewDuck(q, f)
	return &DecoyDuck{
		Duck: *duck,
	}
}

func (d DecoyDuck) Display() {
	fmt.Println("I am a decoy duck")
}

type MagicDuck struct {
	Duck
}

func NewMagicDuck(q QuackBehavior, f FlyBehavior) *MagicDuck {
	duck := NewDuck(q, f)
	return &MagicDuck{
		Duck: *duck,
	}
}

func (d MagicDuck) Display() {
	fmt.Println("I am a magic duck")
}

func (d *MagicDuck) Fly() {
	d.Duck.Fly()
	d.SetQuackBehavior(canQuack{})
}

func main() {
	duck := NewDuck(canQuack{}, canFly{})
	mallardDuck := NewMallardDuck(canQuack{}, canFly{})
	redheadDuck := NewRedheadDuck(canQuack{}, canFly{})
	rubberDuck := NewRubberDuck(canQuack{}, cannotFly{})
	decoyDuck := NewDecoyDuck(cannotQuack{}, cannotFly{})
	magicDuck := NewMagicDuck(cannotQuack{}, canFly{})
	ducks := []Ducker{duck, mallardDuck, redheadDuck, rubberDuck, decoyDuck, magicDuck}

	fmt.Printf("\n===== Beginning of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		d.Fly()
		fmt.Println()
	}

	redheadDuck.SetFlyBehavior(cannotFly{})

	fmt.Printf("\n===== Middle of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		d.Fly()
		fmt.Println()
	}

	magicDuck.SetFlyBehavior(rocketPowerFly{})
	magicDuck.SetQuackBehavior(veryLoudQuack{})

	fmt.Printf("\n===== Ending of the Game =====\n\n")
	for _, d := range ducks {
		d.Display()
		d.Swim()
		d.Quack()
		d.Fly()
		fmt.Println()
	}
}
