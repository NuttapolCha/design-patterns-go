# Design Patterns in Golang

In each pattern, please read the *Setting the Scene*.
Once you understand your context, then you can go to the problem statement.

## Strategy Pattern

### Setting the Scene

We are making a simple duck simulator game.
Given `MullardDuck`, `RedheadDuck`, `RubberDuck`, all of them is `Duck` and have behaviors as the following

```go
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

type MullardDuck struct {
    Duck
}

func (d MullardDuck) Display() {
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
```

In the beginning of the game, you are going to make 4 kinds of ducks perform their 3 behaviors.

```go
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
```

You can run the program above by `make`.

```sh
make strategy
```

### Problem Statements

From now on, the business will requires you to add awesome features to the game.

Please add the following features **step by step**, **NOT SKIPPING OR READ OVER THE STEPS BEFORE YOU HAVE DONE EACH ONE** or you might not get the best result from this exercise.

1. First of all, business want you to add the `Fly()` behavior to all ducks.
How do you going to implement this? Asuume that the method only print the "I am flying".

    ```go
    type Ducker interface {
        Display()
        Swim()
        Quack()
        Fly() // introducing fly behavior
    }
    ```

2. The business see that `RubberDuck` should not flyable, can you make the `RubberDuck` print *"I cannot fly" instead of "I am flying"*?

3. The business ask you to add new duck type, a `DecoyDuck` which is *cannot fly and cannot quack*. Let's see how you can add a new duck type.

4. Now, we are in the middle game! The ReadheadDuck is now tried and no longer flyable. Business wants the `ReadheadDuck` is flyable at the beginning of the game but cannot fly at the middle game. How you gonna implement this? Please notice that whether you are violating the *Open/Close Principle* or not?

    ```go
    func main() {
        duck := &Duck{}
        mullardDuck := &MallardDuck{}
        redheadDuck := &RedheadDuck{}
        rubberDuck := &RubberDuck{}
        decoyDuck := &DecoyDuck{}
        ducks := []Ducker{duck, mullardDuck, redheadDuck, rubberDuck, decoyDuck}

        fmt.Printf("\n===== Beginning of the Game =====\n\n")
        for _, d := range ducks {
            d.Display()
            d.Swim()
            d.Quack()
            d.Fly() // ReadheadDuck can fly here
            fmt.Println()
        }

        // you can add logic here before entering the middle game

        fmt.Printf("\n===== Middle of the Game =====\n\n")
        for _, d := range ducks {
            d.Display()
            d.Swim()
            d.Quack()
            d.Fly() // ReadheadDuck cannot fly here
            fmt.Println()
        }
    }
    ```

5. Business is going crazy now. They require you to add `MagicDuck` which initialy cannot quack, but after it flying, they can.
By implementing this, you will see that the MagicDuck cannot quack at the beginning of the game but can quack at the middle game wihout additional logic in main function.

    ```go
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
            d.Quack() // magic duck cannot quack here
            d.Fly() // ReadheadDuck can fly here
            fmt.Println()
        }

        // you can add logic here before entering the middle game
        // no additional logic about magic duck is needed

        fmt.Printf("\n===== Middle of the Game =====\n\n")
        for _, d := range ducks {
            d.Display()
            d.Swim()
            d.Quack() // magic duck can quack here because it has been fly at the beginning of the game
            d.Fly() // ReadheadDuck cannot fly here
            fmt.Println()
        }
    }
    ```

6. In the ending of the game, `MagicDuck` will fly with rocket power and quacking very loud! Can you implement this?
Just printing the *"I am flying with rocket power!"* and *"I am quacking very loud"* for `Fly()` and `Quack()` methods respectively.

7. That's enough for today. Now thinking about what you've done, think about problems you have faced.
Are your code easy to changes (i.e. maintainable and readable) as the complexity is growing?
