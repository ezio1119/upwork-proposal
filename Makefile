TOPIC=example-topic
RUNTIME = docker

create-topic:
	${RUNTIME} compose exec kafka \
		kafka-topics.sh \
		--create \
		--topic $(TOPIC) \
		--bootstrap-server localhost:9092

build:
	${RUNTIME} build -t app .

run-consumer: build
	${RUNTIME} run --rm \
		--name consumer \
		--entrypoint /app/consumer \
		--net upwork-proposal_default \
		app

run-producer: build
	${RUNTIME} run --rm \
		--name producer \
		--entrypoint /app/producer \
		--net upwork-proposal_default \
		app