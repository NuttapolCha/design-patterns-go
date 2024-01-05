package main

import (
	"fmt"

	"github.com/NuttapolCha/design-patterns/StrategyPattern/duck"
)

func main() {
	fmt.Printf("==Mullard Duck==\n")
	mallardDuck := duck.NewMallardDuck()
	mallardDuck.Display()
	mallardDuck.PerformFly()
	mallardDuck.PerformQuack()
	fmt.Println()

	fmt.Printf("==Redhead Duck==\n")
	redheadDuck := duck.NewRedheadDuck()
	redheadDuck.Display()
	redheadDuck.PerformFly()
	redheadDuck.PerformQuack()
	fmt.Println()

	fmt.Printf("==Rubber Duck==\n")
	rubberDuck := duck.NewRubberDuck()
	rubberDuck.Display()
	rubberDuck.PerformFly()
	rubberDuck.PerformQuack()
	fmt.Println()
}
