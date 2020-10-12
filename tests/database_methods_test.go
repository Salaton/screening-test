package tests

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/Salaton/screening-test/graph/model"
	postgresdatabase "github.com/Salaton/screening-test/postgres"
	"github.com/go-test/deep"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	// "gorm.io/gorm"
)

// Setup our suite..
type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	database postgresdatabase.DBClient
	customer *model.Customer
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	// s.database = CreateRepository(s.DB)
}
func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
func (s *Suite) Test_repository_GetUserByID() {
	var (
		// id       = uuid.NewV4()
		id       = 1
		username = "Elvis"
		password = "password12345"
	)
	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT (.+) FROM `users`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
			AddRow(id, username, password))

	// s.mock.ExpectQuery(regexp.QuoteMeta(
	// 	`SELECT * FROM "user" WHERE (id = $1)`)).
	// 	WithArgs(id).
	// 	WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
	// 		AddRow(id, username, password))
	s.mock.ExpectCommit()

	res, err := s.database.GetUserID(username)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&model.User{ID: id, Username: username, Password: password}, res))
}
