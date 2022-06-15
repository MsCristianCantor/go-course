package main

import "fmt"

func main() {
	//unbuffered channel (canal con bufer)
	ca := make(chan int, 1)

	ca <- 42

	fmt.Println(<-ca)
}
