package postgres

// // Here we declare our models --> normal structs with basic Go types..
import (
	"time"
)

type Customer struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Phonenumber int      `json:"Phonenumber"`
	Email       string   `json:"Email"`
	Orders      []*Order `json:"orders"`
}

type Item struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	OrderID  uint   `json:"-"`
}

type ItemInput struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Order struct {
	ID                  int       `json:"id" gorm:"primaryKey"`
	CustomerName        string    `json:"customerName"`
	CustomerPhoneNumber int       `json:"customerPhoneNumber"`
	Item                []*Item   `json:"item"`
	Price               float64   `json:"price"`
	DateOrderPlaced     time.Time `json:"date_order_placed"`
	CustomerID          int       `json:"-"`
}

type CustomerInput struct {
	Name        string `json:"name"`
	Phonenumber int    `json:"Phonenumber"`
	Email       string `json:"Email"`
}

type OrderInput struct {
	Customername        string       `json:"customername"`
	CustomerPhoneNumber int          `json:"customerPhoneNumber"`
	Item                []*ItemInput `json:"item"`
	Price               float64      `json:"price"`
	DateOrderPlaced     time.Time    `json:"date_order_placed"`
}
