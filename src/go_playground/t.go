package main

import (
	"fmt"
)

type I1 []int
func (i I1) Show() {
	fmt.Println(i[0])
}

type I2 struct {
	I1
}
func (i I2) Close() {
	fmt.Println("I2 close")
}

type In interface {
	Show()
	Close()
}

type I3 struct {
	I1
}
func (i I3) Close() {
	fmt.Println("I3 close")
}

func main() {
	var a2 I2
	a2.I1 = append(a2.I1, 3)
	a2.I1 = append(a2.I1, 2)
	a2.Show()
	a2.Close()
	var a3 I3
	a3.I1 = append(a3.I1, 3)
	a3.I1 = append(a3.I1, 2)
	a3.Show()
	a3.Close()

	var v2 In = a2
	var v3 In = a3
	v2.Show()
	v2.Close()
	v3.Show()
	v3.Close()
}

