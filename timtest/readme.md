#### 测试数据与长度

	MsType:    1,                                                                         //1
	OdType:    1,                                                                         //1
	Id:        1 << 60,                                                                   //8
	Mid:       1 << 61,                                                                   //8
	BnType:    1 << 30,                                                                   //4
	FromTid:   &TidBean{Node: "tom"},                                                     //3
	ToTid:     &TidBean{Node: "jerry"},                                                   //5
	Body:      []byte("A good programmer"),                                               //17
	IsOffline: true,                                                                      //1
	Timestamp: 1 << 60,                                                                   //8    ->    56
	Extend:    map[string]string{"ex1": "A good programmer", "ex2": "A good programmer"}, //40   ->    96

-------

##### 测试输入数据 总长度 96字节：

##### 序列化后数据长度：
* Msgpack ：254 字节
* Json    ：366 字节
* Thrift  ：124 字节
* ProtoBuf：125 字节

```text
pkg: test/timtest
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkSerialMsgpackEncode
BenchmarkSerialMsgpackEncode-4            883100              1411 ns/op             616 B/op          8 allocs/op
BenchmarkSerialMsgpackEncode-8            888198              1400 ns/op             616 B/op          8 allocs/op
BenchmarkSerialMsgpackDecode
BenchmarkSerialMsgpackDecode-4            526042              2206 ns/op             432 B/op         28 allocs/op
BenchmarkSerialMsgpackDecode-8            515901              2192 ns/op             432 B/op         28 allocs/op
BenchmarkSerialGojsonEncode
BenchmarkSerialGojsonEncode-4             741568              1622 ns/op             712 B/op          9 allocs/op
BenchmarkSerialGojsonEncode-8             703032              1694 ns/op             712 B/op          9 allocs/op
BenchmarkSerialGoJsonDecode
BenchmarkSerialGoJsonDecode-4             220232              5560 ns/op             536 B/op         22 allocs/op
BenchmarkSerialGoJsonDecode-8             215497              5564 ns/op             536 B/op         22 allocs/op
BenchmarkSerialThriftEncode
BenchmarkSerialThriftEncode-4             980503              1167 ns/op             456 B/op          7 allocs/op
BenchmarkSerialThriftEncode-8            1028618              1186 ns/op             456 B/op          7 allocs/op
BenchmarkSerialThriftDecode
BenchmarkSerialThriftDecode-4             879506              1449 ns/op             912 B/op         22 allocs/op
BenchmarkSerialThriftDecode-8             773559              1548 ns/op             912 B/op         22 allocs/op
BenchmarkSerialProtoBufEncode
BenchmarkSerialProtoBufEncode-4           997920              1234 ns/op             256 B/op          9 allocs/op
BenchmarkSerialProtoBufEncode-8          1000000              1234 ns/op             256 B/op          9 allocs/op
BenchmarkSerialProtoBufDecode
BenchmarkSerialProtoBufDecode-4           696997              1669 ns/op             928 B/op         29 allocs/op
BenchmarkSerialProtoBufDecode-8           696936              1737 ns/op             928 B/op         29 allocs/op
```


##### 序列化与反序列化 压测 结果：

##### 串行压测
```text
序列化                                  反序列化
Msgpack : 1400 ns/op                    Msgpack : 2206 ns/op
Json    : 1622 ns/op                    Json    : 5560 ns/op
Thrift  : 1167 ns/op                    Thrift  : 1449 ns/op
ProtoBuf: 1234 ns/op                    ProtoBuf: 1669 ns/op
```

-------

```text
BenchmarkParallelMsgpackEncode
BenchmarkParallelMsgpackEncode-4         2258724               492.1 ns/op           616 B/op          8 allocs/op
BenchmarkParallelMsgpackEncode-8         2138186               557.9 ns/op           616 B/op          8 allocs/op
BenchmarkParallelMsgpackDecode
BenchmarkParallelMsgpackDecode-4         1230190               944.8 ns/op           1015 B/op         34 allocs/op
BenchmarkParallelMsgpackDecode-8         1000000               1159 ns/op            1015 B/op         34 allocs/op
BenchmarkParallelGojsonEncode
BenchmarkParallelGojsonEncode-4          2028855               582.4 ns/op           712 B/op          9 allocs/op
BenchmarkParallelGojsonEncode-8          1777850               686.1 ns/op           712 B/op          9 allocs/op
BenchmarkParallelGoJsonDecode
BenchmarkParallelGoJsonDecode-4           578404               2079 ns/op            1096 B/op         27 allocs/op
BenchmarkParallelGoJsonDecode-8           595042               2381 ns/op            1096 B/op         27 allocs/op
BenchmarkParallelThriftEncode
BenchmarkParallelThriftEncode-4          2622094               456.3 ns/op           456 B/op          7 allocs/op
BenchmarkParallelThriftEncode-8          2328318               518.0 ns/op           456 B/op          7 allocs/op
BenchmarkParallelThriftDecode
BenchmarkParallelThriftDecode-4          1905132               756.4 ns/op           912 B/op          22 allocs/op
BenchmarkParallelThriftDecode-8          1000000               1103 ns/op            912 B/op          22 allocs/op
BenchmarkParallelProtoBufEncode
BenchmarkParallelProtoBufEncode-4        1745168               624.9 ns/op           256 B/op          9 allocs/op
BenchmarkParallelProtoBufEncode-8        1854494               652.8 ns/op           256 B/op          9 allocs/op
BenchmarkParallelProtoBufDecode
BenchmarkParallelProtoBufDecode-4        1287536               906.7 ns/op           928 B/op          29 allocs/op
BenchmarkParallelProtoBufDecode-8        1000000               1042 ns/op            928 B/op          29 allocs/op
```


#### 序列化与反序列化 压测 结果：

##### 并行压测

```text
序列化                                   反序列化：
Msgpack : 492.1 ns/op                    Msgpack : 944.8 ns/op 
Json    : 582.4 ns/op                    Json    : 2079 ns/op  
Thrift  : 456.3 ns/op                    Thrift  : 756.4 ns/op 
ProtoBuf: 624.9 ns/op                    ProtoBuf: 906.7 ns/op
```
