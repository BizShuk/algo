package gomock

type IHelper interface {
	CallHelp(content string)
}

var helper IHelper

// [Pattern]: [Go Mock Testing] Use package variable to inject mock object
func CustomerService(content string) {
	helper.CallHelp(content)
}
