package main

import (
	"fmt"
	
	"time"
)

//order struct
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time //nano sec precision
}

//func to change status
//receiver type 
func (o *order) changeStatus(status string){
o.status =status
}

func (o order) getAmount() float32{
	return o.amount
	}

	//constructor
	func newOrder(id string,amount float32, status string) *order{
		//initial setup goes here
		myOrder3 := order{
			id: id,
			amount: amount,
			status: status,
		
		 }
		 return &myOrder3
	}


	//struct embedd
	type customer struct{
		name string
		phone string
	}
	type customerOrder struct{
		id string
	amount float32
	status string
	createdAt time.Time
	customer
	}
func main(){

 myOrder := order{
	id: "1",
	amount: 50.00,
	status: "received",

 }
 myOrder2 := order{
	id: "2",
	amount: 500,
	status: "delevered",
	createdAt: time.Now(),
 }
 myOrder.createdAt =time.Now()

 myOrder.changeStatus("confirmed")


 fmt.Println(myOrder.id)
 fmt.Println(myOrder)
 fmt.Println(myOrder2)

 myOrder3:=newOrder("3",30.50, "recieved")
 fmt.Println(myOrder3)


 //struct without name

 language := struct{
	name string
	isGood bool
 }{"golang",true}

 fmt.Println(language)

 customer1 := customer{
	name: "Alisha",
	phone: "90088888",
 }
 newCustomerOrder :=customerOrder{
	id: "4",
	amount: 300,
	status: "ask",
customer: customer1,
 }

 fmt.Println(newCustomerOrder)
}