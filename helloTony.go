package main

import "fmt"

func helloTony(myChannel chan string) {

	myChannel <- "hello tony"

	x := <-myChannel
	fmt.Println(x)
}
