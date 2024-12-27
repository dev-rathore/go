package main

import (
	"fmt"
	"time"
)

type customer struct {
	name string
	email string
}

type product struct {
	name string
	price float32
}

type order struct {
	orderId int
	amount float32
	status string
	createdAt time.Time
	customer
	// customerInfo customer
	products []product
}

func main() {
	// newCustomer := customer{
	// 	name: "John Doe",
	// 	email: "john.doe@gmail.com",
	// }

	newOrder := order{
		orderId: 1,
		amount: 100.0,
		status: "Received",
		createdAt: time.Now(),
		customer: customer{
			name: "John Doe",
			email: "john.doe@gmail.com",
		},
		// customer: newCustomer,
		products: []product{
			{
				name: "Laptop",
				price: 2000.0,
			},
			{
				name: "Mouse",
				price: 20.0,
			},
		},
	}

	newOrder.customer.name = "Alice Johnson"

	fmt.Println(newOrder)
}
