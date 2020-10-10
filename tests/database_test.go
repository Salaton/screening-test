package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/Salaton/screening-test.git/graph/model"
	database "github.com/Salaton/screening-test.git/postgres"
)

var DB database.DBClient

func TestPostingToDatabase(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Errorf("An error occured: '%v'", err)
	}
	defer db.Close()

	// Create rows in that mocked database
	rows := sqlmock.NewRows([]string{
		"id", "name", "phonenumber", "email",
	}).AddRow(1, "Elvis", "254712345676", "elvis@gmail.com").AddRow(2, "Timothy", "254712345676", "tim@gmail.com")

	mock.ExpectQuery("^SELECT name, phonenumber from customers$").WillReturnRows(rows)

	var customer []model.Customer
	customer1 := model.Customer{
		Name:        "Elvis",
		Phonenumber: "254712345676",
		Email:       "elvis@gmail.com",
	}

	customer = append(customer, customer1)

	// DB.CreateCustomer()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something unexpected happened: %s", err)
	}
}

func TestFetchPhoneNumber(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Errorf("An error occured: '%v'", err)
	}
	defer db.Close()

	// Create rows in that mocked database
	rows := sqlmock.NewRows([]string{
		"id", "name", "phonenumber", "email",
	}).AddRow(1, "Elvis", "254712345676", "elvis@gmail.com").AddRow(2, "Timothy", "254712345676", "tim@gmail.com")

	mock.ExpectQuery("^SELECT phonenumber from customers$").WillReturnRows(rows)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something unexpected happened: %s", err)
	}
}

// THIS TEST WORKS WELL..
// func TestSendingNotification(t *testing.T) {
// 	got := notification.SendNotification("Elvis Salaton", "254719158559")
// 	want := "Hello there Elvis Salaton Your order has been received"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }
