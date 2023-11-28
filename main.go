package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dorman99/point-of-sales/model"
	"github.com/dorman99/point-of-sales/service"
)

type ResponseData struct {
	Success string      `json:"success"`
	Data    interface{} `json:"data"`
}

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
		responsePayload := ResponseData{
			Success: "oke!",
			Data:    items,
		}
		response, err := json.Marshal(responsePayload)
		if err != nil {
			log.Fatalln("failed to retrived items")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jsonƒ")
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatalln("error read body", r.URL)
		}

		var orderItem []model.OrderItem
		err = json.Unmarshal(body, &orderItem)

		if err != nil {
			log.Fatalln("error", r.URL, err)
		}

		orders := service.CreateOrder(orderItem)

		w.WriteHeader(http.StatusOK)

		responsePayload := ResponseData{
			Success: "oke!",
			Data:    orders,
		}

		response, _ := json.Marshal(responsePayload)
		w.Write(response)
	}
}

func main() {
	http.HandleFunc("/health-check", HealthCheck)
	http.HandleFunc("/items", ItemsHandler)
	http.HandleFunc("/orders", OrdersHandler)
	log.Println("Application is Running: 8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
