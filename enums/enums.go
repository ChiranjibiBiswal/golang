package main

import "fmt"

// OrderStatus represents the stages of an order's lifecycle.
type OrderStatus int

const (
	Received OrderStatus = iota
	Process
	Dispatched
	Delivered
)

func checkOrderStatus(status OrderStatus) {
	fmt.Println("Order status is:", status)
}

// PaymentStatus represents the result of a payment operation.
type PaymentStatus string

const (
	Processed PaymentStatus = "processed"
	Paid                    = "paid"
	Failed                  = "failed"
)

func checkPaymentStatus(status PaymentStatus) {
	fmt.Println("Payment status is:", status)
}

func main() {
	checkOrderStatus(Delivered)
	checkPaymentStatus(Failed)

}
