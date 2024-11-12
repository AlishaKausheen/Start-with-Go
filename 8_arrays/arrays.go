package main

import "fmt"

func main() {
	var nums [4]int

	//zeroed value
	//int->0 string->"" bool -> false

	//array length
	fmt.Println(len(nums))

	nums[0] =1
	fmt.Println(nums[0])
	fmt.Println(nums)


	//to declare array in single line
	arr :=[3]int{1,2,3}

	fmt.Println(arr)

	// 2d arrays
	num :=[2][2]int{{2,5},{3,7}}
	fmt.Println(num)
}