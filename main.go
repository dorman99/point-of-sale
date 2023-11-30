package main

import (
	"log"
	"net/http"

	"github.com/dorman99/point-of-sales/common"
	"github.com/dorman99/point-of-sales/model"
	"github.com/dorman99/point-of-sales/service"
	"github.com/dorman99/point-of-sales/util"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "application is healty"}`))
	}
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)

		items := service.FindAllItems()
		response := util.ParseResponseBody(items)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jsonƒ")
	switch r.Method {
	case "POST":
		var orderItem []model.OrderItem
		util.ParseRequestBody(r, &orderItem)

		order := service.CreateOrder(orderItem)
		response := util.ParseResponseBody(order)

		go service.ProcessIncomingOrder(order)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func OrderAdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var requestUpdateStatus common.UpdateOrderRequestDto
		util.ParseRequestBody(r, &requestUpdateStatus)
		service.UpdateOrder(requestUpdateStatus)

		response := util.ParseResponseBody(nil)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func main() {
	http.HandleFunc("/health-check", HealthCheck)
	http.HandleFunc("/items", ItemsHandler)
	http.HandleFunc("/orders", OrdersHandler)
	http.HandleFunc("/admin/orders", OrderAdminHandler)
	log.Println("Application is Running: 8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
