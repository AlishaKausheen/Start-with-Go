package main

import "fmt"

func main() {
	//	const name="alisha"
	//cannot be reassigned
	//name="xyz"

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port,host)
}