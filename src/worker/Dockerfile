FROM golang:1.21.4

WORKDIR /app

COPY ./worker/go.mod ./worker/go.sum ./worker/
COPY ./common ./common
RUN cd worker; go mod download

COPY ./worker/src/ ./worker/src/

RUN cd worker/src; CGO_ENABLED=0 GOOS=linux go build -o /app/worker/worker-gs

CMD cd /app/worker/; ./worker-gs