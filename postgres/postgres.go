package postgres

import (
	"log"

	model "github.com/Salaton/screening-test.git/graph/model"
	notification "github.com/Salaton/screening-test.git/notification"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBClient interface to hold Methods that interact with our DB
type DBClient interface {
	Open(dbConnString string) error
	// CreateOrder(model.OrderInput)
	CreateCustomer(model.CustomerInput)
	// FetchPhoneNumber()
	// FetchAllOrdersForCustomer()
	// CreateItem(model.ItemInput)
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

func loopOverOrders(ordersInput []*model.OrderInput) []*model.Order {
	var orders []*model.Order
	for _, orderInput := range ordersInput {
		orders = append(orders, &model.Order{
			Item:            loopOverItems(orderInput.Item),
			Price:           orderInput.Price,
			DateOrderPlaced: orderInput.DateOrderPlaced,
		})
	}
	return orders
}

func (ps *PostgresClient) CreateCustomer(customer model.CustomerInput) {
	if err := ps.db.Create(&model.Customer{
		// ID:          "xtfg", -->Automatically generated
		Name:        customer.Name,
		Phonenumber: customer.Phonenumber,
		Email:       customer.Email,
		Orders:      loopOverOrders(customer.Orders),
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
	// After the customer created the order, send a notification to them
	notification.SendNotification(customer.Name, customer.Phonenumber)
}

func (ps *PostgresClient) FetchPhoneNumber(customerID int) {
	var customer Customer
	ps.db.Raw("SELECT phonenumber FROM customers WHERE id=?", customerID).Scan(&customer)

	// ps.db.First(&customer.Phonenumber, 1)

	// if err := ps.db.Select("phonenumber").Find(&customer).Error; err != nil {
	// 	log.Printf("Something happened %v", err.Error())
	// }

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
