package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	money int32 = 100
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, 10)
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, -10)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep((3 * time.Second))
	fmt.Println(money)
}
