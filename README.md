# Redis Monitor Prometheus
Executes the `redis-cli monitor`, parse the query and export as a prometheus metrics on port `8080`.

The metrics will look like this:
```
# HELP monitor_processed_ops_total The total number of processed events
# TYPE monitor_processed_ops_total counter
monitor_processed_ops_total 13
# HELP monitor_query_executed Total number that a query was executed
# TYPE monitor_query_executed counter
monitor_query_executed{query="hgetall my_executed_keys"} 13
```
### How to run it with docker (easiest way)
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

