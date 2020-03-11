package main

import "fmt"
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

//	Sends to a buffered channel block only when the buffer is full.
//	Receives block when the buffer is empty.
func main() {
	ch := make(chan int, 2)
	s := []int{1, 2, 3, 4}
	go sum(s[:2], ch)
	go sum(s[2:], ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}
