package main

import (
	"fmt"
	"sync"
)
var num int
var wg sync.WaitGroup
var channel1 = make(chan int,1)

func add(){
	defer wg.Done()
	for i:=0; i < 50000 ;i++{
		channel1 <- 1

		num += 1
		<-channel1
 	}

}

func main() {


	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println("end" ,num)
}