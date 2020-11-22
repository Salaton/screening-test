package postgres

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	model "github.com/Salaton/screening-test/graph/model"
	"github.com/go-test/deep"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository DBClient
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	dbURI := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		"sala", "salaton", "testdatabase", "5432", "disable", "Africa/Nairobi")
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	defer db.Close()

	s.DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	require.NoError(s.T(), err)

	// s.db.LogMode(true)

	s.repository = CreateRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	(s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Create_Customer() {
	CustomerDetails := model.CustomerInput{
		Name:        "Joseph",
		Phonenumber: "254719158559",
		Email:       "kosen@gmail.com",
		Password:    HashPassword("password12345"),
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "customer" ("id","name","phonenumber","email","password")
			VALUES ($1,$2,$3,$4,$5)`))
	s.repository.CreateCustomer(CustomerDetails)
}

func (s *Suite) Test_Get_Customer() {
	var (
		id          = 1
		name        = "test-name"
		phonenumber = "254719158559"
		email       = "elvisdenis056@gmail.com"
	)
	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "customer" WHERE (id = $1)`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phonenumber", "email"}).
			AddRow(id, name, phonenumber, email))
	s.mock.ExpectCommit()

	res, err := s.repository.GetCustomer(email)

	require.NoError(s.T(), err)
	deep.Equal(&model.Customer{ID: id, Name: name, Phonenumber: phonenumber, Email: email}, res)
}

func (s *Suite) Test_Get_Customer_ID() {
	var (
		id          = 1
		name        = "test-name"
		phonenumber = "254719158559"
		email       = "elvisdenis056@gmail.com"
	)
	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id FROM "customer" WHERE (id = $1)`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phonenumber", "email"}).
			AddRow(id, name, phonenumber, email))
	s.mock.ExpectCommit()

	res, err := s.repository.GetCustomerID(email)

	require.NoError(s.T(), err)
	deep.Equal(&model.Customer{ID: id, Name: name, Phonenumber: phonenumber, Email: email}, res)
}

func (s *Suite) Test_FindAllCustomers() {

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "customer" `))
	s.mock.ExpectCommit()

	s.repository.FindCustomers()
}

func (s *Suite) Test_FindAllOrders() {

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "order" `))
	s.mock.ExpectCommit()

	s.repository.FindOrders()
}

func (s *Suite) Test_Create_Order() {
	var items []*model.ItemInput
	item := &model.ItemInput{
		ProductID: 1,
		Quantity:  1,
	}

	items = append(items, item)
	OrderDetails := model.OrderInput{
		CustomerID: 10,
		Item:       items,
		Price:      5000.00,
	}

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "order" ("id","customer_id","price","date_order_placed")
			VALUES ($1,$2,$3,$4)`))

	s.mock.ExpectCommit()
	s.repository.CreateOrder(OrderDetails)
}

func (s *Suite) Test_Update_Order() {
	var items []*model.ItemInput
	item := &model.ItemInput{
		ProductID: 1,
		Quantity:  2,
	}

	items = append(items, item)

	OrderDetails := model.OrderInput{
		CustomerID: 10,
		Item:       items,
		Price:      5000.00,
	}

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`UPDATE "order" ("id","customer_id","price","date_order_placed")
			SET Item=$2, Price=$3,DateOrderPlaced=$3 WHERE customer_id=1`))

	s.mock.ExpectCommit()
	s.repository.UpdateOrder(22, OrderDetails)
}

func (s *Suite) Test_DeleteOrder() {

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`DELETE FROM "order" WHERE (id = $1)`))
	s.mock.ExpectCommit()

	s.repository.DeleteOrder(2)
}

func (s *Suite) Test_Authentication() {
	logindetails := model.LoginDetails{
		Email:    "testemail@gmail.com",
		Password: "password12345",
	}
	var Password string

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT password FROM "customers" WHERE (email = $1)`)).WillReturnRows(sqlmock.NewRows([]string{"password"}).
		AddRow(Password))
	s.mock.ExpectCommit()
	s.repository.Authenticate(logindetails)
}

func (s *Suite) Test_DBConnection() {
	dbconnstring := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		"sala", "salaton", "savannahtest", "5432", "disable", "Africa/Nairobi")
	s.repository.Open(dbconnstring)
}

func TestHashPassword(t *testing.T) {

	t.Run("compare password and hash", func(t *testing.T) {
		password := "password12345"
		hash := "$2a$14$fNPi4m0o8ooKCUYS4TlU3erQst453fiF.QvtFyKu2EtJGLDPG4kLG"
		got := CheckPasswordHash(password, hash)
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

}

func (s *Suite) Test_Create_Product() {
	ProductDetails := model.ProductInput{
		Name:  "Xbox One S",
		Price: 5000.00,
	}
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "product" ("id","name","price")
			VALUES ($1,$2,$3)`))
	s.repository.CreateProduct(ProductDetails)
}

func (s *Suite) Test_FindAllProducts() {

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "product" `))
	s.mock.ExpectCommit()

	s.repository.FindProducts()
}
