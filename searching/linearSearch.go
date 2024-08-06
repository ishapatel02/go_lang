package main

type searchSlice []int

func (a searchSlice) linearSearch(x int) bool {

	for i := range a {
		if a[i] == x {
			return true
		}
		i++
	}
	return false
}
