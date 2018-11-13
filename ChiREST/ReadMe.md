# Kiểm thử tốc độ trả về của các framework: go-micro, iris, chi, gin

```
hey -H 'Content-Type: application/json' -d '{}' http://localhost:8080/blog/getAllArticle
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/blog/getAllArticle
hey -n 5000 -c 120 http://localhost:8000/articles
hey -n 5000 -c 120 -H 'Content-Type: application/json' -d '{}' http://localhost:8080/blog/getAllArticle

k6 run --vus 200 --duration 10s benchrest.js
k6 run --vus 200 --duration 10s benchmicro.js
```

# Kiểm thử bằng K6 - Grafana - InfluxDB
```
cd k6
docker-compose up -d influxdb grafana
k6 run --vus 200 --duration 10s --out influxdb=http://localhost:8086/k6db benchrest.js

```