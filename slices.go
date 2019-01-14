package main

import "fmt"

/*
Go Slice is an abstraction over Go Array.
Go Array allows you to define variables that can hold several data items of the same kind
but it does not provide any inbuilt method to increase its size dynamically or get a sub-array of its own.
Slices overcome this limitation.
It provides many utility functions required on Array and is widely used in Go programming.
*/

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	main_range_map()
}

/*
len() and cap() functions

A slice is an abstraction over array. It actually uses arrays as an underlying structure.
The len() function returns the elements presents in the slice where cap() function returns the capacity of the slice
(i.e., how many elements it can be accommodate).
*/

func printSlice(x []int) {
	fmt.Printf("type=%T len=%d cap=%d slice=%v\n", x, len(x), cap(x), x)
}

/*
The range keyword is used in for loop to iterate over items of an array, slice, channel or map.
With array and slices, it returns the index of the item as integer.
With maps, it returns the key of the next key-value pair. Range either returns one value or two.
If only one value is used on the left of a range expression, it is the 1st value in the following table.
*/

func main_range_map() {
	/* create a slice */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	/* print the numbers */
	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}

	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}
}

/*
Go provides another important data type named map which maps unique keys to values.
A key is an object that you use to retrieve a value at a later date.
Given a key and a value, you can store the value in a Map object.
After the value is stored, you can retrieve it by using its key.
*/
