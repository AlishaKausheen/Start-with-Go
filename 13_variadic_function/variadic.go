package main

import "fmt"


func sum(nums ...int)int{
	total :=0
	for _, num:=range nums{
		total =total+num
	}

	return total
}

//if any type then use interface
func random (x ...interface{}){

}
//functions were we can pass n number of parameter
func main(){

	result :=sum(3,4,5,9)
	//numbers are in slice
	nums:=[]int{1,2,3,4,5,6}
	fmt.Println(result)
	fmt.Println(sum(nums...))
}