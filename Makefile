.PHONY: docker-up docker-down migrate

docker-up:
	docker-compose -f docker/docker-compose.yml up -d

docker-down:
	docker-compose -f docker/docker-compose.yml down

migrate:
	cat migrations/create_tasks_table.sql | docker-compose -f docker/docker-compose.yml exec -T db psql -U postgres -d todo
