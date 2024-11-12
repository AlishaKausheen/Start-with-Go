package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch

	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	//multiple condition switch

	switch time.Now().Weekday(){
 
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("weekday")
	}

	//type switch
	whoAmI :=func ( i interface{})  {
		switch t:=i.(type){
		case int :
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("its boolean")
		default:
			fmt.Println("other",t)
		}
	}
	whoAmI(10)
}