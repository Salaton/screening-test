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
	CreateOrder(model.OrderInput)
	CreateCustomer(model.CustomerInput)
	CreateItem(model.ItemInput)
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
	// Create Customer, Item and Order tables..
	ps.db.AutoMigrate(&model.Customer{}, &model.Order{}, &model.Item{})

	return nil
}

func (ps *PostgresClient) CreateItem(item model.ItemInput) {
	if err := ps.db.Create(&model.Item{
		Name:     item.Name,
		Quantity: item.Quantity,
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
}

func loopOverItems(itemsInput []*model.ItemInput) []*model.Item {
	var items []*model.Item
	for _, itemInput := range itemsInput {
		items = append(items, &model.Item{
			Name:     itemInput.Name,
			Quantity: itemInput.Quantity,
		})
	}
	return items
}

func (ps *PostgresClient) CreateOrder(order model.OrderInput) {
	log.Printf("Running")
	if err := ps.db.Create(&model.Order{

		// ID:              "texrtxf",
		CustomerName:        order.Customername,
		CustomerPhoneNumber: order.CustomerPhoneNumber,
		//Should return a model.Item type..
		Item:            loopOverItems(order.Item),
		Price:           order.Price,
		DateOrderPlaced: time.Now(),

		// CustomerID:      Customer,
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}

}

func (ps *PostgresClient) CreateCustomer(customer model.CustomerInput) {
	if err := ps.db.Create(&model.Customer{
		// ID:          "xtfg", -->Automatically generated
		Name:        customer.Name,
		Phonenumber: customer.Phonenumber,
		Email:       customer.Email,
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
}

// func (ps *PostgresClient) FetchAllOrdersForCustomer(custID string) {
// 	var customer Customer
// 	if err := ps.db.Where("id = ?", custID).Preload("Orders").First(&customer).Error; err != nil {
// 		log.Printf("Something bad happened %v", err.Error())
// 	}

// 	for _, order := range customer.Orders {
// 		log.Printf("ORDER %v", order)
// 	}
// }
