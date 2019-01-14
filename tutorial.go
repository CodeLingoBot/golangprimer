package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main_swap() {
	var a, b string
	a, b = swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}

func main_string() {
	/* local variable declaration */
	var a, b, c int

	/* actual initialization */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("value of a = %d, b = %d and c = %d\n", a, b, c)
}

// arrays

/*
	decleration:
	  var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
*/

func main_arr() {
	var n [10]int /* n is an array of 10 integers */
	var i, j int

	/* initialize elements of array n to 0 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* set element at location i to i + 100 */
	}

	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

// pointers

/*
func main_pointer() {
   var  ptr *int

   fmt.Printf("The value of ptr is : %x\n", ptr  )
}
*/

func main() {
	main_swap()
	main_string()
	main_arr()
}
