package gomock

import "fmt"

type Hi func(name string)

func SayHi(name string) {
	fmt.Println("Hi", name)
}

// [Pattern]: [Go Mock Testing] by passing mock function as a parameter into test function. Could use interface instead of exact struct/func
func CallSayHi(name string, sayHi Hi) {
	sayHi(name)
}
