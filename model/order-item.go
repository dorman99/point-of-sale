package model

type OrderItem struct {
	Item     Item `json:"item"`
	Quantity int  `json:"quantity"`
}
