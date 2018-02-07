package main

import "fmt"

func helloTony(myChannel chan string) {

	go func() {
		myChannel <- "hello tony"
	}()

	x := <-myChannel
	fmt.Println(x)
}
