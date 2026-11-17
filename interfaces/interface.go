package main

import "fmt"

// paymenter describes the operations a payment gateway must provide.
type paymenter interface {
	pay(amount float32)
	printDetails(name string, isPaid bool)
}

// payment uses any gateway that implements paymenter.
type payment struct {
	gateway paymenter
}

// makePayment charges the amount and then displays the payment details.
func (p payment) makePayment(amount float32, name string, isPaid bool) {
	// gateway := razorPay{}
	// gateway.pay(amount)
	p.gateway.pay(amount)
	p.gateway.printDetails(name, isPaid)
}

// razorPay represents a Razorpay payment gateway.
type razorPay struct{}

// pay processes a payment through Razorpay.
func (r razorPay) pay(amount float32) {
	fmt.Println("Making payment using razorpay: ", amount)
}

// printDetails displays the customer and payment status for a Razorpay payment.
func (r razorPay) printDetails(name string, isPaid bool) {
	fmt.Println("Razorpay payment for", name, "paid:", isPaid)
}

// stripe represents a Stripe payment gateway.
type stripe struct{}

// pay processes a payment through Stripe.
func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe: ", amount)
}

// printDetails displays the customer and payment status for a Stripe payment.
func (s stripe) printDetails(name string, isPaid bool) {
	fmt.Println("Stripe payment for", name, "paid:", isPaid)
}

// fakePayment is a simple gateway implementation for demonstration or testing.
type fakePayment struct{}

// pay simulates processing a payment.
func (f fakePayment) pay(amount float32) {
	fmt.Println("Making fake payment for testing")
}

// printDetails displays the supplied customer and payment status.
func (f fakePayment) printDetails(name string, isPaid bool) {
	fmt.Println("for fake payment", name, isPaid)
}

func main() {
	// gw := razorPay{}
	// Use fakePayment to demonstrate a type satisfying the paymenter interface.
	gw := fakePayment{}
	paying := payment{
		gateway: gw,
	}
	paying.makePayment(100, "Raju", true)
}
