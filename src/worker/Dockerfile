FROM golang:1.21.4 as builder
LABEL stage=builder

WORKDIR /app

COPY ./worker/go.mod ./worker/go.sum ./worker/
COPY ./common ./common
RUN cd worker; go mod download

COPY ./worker/src/ ./worker/src/

RUN cd worker/src; CGO_ENABLED=0 GOOS=linux go build -o /app/worker/worker-gs

FROM alpine

WORKDIR /app

COPY --from=builder /app /app

CMD cd /app/worker/; ./worker-gs