package main

import "fmt"

func main() {
//	age := 16
//	if age >= 18 {
//		fmt.Println("adult")
//	}else if age>=12{
//		fmt.Println("teenage")
//	}else{
//		fmt.Println("child")
//	}
	var role = "admin"
	var hasPermission=true

	if role=="admin" || hasPermission{
		fmt.Println("yess")
	}
	//can declare variable in if else
	if age:=15; age>=18{
		fmt.Println("adult",age)
	}else{
		fmt.Println("not adult",age)
	}

	//go doesnt have ternary operator, you will have to use normal if else
}