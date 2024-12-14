package packageB

import "example.com/common"

type B struct {
	A common.AInterface // Depends on the interface
}

func NewB(a common.AInterface) B {
	return B{A: a}
}
