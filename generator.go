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

func main() {

	var Iterations int = 10

	if len(os.Args) > 1 {
		Iterations, _ = strconv.Atoi(os.Args[1])
	}

	Bob := boring("Bob")
	Alice := boring("Alice")

	for i := 0; i < Iterations; i++ {
		fmt.Println(<-Bob)
		fmt.Println(<-Alice)
	}
}
