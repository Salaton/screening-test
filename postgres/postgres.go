package postgres

import (
	"log"
	"time"

	model "github.com/Salaton/screening-test/graph/model"
	notification "github.com/Salaton/screening-test/notification"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBClient interface to hold Methods that interact with our DB
type DBClient interface {
	Open(dbConnString string) error
	CreateOrder(model.OrderInput)
	CreateCustomer(model.CustomerInput)
	Authenticate(model.LoginDetails) bool
	GetCustomerID(email string) (int, error)
	GetCustomer(email string) (model.Customer, error)
	FindCustomers() []*model.Customer
	FindOrders() []*model.Order
	DeleteOrder(id int)
	UpdateOrder(orderID int, order model.OrderInput)
}

// PostgresClient exposes reference to the DB
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

func (ps *PostgresClient) GetCustomer(email string) (model.Customer, error) {
	var customer model.Customer
	row := ps.db.Table("customers").Where("email = ?", email).Select("id,name,phonenumber,email,").Row()
	row.Scan(&customer.ID, &customer.Name, &customer.Phonenumber, &customer.Email)

	return customer, nil
}

// GetCustomerID -> Check if a customer exists in a DB using the Email
func (ps *PostgresClient) GetCustomerID(email string) (int, error) {
	var customer model.Customer
	row := ps.db.Table("customers").Where("email = ?", email).Select("id").Row()
	var id int
	row.Scan(&customer.ID)

	return id, nil
}

func (ps *PostgresClient) CreateOrder(order model.OrderInput) {
	var customer model.Customer
	if err := ps.db.Create(&model.Order{
		CustomerID:      order.CustomerID,
		Item:            loopOverItems(order.Item),
		Price:           order.Price,
		DateOrderPlaced: time.Now(),
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
	// Use GORM API build SQL
	row := ps.db.Table("customers").Where("id = ?", order.CustomerID).Select("phonenumber", "name").Row()
	row.Scan(&customer.Phonenumber, &customer.Name)
	notification.SendNotification(customer.Name, customer.Phonenumber)

}

func (ps *PostgresClient) UpdateOrder(orderID int, order model.OrderInput) {
	var customer model.Customer
	if err := ps.db.Save(&model.Order{
		ID:              orderID,
		CustomerID:      order.CustomerID,
		Item:            loopOverItems(order.Item),
		Price:           order.Price,
		DateOrderPlaced: time.Now(),
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
	// Use GORM API build SQL
	row := ps.db.Table("customers").Where("id = ?", order.CustomerID).Select("phonenumber", "name").Row()
	row.Scan(&customer.Phonenumber, &customer.Name)
	notification.SendNotification(customer.Name, customer.Phonenumber)

}

func (ps *PostgresClient) DeleteOrder(id int) {
	var order model.Order
	ps.db.Where("id = ?", order.ID).Delete(&model.Item{})
	ps.db.Delete(&model.Order{}, id)
}

func (ps *PostgresClient) CreateCustomer(customer model.CustomerInput) {
	if err := ps.db.Create(&model.Customer{
		Name:        customer.Name,
		Phonenumber: customer.Phonenumber,
		Email:       customer.Email,
		Password:    HashPassword(customer.Password),
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
}

func (ps *PostgresClient) FindCustomers() []*model.Customer {
	var customer []*model.Customer
	if err := ps.db.Table("customers").Find(&customer).Error; err != nil {
		log.Printf("Something bad happened %v", err.Error())
	}
	return customer
}

// FindOrders method to return all orders from the database
func (ps *PostgresClient) FindOrders() []*model.Order {
	var order []*model.Order
	if err := ps.db.Table("orders").Preload("Customer").Preload("Item").Find(&order).Error; err != nil {
		log.Printf("Something bad happened %v", err.Error())

	}
	return order
}

//Authenticate Method to authenticate customers
func (ps *PostgresClient) Authenticate(customer model.LoginDetails) bool {
	row := ps.db.Table("customers").Where("email = ?", customer.Email).Select("password").Row()
	var hashedPassword string
	row.Scan(&hashedPassword)

	return CheckPasswordHash(customer.Password, hashedPassword)
}
