package main

import (
	"fmt"
	"time"
)

type order struct {
	orderId int
	customer string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
}

// receiver type
func (o *order) changeStatus(status string) {
	// Struct do the dereferencing automatically
	o.status = status
}

// Works without *
func (o order) getAmount() float32 {
	return o.amount
}

func newOrder(
	orderId int,
	customer string,
	amount float32,
	status string,
) *order {
	// Initialization goes here

	order := order{
		orderId: orderId,
		customer: customer,
		amount: amount,
		status: status,
		createdAt: time.Now(),
	}

	return &order
}

func main() {
	var order1 = order{
		orderId: 1,
		customer: "John Doe",
		amount: 100.0,
		status: "Received",
		createdAt: time.Now(),
	}

	fmt.Println(order1)
	fmt.Println(order1.createdAt)

	order2 := order{
		orderId: 2,
		customer: "Jane Doe",
		amount: 200.0,
		status: "Shipped",
		createdAt: time.Now(),
	}

	fmt.Println(order2)
	fmt.Println(order2.createdAt)

	order1.changeStatus("Shipped")
	fmt.Println(order1)
	fmt.Println(order1.getAmount())

	// If we don't set any field, default value is zero value
	// int: 0, float: 0.0, string: "", bool: false, pointer: nil

	order3 := newOrder(3, "Alice", 300.0, "Delivered")
	order3.changeStatus("Completed")
	fmt.Println(order3)
	fmt.Println(order3.amount)

	// Inline struct
	order4 := struct {
		orderId int
		customer string
		amount float32
		status string
		createdAt time.Time
	}{
		orderId: 4,
		customer: "Bob",
		amount: 400.0,
		status: "Received",
		createdAt: time.Now(),
	}

	fmt.Println(order4)
}
