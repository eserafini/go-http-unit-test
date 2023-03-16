package client

import (
	"context"
	"eserafini/go-http-unit-test/helpers"
	"net/http"
	"time"
)

const OrderStatusApproved string = "Aprovado"

type ListModel struct {
	StatusCode       int               `json:"status_code"`
	Success          bool              `json:"sucesso"`
	Data             ListModelData     `json:"dados"`
	PayOrderData     []PayOrderModel   `json:"pay_order_data"`
	OrderStatusData  OrderStatusModel  `json:"order_status_data"`
	ConfirmOrderData ConfirmOrderModel `json:"confirm_order_data"`
}

type ListModelData struct {
	OrderStatus        string `json:"statusPedido"`
	OrderPartnerNumber []int  `json:"numeroPedido"`
}

type PayOrderModel struct {
	OrderUUID string `json:"order_uuid"`
}

type OrderStatusModel struct {
	LeadUUID string `json:"lead_uuid"`
}

type ConfirmOrderModel struct {
	LeadUUID string `json:"lead_uuid"`
}

type API struct {
	URL string
}

func (api *API) PayOrder(data []PayOrderModel) (*ListModel, error) {
	order := &ListModel{
		PayOrderData: data,
	}

	to := time.Duration(10)
	opt := &helpers.HttpOptions{
		Ctx:    context.Background(),
		Url:    api.URL + "/pedidos/pci",
		TO:     &to,
		Method: http.MethodPost,
	}

	statusCode, err := helpers.DoRequest(opt, order)
	order.StatusCode = statusCode
	return order, err
}

func (api *API) GetOrderStatus(data OrderStatusModel) (*ListModel, error) {
	order := &ListModel{
		OrderStatusData: data,
	}

	to := time.Duration(10)
	opt := &helpers.HttpOptions{
		Ctx:    context.Background(),
		Url:    api.URL + "/paymentstatus",
		TO:     &to,
		Method: http.MethodGet,
	}

	statusCode, err := helpers.DoRequest(opt, order)
	order.StatusCode = statusCode
	return order, err
}

func (api *API) ConfirmOrder(data ConfirmOrderModel) (*ListModel, error) {
	order := &ListModel{
		ConfirmOrderData: data,
	}

	to := time.Duration(10)
	opt := &helpers.HttpOptions{
		Ctx:    context.Background(),
		Url:    api.URL + "/pedidos/confirmar",
		TO:     &to,
		Method: http.MethodPost,
	}

	statusCode, err := helpers.DoRequest(opt, order)
	order.StatusCode = statusCode
	return order, err
}
