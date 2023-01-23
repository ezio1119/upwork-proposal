# upwork-proposal

This is example in kafka consumer and producer in golang.

## Usage

### Launch kafka and zookeeper

Run kafka and zookeeper using compose.

```bash
docker compose up -d
```

### Create topic

Create a topic named example-topic using kafka-topics.sh in the kafka container.

```bash
make create-topic
```

### Run Consumer

Run the built image of /cmd/consumer/consumer.go on the compose network.
Wait for a message from the producer.

```bash
make run-consumer
```

### Run Producer

Run the built image of /cmd/consumer/producer.go on the compose network.
Text is retrieved from `/cmd/producer/text.csv` and pushed to a kafka topic.

```bash
make run-producer
```

### See the output of Consumer

You should receive the same text as in `/cmd/producer/text.csv`.

```bash
message at offset 0: 0 = sametext0
message at offset 1: 1 = sametext1
message at offset 2: 2 = sametext2
message at offset 3: 3 = sametext3
```

### tree

```bash
$ tree
.
├── cmd
│   ├── consumer
│   │   └── consumer.go
│   └── producer
│       ├── producer.go
│       └── text.csv
├── compose.yml
├── config
│   └── config.go
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
