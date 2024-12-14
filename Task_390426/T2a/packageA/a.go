package packageA

import "example.com/project/common"

type A struct {
}

func (a A) DoSomething() {
    // Implementation for A
}

func NewA() common.AInterface {
    return &A{}
}