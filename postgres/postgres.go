package postgres

import (
	"log"
	"time"

	model "github.com/Salaton/screening-test.git/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBClient interface to hold Methods that interact with our DB
type DBClient interface {
	Open(dbConnString string) error
	CreateOrder(order model.OrderInput)
	CreateCustomer(model.CustomerInput)
}

// PostgresClient for db references
type PostgresClient struct {
	db *gorm.DB
}

// Open method to create connection with our DB
func (ps *PostgresClient) Open(dbConnString string) error {
	var err error
	ps.db, err = gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		return err
	}
	ps.db.AutoMigrate(&Customer{})
	ps.db.AutoMigrate(&Order{})

	return nil
}

func (ps *PostgresClient) CreateOrder(order model.OrderInput) {
	log.Printf("Running")
	if err := ps.db.Create(&Order{
		// ID:              "texrtxf",
		Item:            order.Item,
		Price:           order.Price,
		DateOrderPlaced: time.Now().String(),
		// CustomerID:      Customer,
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}

}

func (ps *PostgresClient) CreateCustomer(customer model.CustomerInput) {
	if err := ps.db.Create(&Customer{
		// ID:          "xtfg", -->Automatically generated
		Name:        customer.Name,
		Phonenumber: customer.Phonenumber,
		Email:       customer.Email,
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
}

func (ps *PostgresClient) FetchAllOrdersForCustomer(custID string) {
	var customer Customer
	if err := ps.db.Where("id = ?", custID).Preload("Orders").First(&customer).Error; err != nil {
		log.Printf("Something bad happened %v", err.Error())
	}

	for _, order := range customer.Orders {
		log.Printf("ORDER %v", order)
	}
}
