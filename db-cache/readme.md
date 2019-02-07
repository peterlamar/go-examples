# db-cache

This is an example using redis and postgres in a db caching pattern. This uses
a cache (in the case redis) to as a first stop in getting data from a database.
This builds on the traditional relational database pattern of retrieving data
from the database with sql queries and adds a new layer of cache between the
application and the database

### Cache pattern design

backendapp (golang) <--> cache <--> postgres


* Case 1 - No data in cache
The application checks the cache if it has a sql query result. Since a cache may
be accessed using a [key/value](https://redis.io/topics/data-types-intro) pair,
the cache can be checked using a unique key (the key can be the entire sql query
or a string comprising of the table name and its arguments). If the key does not
exist in the cache, then the Sql query is run against the Postgres db and the
data is retrieved per the normal pattern. The data is then stored into the cache
for next time the query must be run using a unique key as discussed and the data
as the corresponding value to the key. The cache may be set with a timer that
[auto expires](https://redislabs.com/ebook/part-2-core-concepts/chapter-3-commands-in-redis/3-7-other-commands/3-7-3-expiring-keys/) the data after a certain amount of time. Depending on your
application this may be appropriate or not. If there are lots of reads with
infrequent writes then a longer auto-expire time may be appropriate (such as 4
or 12 hours for example). Lots of writes may deserve a shorter time or may
justify not using a cache at all.

1. backendapp <- postgres
2. backendapp -> cache


* Case 2 -
If the database query result has been previously stored in the cache, then the
data is retrieved (much faster than a traditional db query since the database
access and compares values on disk and the cache does so from RAM). Then the db
query does not occur. Caches can be easier to horizontally scale than relational
databases and this strategy can be used to alleviate the burden on a single
relational database for a monolith for example.

1. backendapp <- cache


### Performance

Cache, typically being in Ram and consisting of a simple key/value query is
usually an order of [magnitude faster](https://redis.io/topics/benchmarks) than
a database query. Your mileage may vary depending on your stack. At the risk of
offending the internet I'll share antidotal evidence. My observations (network 
trip included) on a recent application had cache results being returned in
approx 5-20 miliseconds and Sql database queries typically between 100
miliseconds and 2000 miliseconds. However, it is worth mentioning that many of
these queries were quite complex so again you will have to benchmark your own
app to be certain.

## Requirements

Run the following commands to set the environment variables required to run the
service

```
export POSTGRES_HOST=localhost
export POSTGRES_DB=database-name
export POSTGRES_USER=database-user-name
export POSTGRES_PASSWORD=super-secret-password
export CONNECTION_TIMEOUT=duration before connection timeout when starting
 connection to db
```

### Optional

Port is optional and may be set as follows

```
export POSTGRES_PORT=database-port
```

### Redis

Redis can be run with the following:

```
docker run --name some-redis -p 6379:6379 -d redis
```

Custom Redis ports can be setup with:

```
export REDIS_PORT=redis-port
```

## Run the code

```
go build && ./twotiercache
```


### References

[db caching layer](https://www.reddit.com/r/golang/comments/6o8rzt/how_do_you_organize_db_and_caching_layer/)
