package main

import "fmt"

// enumerated types

type MyType string

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Received OrderStatus = "Received"
	Confirmed = "Confirmed"
	Prepared = "Prepared"
	Delivered = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}

func main() {
	changeOrderStatus(Received)
	changeOrderStatus(Prepared)
	changeOrderStatus(Delivered)
}
