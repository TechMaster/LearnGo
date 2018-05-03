#Benchmark
#iris
wrk -t4 -c100 -d30s -R2000 http://localhost:8080/ 
Running 30s test @ http://localhost:8080/
  4 threads and 100 connections
  Thread calibration: mean lat.: 2.719ms, rate sampling interval: 11ms
  Thread calibration: mean lat.: 2.666ms, rate sampling interval: 11ms
  Thread calibration: mean lat.: 1.537ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.364ms, rate sampling interval: 10ms
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.53ms  652.54us   4.36ms   65.80%
    Req/Sec   534.56    647.17     2.50k    88.81%
  58968 requests in 30.00s, 6.86MB read
Requests/sec:   1965.39
Transfer/sec:    234.16KB

---
wrk -t4 -c200 -d30s -R2000 http://localhost:8080/          30s  ~
Running 30s test @ http://localhost:8080/
  4 threads and 200 connections
  Thread calibration: mean lat.: 1.446ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.460ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.476ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 3.780ms, rate sampling interval: 10ms
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.49ms  663.92us   4.87ms   67.89%
    Req/Sec   532.13    757.39     5.55k    97.28%
  59390 requests in 30.00s, 6.91MB read
Requests/sec:   1979.44
Transfer/sec:    235.83KB
---------------------
# fasthttp
wrk -t4 -c100 -d30s -R2000 http://localhost:8081/          30s  ~
Running 30s test @ http://localhost:8081/
  4 threads and 100 connections
  Thread calibration: mean lat.: 1.301ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.339ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.234ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 2.948ms, rate sampling interval: 10ms
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.30ms  599.40us   4.10ms   66.29%
    Req/Sec   534.87    510.24     2.78k    76.58%
  59447 requests in 30.01s, 7.88MB read
Requests/sec:   1981.21
Transfer/sec:    268.93KB

wrk -t4 -c200 -d30s -R2000 http://localhost:8081/          30s  ~
Running 30s test @ http://localhost:8081/
  4 threads and 200 connections
  Thread calibration: mean lat.: 1.425ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.427ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.426ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.783ms, rate sampling interval: 10ms
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.34ms  570.73us   3.81ms   65.73%
    Req/Sec   531.06    716.39     5.55k    96.97%
  59386 requests in 30.00s, 40.10MB read
Requests/sec:   1979.41
Transfer/sec:      1.34MB