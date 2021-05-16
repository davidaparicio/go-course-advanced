package main

import (
	"fmt"
	"sync"
)

type StateMachine struct {
	mu    sync.Mutex
	count int
}

func NewStateMachine() *StateMachine {
	return new(StateMachine)
}

func (sm *StateMachine) Increment() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.count++
	return sm.count
}

func (sm *StateMachine) Decrement() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.count--
	return sm.count
}

func main() {
	sm := NewStateMachine()
	fmt.Println(sm.Increment())
	fmt.Println(sm.Increment())
	fmt.Println(sm.Decrement())
	fmt.Println(sm.Decrement())
}
