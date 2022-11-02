package main

import "fmt"

func main() {
	over := make(chan bool)
	num := 0
	for i := 0; i < 10; i++ {

		go func() {

			fmt.Println(i)
			if i == num {
				over <- true
				num += 1
			}
		}()
		<-over

	}

	fmt.Println("over!!!")
}