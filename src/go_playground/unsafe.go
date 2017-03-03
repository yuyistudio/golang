package main

import (
	"fmt"
	"unsafe"
)

func toBits(v float64) string {
	var vn uint64 = *(*uint64)(unsafe.Pointer(&v))
	return fmt.Sprintf("%#016x", vn)
}

func main() {
	fmt.Println(toBits(3.14))
}
