package main

import (
	"fmt"
	"io"
	"os"
)

func CreatFile() {
	fmt.Println("创建文件")
	f, err := os.Create("plan.txt")
	if err == nil {
		fmt.Println(f.Name())
	} else {
		fmt.Println(err)
	}
	fmt.Println("创建完毕")

	f.Close()
}

func readFile(){
	fmt.Println("读文件:")

	f,_ := os.Open("plan.txt")
	for {
		str := make([]byte, 100)
		_,err1 := f.Read(str)
		if err1 == io.EOF {
			break
		}
		fmt.Println(string(str))
	}
	f.Close()
	fmt.Println("读完毕")

}


func writeFile(){
	fmt.Println("写文件")

	f,_ := os.OpenFile("plan.txt", os.O_RDWR|os.O_CREATE,0755)
	f.WriteString("I’m not afraid of difficulties and insist on learning programming")
	f.Close()
	fmt.Println("写完毕")

}
func main()  {
	CreatFile()
	writeFile()
	readFile()

}