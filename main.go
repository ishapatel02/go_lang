package main

import "fmt"

func main() {

	myArr := mySlice{1, 2, 3, 4, 5, 6}

	length := myArr.GetLength()

	// Use Printf for formatted output
	fmt.Printf("Length of arr is: %d\n", length)

	myArr.ReverseAnArray()

	fmt.Println("My Reversed Array")
	myArr.print()

	myArr.cyclicRotateByOne()
	fmt.Println("My Cyclic Rotate Array")
	myArr.print()

	inversionArr := mySlice{8, 4, 2, 1}
	fmt.Println("Count Inversion of an Array is", inversionArr.countInversion())
}
