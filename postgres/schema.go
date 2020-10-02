package postgres

import (
	model "github.com/Salaton/screening-test.git/graph/model"
)

// Using a custom schema for the one-to-many relationship..
type Customer struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Phonenumber int      `json:"Phonenumber"`
	Email       string   `json:"Email"`
	Orders      []*Order `json:"orders" gorm:"foreignKey:CustomerID;"`
}

type Order struct {
	ID              int64              `json:"id" gorm:"primaryKey"`
	Item            []*model.ItemInput `json:"item" gorm:"-"`
	Price           float64            `json:"price"`
	DateOrderPlaced string             `json:"date_order_placed"`
	CustomerID      int
}
