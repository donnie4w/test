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

--------

```text
Benchmark_Serial_NativeLog
Benchmark_Serial_NativeLog-4              570457              3956 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_NativeLog-8              610975              4044 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_Zap
Benchmark_Serial_Zap-4                    478735              4815 ns/op             352 B/op          6 allocs/op
Benchmark_Serial_Zap-8                    482280              4933 ns/op             353 B/op          6 allocs/op
Benchmark_Serial_GoLogger
Benchmark_Serial_GoLogger-4               576457              4010 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_GoLogger-8               513609              3966 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_Slog
Benchmark_Serial_Slog-4                   482364              4914 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_Slog-8                   479356              4921 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger
Benchmark_Serial_SlogAndGoLogger-4        461556              5058 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger-8        489222              5046 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_NativeLog
Benchmark_Parallel_NativeLog-4            475184              4916 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_NativeLog-8            485504              5026 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_Zap
Benchmark_Parallel_Zap-4                  349200              6773 ns/op             352 B/op          6 allocs/op
Benchmark_Parallel_Zap-8                  357151              6610 ns/op             353 B/op          6 allocs/op
Benchmark_Parallel_GoLogger
Benchmark_Parallel_GoLogger-4            4629242               568.1 ns/op           165 B/op          3 allocs/op
Benchmark_Parallel_GoLogger-8            4347580               576.0 ns/op           128 B/op          3 allocs/op
Benchmark_Parallel_SLog
Benchmark_Parallel_SLog-4                 478394              4952 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLog-8                 470799              5075 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger
Benchmark_Parallel_SLogAndGoLogger-4      449910              5150 ns/op             345 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger-8      439743              5250 ns/op             345 B/op          6 allocs/op
```

```text
Benchmark_Serial_NativeLog
Benchmark_Serial_NativeLog-4              578982              4068 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_NativeLog-8              527059              4032 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_Zap
Benchmark_Serial_Zap-4                    471318              4866 ns/op             352 B/op          6 allocs/op
Benchmark_Serial_Zap-8                    492363              4880 ns/op             353 B/op          6 allocs/op
Benchmark_Serial_GoLogger
Benchmark_Serial_GoLogger-4               576114              4096 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_GoLogger-8               535795              4084 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_Slog
Benchmark_Serial_Slog-4                   477541              4977 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_Slog-8                   478213              4970 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger
Benchmark_Serial_SlogAndGoLogger-4        465369              5056 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger-8        406258              5085 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_NativeLog
Benchmark_Parallel_NativeLog-4            467674              4922 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_NativeLog-8            475890              4963 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_Zap
Benchmark_Parallel_Zap-4                  352056              6703 ns/op             352 B/op          6 allocs/op
Benchmark_Parallel_Zap-8                  364632              6489 ns/op             353 B/op          6 allocs/op
Benchmark_Parallel_GoLogger
Benchmark_Parallel_GoLogger-4            4334329               545.7 ns/op           165 B/op          3 allocs/op
Benchmark_Parallel_GoLogger-8            4512244               541.5 ns/op           128 B/op          3 allocs/op
Benchmark_Parallel_SLog
Benchmark_Parallel_SLog-4                 483408              5012 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLog-8                 462290              5060 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger
Benchmark_Parallel_SLogAndGoLogger-4      456644              5372 ns/op             345 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger-8      442698              5264 ns/op             345 B/op          6 allocs/op
```

```text
Benchmark_Serial_NativeLog
Benchmark_Serial_NativeLog-4              618345              4051 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_NativeLog-8              600762              4045 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_Zap
Benchmark_Serial_Zap-4                    466516              4784 ns/op             352 B/op          6 allocs/op
Benchmark_Serial_Zap-8                    478645              4803 ns/op             353 B/op          6 allocs/op
Benchmark_Serial_GoLogger
Benchmark_Serial_GoLogger-4               602014              4031 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_GoLogger-8               573824              4078 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_Slog
Benchmark_Serial_Slog-4                   418012              4927 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_Slog-8                   472023              4925 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger
Benchmark_Serial_SlogAndGoLogger-4        476856              4988 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger-8        500853              5060 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_NativeLog
Benchmark_Parallel_NativeLog-4            468685              4907 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_NativeLog-8            479719              4908 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_Zap
Benchmark_Parallel_Zap-4                  342488              6929 ns/op             352 B/op          6 allocs/op
Benchmark_Parallel_Zap-8                  350866              6710 ns/op             353 B/op          6 allocs/op
Benchmark_Parallel_GoLogger
Benchmark_Parallel_GoLogger-4            4308105               547.1 ns/op           164 B/op          3 allocs/op
Benchmark_Parallel_GoLogger-8            4498796               543.2 ns/op           128 B/op          3 allocs/op
Benchmark_Parallel_SLog
Benchmark_Parallel_SLog-4                 474552              5020 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLog-8                 454094              5157 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger
Benchmark_Parallel_SLogAndGoLogger-4      431023              5252 ns/op             345 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger-8      414043              5241 ns/op             345 B/op          6 allocs/op
```

```text
Benchmark_Serial_NativeLog
Benchmark_Serial_NativeLog-4              573181              4019 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_NativeLog-8              555907              4188 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_Zap
Benchmark_Serial_Zap-4                    455720              4921 ns/op             352 B/op          6 allocs/op
Benchmark_Serial_Zap-8                    487197              5000 ns/op             353 B/op          6 allocs/op
Benchmark_Serial_GoLogger
Benchmark_Serial_GoLogger-4               620002              4083 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_GoLogger-8               598402              4039 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_Slog
Benchmark_Serial_Slog-4                   477624              4838 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_Slog-8                   520885              5031 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger
Benchmark_Serial_SlogAndGoLogger-4        478291              4980 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger-8        495639              5065 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_NativeLog
Benchmark_Parallel_NativeLog-4            456236              4928 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_NativeLog-8            485220              4928 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_Zap
Benchmark_Parallel_Zap-4                  353738              6668 ns/op             352 B/op          6 allocs/op
Benchmark_Parallel_Zap-8                  357856              6689 ns/op             353 B/op          6 allocs/op
Benchmark_Parallel_GoLogger
Benchmark_Parallel_GoLogger-4            4386367               550.2 ns/op           157 B/op          3 allocs/op
Benchmark_Parallel_GoLogger-8            4308735               640.8 ns/op           128 B/op          3 allocs/op
Benchmark_Parallel_SLog
Benchmark_Parallel_SLog-4                 472104              4983 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLog-8                 474613              4989 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger
Benchmark_Parallel_SLogAndGoLogger-4      447592              5200 ns/op             347 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger-8      454953              5213 ns/op             345 B/op          6 allocs/op
```

```text
Benchmark_Serial_NativeLog
Benchmark_Serial_NativeLog-4              543579              4068 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_NativeLog-8              584174              4058 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_Zap
Benchmark_Serial_Zap-4                    466582              4867 ns/op             352 B/op          6 allocs/op
Benchmark_Serial_Zap-8                    485359              4829 ns/op             353 B/op          6 allocs/op
Benchmark_Serial_GoLogger
Benchmark_Serial_GoLogger-4               531115              3983 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_GoLogger-8               547891              4099 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_Slog
Benchmark_Serial_Slog-4                   491005              4981 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_Slog-8                   498840              4970 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger
Benchmark_Serial_SlogAndGoLogger-4        454028              5124 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger-8        422102              5207 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_NativeLog
Benchmark_Parallel_NativeLog-4            476598              5041 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_NativeLog-8            486174              4997 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_Zap
Benchmark_Parallel_Zap-4                  356665              6676 ns/op             352 B/op          6 allocs/op
Benchmark_Parallel_Zap-8                  356013              6532 ns/op             353 B/op          6 allocs/op
Benchmark_Parallel_GoLogger
Benchmark_Parallel_GoLogger-4            4573236               530.2 ns/op           157 B/op          3 allocs/op
Benchmark_Parallel_GoLogger-8            4565874               519.8 ns/op           128 B/op          3 allocs/op
Benchmark_Parallel_SLog
Benchmark_Parallel_SLog-4                 462074              4980 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLog-8                 453877              5037 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger
Benchmark_Parallel_SLogAndGoLogger-4      450961              5205 ns/op             345 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger-8      442342              5354 ns/op             345 B/op          6 allocs/op
```

```text
Benchmark_Serial_NativeLog
Benchmark_Serial_NativeLog-4              514654              4323 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_NativeLog-8              570302              4145 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_Zap
Benchmark_Serial_Zap-4                    514950              4890 ns/op             352 B/op          6 allocs/op
Benchmark_Serial_Zap-8                    465789              4933 ns/op             353 B/op          6 allocs/op
Benchmark_Serial_GoLogger
Benchmark_Serial_GoLogger-4               615319              4088 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_GoLogger-8               594363              4088 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_Slog
Benchmark_Serial_Slog-4                   461121              4872 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_Slog-8                   499617              5059 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger
Benchmark_Serial_SlogAndGoLogger-4        459706              5031 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger-8        482911              5062 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_NativeLog
Benchmark_Parallel_NativeLog-4            431725              5079 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_NativeLog-8            479290              5022 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_Zap
Benchmark_Parallel_Zap-4                  358591              6675 ns/op             352 B/op          6 allocs/op
Benchmark_Parallel_Zap-8                  324229              6430 ns/op             353 B/op          6 allocs/op
Benchmark_Parallel_GoLogger
Benchmark_Parallel_GoLogger-4            4533364               557.3 ns/op           162 B/op          3 allocs/op
Benchmark_Parallel_GoLogger-8            4173549               739.4 ns/op           128 B/op          3 allocs/op
Benchmark_Parallel_SLog
Benchmark_Parallel_SLog-4                 476944              4989 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLog-8                 446238              5036 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger
Benchmark_Parallel_SLogAndGoLogger-4      428470              5196 ns/op             345 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger-8      434781              5288 ns/op             345 B/op          6 allocs/op
```

```text
Benchmark_Serial_NativeLog
Benchmark_Serial_NativeLog-4              598425              4095 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_NativeLog-8              589526              4272 ns/op             248 B/op          2 allocs/op
Benchmark_Serial_Zap
Benchmark_Serial_Zap-4                    485172              4943 ns/op             352 B/op          6 allocs/op
Benchmark_Serial_Zap-8                    491910              4851 ns/op             353 B/op          6 allocs/op
Benchmark_Serial_GoLogger
Benchmark_Serial_GoLogger-4               527454              3987 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_GoLogger-8               574303              4083 ns/op              80 B/op          2 allocs/op
Benchmark_Serial_Slog
Benchmark_Serial_Slog-4                   498553              4952 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_Slog-8                   466743              4942 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger
Benchmark_Serial_SlogAndGoLogger-4        443798              5149 ns/op             344 B/op          6 allocs/op
Benchmark_Serial_SlogAndGoLogger-8        460762              5208 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_NativeLog
Benchmark_Parallel_NativeLog-4            424681              5176 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_NativeLog-8            479988              5045 ns/op             248 B/op          2 allocs/op
Benchmark_Parallel_Zap
Benchmark_Parallel_Zap-4                  341937              6736 ns/op             352 B/op          6 allocs/op
Benchmark_Parallel_Zap-8                  353247              6517 ns/op             353 B/op          6 allocs/op
Benchmark_Parallel_GoLogger
Benchmark_Parallel_GoLogger-4            4240896               549.9 ns/op           163 B/op          3 allocs/op
Benchmark_Parallel_GoLogger-8            4441388               550.4 ns/op           128 B/op          3 allocs/op
Benchmark_Parallel_SLog
Benchmark_Parallel_SLog-4                 477423              4972 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLog-8                 447642              5064 ns/op             344 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger
Benchmark_Parallel_SLogAndGoLogger-4      424813              5242 ns/op             345 B/op          6 allocs/op
Benchmark_Parallel_SLogAndGoLogger-8      425070              5215 ns/op             345 B/op          6 allocs/op
```

| 库              | 测试类型       | 并发数 | 平均时间(ns/op) | 内存分配(B/op) | 内存分配次数(allocs/op) |
|------------------|----------------|--------|------------------|------------------|--------------------------|
| **NativeLog**    | Serial         | 4      | 3956             | 248              | 2                        |
|                  |                | 8      | 4044             | 248              | 2                        |
|                  | Parallel       | 4      | 4916             | 248              | 2                        |
|                  |                | 8      | 5026             | 248              | 2                        |
| **Zap**          | Serial         | 4      | 4815             | 352              | 6                        |
|                  |                | 8      | 4933             | 353              | 6                        |
|                  | Parallel       | 4      | 6773             | 352              | 6                        |
|                  |                | 8      | 6610             | 353              | 6                        |
| **GoLogger**     | Serial         | 4      | 4010             | 80               | 2                        |
|                  |                | 8      | 3966             | 80               | 2                        |
|                  | Parallel       | 4      | 568.1            | 165              | 3                        |
|                  |                | 8      | 576.0            | 128              | 3                        |
| **slog**         | Serial         | 4      | 4914             | 344              | 6                        |
|                  |                | 8      | 4921             | 344              | 6                        |
|                  | Parallel       | 4      | 4952             | 344              | 6                        |
|                  |                | 8      | 5075             | 344              | 6                        |
| **slog + GoLogger** | Serial      | 4      | 5058             | 344              | 6                        |
|                  |                | 8      | 5046             | 344              | 6                        |
|                  | Parallel       | 4      | 5150             | 345              | 6                        |
|                  |                | 8      | 5250             | 345              | 6                        |

### 性能分析

1. **NativeLog（log库）**:
    - **串行性能**: 具有相对较低的延迟（3956 ns/op 和 4044 ns/op），且内存占用较少（248 B/op）。
    - **并行性能**: 在并发测试中，NativeLog 的性能也保持稳定，延迟（4916 ns/op 和 5026 ns/op）仍然低于其他库。

2. **Zap（zap库）**:
    - **串行性能**: Zap 的串行性能稍逊色于 log，延迟略高（4815 ns/op 和 4933 ns/op），并且内存占用较高（352 B/op）。
    - **并行性能**: Zap 在并行测试中表现较差，延迟明显高于其他库，达到 6773 ns/op 和 6610 ns/op，显示出其在高并发情况下的瓶颈。

3. **GoLogger（go-logger）**:
    - **串行性能**: 在串行性能上表现良好，延迟（4010 ns/op 和 3966 ns/op），内存使用最低（80 B/op）。
    - **并行性能**: 在并行测试中表现优异，延迟显著低于其他库，仅为 568.1 ns/op 和 576.0 ns/op，显示了其极高的并发处理能力。

4. **slog（slog库）**:
    - **串行性能**: slog 的串行性能在所有库中属于中等水平（4914 ns/op 和 4921 ns/op），内存占用相对较高（344 B/op）。
    - **并行性能**: 在并行情况下，slog 的性能表现中规中矩（4952 ns/op 和 5075 ns/op）。

5. **slog + GoLogger（slog+go-logger）**:
    - **串行性能**: 当结合 slog 和 GoLogger 时，性能表现为（5058 ns/op 和 5046 ns/op），内存占用（344 B/op）与单独使用slog库相同。
    - **并行性能**: 并行测试中，组合使用的性能（5150 ns/op 和 5250 ns/op）。