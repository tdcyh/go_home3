package main

import (
	"fmt"
	"sync"
)
var wg1 sync.WaitGroup
var channel21 = make(chan int,1)
var channel22 = make(chan int,1)

func pri1(i int){
	defer wg1.Done()
	for 	; i <= 100 ;i += 2{
		channel21 <- 1
		fmt.Println(i)
		channel22 <- 1
	}

}
func pri2(i int){
	defer wg1.Done()
	for 	; i <= 100 ;i += 2{
		<-channel22
		fmt.Println(i)
		<-channel21
	}

}
func main() {


	wg1.Add(2)

	go pri1(1)
	go pri2(2)

	wg1.Wait()
}