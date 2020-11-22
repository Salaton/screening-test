package postgres

import (
	"log"
	"time"

	model "github.com/Salaton/screening-test/graph/model"
	"github.com/Salaton/screening-test/notification"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBClient interface to hold Methods that interact with our DB
type DBClient interface {
	Open(dbConnString string) error
	CreateOrder(model.OrderInput)
	CreateCustomer(model.CustomerInput)
	CreateProduct(model.ProductInput)
	Authenticate(model.LoginDetails) bool
	GetCustomerID(email string) (int, error)
	GetCustomer(email string) (model.Customer, error)
	FindCustomers() []*model.Customer
	FindOrders() []*model.Order
	FindProducts() []*model.Product
	DeleteOrder(id int)
	UpdateOrder(orderID int, order model.OrderInput)
}

// PostgresClient exposes reference to the DB
type PostgresClient struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) DBClient {
	return &PostgresClient{
		db: db,
	}
}

// Open method to create connection with our DB
func (ps *PostgresClient) Open(dbConnString string) error {
	var err error
	ps.db, err = gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		return err
	}
	// Create Customer, Item and Order tables..
	ps.db.AutoMigrate(&model.Customer{}, &model.Order{}, &model.Product{}, &model.OrderItem{})

	return nil
}

// GetCustomer method to find a particular customer based on their email
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

// CreateOrder method that creates an order and sends a notification to the user
func (ps *PostgresClient) CreateOrder(order model.OrderInput) {
	var customer model.Customer

	if err := ps.db.Create(&model.Order{
		CustomerID:      order.CustomerID,
		Item:            loopOverItems(order.Item),
		TotalPrice:      order.Price,
		DateOrderPlaced: time.Now(),
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}

	row := ps.db.Table("customers").Where("id = ?", order.CustomerID).Select("phonenumber", "name").Row()
	row.Scan(&customer.Phonenumber, &customer.Name)
	notification.SendNotification(customer.Name, customer.Phonenumber)

}

// UpdateOrder method that takes in an order ID and updates the selected order
// then sends a notification
func (ps *PostgresClient) UpdateOrder(orderID int, order model.OrderInput) {
	var customer model.Customer

	if err := ps.db.Updates(&model.Order{
		ID:              orderID,
		CustomerID:      order.CustomerID,
		Item:            loopOverItems(order.Item),
		TotalPrice:      order.Price,
		DateOrderPlaced: time.Now(),
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
	// Use GORM API build SQL
	row := ps.db.Table("customers").Where("id = ?", order.CustomerID).Select("phonenumber", "name").Row()
	row.Scan(&customer.Phonenumber, &customer.Name)
	notification.SendNotification(customer.Name, customer.Phonenumber)

}

// DeleteOrder to delete an order based on the passed order id
func (ps *PostgresClient) DeleteOrder(id int) {
	ps.db.Where("order_id = ?", id).Delete(&model.OrderItem{})
	ps.db.Delete(&model.Order{}, id)
}

// CreateProduct to create an order
func (ps *PostgresClient) CreateProduct(product model.ProductInput) {
	if err := ps.db.Create(&model.Product{
		Name:  product.Name,
		Price: product.Price,
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
}

// CreateCustomer to create and store a customer in the DB
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

// FindProducts method to return all products stores in the db
func (ps *PostgresClient) FindProducts() []*model.Product {
	var product []*model.Product
	if err := ps.db.Table("products").Find(&product).Error; err != nil {
		log.Printf("Something bad happened %v", err.Error())
	}
	return product
}

// FindCustomers method to query and return all stored customers
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
