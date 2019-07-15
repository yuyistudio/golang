package main


/*
测试继承多态相关的特性
 */

import (
	"fmt"
)


type A struct {
	Name string
}

func (a *A) GetName() string {
	return a.Name
}

type A1 struct {
	A
	Age int
}

type A2 struct {
	A
	NewName string
}

func (a2 *A2) GetName() string {
	return a2.NewName + a2.A.GetName()
}

type IA interface {
	GetName() string
}

func main() {
	a1o := new(A1)
	a1o.Name = "a1"
	a2o := new(A2)
	a2o.Name = "a2"
	a2o.NewName = "a2_new"
	var a1 IA = a1o
	var a2 IA = a2o
	fmt.Printf("a1:%v a2:%v\n", a1.GetName(), a2.GetName())
}
