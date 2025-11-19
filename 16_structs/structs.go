package main

import (
	"fmt"
	"time"
)

// order struct
// type orders struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time //nanoseconds precision

// }

type customer struct {
	name  string
	phone int
}

type order struct {
	id       string
	amount   float32
	status   string
	cretedAt time.Time
	customer
}

// func newOrder(id string, amount float32, status string) *orders {
// 	// Create a variable named myOrder.
// 	// It is of type 'order' (a struct).
// 	// We fill the struct fields with the values passed to the function.
// 	myOrder := orders{
// 		id:     id,     // set the struct's id field to the id argument
// 		amount: amount, // set the amount field to the amount argument
// 		status: status, // set the status field to the status argument
// 	}

// 	// Return the address of myOrder.
// 	// This means the function returns a pointer (*order).
// 	return &myOrder
// }

// receiver type
//
//	func (o *orders) changeStatus(status string) {
//		o.status = status
//	}
//
//	func (o *orders) getAmount() float32 {
//		return o.amount
//	}
func main() {

	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: 123456789,
	// }
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		// customer: newCustomer,
		customer: customer{
			name:  "john",
			phone: 123456789},
	}
	newOrder.customer.name = "robin"
	fmt.Println(newOrder)
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}

	// fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder.amount)

	// myOrder := orders{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder.getAmount())

	// myOrder.createdAt = time.Now()

	// fmt.Println(myOrder.status)

	// myOrder2 := orders{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"
	// fmt.Println("Order struct", myOrder2)
	// // fmt.Println(myOrder.status)
	// fmt.Println("Order struct", myOrder)

}
