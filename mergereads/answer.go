package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func dataSource1() string {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return fmt.Sprintf("1-%d", rand.Intn(100))
}

func dataSource2() string {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return fmt.Sprintf("2-%d", rand.Intn(100))
}

func main() {
	c1 := make(chan string)
	go func() {
		for {
			c1 <- dataSource1()
		}
	}()

	c2 := make(chan string)
	go func() {
		for {
			c2 <- dataSource2()
		}
	}()

	for {
		select {
		case s := <-c1:
			fmt.Println(s)
		case s := <-c2:
			fmt.Println(s)
		}
	}
}
