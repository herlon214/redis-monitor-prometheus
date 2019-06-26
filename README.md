# Redis Monitor Prometheus
Executes the `redis-cli monitor`, parse the query and export as a prometheus metrics on port `8080`.

The metrics will look like this:
```
# HELP redis_processed_queries_total The total number of processed events
# TYPE redis_processed_queries_total counter
redis_processed_queries_total 13
# HELP redis_query_executed Total number that a query was executed
# TYPE redis_query_executed counter
redis_query_executed{query="hgetall my_executed_keys"} 13
```

### Environment

```
* PORT=8080 // Webserver port
* REDIS_URI=redis://your-redis-server-uri:6379;redis://your-redis-server2-uri:6379 // It's also supported multiple servers
* GRANULARITY=1 // Check the section about granularity below
```

#### Granularity
It's how many parts of the executed commands do you want. 

E.g: The command `sismember firstpart secondpart`

```
GRANULARITY = 0 = sismember firstpart secondpart // The whole command
GRANULARITY = 1 = sismember 
GRANULARITY = 2 = sismember firstpart
GRANULARITY = 3 = sismember firstpart secondpart
```


### How to run it with docker (easiest way)
The image is a `linux:alpine` with `redis-cli` and the build binary installed. It's only `16.8MB` :smile:
```
$ docker run -e PORT=8080 \
    -e REDIS_URI=redis://your-redis-server-uri:6379 \
    -p 8080:8080 \
    herlon214/redis-monitor-prometheus
```

You can check the `docker-compose.yml` file, it's a working example.

### How to run it without docker
Environment:
* PORT=8080
* REDIS_URI=redis://your-redis-server-uri:6379

You need the `redis-cli` installed on your machine.

```
$ go get -u github.com/herlon214/redis-monitor-prometheus
$ redis-monitor-prometheus
2019/06/24 22:47:20 [ Redis Monitor Prometheus ]
2019/06/24 22:47:20 -> Running watcher...
2019/06/24 22:47:20 -> Starting webserver...
```

