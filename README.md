## ApiLoadTesting

1. This code illustrates how to perform load testing on APIs using Go lang.
2. It uses this library: https://github.com/tsenart/vegeta
3. I used this to debug and tune up some nginx config, which was crashing under load 
   (although some times)
4. This should be simple to extend depending on the use case.
5. It also generates a very informative statistic report.
```
Sample Report:
--------------
Requests      [total, rate]            30000, 250.01
Duration      [total, attack, wait]    2m0.548248442s, 1m59.995999932s, 552.24851ms
Latencies     [mean, 50, 95, 99, max]  74.437595ms, 54.501095ms, 132.622449ms, 564.452081ms, 1.574880372s
Bytes In      [total, mean]            196785000, 6559.50
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:30000  
Error Set:
```
