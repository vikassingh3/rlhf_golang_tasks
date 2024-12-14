package packageA

import "example.com/common"

type A struct {
	// Implements common.AInterface
}

func (a A) DoSomething() {
	// Implementation of the interface method
}

func NewA() common.AInterface {
	return A{}
}
