package packageB

import "example.com/project/common"

type B struct {
	A common.AInterface // Use interface instead of the concrete type
}

func NewB(a common.AInterface) B {
	return B{A: a} // Accepts AInterface as a parameter
}
