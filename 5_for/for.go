package main

import "fmt"

//for-> only construct in go fo looping
func main(){
//while loop
	i:=1
	for i<=3{
		fmt.Println(i)
		i=i+1
	}

	//infinite loop
	//for{
	//	println("1")
	//}

	//classic for loop
	for i := 0; i < 3; i++ {
		//break
		//continue
		fmt.Println(i)
	}

	//1.22range
	for i:=range 3{
		fmt.Println(i)
	}

}