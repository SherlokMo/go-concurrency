package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func boring(name string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func main() {
	var Iterations int = 10

	if len(os.Args) > 1 {
		Iterations, _ = strconv.Atoi(os.Args[1])
	}

	fan := FanIn(boring("Bob"), boring("Alice"))
	/*
		| Notice how we set the limit to Iterations * 2
		| Instead of the Just iteration like the generator example
	*/
	for i := 0; i < Iterations*2; i++ {
		fmt.Println(<-fan)
	}

}
