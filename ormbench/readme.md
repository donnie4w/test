```bash
go test -v -run=Benchmark -bench=Benchmark -benchmem -benchtime 2s -cpu=4,8 -parallel=8 ./
```

#### 连接mysql场景的压测数据

```text
goarch: amd64
pkg: test/gdaobench
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
Benchmark_Native
Benchmark_Native-4                          3896            270091 ns/op            5200 B/op        105 allocs/op
Benchmark_Native-8                          4480            280356 ns/op            5200 B/op        105 allocs/op
Benchmark_gdao_struct
Benchmark_gdao_struct-4                     5524            201214 ns/op           20482 B/op        487 allocs/op
Benchmark_gdao_struct-8                     5959            201914 ns/op           20489 B/op        487 allocs/op
Benchmark_gdao_sql
Benchmark_gdao_sql-4                        6686            184735 ns/op           17664 B/op        391 allocs/op
Benchmark_gdao_sql-8                        6445            185814 ns/op           17671 B/op        391 allocs/op
Benchmark_gdao_databean
Benchmark_gdao_databean-4                   6716            176825 ns/op           10898 B/op        312 allocs/op
Benchmark_gdao_databean-8                   7009            183525 ns/op           10901 B/op        312 allocs/op
Benchmark_gdao_mapper_struct
Benchmark_gdao_mapper_struct-4              4930            214244 ns/op           19838 B/op        505 allocs/op
Benchmark_gdao_mapper_struct-8              5661            205118 ns/op           19841 B/op        505 allocs/op
Benchmark_gdao_mapper_databean
Benchmark_gdao_mapper_databean-4            6152            186940 ns/op           15961 B/op        407 allocs/op
Benchmark_gdao_mapper_databean-8            6427            197207 ns/op           15969 B/op        407 allocs/op
Benchmark_gdao_cache_struct
Benchmark_gdao_cache_struct-4             758840              1631 ns/op            1472 B/op         21 allocs/op
Benchmark_gdao_cache_struct-8             751695              1678 ns/op            1472 B/op         21 allocs/op
Benchmark_gdao_cache_mapper
Benchmark_gdao_cache_mapper-4            2204226               549.0 ns/op           195 B/op          6 allocs/op
Benchmark_gdao_cache_mapper-8            2220657               549.8 ns/op           195 B/op          6 allocs/op
Benchmark_gdao_cache_mapper_sql
Benchmark_gdao_cache_mapper_sql-4        2301375               514.4 ns/op           232 B/op          8 allocs/op
Benchmark_gdao_cache_mapper_sql-8        2281491               536.6 ns/op           232 B/op          8 allocs/op
Benchmark_gdao_struct_reflect
Benchmark_gdao_struct_reflect-4             5845            208120 ns/op           10148 B/op        445 allocs/op
Benchmark_gdao_struct_reflect-8             5082            205184 ns/op           10152 B/op        445 allocs/op
Benchmark_gdao_sql_reflect
Benchmark_gdao_sql_reflect-4                5896            201146 ns/op           14518 B/op        456 allocs/op
Benchmark_gdao_sql_reflect-8                5798            201342 ns/op           14523 B/op        456 allocs/op
Benchmark_gorm_struct
Benchmark_gorm_struct-4                     3465            335664 ns/op           16654 B/op        251 allocs/op
Benchmark_gorm_struct-8                     3429            331054 ns/op           16664 B/op        251 allocs/op
Benchmark_gorm_sql
Benchmark_gorm_sql-4                        3506            311013 ns/op           14013 B/op        208 allocs/op
Benchmark_gorm_sql-8                        4014            318121 ns/op           14019 B/op        208 allocs/op
Benchmark_sqlx_sql
Benchmark_sqlx_sql-4                        3852            294354 ns/op            8601 B/op        123 allocs/op
Benchmark_sqlx_sql-8                        4208            306885 ns/op            8602 B/op        123 allocs/op
```

------

#### 连接postgreSql场景的压测数据

```text

goarch: amd64
pkg: test/gdaobench
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
Benchmark_Native
Benchmark_Native-4                          4276            241257 ns/op            2800 B/op         97 allocs/op
Benchmark_Native-8                          5088            242505 ns/op            2800 B/op         97 allocs/op
Benchmark_gdao_struct
Benchmark_gdao_struct-4                     8365            141638 ns/op           19719 B/op        448 allocs/op
Benchmark_gdao_struct-8                     8569            139506 ns/op           19727 B/op        448 allocs/op
Benchmark_gdao_sql
Benchmark_gdao_sql-4                        7666            133117 ns/op           16222 B/op        376 allocs/op
Benchmark_gdao_sql-8                        9184            131303 ns/op           16229 B/op        376 allocs/op
Benchmark_gdao_databean
Benchmark_gdao_databean-4                   9296            127276 ns/op           11610 B/op        328 allocs/op
Benchmark_gdao_databean-8                   9146            124861 ns/op           11614 B/op        328 allocs/op
Benchmark_gdao_mapper_struct
Benchmark_gdao_mapper_struct-4              8725            142578 ns/op           13552 B/op        396 allocs/op
Benchmark_gdao_mapper_struct-8              8104            137572 ns/op           13554 B/op        396 allocs/op
Benchmark_gdao_mapper_databean
Benchmark_gdao_mapper_databean-4            9306            129995 ns/op           11051 B/op        341 allocs/op
Benchmark_gdao_mapper_databean-8            9505            130523 ns/op           11053 B/op        341 allocs/op
Benchmark_gdao_cache_struct
Benchmark_gdao_cache_struct-4             745035              1662 ns/op            1472 B/op         21 allocs/op
Benchmark_gdao_cache_struct-8             738948              1680 ns/op            1472 B/op         21 allocs/op
Benchmark_gdao_cache_mapper
Benchmark_gdao_cache_mapper-4            2233680               542.6 ns/op           195 B/op          6 allocs/op
Benchmark_gdao_cache_mapper-8            2187944               542.1 ns/op           195 B/op          6 allocs/op
Benchmark_gdao_cache_mapper_sql
Benchmark_gdao_cache_mapper_sql-4        2236605               519.0 ns/op           232 B/op          8 allocs/op
Benchmark_gdao_cache_mapper_sql-8        2217334               533.0 ns/op           232 B/op          8 allocs/op
Benchmark_gdao_struct_reflect
Benchmark_gdao_struct_reflect-4             8601            139569 ns/op           10924 B/op        408 allocs/op
Benchmark_gdao_struct_reflect-8             8404            138996 ns/op           10928 B/op        408 allocs/op
Benchmark_gdao_sql_reflect
Benchmark_gdao_sql_reflect-4                7766            149846 ns/op           13044 B/op        416 allocs/op
Benchmark_gdao_sql_reflect-8                8704            144306 ns/op           13049 B/op        416 allocs/op
Benchmark_gorm_struct
Benchmark_gorm_struct-4                     8085            147442 ns/op           13253 B/op        231 allocs/op
Benchmark_gorm_struct-8                     7827            149199 ns/op           13261 B/op        231 allocs/op
Benchmark_gorm_sql
Benchmark_gorm_sql-4                        8275            147410 ns/op           11466 B/op        194 allocs/op
Benchmark_gorm_sql-8                        8739            144095 ns/op           11605 B/op        194 allocs/op
Benchmark_sqlx_sql
Benchmark_sqlx_sql-4                        4758            260370 ns/op            3080 B/op         65 allocs/op
Benchmark_sqlx_sql-8                        4665            267808 ns/op            3080 B/op         65 allocs/op
```

-------

#### 连接postgreSql场景的压测数据

```text
pkg: github.com/donnie4w/test/ormbench
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkSerial_Native
BenchmarkSerial_Native-4                            4800            247637 ns/op            2816 B/op         97 allocs/op
BenchmarkSerial_Native-8                            4441            250316 ns/op            2816 B/op         97 allocs/op
BenchmarkSerial_gdao_struct
BenchmarkSerial_gdao_struct-4                       8030            146417 ns/op           19997 B/op        449 allocs/op
BenchmarkSerial_gdao_struct-8                       8245            150725 ns/op           20023 B/op        449 allocs/op
BenchmarkSerial_gdao_sql
BenchmarkSerial_gdao_sql-4                          7287            141983 ns/op           16481 B/op        378 allocs/op
BenchmarkSerial_gdao_sql-8                          8625            138804 ns/op           16502 B/op        378 allocs/op
BenchmarkSerial_gdao_databean
BenchmarkSerial_gdao_databean-4                     9358            131579 ns/op           11859 B/op        330 allocs/op
BenchmarkSerial_gdao_databean-8                     9093            131845 ns/op           11872 B/op        330 allocs/op
BenchmarkSerial_gdao_mapper_struct
BenchmarkSerial_gdao_mapper_struct-4                6784            150160 ns/op           13716 B/op        398 allocs/op
BenchmarkSerial_gdao_mapper_struct-8                8329            148866 ns/op           13731 B/op        398 allocs/op
BenchmarkSerial_gdao_mapper_databean
BenchmarkSerial_gdao_mapper_databean-4              8152            134157 ns/op           11205 B/op        343 allocs/op
BenchmarkSerial_gdao_mapper_databean-8              8794            138595 ns/op           11217 B/op        343 allocs/op
BenchmarkSerial_gorm_struct
BenchmarkSerial_gorm_struct-4                       6547            155387 ns/op           13295 B/op        231 allocs/op
BenchmarkSerial_gorm_struct-8                       7963            154454 ns/op           13323 B/op        231 allocs/op
BenchmarkSerial_gorm_sql
BenchmarkSerial_gorm_sql-4                          7987            149439 ns/op           11772 B/op        194 allocs/op
BenchmarkSerial_gorm_sql-8                          7690            152698 ns/op           12011 B/op        194 allocs/op
BenchmarkSerial_sqlx_sql
BenchmarkSerial_sqlx_sql-4                          4466            276594 ns/op            3097 B/op         65 allocs/op
BenchmarkSerial_sqlx_sql-8                          4030            270020 ns/op            3098 B/op         65 allocs/op
BenchmarkSerial_ent
BenchmarkSerial_ent-4                               3814            283573 ns/op           10136 B/op        280 allocs/op
BenchmarkSerial_ent-8                               4154            280818 ns/op           10136 B/op        280 allocs/op
```

#### 连接mysql场景的压测数据

```text
pkg: github.com/donnie4w/test/ormbench
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkSerial_Native
BenchmarkSerial_Native-4                            3916            290730 ns/op            5216 B/op        105 allocs/op
BenchmarkSerial_Native-8                            4207            292908 ns/op            5216 B/op        105 allocs/op
BenchmarkSerial_gdao_struct
BenchmarkSerial_gdao_struct-4                       5109            200813 ns/op           20730 B/op        488 allocs/op
BenchmarkSerial_gdao_struct-8                       5972            203126 ns/op           20758 B/op        488 allocs/op
BenchmarkSerial_gdao_sql
BenchmarkSerial_gdao_sql-4                          5436            187344 ns/op           17893 B/op        393 allocs/op
BenchmarkSerial_gdao_sql-8                          6284            196396 ns/op           17918 B/op        393 allocs/op
BenchmarkSerial_gdao_databean
BenchmarkSerial_gdao_databean-4                     6566            178864 ns/op           11114 B/op        314 allocs/op
BenchmarkSerial_gdao_databean-8                     6204            183293 ns/op           11124 B/op        314 allocs/op
BenchmarkSerial_gdao_mapper_struct
BenchmarkSerial_gdao_mapper_struct-4                5756            212888 ns/op           20006 B/op        507 allocs/op
BenchmarkSerial_gdao_mapper_struct-8                5682            207857 ns/op           20036 B/op        508 allocs/op
BenchmarkSerial_gdao_mapper_databean
BenchmarkSerial_gdao_mapper_databean-4              6091            202900 ns/op           16118 B/op        409 allocs/op
BenchmarkSerial_gdao_mapper_databean-8              6054            204899 ns/op           16135 B/op        409 allocs/op
BenchmarkSerial_gorm_struct
BenchmarkSerial_gorm_struct-4                       3116            355411 ns/op           16700 B/op        251 allocs/op
BenchmarkSerial_gorm_struct-8                       3550            356480 ns/op           16734 B/op        251 allocs/op
BenchmarkSerial_gorm_sql
BenchmarkSerial_gorm_sql-4                          3742            314655 ns/op           14044 B/op        208 allocs/op
BenchmarkSerial_gorm_sql-8                          3735            327450 ns/op           14066 B/op        208 allocs/op
BenchmarkSerial_sqlx_sql
BenchmarkSerial_sqlx_sql-4                          3524            295426 ns/op            8620 B/op        123 allocs/op
BenchmarkSerial_sqlx_sql-8                          4105            303361 ns/op            8623 B/op        123 allocs/op
BenchmarkSerial_ent
BenchmarkSerial_ent-4                               3543            329408 ns/op           14144 B/op        319 allocs/op
BenchmarkSerial_ent-8                               3548            331312 ns/op           14144 B/op        319 allocs/op
```