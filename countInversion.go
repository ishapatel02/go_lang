package main

import "log"

func (a mySlice) countInversion() int {

	log.Print("Inside inversion")
	a.print()

	countInversion := 0

	length := a.GetLength()
	i := 0

	for i < length-1 {
		j := i + 1
		for j < length {
			if a[i] > a[j] {
				countInversion++
			}
			j++
		}
		i++
	}
	return countInversion
}
