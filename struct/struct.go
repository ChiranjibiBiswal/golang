package main

import (
	"fmt"
	"time"
)

// customer stores basic information for a buyer.
type customer struct {
	name    string
	age     int
	isAdult bool
}

// order is a struct that contains both primitive and nested values.
type order struct {
	id              string
	amount          float32
	status          string
	quantity        int
	createdAt       time.Time
	customerDetails customer
}

// newOrder creates a new order with a few required fields.
func newOrder(id string, amount float32, status string) *order {
	userOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &userOrder
}

// changeStatus updates the status using a pointer receiver.
func (o *order) changeStatus(status string) {
	o.status = status
}

// getAmount returns the total amount for the order.
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// Struct literal with named fields.
	user1Order := order{
		id:     "21",
		amount: 100.00,
		status: "RECEIVED",
	}
	fmt.Println("user1Order:", user1Order)

	// You can also initialize a struct and set values later.
	user1Order.createdAt = time.Now()
	user1Order.quantity = 3
	fmt.Println("Updated user1Order:", user1Order)

	// Composite literal with all fields specified.
	user2Order := order{
		id:        "22",
		amount:    200,
		status:    "PAID",
		quantity:  5,
		createdAt: time.Now(),
	}
	fmt.Println("user2Order:", user2Order)

	// A method on a pointer receiver can modify the original struct.
	user1Order.changeStatus("PROCESSED")
	fmt.Println("After status change:", user1Order)
	fmt.Println("Amount of user2Order:", user2Order.getAmount())

	// Constructor function returns a pointer to a struct.
	user3Order := newOrder("23", 300.00, "PROGRESS")
	fmt.Println("user3Order:", *user3Order)

	// Anonymous struct type can be defined inline.
	language := struct {
		name   string
		isGood bool
	}{
		name:   "English",
		isGood: true,
	}
	fmt.Println("Language:", language.name, language.isGood)

	// Nested structs are used to organize related data.
	user2Order.customerDetails = customer{name: "rishi", age: 25, isAdult: true}
	fmt.Println("Customer details:", user2Order.customerDetails)

	// A zero-value struct can be created without explicit initialization.
	var defaultOrder order
	fmt.Println("Zero-value order:", defaultOrder)
}
