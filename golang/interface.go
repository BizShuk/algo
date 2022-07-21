package main

import (
	"fmt"
	"sort"
)

// Notice: type alias for underlying method
// No underlying type method call unless you declar it
type x sort.IntSlice

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	var b x = x(sort.IntSlice(a))
	fmt.Println(b)
	var c x = x(a)
	fmt.Println(c)
	d := append(c, 5)
	fmt.Println(d)
	fmt.Println(a)

	// w1 : sort.IntSlice
	w1 := sort.IntSlice{4, 3, 2, 1}
	fmt.Println(w1)

	// w2 : sort.Interface
	w2 := convertY(w1)
	fmt.Println(w2)

	var w3 y // w3 : y
	w3 = convertY(w1).(y)
	fmt.Println(w3)

	var w4 sort.Interface = w1 // w5 : y from sort.Interface
	var w5 y = y{w4}
	fmt.Println(w5)

	var w6 y = y{w1} // w6 : y from sort.IntSlice
	fmt.Println(w6)

	// w7 : sort.IntSlice from y  // Notice: w1 is assigned to y.Interface, not y
	var w7 sort.IntSlice = w6.Interface.(sort.IntSlice)
	fmt.Println(w7)
}

// Notice: empty concrete type struct with interface
type y struct {
	sort.Interface
}

func convertY(data sort.Interface) sort.Interface {
	return y{data}
}
