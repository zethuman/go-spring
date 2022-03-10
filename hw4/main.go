package main

import (
	"fmt"
)

func a(done chan struct{}) {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
	done <- struct{}{}
}

func b(done chan struct{}) {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go a(done)
	go b(done)
	<-done
	<-done
	fmt.Println("\nend main()")
}
