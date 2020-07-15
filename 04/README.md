# Homework 04

## Benchmark

Use `docker-compose` for testing

```bash
docker-compose run siege -b -r1000 -c $CONC -f /tmp/urls.txt
```

### Without Cache

| Concurrency           | 10  | 25  | 50  | 100 |
|:---------------------:|:---:|:---:|:---:|:---:|
| Resource availability | 100 | 100 | 100 | 100 |
| Avg response time     |  0  |  0  |  0  |  0  |
| Throughput            |  0  |  0  |  0  |  0  |
| Transaction rate      |  0  |  0  |  0  |  0  |

### With Cache

| Concurrency           | 10  | 25  | 50  | 100 |
|:---------------------:|:---:|:---:|:---:|:---:|
| Resource availability | 100 | 100 | 100 | 100 |
| Avg response time     |  0  |  0  |  0  |  0  |
| Throughput            |  0  |  0  |  0  |  0  |
| Transaction rate      |  0  |  0  |  0  |  0  |

### Basic

| Concurrency           | 10  | 25  | 50  | 100 |
|:---------------------:|:---:|:---:|:---:|:---:|
| Resource availability | 100 | 100 | 100 | 100 |
| Avg response time     |  0  |  0  |  0  |  0  |
| Throughput            |  0  |  0  |  0  |  0  |
| Transaction rate      |  0  |  0  |  0  |  0  |

### Probabilistic cache flushing

| Concurrency           | 10  | 25  | 50  | 100 |
|:---------------------:|:---:|:---:|:---:|:---:|
| Resource availability | 100 | 100 | 100 | 100 |
| Avg response time     |  0  |  0  |  0  |  0  |
| Throughput            |  0  |  0  |  0  |  0  |
| Transaction rate      |  0  |  0  |  0  |  0  |
