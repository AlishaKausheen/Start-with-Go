package main

import "fmt"

func add(a int, b int) int {

	return a + b
}
// if params are of same type we can write like

func sub(a, b int) int {

	return a - b
}

//can return many things

func getLanguages() (string,string,string){
    return "golang", "c", "js"
}

//passing function as an argument in function

func process () func (a int ) int{
    return func(a int) int{
        return 5;
    }
}
func main() {

fn:= process()
fn(6)
fmt.Println(fn(7))

	result := add(3, 5)

	fmt.Println(result)
    //to supress lang3 use underscore
    lang1, lang2, _ := getLanguages()
    fmt.Println(lang1,lang2)
    fmt.Println(getLanguages())
}