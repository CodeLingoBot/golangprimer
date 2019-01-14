package main

import "errors"
import "fmt"
import "math"

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
func main() {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

/*
	https://www.tutorialspoint.com/go/go_quick_guide.htm
*/

/*
 Go language is a programming language initially developed at Google in the year 2007
 by Robert Griesemer, Rob Pike, and Ken Thompson.
 It is a statically-typed language having syntax similar to that of C.
 It provides garbage collection, type safety, dynamic-typing capability, many advanced built-in types
 such as variable length arrays and key-value maps. It also provides a rich standard library.
 The Go programming language was launched in November 2009 and is used in some of the Google's production systems.
*/
