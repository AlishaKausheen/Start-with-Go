package main

import (
	"fmt"
	"slices"
)

//slices->dynamic array
//most used construct in go
func main(){
	//uninitialised slice is nil
var nums [] int 
fmt.Println(nums)
fmt.Println(len(nums))


var a = make([] int , 2,5)
//first argument is initial capacity 0 and second argument is initial clength

//capacity
fmt.Println(cap(a))


nums = append(nums, 1)
nums = append(nums, 2)
nums = append(nums, 3)

a = append(a, 0)
a = append(a, 1)
a = append(a, 3)
a= append(a, 9)
fmt.Println(nums)
fmt.Println(a)

//copy function
nums= append(nums,2)

var nums2 = make([]int , len(nums))

copy(nums2,nums)

//slice operator

var nums3=[]int{1,2,3}

//slice operator
fmt.Println(nums3[0:2])


//slice package
var a1 = [] int {1,2}
var a2 = [] int {1,2}

fmt.Println(slices.Equal(a1,a2))

}