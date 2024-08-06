package main

func (a mySlice) cyclicRotateByOne() {
	length := a.GetLength()

	i := length - 1

	last_ele := a[i]

	for i > 0 {
		a[i] = a[i-1]
		i--
	}

	a[0] = last_ele
}
