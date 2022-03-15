package main

import (
	"fmt"
	"runtime"
)

func printMsg(msg chan string) {
	fmt.Println(<-msg)
}

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			message <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMsg(message)
	}
}
