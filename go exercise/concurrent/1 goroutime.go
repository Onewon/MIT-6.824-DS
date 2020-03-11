package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
// func say(string s){ for i :=0; i<10; i++ { time.sleep(100); fmt.Println(s);  }   }

func main() {
	go say("world")
	say("hello")
}
