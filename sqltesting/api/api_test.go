package api

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/peterlamar/go-examples/sqltesting/database"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// Testing the GetDifference function
func TestGetDifference(t *testing.T) {

	expectedResult := 2
	movieID := 1

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	database.SetDB(sqlx.NewDb(mockDB, "sqlmock"))

	baseRate := sqlmock.NewRows([]string{"table_id", "movie_title",
		"domestic_box", "worldwide_box"}).
		AddRow(1, "monty movie", 10, 12)

	mock.ExpectPrepare(
		"^select table_id, movie_title.*$").
		ExpectQuery().WillReturnRows(baseRate)

	rtn := GetDifference(movieID)

	assert.Equal(t, expectedResult, rtn, "The return value should equal the expected value.")
}
