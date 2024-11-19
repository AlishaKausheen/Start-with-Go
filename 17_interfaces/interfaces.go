package main

import "fmt"

type paymenter interface{
	pay(amount float32)

}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
//razorpayPaymentGw :=razorpay{}
//razorpayPaymentGw.pay(amount)
//stripePaymentGw := stripe{}
p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment uaing razorpay", amount)
}

//type stripe struct{}

//func (s stripe) pay(amount float32){
//	fmt.Println("making payment using stripe", amount)
//}

func main() {
	//stripePaymentGw := stripe{}
	razorpayPaymentGw := razorpay{}
   newPayment:=payment{
	gateway: razorpayPaymentGw ,
   }
   newPayment.makePayment(100)

}