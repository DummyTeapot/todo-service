.PHONY: docker-up docker-down migrate build run clean

build:
	go build -o bin/todo-service cmd/main.go

run:
	go run cmd/main.go

docker-up:
	docker-compose -f docker-compose.yml up -d

docker-down:
	docker-compose -f docker-compose.yml down

migrate:
	cat migrations/create_tasks_table.sql | docker-compose -f docker-compose.yml exec -T db psql -U postgres -d todo

clean:
	rm -rf bin && rm logs/app.log
