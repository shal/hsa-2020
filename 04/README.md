# Homework 04

## Benchmark

Use `docker-compose` for testing

```bash
docker-compose run siege -b -r1000 -c $CONC -f /tmp/urls.txt
```

### Without Cache

| Concurrency           | 10     | 25      | 50      | 100     |
|:---------------------:|:------:|:-------:|:-------:|:-------:|
| Resource availability | 100    | 100     | 100     | 100     |
| Avg response time     | 0.01   | 0.02    | 0.03    | 0.05    |
| Throughput            | 0.01   | 0.02    | 0.03    | 0.03    |
| Transaction rate      | 750.31 | 1319.79 | 1706.98 | 1864.92 |

### With Cache

### Basic

| Concurrency           | 10      | 25      | 50      | 100     |
|:---------------------:|:-------:|:-------:|:-------:|:-------:|
| Resource availability | 100     | 100     | 100     | 100     |
| Avg response time     | 0.02    | 0.03    | 0.04    | 0.04    |
| Throughput            | 0.03    | 0.03    | 0.04    | 0.06    |
| Transaction rate      | 1634.09 | 1799.32 | 1931.95 | 2243.12 |

### Probabilistic cache flushing

| Concurrency           | 10  | 25  | 50  | 100 |
|:---------------------:|:---:|:---:|:---:|:---:|
| Resource availability | 100 | 100 | 100 | 100 |
| Avg response time     |  0  |  0  |  0  |  0  |
| Throughput            |  0  |  0  |  0  |  0  |
| Transaction rate      |  0  |  0  |  0  |  0  |
