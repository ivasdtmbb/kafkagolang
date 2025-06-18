# kafkagolang
Kafka project for Golang

# Libraries for Kafka
github.com/confluentinc/confluent-kafka-go/v2/kafka
github.com/IBM/sarama
github.com/segmentio/kafka-go
github.com/lovoo/goka

# CLI Commands

sudo systemctl start docker
sudo systemctl enable docker
systemctl status docker
docker info
docker-compose --version

# rebuild images
docker-compose up -d --build

# -d detach mode (фоновый режим запуска), up запуск контейнера
docker compose up -d 

# stop the containers
docker-compose down

# docker compose main commands
docker-compose ps
Список работающих контейнеров

docker-compose logs
Логи всех сервисов

docker-compose logs <service>
Логи конкретного сервиса

docker-compose exec <service> sh
Выполнить команду внутри контейнера

docker-compose stop
Остановить все сервисы

docker-compose start
Запустить остановленные сервисы

docker-compose build
Собрать образы

# Подключение к Kafka UI
localhost:9020/