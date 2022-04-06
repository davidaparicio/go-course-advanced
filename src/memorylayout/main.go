package main

import (
	"fmt"
	"unsafe"
)

const uintSizeInBytes = 4 << ((^uint(0)) >> 32 & 1)

type User struct {
	admin       bool   // 1
	nbFollowers uint64 // 8
	active      bool   // 1
	age         uint64 // 8
	retired     bool   // 1
}

func main() {
	fmt.Println(uintSizeInBytes)
	fmt.Println(unsafe.Sizeof(User{}))
}
