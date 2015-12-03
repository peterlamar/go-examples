# cassandra-demo

https://hub.docker.com/_/cassandra/

https://github.com/gocql/gocql

Stand up Single Cassandra (cluster is a few more calls and is explained on the docker hub page)

```
docker run --name some-cassandra -d -p 9042:9042 cassandra:latest
```

Connect to Cassandra with cqlsh

```
docker run -it --link some-cassandra:cassandra --rm cassandra sh -c 'exec cqlsh "$CASSANDRA_PORT_9042_TCP_ADDR"'
```

Seed Cassandra for the demo

```
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
```

Exit the cqlsh mode (exit)

Get dependencies

```
./goget.sh
```

Build and Run  Code

```
go build hellocassandra.go
./hellocassandra
```                      

Expected output

```
Tweet: 1192b59f-6d7e-11e5-899b-fa163e0b3aef hello world
Tweet: 1192b59f-6d7e-11e5-899b-fa163e0b3aef hello world
```
