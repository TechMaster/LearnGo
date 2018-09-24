# Demo REST Microservice using Go-Chi

3rd parties libraries used:
1. [Go Chi](https://github.com/go-chi/chi)
2. [sJSON](https://github.com/tidwall/sjson)

```
hey -H 'Content-Type: application/json' -d '{}' http://localhost:8080/blog/getAllArticle
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/blog/getAllArticle
hey -n 5000 -c 120 http://localhost:8000/articles
hey -n 5000 -c 120 -H 'Content-Type: application/json' -d '{}' http://localhost:8080/blog/getAllArticle

k6 run --vus 200 --duration 10s benchchi.js
k6 run --vus 200 --duration 10s benchmicro.js
```

# Kiểm thử bằng K6 - Grafana - InfluxDB
```
cd k6
docker-compose up -d influxdb grafana
k6 run --vus 200 --duration 10s --out influxdb=http://localhost:8086/k6db benchchi.js

```