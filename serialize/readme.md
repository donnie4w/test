```bash
go test -v -run=Benchmark -bench=Benchmark -benchmem -cpu=4,8 -parallel=8 ./
```

```text
goarch: amd64
pkg: test/serialize
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkMsgpackEncode
BenchmarkMsgpackEncode-4                  548551              1931 ns/op             856 B/op         14 allocs/op
BenchmarkMsgpackEncode-8                  764978              1927 ns/op             856 B/op         14 allocs/op
BenchmarkMsgpackDecode
BenchmarkMsgpackDecode-4                  472929              2505 ns/op             520 B/op         24 allocs/op
BenchmarkMsgpackDecode-8                  478869              3875 ns/op             520 B/op         24 allocs/op
BenchmarkGojsonEncode
BenchmarkGojsonEncode-4                  1000000              1513 ns/op             832 B/op         12 allocs/op
BenchmarkGojsonEncode-8                   765394              1558 ns/op             832 B/op         12 allocs/op
BenchmarkGoJsonDecode
BenchmarkGoJsonDecode-4                   324608              4424 ns/op             608 B/op         26 allocs/op
BenchmarkGoJsonDecode-8                   291723              4210 ns/op             608 B/op         26 allocs/op
BenchmarkThriftEncode
BenchmarkThriftEncode-4                   855663              1181 ns/op             616 B/op          9 allocs/op
BenchmarkThriftEncode-8                  1000000              1203 ns/op             616 B/op          9 allocs/op
BenchmarkThriftDecode
BenchmarkThriftDecode-4                  1000000              1333 ns/op             976 B/op         18 allocs/op
BenchmarkThriftDecode-8                  1000000              1393 ns/op             976 B/op         18 allocs/op
BenchmarkProtoBufEncode
BenchmarkProtoBufEncode-4                 739744              1676 ns/op             512 B/op         16 allocs/op
BenchmarkProtoBufEncode-8                 728253              1655 ns/op             512 B/op         16 allocs/op
BenchmarkProtoBufDecode
BenchmarkProtoBufDecode-4                 624775              2074 ns/op            1024 B/op         30 allocs/op
BenchmarkProtoBufDecode-8                 593330              2075 ns/op            1024 B/op         30 allocs/op
BenchmarkParallelMsgpackEncode
BenchmarkParallelMsgpackEncode-4         1847612               669.4 ns/op           856 B/op         14 allocs/op
BenchmarkParallelMsgpackEncode-8         1446364               814.3 ns/op           856 B/op         14 allocs/op
BenchmarkParallelMsgpackDecode
BenchmarkParallelMsgpackDecode-4         1000000              1132 ns/op            1264 B/op         34 allocs/op
BenchmarkParallelMsgpackDecode-8         1000000              1344 ns/op            1264 B/op         34 allocs/op
BenchmarkParallelGojsonEncode
BenchmarkParallelGojsonEncode-4          1989336               568.0 ns/op           832 B/op         12 allocs/op
BenchmarkParallelGojsonEncode-8          1362235               905.0 ns/op           832 B/op         12 allocs/op
BenchmarkParallelGoJsonDecode
BenchmarkParallelGoJsonDecode-4           548776              2020 ns/op            1239 B/op         32 allocs/op
BenchmarkParallelGoJsonDecode-8           665671              2357 ns/op            1239 B/op         32 allocs/op
BenchmarkParallelThriftEncode
BenchmarkParallelThriftEncode-4          2087115               555.6 ns/op           616 B/op          9 allocs/op
BenchmarkParallelThriftEncode-8          1826240               678.1 ns/op           616 B/op          9 allocs/op
BenchmarkParallelThriftDecode
BenchmarkParallelThriftDecode-4          1757334               697.9 ns/op           976 B/op         18 allocs/op
BenchmarkParallelThriftDecode-8          1475140               818.3 ns/op           976 B/op         18 allocs/op
BenchmarkParallelProtoBufEncode
BenchmarkParallelProtoBufEncode-4        1551628               750.0 ns/op           512 B/op         16 allocs/op
BenchmarkParallelProtoBufEncode-8        1369939               860.6 ns/op           512 B/op         16 allocs/op
BenchmarkParallelProtoBufDecode
BenchmarkParallelProtoBufDecode-4        1000000              1076 ns/op            1136 B/op         31 allocs/op
BenchmarkParallelProtoBufDecode-8        1000000              1227 ns/op            1136 B/op         31 allocs/op
```
