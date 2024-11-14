package main

import (
	"fmt"
	"maps"
)

// maps are associative data structure
//maps -> hash, object, dictionary
func main(){
//creating map

m := make(map[string]string)

//setting an element
m["name"] ="golang"
m["x"] ="y"
//getting an element
fmt.Println(m["name"])

//imp: if key doesnt exist in the map then it returns zero value

//map length
fmt.Println(len(m))


//delete 
delete(m,"name")


//clear map
clear(m)

mpp := map[string]int{"a":1,"b":2,"c":3}
fmt.Println(mpp)


// param1 gives value param 2 checks if it exists or not
v,ok :=mpp["a"]
fmt.Println(v)
if ok{
	fmt.Println("ok")
}else{
	fmt.Println("not ok")
}


//check if 2 maps are equal or not
mpp1 := map[string]int{"a":1,"b":2,"c":3}
mpp2 := map[string]int{"a":1,"b":2,"c":3}

fmt.Println(maps.Equal(mpp1,mpp2))
}