package main 

import (
	"fmt"
	"time"
)

type ticketOrder struct{
	customerName string
	ticketRequested int
	orderStatus string
}
//order placement
func placeOrder() <-chan ticketOrder{
	out := make(chan ticketOrder)
	go func (){
		customers := []string{"Alice", "zack", "franc", "fred", "janet"}
		tickets := []int{1, 2, 3, 4, 5}

		for i := 0; i < len(customers); i++{
			order := ticketOrder{customerName: customers[i], ticketRequested: tickets[i], orderStatus: "Pending"}
			fmt.Printf("order placed: %s wants %d tickets\n", order.customerName, order.ticketRequested)
			out <-order
			time.Sleep(500 * time.Millisecond) //simulate time delay

	}
	close(out)}()
	return out
}
//Order Validation
func validateOrder(in <-chan ticketOrder) <-chan ticketOrder{
	out := make(chan ticketOrder)
	go func(){
		for order := range in{
			if order.ticketRequested <= totalTickets{
				order.orderStatus = "Valid"
				totalTickets -= order.ticketRequested
			}else{
				order.orderStatus = "Rejected(not enough tickets)"

			}
				out <- order
				time.Sleep(500 * time.Millisecond) //simulate processing time
		}
		close(out)
	}()
	return out
			
}