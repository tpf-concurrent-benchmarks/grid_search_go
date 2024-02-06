## Measurements

The system was run on the designated server, using the [Griewank function](https://www.sfu.ca/~ssurjano/griewank.html), with 4,8 and 16 nodes; with the following parameters:

```json
{
  "data": [
    [-600, 600, 0.2, 5],
    [-600, 600, 0.2, 5],
    [-600, 600, 0.2, 5]
  ],
  "agg": "MIN",
  "maxItemsPerBatch": 10800000
}
```

### Average Summary

#### FaMAF Server

| Measurement         | 4 Nodes             | 8 Nodes             | 16 Nodes            |
|---------------------|---------------------|---------------------|---------------------|
| Worker Throughput   | 2.07 Results/Second | 1.94 Results/Second | 1.92 Results/Second |
| Combined Throughput | 8.25 Results/Second | 14.5 Results/Second | 30.2 Results/Second |
| Work-time Variation | 0.686%              | 0.466%              | 0.0840%             |
| Memory Usage        | 6.8-8.4 MB/Worker   | 4.1-8.7 MB/Worker   | 4.0-6.2 MB/Worker   |
| Network Usage (Tx)  | 637 B/(s * Worker)  | 598 B/(s * Worker)  | 578 B/(s * Worker)  |
| Network Usage (Rx)  | 150 B/(s * Worker)  | 130 B/(s * Worker)  | 126 B/(s * Worker)  |
| CPU Usage           | 100%/Worker         | 100%/Worker         | 100%/Worker         |
| Completion Time     | 49.9 Minutes        | 27.0 Minutes        | 13.4 Minutes        |

#### Cloud (GCP)

| Measurement         | 4 Nodes             | 8 Nodes             | 16 Nodes            |
|---------------------|---------------------|---------------------|---------------------|
| Worker Throughput   | 1.52 Results/Second | 1.54 Results/Second | 1.51 Results/Second |
| Combined Throughput | 5.91 Results/Second | 11.7 Results/Second | 23.5 Results/Second |
| Work-time Variation | 0.275%              | 5.21%               | 0.644%              |
| Memory Usage        | 2.4-4.8 MB/Worker   | 1.8-4.4 MB/Worker   | 1.4-2.8 MB/Worker   |
| Network Usage (Tx)  | 462 B/(s * Worker)  | 490 B/(s * Worker)  | 480 B/(s * Worker)  |
| Network Usage (Rx)  | 102 B/(s * Worker)  | 104 B/(s * Worker)  | 100 B/(s * Worker)  |
| CPU Usage           | 100%/Worker         | 100%/Worker         | 100%/Worker         |
| Completion Time     | 67.2 Minutes        | 34.2 Minutes        | 17.2 Minutes        |
