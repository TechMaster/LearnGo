Tham khảo bài viết ở đây [Lập trình Golang kết nối Kafja](https://techmaster.vn/posts/34654/lap-trinh-golang-ket-noi-kafka)

```
docker run --net=host -it --name gokafka -v /Volumes/CODE/GoProjects/src/github.com/TechMaster/LearnGo/Kafka:/foo -w /foo golang:alpine /bin/sh

docker run --net=host -it --name gokafka -v /Volumes/CODE/GoProjects/src/github.com/TechMaster/LearnGo/Kafka:/foo -w /foo gokafka:latest /bin/sh

docker-compose exec kafka  \
kafka-topics --create --topic bar --partitions 1 --replication-factor 1 --if-not-exists --zookeeper localhost:32181

docker exec kafka  \
kafka-topics --create --topic bar --partitions 1 --replication-factor 1 --if-not-exists --zookeeper localhost:32181

docker exec kafka  \
  kafka-topics --describe --topic foo --zookeeper localhost:32181

docker exec -it kafka  \
  kafka-console-producer --request-required-acks 1 --broker-list localhost:29092 --topic foo
docker exec kafka  \
  kafka-console-consumer --bootstrap-server localhost:29092 --topic foo --from-beginning --max-messages 1

docker exec kafka  \
  kafka-console-consumer --bootstrap-server localhost:29092 --topic foo --partition 0 --offset 4

docker run -it --rm --net=host \
confluentinc/cp-kafka:latest
```