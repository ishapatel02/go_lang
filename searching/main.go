package main

import "fmt"

func main() {
	myArr := searchSlice{8, 4, 2, 1, 6}
	valuePresent := myArr.linearSearch(10)
	fmt.Println("Linear search Value Present:", valuePresent)
}
