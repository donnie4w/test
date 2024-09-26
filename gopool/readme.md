```bash
go test -v -run ^Test_pool$ .
```
```text
=== RUN   Test_pool
9.3173979s
1773304 578693904 165
10.5804948s
556486136 1767823032 175
--- PASS: Test_pool (19.92s)
PASS
ok      test/gopool     20.414s
```

------


```bash
go test -v -run ^Test_pool$ .
```
```text
=== RUN   Test_pool
9.8880305s
3479872 578174336 167
8.5609988s
668965496 1549807680 177
--- PASS: Test_pool (18.47s)
PASS
ok      test/gopool     19.011s
```


------


```bash
go test -v -run ^Test_pool$ .
```
```text
=== RUN   Test_pool
9.9579494s
1669864 578493344 167
9.1508301s
677125960 1553627912 176
--- PASS: Test_pool (19.13s)
PASS
ok      test/gopool     19.675s
```
-------

#### 结果说明：

* 9.9579494s  gopool 执行时间
* 1669864 578493344 167   ：  Allo:1669864    totalAlloc:578493344    NumGC:167
* 9.1508301s  原生go 执行时间
* 677125960 1553627912 176 ： Allo:677125960  totalAlloc:1553627912   NumGC:176


测试中：NewPool(50, 100) 最小活跃协程50，最大100.

测试中100个协程与100万个协程工作效率基本一致
内存消耗上节省大约 2/3，相应触发gc的次数也少一些

根据实际业务调整 gopool池的大小，才能得到更好的结果
并非处处要用协程池，应该根据实际情况考虑
大部分情况下，调用go创建协程就可以了