package main

import "fmt"


//change by value
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum func", num)
}

//change by ref
func byRef(num *int){
*num = 5
fmt.Println("by referrence", *num)
}

func main() {
num :=1
changeNum(num)
fmt.Println("Ater changeNum in main",num)

//knowing address
fmt.Println("Memory address", &num)
//by referrence
byRef(&num)
fmt.Println("after by ref func",num)

}