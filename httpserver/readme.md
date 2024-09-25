
```bash
go test -v -run=Benchmark -bench=Benchmark -benchmem -cpu=4,8 -parallel=8 ./
```

### 找到URL的压测数据:

```text
pkg: test/httpserver
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkGin
BenchmarkGin-4           2356797               507.6 ns/op          1040 B/op          9 allocs/op
BenchmarkGin-8           2428094               500.5 ns/op          1040 B/op          9 allocs/op
BenchmarkHttp
BenchmarkHttp-4          3405583               346.5 ns/op           344 B/op          9 allocs/op
BenchmarkHttp-8          3551180               330.7 ns/op           344 B/op          9 allocs/op
BenchmarkTlnet
BenchmarkTlnet-4         6224007               187.6 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-8         6570204               184.3 ns/op           288 B/op          5 allocs/op
BenchmarkEcho
BenchmarkEcho-4          2333074               535.3 ns/op          1024 B/op         10 allocs/op
BenchmarkEcho-8          2366791               528.6 ns/op          1024 B/op         10 allocs/op
PASS
```

```text
goarch: amd64
pkg: test/httpserver
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkGin
BenchmarkGin-4           2290322               494.4 ns/op          1040 B/op          9 allocs/op
BenchmarkGin-8           2491639               487.8 ns/op          1040 B/op          9 allocs/op
BenchmarkGin-16          2560693               493.1 ns/op          1041 B/op          9 allocs/op
BenchmarkHttp
BenchmarkHttp-4          3565180               323.4 ns/op           344 B/op          9 allocs/op
BenchmarkHttp-8          3544458               339.9 ns/op           344 B/op          9 allocs/op
BenchmarkHttp-16         3596947               307.0 ns/op           344 B/op          9 allocs/op
BenchmarkTlnet
BenchmarkTlnet-4         5572468               189.7 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-8         6420810               189.5 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-16        5862798               197.4 ns/op           288 B/op          5 allocs/op
BenchmarkEcho
BenchmarkEcho-4          2356166               504.5 ns/op          1024 B/op         10 allocs/op
BenchmarkEcho-8          2238975               540.4 ns/op          1024 B/op         10 allocs/op
BenchmarkEcho-16         1740646               639.0 ns/op          1024 B/op         10 allocs/op
PASS
```

```text
goarch: amd64
pkg: test/httpserver
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkGin
BenchmarkGin-4           2445027               484.4 ns/op          1040 B/op          9 allocs/op
BenchmarkGin-8           2465703               489.2 ns/op          1040 B/op          9 allocs/op
BenchmarkGin-16          2567120               462.3 ns/op          1040 B/op          9 allocs/op
BenchmarkHttp
BenchmarkHttp-4          3700851               332.5 ns/op           344 B/op          9 allocs/op
BenchmarkHttp-8          3507562               385.1 ns/op           344 B/op          9 allocs/op
BenchmarkHttp-16         3458373               318.5 ns/op           344 B/op          9 allocs/op
BenchmarkTlnet
BenchmarkTlnet-4         7283329               166.9 ns/op           288 B/op          6 allocs/op
BenchmarkTlnet-8         7045941               163.3 ns/op           288 B/op          6 allocs/op
BenchmarkTlnet-16        7091187               162.7 ns/op           288 B/op          6 allocs/op
BenchmarkEcho
BenchmarkEcho-4          2395107               501.3 ns/op          1024 B/op         10 allocs/op
BenchmarkEcho-8          2230350               541.2 ns/op          1024 B/op         10 allocs/op
BenchmarkEcho-16         2118523               571.2 ns/op          1025 B/op         10 allocs/op
```
------

### 找不到URL的压测数据:

```text
goarch: amd64
pkg: test/httpserver
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkGin
BenchmarkGin-4           2695962               436.6 ns/op           992 B/op          8 allocs/op
BenchmarkGin-8           2682855               448.6 ns/op           992 B/op          8 allocs/op
BenchmarkGin-16          2692042               456.7 ns/op           993 B/op          8 allocs/op
BenchmarkHttp
BenchmarkHttp-4          1462479               806.5 ns/op          1200 B/op         19 allocs/op
BenchmarkHttp-8          1378512               868.7 ns/op          1200 B/op         19 allocs/op
BenchmarkHttp-16         1359558               898.3 ns/op          1201 B/op         19 allocs/op
BenchmarkTlnet
BenchmarkTlnet-4         6343191               182.1 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-8         6371780               193.3 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-16        6292598               182.6 ns/op           288 B/op          5 allocs/op
BenchmarkEcho
BenchmarkEcho-4          1387812               854.2 ns/op          1489 B/op         16 allocs/op
BenchmarkEcho-8          1000000              1031 ns/op            1489 B/op         16 allocs/op
BenchmarkEcho-16          871554              1220 ns/op            1491 B/op         16 allocs/op
```

```text
goarch: amd64
pkg: test/httpserver
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkGin
BenchmarkGin-4           2695962               436.6 ns/op           992 B/op          8 allocs/op
BenchmarkGin-8           2682855               448.6 ns/op           992 B/op          8 allocs/op
BenchmarkGin-16          2692042               456.7 ns/op           993 B/op          8 allocs/op
BenchmarkHttp
BenchmarkHttp-4          1462479               806.5 ns/op           1200 B/op         19 allocs/op
BenchmarkHttp-8          1378512               868.7 ns/op           1200 B/op         19 allocs/op
BenchmarkHttp-16         1359558               898.3 ns/op           1201 B/op         19 allocs/op
BenchmarkTlnet
BenchmarkTlnet-4         6343191               182.1 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-8         6371780               193.3 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-16        6292598               182.6 ns/op           288 B/op          5 allocs/op
BenchmarkEcho
BenchmarkEcho-4          1387812               854.2 ns/op           1489 B/op         16 allocs/op
BenchmarkEcho-8          1000000               1031 ns/op            1489 B/op         16 allocs/op
BenchmarkEcho-16          871554               1220 ns/op            1491 B/op         16 allocs/op
```

```text
goarch: amd64
pkg: test/httpserver
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkGin
BenchmarkGin-4           2758132               436.3 ns/op           992 B/op          8 allocs/op
BenchmarkGin-8           2667070               447.5 ns/op           992 B/op          8 allocs/op
BenchmarkGin-16          2778464               423.7 ns/op           992 B/op          8 allocs/op
BenchmarkHttp
BenchmarkHttp-4          1484661               809.1 ns/op          1200 B/op         19 allocs/op
BenchmarkHttp-8          1470441               831.8 ns/op          1200 B/op         19 allocs/op
BenchmarkHttp-16         1466318               841.5 ns/op          1201 B/op         19 allocs/op
BenchmarkTlnet
BenchmarkTlnet-4         6568359               178.0 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-8         6523093               185.5 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-16        6733190               172.9 ns/op           288 B/op          5 allocs/op
BenchmarkEcho
BenchmarkEcho-4          1412888               828.2 ns/op          1488 B/op         16 allocs/op
BenchmarkEcho-8          1404442               892.6 ns/op          1489 B/op         16 allocs/op
BenchmarkEcho-16         1306354               932.2 ns/op          1491 B/op         16 allocs/op
```

```text
goarch: amd64
pkg: test/httpserver
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkGin
BenchmarkGin-4           2715367               441.7 ns/op           992 B/op          8 allocs/op
BenchmarkGin-8           2552917               463.5 ns/op           992 B/op          8 allocs/op
BenchmarkGin-16          2607180               448.3 ns/op           993 B/op          8 allocs/op
BenchmarkHttp
BenchmarkHttp-4          1500954               808.5 ns/op          1200 B/op         19 allocs/op
BenchmarkHttp-8          1417261               861.2 ns/op          1200 B/op         19 allocs/op
BenchmarkHttp-16         1334204               899.6 ns/op          1201 B/op         19 allocs/op
BenchmarkTlnet
BenchmarkTlnet-4         6459314               182.8 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-8         5938984               195.5 ns/op           288 B/op          5 allocs/op
BenchmarkTlnet-16        6510511               191.6 ns/op           288 B/op          5 allocs/op
BenchmarkEcho
BenchmarkEcho-4          1414153               840.8 ns/op          1489 B/op         16 allocs/op
BenchmarkEcho-8          1344892               896.2 ns/op          1489 B/op         16 allocs/op
BenchmarkEcho-16         1244558              1139 ns/op            1490 B/op         16 allocs/op
```