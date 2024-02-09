# Benchmarks

## Measurements

### FaMAF Server

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

### Cloud (GCP)

| Measurement         | 4 Nodes        | 8 Nodes        | 16 Nodes       |
|---------------------|----------------|----------------|----------------|
| Worker Throughput   | Results/Second | Results/Second | Results/Second |
| Combined Throughput | Results/Second | Results/Second | Results/Second |
| Work-time Variation | %              | %              | %              |
| Memory Usage        | MB/Worker      | MB/Worker      | MB/Worker      |
| Network Usage (Tx)  | B/(s * Worker) | B/(s * Worker) | B/(s * Worker) |
| Network Usage (Rx)  | B/(s * Worker) | B/(s * Worker) | B/(s * Worker) |
| CPU Usage           | %/Worker       | %/Worker       | %/Worker       |
| Completion Time     | Minutes        | Minutes        | Minutes        |

Average measurements using the [specified configuration](measurements/README.md)
