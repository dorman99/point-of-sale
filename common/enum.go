package common

type TOrderStatus string

var OrderStatus = map[string]TOrderStatus{
	"Todo":       "Todo",
	"InProgress": "InProgress",
	"Done":       "Done",
}
