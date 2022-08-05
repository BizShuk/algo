package golang

import "fmt"

type S1 struct {
	f1 string
	f2 int
}

type S2 struct {
	f1 string
	f2 int
}

type S3 struct {
	S1
}

// type S1alias S1

// type S4 struct {
// }

func TypeCastingMain() {
	a := S1{"10", 10} // S1 init
	fmt.Println("a", a)
	b := S2(a) // cast S1 to S2
	fmt.Println("b", b)
	c := S3{a} // cast S1 to S3
	fmt.Println("c", c, c.f1)

	// Notice: Struct only can assign to the struct has the same fields
	// But with both interface, it will work on interface level

	// var aAlias S1alias = S1alias(a)
	// var d S4 = aAlias // cast S1 to S4
	// fmt.Println("d", d)

	a = S1(b) // cast S2 to S1
	fmt.Println("a", a)
}
