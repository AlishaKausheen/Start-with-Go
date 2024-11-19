package main

import "fmt"

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

// enumerated types

type OrderStatus int

const (
	Received OrderStatus=iota
	Confirmed 
	Prepared
	Delivered
)

func main() {
changeOrderStatus(Prepared)
}