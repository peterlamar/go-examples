# sqltesting

This example takes the previous db-cache example and adds sqltesting. Testing
logic that references sql queries can be awkward to mock. This strategy takes
advantage of the [SQL Mock](https://github.com/DATA-DOG/go-sqlmock) library to
make this more elegant.

## Usage

```golang
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

```

## Run the tests

These tests can be run without spinning up the redis or postgres container. 
This makes another argument for test driven development as its actually 
easier to get started because dependencies can all be mocked out. 

```
go test ./...
```

## Setup

### Build and run postgres db

1. This Builds the local postgres container with two local SQL files that create a table
and insert some sample data into the table

```
docker build -t postgres . -f deploy/Dockerfile
```

2. Run the database. DON'T deploy anything with this simple default password,
this is for local learning ONLY.

```
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword  -p 5432:5432 -d  postgres
```

### Run Redis

Redis can be run with the following:

```
docker run --name some-redis -p 6379:6379 -d redis
```


### Build and run the code

```
go build && ./sqltesting
```


#### Optional

Hit the endpoint

```
curl -X GET "http://localhost:8080/getboxdifference/2" -H "accept: application/json"
```

### References
