package main

import (
	"eserafini/go-http-unit-test/client"
	"fmt"
)

func main() {
	fmt.Println("starting")

	respPayOrder := &client.ListModel{
		PayOrderData: []client.PayOrderModel{
			{
				OrderUUID: "asdasd-dasdasd-asdada-dasdas",
			},
		},
	}

	respOrderStatus := &client.ListModel{
		OrderStatusData: client.OrderStatusModel{
			LeadUUID: "asdasd-dasdasd-asdada-dasdas",
		},
	}

	respConfirmOrder := &client.ListModel{
		ConfirmOrderData: client.ConfirmOrderModel{
			LeadUUID: "asdasd-dasdasd-asdada-dasdas",
		},
	}

	api := client.API{URL: "http://localhost:3000"}

	order, err := api.PayOrder(respPayOrder.PayOrderData)
	if err != nil {
		fmt.Println("pay order error:", err)
	}

	fmt.Println("pay order response", order)

	order, err = api.GetOrderStatus(respOrderStatus.OrderStatusData)
	if err != nil {
		fmt.Println("get order error:", err)
	}

	fmt.Println("get order response", order)

	order, err = api.ConfirmOrder(respConfirmOrder.ConfirmOrderData)
	if err != nil {
		fmt.Println("confirm order error:", err)
	}

	fmt.Println("confirm order response", order)
}
