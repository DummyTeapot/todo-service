### Микросервис управления задачами

Порядок действий для Linux / Mac с установленным Docker Compose:

1. ```make docker-up```
2. ```make migrate```
3. ```make build && make run```

Для проверки HTTP через curl:

4. Создание задачи: ```curl -X POST http://localhost:8080/tasks \ -H "Content-Type: application/json" \ -d '{ "title": "Новая задача", "description": "Описание новой задачи" }'```
5. Получение списка задач: ```curl http://localhost:8080/tasks```
6. Получение задачи по ID: ```curl http://localhost:8080/tasks/1```
7. Обновление задачи: ```curl -X PUT http://localhost:8080/tasks/1 \ -H "Content-Type: application/json" \ -d '{ "title": "Обновленная задача", "description": "Обновленное описание", "completed": true }'```
8. Удаление задачи: ```curl -X DELETE http://localhost:8080/tasks/1```

Для проверки gRPC:

9. Получение списка задач: ```grpcurl -plaintext localhost:50051 todo.TaskService/GetTasks```
10. Обновление статуса задачи ```grpcurl -plaintext -d '{"id":1, "completed":true}' localhost:50051 todo.TaskService/UpdateTaskStatus```
