package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	database "github.com/Salaton/screening-test/postgres"
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

	mock.ExpectQuery("INSERT INTO customers name, phonenumber,email").WillReturnRows(rows)
	// var customer []model.Customer
	// customer1 := model.Customer{
	// 	Name:        "Elvis",
	// 	Phonenumber: "254712345676",
	// 	Email:       "elvis@gmail.com",
	// }

	// customer = append(customer, customer1)

	// DB.CreateCustomer(customer)

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

func TestGettingUserID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Errorf("An error occured: '%v'", err)
	}
	defer db.Close()

	// DB.Open("postgres")

	rows := sqlmock.NewRows([]string{
		"id", "name", "phonenumber", "email",
	}).AddRow(1, "Elvis", "254712345676", "elvis@gmail.com").AddRow(2, "Timothy", "254712345676", "tim@gmail.com")
	mock.ExpectBegin()

	// Query to fetch from table 'users'
	mock.ExpectQuery("SELECT (.+) FROM `users`").WillReturnRows(rows)
	mock.ExpectRollback()
	_, err = DB.GetUserID("Elvis")

	if err != nil {
		t.Errorf("Expected no error, but got %s instead", err)
	}

	// we make sure that all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestShouldGetPosts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id", "name", "phonenumber", "email",
	}).AddRow(1, "Elvis", "254712345676", "elvis@gmail.com").AddRow(2, "Timothy", "254712345676", "tim@gmail.com")

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT (.+) from users").WillReturnRows(rows)
	mock.ExpectCommit()
	DB.GetUserID("Elvis") // to call the function itself

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestShouldFindCustomers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id", "name", "phonenumber", "email",
	}).AddRow(1, "Elvis", "254712345676", "elvis@gmail.com").AddRow(2, "Timothy", "254712345676", "tim@gmail.com")

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT (.+) from users").WillReturnRows(rows)
	mock.ExpectCommit()
	// DB.FindCustomers() // to call the function itself

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
