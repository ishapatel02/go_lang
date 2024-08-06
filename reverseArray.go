package main

import "fmt"

type mySlice []int

func (a mySlice) GetLength() int {
	return len(a)
}

func (a mySlice) ReverseAnArray() {
	start := 0
	end := a.GetLength()

	for start < end {
		temp := a[start]
		a[start] = a[end-1]
		a[end-1] = temp

		start++
		end--
	}
}

func (a mySlice) print() {
	for _, ele := range a {
		fmt.Println(ele)
	}
}
