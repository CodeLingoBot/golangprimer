package main

/*
	The first line of the program package main defines the package name in which this program should lie.
	It is a mandatory statement, as Go programs run in packages.
	The main package is the starting point to run the program.
	Each package has a path and name associated with it.
*/

import "fmt"

// https://github.com/golang

/* structure is:
   Package Declaration
   Import Packages
   Functions
   Variables
   Statements and Expressions
   Comments
*/

/* key words:
 	break 	default
 	func 	interface 	select
	case
 	defer 	Go
 	map 	Struct
	chan 	else 	Goto
 	package 	Switch
	const 	fallthrough 	if
 	range 	Type
	continue 	for 	import 	return 	Var
*/

func main() {

	/*
		Go program consists of various tokens.
		A token is either a keyword, an identifier, a constant, a string literal, or a symbol.
		For example, the following Go statement consists of six tokens
	*/

	fmt.Println("Hello, World!")
}

// go run test.go
// go build test.go
// ./test // 2mb binary...
