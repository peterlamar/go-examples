# Redis example

This example explores the two popular golang redis clients and compares their access methods. There were the two more popular from the redis goclients offical [list](https://redis.io/clients#go)

Redis clients

[redigo](https://github.com/gomodule/redigo) - Ultimately more powerful but
harder to get started. Passes through calls to redis with strings, using
```
c.Do("GET", key)
```
and
```
c.Do("SET", key)
```
for example

[go-redis](https://github.com/go-redis/redis) - Slightly less functionality (supports almost all the commands of redis, which was fine for my project that just used get/set) than
Redigo but more natural to use. Functions for

```
client.Get("key2")
```

and

```
client.Set(query, "value1", 0)
```

are more intuitive when starting out.

## Prerequisites

Get dependencies
```
go get ./...
```

Spin up a Redis container locally with Docker

```
docker run --name some-redis -p 6379:6379 -d redis
```

Then run the selected example

```
go run maingoredis.go
```

Or

```
go run maingoredigo.go
```
