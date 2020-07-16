# Homework 04

## Benchmark

Use `docker-compose` for testing

```bash
docker-compose run siege -b -t60S -c100 'http://app:8080/api/v1/04/transactions'
```

```bash

```

### Without Cache

| Concurrency           | 10     | 25      | 50      | 100     |
|:---------------------:|:------:|:-------:|:-------:|:-------:|
| Resource availability | 100    | 100     | 100     | 100     |
| Avg response time     | 0.45   | 0.59    | 0.73    | 0.89    |
| Throughput            | 0.00   | 0.00    | 0.00    | 0.01    |
| Transaction rate      | 22.10  | 42.10   | 67.42   | 110.76  |

### With Cache

### Basic

| Concurrency           | 10      | 25      | 50      | 100     |
|:---------------------:|:-------:|:-------:|:-------:|:-------:|
| Resource availability | 100     | 100     | 100     | 100     |
| Avg response time     | 0.03    | 0.04    | 0.07    | 0.08    |
| Throughput            | 0.02    | 0.03    | 0.04    | 0.12    |
| Transaction rate      | 365.00  | 579.51  | 733.33  | 1244.29 |

### Probabilistic cache flushing

| Concurrency           | 10      | 25      | 50      | 100    |
|:---------------------:|:-------:|:-------:|:-------:|:------:|
| Resource availability | 100     | 100     | 100     | 100    |
| Avg response time     | 0.03    | 0.05    | 0.07    | 0.11   |
| Throughput            | 0.01    | 0.02    | 0.03    | 0.04   |
| Transaction rate      | 356.00  | 525.17  | 755.14  | 897.68 |
