# Distributed Calculator 🥭

Это сервис, который выполняет роль калькулятора. Данный сервис использует [Shunting yard algorithm](https://en.wikipedia.org/wiki/Shunting_yard_algorithm) для подсчета выражений.

## Инструкция по запуску:

### Запуск локально:

```bash
./runLocal.sh
```

### Запуск через докер:

```bash
docker-compose up -d
```

Логи сервера можно посмотреть в `Docker Desktop`
Команда для остановки сервиса:

```bash
docker compose down
```

## Пример curl запроса:

```bash
curl -X 'POST' -w "%{http_code}"\
'http://localhost:8080/api/v1/calculate' \
-H 'Content-Type: application/json' \
-d '{
  "expression": "2+2*2"
}'
```
