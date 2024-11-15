package main

import "fmt"

// range is used for iterating over data structures
func main(){

nums := []int{6,7,8}

sum :=0


//first argument is index second argument is value at that index
for _,num := range nums{
	sum += num
}
fmt.Println(sum)


for i,num := range nums{
   fmt.Println(i,num)
}

//map
mpp := map[string]int{"a":1,"b":2,"c":3}

for k,v :=range mpp{
	fmt.Println(k,v)
}


//unicode code point rune
//starting byte of rune

for i,c := range "golang"{
	fmt.Println(i,c)
}

for i,c := range "golang"{
	fmt.Println(i,string(c))
}




}