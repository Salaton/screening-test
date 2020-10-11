package postgres

import (
	"log"
	"time"

	model "github.com/Salaton/screening-test.git/graph/model"
	notification "github.com/Salaton/screening-test.git/notification"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBClient interface to hold Methods that interact with our DB
type DBClient interface {
	Open(dbConnString string) error
	CreateOrder(model.OrderInput)
	CreateCustomer(model.CustomerInput)
	CreateUser(model.CreatedUser)
	Authenticate(model.LoginDetails) bool
	GetUserID(username string) (int, error)
	FindCustomers()
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
	ps.db.AutoMigrate(&model.Customer{}, &model.Order{}, &model.Item{}, &model.User{})

	return nil
}

func (ps *PostgresClient) CreateUser(user model.CreatedUser) {
	if err := ps.db.Create(&model.User{
		Username: user.Username,
		Password: HashPassword(user.Password),
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err)
	}
}

// GetUserID -> Check if a user exists in a DB using the username
func (ps *PostgresClient) GetUserID(username string) (int, error) {
	var user model.User
	row := ps.db.Table("users").Where("username = ?", username).Select("id").Row()
	var id int
	row.Scan(&user.ID)

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

func (ps *PostgresClient) CreateCustomer(customer model.CustomerInput) {
	if err := ps.db.Create(&model.Customer{
		Name:        customer.Name,
		Phonenumber: customer.Phonenumber,
		Email:       customer.Email,
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err.Error())
	}
}

func (ps *PostgresClient) FindCustomers() model.Customer {
	var customer model.Customer
	if err := ps.db.Find(&customer).Error; err != nil {
		log.Printf("Something bad happened %v", err.Error())
	}
	return customer
}

// Method to authenticate users
func (ps *PostgresClient) Authenticate(user model.LoginDetails) bool {
	row := ps.db.Table("users").Where("username = ?", user.Username).Select("password").Row()
	var hashedPassword string
	row.Scan(&hashedPassword)

	return CheckPasswordHash(user.Password, hashedPassword)
}
