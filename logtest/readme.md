```bash
go test -v -run=Benchmark -bench=Benchmark -benchmem -benchtime 10s -cpu=4,8 -parallel=8 ./
```

```text
goarch: amd64
pkg: test/logtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkZap
BenchmarkZap-4                   1822892              6876 ns/op             336 B/op          6 allocs/op
BenchmarkZap-8                   1730490              7037 ns/op             336 B/op          6 allocs/op
BenchmarkLogger
BenchmarkLogger-4                1732777              6461 ns/op             296 B/op          3 allocs/op
BenchmarkLogger-8                1758446              6419 ns/op             296 B/op          3 allocs/op
BenchmarkLoggerNoFORMAT
BenchmarkLoggerNoFORMAT-4        2670556              4340 ns/op             112 B/op          1 allocs/op
BenchmarkLoggerNoFORMAT-8        2694154              4192 ns/op             112 B/op          1 allocs/op
BenchmarkLoggerWrite
BenchmarkLoggerWrite-4           2949058              4087 ns/op             112 B/op          1 allocs/op
BenchmarkLoggerWrite-8           2843649              4093 ns/op             112 B/op          1 allocs/op
BenchmarkNativeGoLog
BenchmarkNativeGoLog-4           2162052              5551 ns/op             296 B/op          3 allocs/op
BenchmarkNativeGoLog-8           2139168              5715 ns/op             296 B/op          3 allocs/op
```

```text
goarch: amd64
pkg: test/logtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkParallelZap
BenchmarkParallelZap-4                   1000000             10572 ns/op             336 B/op          6 allocs/op
BenchmarkParallelZap-8                   1000000             10414 ns/op             337 B/op          6 allocs/op
BenchmarkParallelLogger
BenchmarkParallelLogger-4                1330300              8803 ns/op             296 B/op          3 allocs/op
BenchmarkParallelLogger-8                1363034              8945 ns/op             296 B/op          3 allocs/op
BenchmarkParallelLoggerNoFORMAT
BenchmarkParallelLoggerNoFORMAT-4        2053911              7076 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerNoFORMAT-8        1677360              6888 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite
BenchmarkParallelLoggerWrite-4           1939933              6304 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite-8           1922352              6938 ns/op             112 B/op          1 allocs/op
BenchmarkParallelNativeGoLog
BenchmarkParallelNativeGoLog-4           1204039              9612 ns/op             296 B/op          3 allocs/op
BenchmarkParallelNativeGoLog-8           1362807              8875 ns/op             296 B/op          3 allocs/op
```

```text
goarch: amd64
pkg: test/logtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkParallelZap
BenchmarkParallelZap-4                   1000000             10331 ns/op             336 B/op          6 allocs/op
BenchmarkParallelZap-8                   1000000             10595 ns/op             337 B/op          6 allocs/op
BenchmarkParallelLogger
BenchmarkParallelLogger-4                1352834              8838 ns/op             296 B/op          3 allocs/op
BenchmarkParallelLogger-8                1411458              8754 ns/op             296 B/op          3 allocs/op
BenchmarkParallelLoggerNoFORMAT
BenchmarkParallelLoggerNoFORMAT-4        2266597              5331 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerNoFORMAT-8        2090455              5631 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite
BenchmarkParallelLoggerWrite-4           2062870              5746 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite-8           2037792              5963 ns/op             112 B/op          1 allocs/op
BenchmarkParallelNativeGoLog
BenchmarkParallelNativeGoLog-4           1260445              9398 ns/op             280 B/op          3 allocs/op
BenchmarkParallelNativeGoLog-8           1272560              9123 ns/op             280 B/op          3 allocs/op
```

```text
goarch: amd64
pkg: test/logtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkParallelZap
BenchmarkParallelZap-4                   1000000             10230 ns/op             336 B/op          6 allocs/op
BenchmarkParallelZap-8                   1000000             10276 ns/op             337 B/op          6 allocs/op
BenchmarkParallelLogger
BenchmarkParallelLogger-4                1332555              8774 ns/op             296 B/op          3 allocs/op
BenchmarkParallelLogger-8                1391256              9226 ns/op             296 B/op          3 allocs/op
BenchmarkParallelLoggerNoFORMAT
BenchmarkParallelLoggerNoFORMAT-4        2154008              5483 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerNoFORMAT-8        2115795              5853 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite
BenchmarkParallelLoggerWrite-4           2059722              6069 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite-8           1968092              6116 ns/op             112 B/op          1 allocs/op
BenchmarkParallelNativeGoLog
BenchmarkParallelNativeGoLog-4           1249767              9930 ns/op             280 B/op          3 allocs/op
BenchmarkParallelNativeGoLog-8           1211719              9822 ns/op             280 B/op          3 allocs/op
```

#### 查询日志内容
```bash
tail -1 /d/cfoldTest/zaplog.txt && tail -1 /d/cfoldTest/golog.txt && tail -1 /d/cfoldTest/gologFormat.txt && tail -1 /d/cfoldTest/gologWrite.txt && tail -1 /d/cfoldTest/nativelog.txt
```

```text
2023-07-10T19:58:15.138+0800    DEBUG   logtest/benchmark_test.go:82    >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2023/07/10 19:58:57 benchmark_test.go:94: >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2023/06/10 01:25:55.028277 log_test.go:46:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2023/06/10 01:25:55.028277 log_test.go:55:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[debug]2023/07/10 20:00:53.634554 benchmark_test.go:125: >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
```


```text
goarch: amd64
pkg: test/logtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkSerialZap
BenchmarkSerialZap-4                      234355              5750 ns/op             336 B/op          6 allocs/op
BenchmarkSerialZap-8                      202875              5732 ns/op             336 B/op          6 allocs/op

BenchmarkSerialLogger
BenchmarkSerialLogger-4                   252968              4840 ns/op              64 B/op          1 allocs/op
BenchmarkSerialLogger-8                   216819              4787 ns/op              64 B/op          1 allocs/op

BenchmarkSerialLoggerNoFORMAT
BenchmarkSerialLoggerNoFORMAT-4           302656              3896 ns/op             112 B/op          1 allocs/op
BenchmarkSerialLoggerNoFORMAT-8           312416              3521 ns/op             112 B/op          1 allocs/op

BenchmarkSerialLoggerWrite
BenchmarkSerialLoggerWrite-4              291919              3519 ns/op             112 B/op          1 allocs/op
BenchmarkSerialLoggerWrite-8              384574              3528 ns/op             112 B/op          1 allocs/op

BenchmarkSerialNativeGoLog
BenchmarkSerialNativeGoLog-4              277941              4463 ns/op             232 B/op          2 allocs/op
BenchmarkSerialNativeGoLog-8              260077              4735 ns/op             232 B/op          2 allocs/op

BenchmarkSerialSlog
BenchmarkSerialSlog-4                     196550              5663 ns/op             328 B/op          6 allocs/op
BenchmarkSerialSlog-8                     212346              5724 ns/op             328 B/op          6 allocs/op

BenchmarkParallelZap
BenchmarkParallelZap-4                    144482              8194 ns/op             336 B/op          6 allocs/op
BenchmarkParallelZap-8                    132722              8217 ns/op             337 B/op          6 allocs/op

BenchmarkParallelLogger
BenchmarkParallelLogger-4                 203913              5590 ns/op              64 B/op          1 allocs/op
BenchmarkParallelLogger-8                 223483              5462 ns/op              64 B/op          1 allocs/op

BenchmarkParallelLoggerNoFORMAT
BenchmarkParallelLoggerNoFORMAT-4         297376              4071 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerNoFORMAT-8         304020              4128 ns/op             112 B/op          1 allocs/op

BenchmarkParallelLoggerWrite
BenchmarkParallelLoggerWrite-4            265011              3900 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite-8            323252              3887 ns/op             112 B/op          1 allocs/op

BenchmarkParallelNativeGoLog
BenchmarkParallelNativeGoLog-4            190694              5370 ns/op             232 B/op          2 allocs/op
BenchmarkParallelNativeGoLog-8            195561              5403 ns/op             232 B/op          2 allocs/op

BenchmarkParallelSLog
BenchmarkParallelSLog-4                   231270              5447 ns/op             328 B/op          6 allocs/op
BenchmarkParallelSLog-8                   194901              5392 ns/op             328 B/op          6 allocs/op
```

```bash
tail -1 /d/cfoldTest/zaplog.txt && tail -1 /d/cfoldTest/golog.txt && tail -1 /d/cfoldTest/gologFormat.txt && tail -1 /d/cfoldTest/gologWrite.txt && tail -1 /d/cfoldTest/nativelog.txt  && tail -1 /d/cfoldTest/slog.txt
```

```text
2023-09-30T17:37:15.008+0800    DEBUG   logtest/benchmark_test.go:103   >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2023/09/30 17:37:19 benchmark_test.go:115: >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2023/06/10 01:25:55.028277 log_test.go:46:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2023/06/10 01:25:55.028277 log_test.go:55:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[debug]2023/09/30 17:37:33.562471 benchmark_test.go:146: >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
time=2023-09-30T17:37:38.395+08:00 level=INFO source=benchmark_test.go:164 msg=>>>aaaaaaaaaaaaaaaaaaaaaaa
```

------

```text
goarch: amd64
pkg: test/logtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkSerialZap
BenchmarkSerialZap-4                     1356222              5284 ns/op             336 B/op          6 allocs/op
BenchmarkSerialZap-8                     1340780              5249 ns/op             337 B/op          6 allocs/op
BenchmarkSerialLogger
BenchmarkSerialLogger-4                  1672040              4244 ns/op              64 B/op          1 allocs/op
BenchmarkSerialLogger-8                  1681812              4226 ns/op              64 B/op          1 allocs/op
BenchmarkSerialLoggerNoFORMAT
BenchmarkSerialLoggerNoFORMAT-4          1976608              3609 ns/op             112 B/op          1 allocs/op
BenchmarkSerialLoggerNoFORMAT-8          2003584              3579 ns/op             112 B/op          1 allocs/op
BenchmarkSerialLoggerWrite
BenchmarkSerialLoggerWrite-4             2084395              3513 ns/op             112 B/op          1 allocs/op
BenchmarkSerialLoggerWrite-8             2087563              3464 ns/op             112 B/op          1 allocs/op
BenchmarkSerialNativeGoLog
BenchmarkSerialNativeGoLog-4             1650951              4372 ns/op             232 B/op          2 allocs/op
BenchmarkSerialNativeGoLog-8             1649140              4384 ns/op             232 B/op          2 allocs/op
BenchmarkSerialSlog
BenchmarkSerialSlog-4                    1335410              5432 ns/op             328 B/op          6 allocs/op
BenchmarkSerialSlog-8                    1348922              5492 ns/op             328 B/op          6 allocs/op
BenchmarkSerialSlogAndLogger
BenchmarkSerialSlogAndLogger-4           1305316              5483 ns/op             328 B/op          6 allocs/op
BenchmarkSerialSlogAndLogger-8           1278594              5499 ns/op             328 B/op          6 allocs/op
```

```text
goarch: amd64
pkg: test/logtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkParallelZap
BenchmarkParallelZap-4                    789764              7392 ns/op             336 B/op          6 allocs/op
BenchmarkParallelZap-8                    795351              7172 ns/op             337 B/op          6 allocs/op
BenchmarkParallelZap-16                   832374              6927 ns/op             338 B/op          6 allocs/op
BenchmarkParallelLogger
BenchmarkParallelLogger-4                1000000              5531 ns/op              64 B/op          1 allocs/op
BenchmarkParallelLogger-8                1000000              5697 ns/op              64 B/op          1 allocs/op
BenchmarkParallelLogger-16                949768              5973 ns/op              64 B/op          1 allocs/op
BenchmarkParallelLoggerNoFORMAT
BenchmarkParallelLoggerNoFORMAT-4        1469342              4149 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerNoFORMAT-8        1301102              4615 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerNoFORMAT-16       1303020              4543 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite
BenchmarkParallelLoggerWrite-4           1388493              4365 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite-8           1349751              4429 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite-16          1362633              4403 ns/op             112 B/op          1 allocs/op
BenchmarkParallelNativeGoLog
BenchmarkParallelNativeGoLog-4            966140              5659 ns/op             232 B/op          2 allocs/op
BenchmarkParallelNativeGoLog-8           1000000              5578 ns/op             232 B/op          2 allocs/op
BenchmarkParallelNativeGoLog-16          1000000              5582 ns/op             232 B/op          2 allocs/op
BenchmarkParallelSLog
BenchmarkParallelSLog-4                   996240              5481 ns/op             328 B/op          6 allocs/op
BenchmarkParallelSLog-8                  1000000              5455 ns/op             328 B/op          6 allocs/op
BenchmarkParallelSLog-16                 1000000              5453 ns/op             329 B/op          6 allocs/op
BenchmarkParallelSLogAndgoLogger
BenchmarkParallelSLogAndgoLogger-4        965965              5501 ns/op             328 B/op          6 allocs/op
BenchmarkParallelSLogAndgoLogger-8       1000000              5507 ns/op             328 B/op          6 allocs/op
BenchmarkParallelSLogAndgoLogger-16      1000000              5531 ns/op             329 B/op          6 allocs/op
```

```bash
tail -1 /d/cfoldTest/zaplog.txt && tail -1 /d/cfoldTest/golog.txt && tail -1 /d/cfoldTest/gologFormat.txt && tail -1 /d/cfoldTest/gologWrite.txt && tail -1 /d/cfoldTest/nativelog.txt  && tail -1 /d/cfoldTest/slog.txt && tail -1 /d/cfoldTest/slogLogger.txt
```

```text
2024-05-13T16:41:27.696+0800    DEBUG   logtest/benchmark_test.go:121   >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2024/05/13 16:41:32 benchmark.go:797: >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2023/06/10 01:25:55.028277 log_test.go:46:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[DEBUG]2023/06/10 01:25:55.028277 log_test.go:55:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
[debug]2024/05/13 16:41:46.839928 benchmark_test.go:164: >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
time=2024-05-13T16:41:51.637+08:00 level=INFO source=benchmark_test.go:182 msg=>>>aaaaaaaaaaaaaaaaaaaaaa
time=2024-05-13T16:41:56.255+08:00 level=INFO source=benchmark_test.go:200 msg=>>>aaaaaaaaaaaaaaaaaaaaaa
```

------

```text
pkg: test/logtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkSerialZap
BenchmarkSerialZap-4                      714796              5469 ns/op             336 B/op          6 allocs/op
BenchmarkSerialZap-8                      675508              5316 ns/op             337 B/op          6 allocs/op
BenchmarkSerialLogger
BenchmarkSerialLogger-4                   749774              4458 ns/op             152 B/op          4 allocs/op
BenchmarkSerialLogger-8                   793208              4321 ns/op             152 B/op          4 allocs/op
BenchmarkSerialLoggerNoFORMAT
BenchmarkSerialLoggerNoFORMAT-4           977128              3767 ns/op             128 B/op          2 allocs/op
BenchmarkSerialLoggerNoFORMAT-8          1000000              3669 ns/op             128 B/op          2 allocs/op
BenchmarkSerialLoggerWrite
BenchmarkSerialLoggerWrite-4              856617              3659 ns/op             112 B/op          1 allocs/op
BenchmarkSerialLoggerWrite-8             1000000              3576 ns/op             112 B/op          1 allocs/op
BenchmarkSerialNativeGoLog
BenchmarkSerialNativeGoLog-4              892172              4488 ns/op             232 B/op          2 allocs/op
BenchmarkSerialNativeGoLog-8              798291              4327 ns/op             232 B/op          2 allocs/op
BenchmarkSerialSlog
BenchmarkSerialSlog-4                     634228              5602 ns/op             328 B/op          6 allocs/op
BenchmarkSerialSlog-8                     646191              5481 ns/op             328 B/op          6 allocs/op
BenchmarkSerialSlogAndLogger
BenchmarkSerialSlogAndLogger-4            626898              5671 ns/op             328 B/op          6 allocs/op
BenchmarkSerialSlogAndLogger-8            657820              5622 ns/op             328 B/op          6 allocs/op
BenchmarkParallelZap
BenchmarkParallelZap-4                    430472              7818 ns/op             336 B/op          6 allocs/op
BenchmarkParallelZap-8                    449402              7771 ns/op             337 B/op          6 allocs/op
BenchmarkParallelLogger
BenchmarkParallelLogger-4                 639826              5398 ns/op             152 B/op          4 allocs/op
BenchmarkParallelLogger-8                 604308              5532 ns/op             152 B/op          4 allocs/op
BenchmarkParallelLoggerNoFORMAT
BenchmarkParallelLoggerNoFORMAT-4         806749              4311 ns/op             128 B/op          2 allocs/op
BenchmarkParallelLoggerNoFORMAT-8         790284              4592 ns/op             128 B/op          2 allocs/op
BenchmarkParallelLoggerWrite
BenchmarkParallelLoggerWrite-4            764610              4141 ns/op             112 B/op          1 allocs/op
BenchmarkParallelLoggerWrite-8            880222              4079 ns/op             112 B/op          1 allocs/op
BenchmarkParallelNativeGoLog
BenchmarkParallelNativeGoLog-4            609134              5652 ns/op             232 B/op          2 allocs/op
BenchmarkParallelNativeGoLog-8            588201              5806 ns/op             232 B/op          2 allocs/op
BenchmarkParallelSLog
BenchmarkParallelSLog-4                   620878              5624 ns/op             328 B/op          6 allocs/op
BenchmarkParallelSLog-8                   636448              5532 ns/op             328 B/op          6 allocs/op
BenchmarkParallelSLogAndgoLogger
BenchmarkParallelSLogAndgoLogger-4        612314              5612 ns/op             328 B/op          6 allocs/op
BenchmarkParallelSLogAndgoLogger-8        633426              5596 ns/op             328 B/op          6 allocs/op
```

### 日志库和方法说明：
* Zap：这是一个uber开发的高性能日志库。
* Logger：go-logger日志。
* Native Go Log： Go 内置 log 包。
* Slog：这是 Go 1.19 引入的新标准日志库。
* Slog 和 Logger 结合：指同时使用go-logger作为slog的日志文件管理库。


##### 基准测试指标解释：
* -4 和 -8: 这些数字表示运行基准测试时使用的 CPU 核心数。-4 表示使用 4 个核心，而 -8 表示使用 8 个核心。
* ns/op: 每次日志记录操作所需的平均时间（以纳秒为单位）。
* B/op: 每次日志记录操作分配的平均内存大小（以字节为单位）。
* allocs/op: 每次日志记录操作产生的分配次数。

#### 串行日志记录结果：

* Zap: 在 4 核心上有 5469 ns/op 的性能，在 8 核心上有 5316 ns/op 的性能。
* go-logger: 在 4 核心上有 4458 ns/op 的性能，在 8 核心上有 4321 ns/op 的性能。
* go-logger(无格式): 在 4 核心上有 3767 ns/op 的性能，在 8 核心上有 3669 ns/op 的性能。
* go-logger(写操作): 在 4 核心上有 3659 ns/op 的性能，在 8 核心上有 3576 ns/op 的性能。
* Native Go Log: 在 4 核心上有 4488 ns/op 的性能，在 8 核心上有 4327 ns/op 的性能。
* Slog: 在 4 核心上有 5602 ns/op 的性能，在 8 核心上有 5481 ns/op 的性能。
* Slog 和 go-logger 结合: 在 4 核心上有 5671 ns/op 的性能，在 8 核心上有 5622 ns/op 的性能。

#### 并行日志记录结果：

* Zap: 在 4 核心上有 7818 ns/op 的性能，在 8 核心上有 7771 ns/op 的性能。
* go-logger: 在 4 核心上有 5398 ns/op 的性能，在 8 核心上有 5532 ns/op 的性能。
* go-logger (无格式): 在 4 核心上有 4311 ns/op 的性能，在 8 核心上有 4592 ns/op 的性能。
* go-logger (写操作): 在 4 核心上有 4141 ns/op 的性能，在 8 核心上有 4079 ns/op 的性能。
* Native Go Log: 在 4 核心上有 5652 ns/op 的性能，在 8 核心上有 5806 ns/op 的性能。
* Slog: 在 4 核心上有 5624 ns/op 的性能，在 8 核心上有 5532 ns/op 的性能。
* Slog 和go-logger 结合: 在 4 核心上有 5612 ns/op 的性能，在 8 核心上有 5596 ns/op 的性能。

#### 结果分析：

* Zap 在串行模式下提供了较好的性能，但在并行模式下的性能有所下降。
* go-logger(写操作) 在串行和并行模式下均表现出了最佳性能。
* go-logger(无格式) 通过移除格式化步骤显著提高了性能。
* Native Go Log 在串行和并行模式下性能接近于 go-logger。
* Slog 的性能与 Zap 和 go-logger 相比略逊一筹。
* Slog 和 go-logger 结合 的性能与 Slog 相近

#### 结论

* 从压测结果可以看到，在相同格式下，无论是串行还是高并发场景中，go-logger均表现出最佳性能和最小的内存分配。
* go-logger作为slog日志文件管理库，无论内存分配还是性能，都与单独使用slog的效果相同，没有额外的性能损耗。


