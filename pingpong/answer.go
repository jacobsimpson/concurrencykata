package main

import (
	"fmt"
	"time"
)

func main() {
	s := make(chan bool)

	go func() {
		for f := range s {
			fmt.Println("ping")
			time.Sleep(1 * time.Second)
			s <- f
		}
	}()

	s <- true
	for f := range s {
		fmt.Println("pong")
		time.Sleep(1 * time.Second)
		s <- f
	}
}
