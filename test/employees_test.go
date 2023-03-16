package test

import (
	"eserafini/go-http-unit-test/client"
	"eserafini/go-http-unit-test/helpers"
	"net/http"
	"testing"
)

func TestAPIGetEmployees(t *testing.T) {
	resp := &client.ListModel{
		Status: "success",
		Data: []client.EmployeeModel{
			{
				ID:             "1",
				EmployeeName:   "Tiger Nixon",
				EmployeeSalary: "320800",
				EmployeeAge:    "61",
				ProfileImage:   "",
			},
		},
	}

	srv := helpers.HttpMock("/api/v1/employees", http.StatusOK, resp)
	defer srv.Close()

	api := client.API{URL: srv.URL}

	employees, err := api.GetEmployees()
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if employees.Status != "success" {
		t.Error("expected status success got:", employees.Status)
	}
	if len(employees.Data) != 1 {
		t.Error("expected 1 data got", len(employees.Data))
	}
}
