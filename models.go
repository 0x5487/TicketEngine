package main

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

type Order struct {
	Id          string
	LineItems   []LineItem
	Subtotal    int64
	Total       int64
	BillingInfo BillingInfo
	CustomForm  string
	State       int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type BillingInfo struct {
	BillingFirstName   string
	BillingLastName    string
	BillingFullName    string
	BillingAddress1    string
	BillingAddress2    string
	BillingCity        string
	BillingPostCode    string
	BillingCountryCode string
	BillingPhoneNumber string
	BillingEmail       string
}

type Ticket struct {
	Id    string
	Name  string
	Price int64
}

type LineItem struct {
	Id            int
	Product       interface{}
	Quantity      int
	UnitPrice     int64
	ExtendedPrice int64
}

type RequestTicket struct {
	TicketId string
	package main int
}
