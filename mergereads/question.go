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

// Given two data sources, read from whichever one is available next. To expand
// on the challenge slightly, assume that each data source represents a stream
// of values and read repeatedly from each.
func main() {
}
