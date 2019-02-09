# sqltesting

This example takes the previous db-cache example and adds sqltesting. Testing
logic that references sql queries can be awkward to mock. This strategy takes
advantage of the [SQL Mock](https://github.com/DATA-DOG/go-sqlmock) library to make this more elegant.

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
