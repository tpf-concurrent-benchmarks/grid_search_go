# Grid Search in Go

## Objective

This is a Go implementation of a system for performing a grid search under [common specifications](https://github.com/tpf-concurrent-benchmarks/docs/tree/main/grid_search) defined for multiple languages.

The objective of this project is to benchmark the language on a real-world distributed system.

## Deployment

### Requirements

- [Docker >3](https://www.docker.com/) (needs docker swarm)
- [Go >1.21.4](https://golang.org/) (for running locally)

### Configuration

- **Number of replicas:** `N_WORKERS` constant is defined in the `Makefile` file.
- **Data config:** in `src/manager/resources/data.json` you can define (this config is built into the container):
  - `data`: Intervals for each parameter, in format: [start, end, step, precision]
  - `agg`: Aggregation function to be used: MIN | MAX | AVG
  - `maxItemsPerBatch`: Maximum number of items per batch (batches are sub-intervals)

### Commands

#### Startup

- `make init` starts docker swarm
- `make build` builds the manager and worker images

#### Run

- `make deploy` deploys the manager and worker services locally, alongside with Graphite, Grafana and cAdvisor.
> The first time you run this command, the manager or worker could time out due to Graphite not being ready. In that case,
> wait for Graphite to be ready and run `make deploy` again.

- `make remove` removes all services created by the `deploy` command.

### Monitoring

- Grafana: [http://127.0.0.1:8081](http://127.0.0.1:8081)
- Graphite: [http://127.0.0.1:8080](http://127.0.0.1:8080)

## Used libraries

- [Statsd client](https://github.com/cactus/go-statsd-client): used to send metrics to Graphite.
- [NATS client](https://github.com/nats-io/nats.go): used to communicate between manager and worker.