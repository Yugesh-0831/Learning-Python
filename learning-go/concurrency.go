package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Concurrency() {
	// with go as prefix, golang sums up a new goroutine that runs concurrently
	go say("world")

	// this is still part of main go routine
	say("hello")
}
