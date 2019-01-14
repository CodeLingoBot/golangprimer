package main

import "fmt"

// https://github.com/golang
// https://www.tutorialspoint.com/go/go_environment.htm

var i_var int

func main() {
	i_var = 43534

	const DO_IT = "do that now"

	fmt.Println("Hello, 世界")
	fmt.Println(i_var)
	fmt.Printf("i_var is of type %T\n", i_var)

	for i := 0; i < 5; i++ {
		fmt.Printf("%s %v \n", DO_IT, i)
	}

	xy := 0
	for xy < 6 {
		xy += 1
		fmt.Printf("while %v\n", xy)
	}

	// https://tour.golang.org/list

	/*
		Channels are the pipes that connect concurrent goroutines.
		You can send values into channels from one goroutine and receive those values into another goroutine.
	*/

	// https://gobyexample.com/channels

	//Channels are typed by the values they convey.

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

}
