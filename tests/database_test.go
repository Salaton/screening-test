package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"

	database "github.com/Salaton/screening-test/postgres"
)

var DB database.DBClient

type PostgresClient struct {
	db *gorm.DB
}

func TestCreatingCustomers(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Errorf("An error occured: '%v'", err)
	}
	defer db.Close()
	mock.ExpectExec("INSERT INTO customers\\(name, phonenumber,email\\)").
		WithArgs("name", "phonenumber", "email").
		WillReturnResult(sqlmock.NewResult(1, 1))
	_, err = db.Exec("INSERT INTO customers(name, phonenumber,email) VALUES (?, ?,?)", "name", "phonenumber", "email")

	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestCreatingOrders(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Errorf("An error occured: '%v'", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO orders\\(customer_id, item,Price,DateOrderPlaced\\)").
		WithArgs("customer_id", "item", "Price", "DateOrderPlaced").
		WillReturnResult(sqlmock.NewResult(1, 1))
	_, err = db.Exec("INSERT INTO orders(customer_id, item,Price,DateOrderPlaced) VALUES (?, ?, ?, ?)", "customer_id", "item", "Price", "DateOrderPlaced")
	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

// func TestGetUser(t *testing.T) {
// 	db, mock, err := sqlmock.New()

// 	if err != nil {
// 		t.Errorf("An error occured: '%v'", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectExec("INSERT INTO users\\(user_id, username\\)").
// 		WithArgs("user_id", "username").
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	_, err = db.Exec("SELECT * FROM users VALUES (?, ?)", "user_id", "username")
// 	if err != nil {
// 		t.Errorf("error '%s' was not expected, while inserting a row", err)
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
