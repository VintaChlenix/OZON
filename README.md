# OZON Тестовое задание

Для проверки необходимо в терминале собрать и поднять образ Docker

```
docker-compose build
docker-compose up
```

В Dockerfile выставлен параметр запуска на PostgreSQL, но по умолчанию он - InMemory.

CURL запросы:
```
POST:
curl --request POST \
  --url http://localhost:8080/ \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data url=http://vk.com
  
GET:
curl --request GET \
  --url http://localhost:8080/bGwXqTTUYK \
  --header 'Content-Type: multipart/form-data; boundary=---011000010111000001101001'
  
```

Тестирование:

```
cd internal/handlers
go test
```
