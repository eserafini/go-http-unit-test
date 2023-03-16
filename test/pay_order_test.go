package test

import (
	"eserafini/go-http-unit-test/client"
	"eserafini/go-http-unit-test/helpers"
	"net/http"
	"testing"
)

func TestAPIPayOrderSuccessResp(t *testing.T) {
	resp := &client.ListModel{
		Success: true,
		PayOrderData: []client.PayOrderModel{
			{
				OrderUUID: "asdasd-dasdasd-asdada-dasdas",
			},
		},
	}

	srv := helpers.HttpMock("/pedidos/pci", http.StatusOK, resp)
	defer srv.Close()

	api := client.API{URL: srv.URL}

	order, err := api.PayOrder(resp.PayOrderData)
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if !order.Success {
		t.Error("expected success true got:", order.Success)
	}
	if len(order.PayOrderData) != 1 {
		t.Error("expected 1 data got", len(order.PayOrderData))
	}
}

func TestAPIGetOrderStatusSuccessResp(t *testing.T) {
	resp := &client.ListModel{
		Success: true,
		Data: client.ListModelData{
			OrderStatus: "Aprovado",
		},
		OrderStatusData: client.OrderStatusModel{
			LeadUUID: "asdasd-dasdasd-asdada-dasdas",
		},
	}

	srv := helpers.HttpMock("/paymentstatus", http.StatusOK, resp)
	defer srv.Close()

	api := client.API{URL: srv.URL}

	order, err := api.GetOrderStatus(resp.OrderStatusData)
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if order.Data.OrderStatus != client.OrderStatusApproved {
		t.Error("expected order status Aprovado got:", order.Data.OrderStatus)
	}
}

func TestAPIConfirmOrderSuccessResp(t *testing.T) {
	resp := &client.ListModel{
		Success: true,
		ConfirmOrderData: client.ConfirmOrderModel{
			LeadUUID: "asdasd-dasdasd-asdada-dasdas",
		},
	}

	srv := helpers.HttpMock("/pedidos/confirmar", http.StatusOK, resp)
	defer srv.Close()

	api := client.API{URL: srv.URL}

	order, err := api.ConfirmOrder(resp.ConfirmOrderData)
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if !order.Success {
		t.Error("expected success true got:", order.Success)
	}
}
