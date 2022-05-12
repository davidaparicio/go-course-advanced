package main

import (
	"fmt"
	"unsafe"
)

const uintSizeInBytes = 4 << ((^uint(0)) >> 32 & 1)

type User struct {
	admin       bool   // 1 + 7 of padding
	age         uint64 // 64 bits = 8 bytes + 0
	active      bool   // 1 + 7 of padding
	nbFollowers uint64 // 8 + 0
	retired     bool   // 1 + 7 of padding
}

type User2 struct {
	age         uint64 // 8 + 0
	nbFollowers uint64 // 8 + 0
	admin       bool   // 1
	active      bool   // 1
	retired     bool   // 1 + 5 of padding
}

func main() {
	fmt.Println(uintSizeInBytes)       // 8 Bytes
	fmt.Println(unsafe.Sizeof(User{})) // 40 = 8*5 (size+padding)
	// Let's re-order the memory
	fmt.Println(unsafe.Sizeof(User2{})) // 40 = 24 (almost half memory usage)
}
