# Design Patterns in Golang

## Strategy Pattern

You will learn Strategy Pattern by doing a simple excercise. Please follow the steps below.

1. Given `Duck` base class and `MallardDuck`, `RedheadDuck`, `RubberDuck` sub classes with some of their behaviors as you can see in `./StrategyPattern/duck/duck.go`. Runs

```sh
make strategy
```

2. Here is new requirement, business told us that `RubberDuck` cannot fly and cannot quack. Let's see how you fix it. (When the `RubbleDuck` fly, it should print `I cannot fly` and when it quack, it should print `I cannot quack`)
