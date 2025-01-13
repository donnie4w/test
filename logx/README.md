#### Logx简介
Logx 是一个基于 Go 语言实现的分布式日志系统基础框架，它使用了 `raftx` 作为一致性协议来确保跨节点的日志数据一致性和顺序性。设计用于高并发、高可用性的场景，并支持多种日志级别（Debug, Info, Warn, Error）的日志记录。

#### 特性
- **一致性**：通过 Raft 协议保证所有节点上的日志条目最终一致。
- **可用性**：即使部分节点失效，系统仍然可以处理日志写入。(注意:宕机节点需要增加丢失数据同步的逻辑)
- **性能**：支持高吞吐量的日志写入和快速的查询响应时间。
- **扩展性**：能够随着数据增长或用户数量增加而水平扩展。
- **多级日志**：支持 Debug, Info, Warn, Error 四种日志级别。
- **并发处理**：提供了并发写入日志的能力，适用于并发场景。


#### 测试示例

```go
var c1 *Logx
var c2 *Logx
var c3 *Logx

func init() {
	c1 = newlog1()
	c2 = newlog2()
	c3 = newlog3()
	time.Sleep(3 * time.Second) //模拟等待leader选举完成
}

func Test_logx(t *testing.T) {
	for i := range 1000 {
		c1.Debug([]byte("hello--------------->" + strconv.Itoa(i)))
	}
	time.Sleep(3 * time.Second)
	t.Log("fileByteEq 1&2:", fileByteEq1())
	t.Log("fileByteEq 1&3:", fileByteEq2())
}

func Benchmark_logx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c1.Debug([]byte("hello--------------->" + strconv.Itoa(i)))
		c2.Info([]byte("world--------------->" + strconv.Itoa(i)))
		c3.Warn([]byte("hello raftx--------------->" + strconv.Itoa(i)))
	}
}

func Test_Parallel(t *testing.T) {
	for i := range 1 << 17 { //这里将模拟每个节点 13万并发写日志数据
		go func() {
			e1 := c1.Debug([]byte("hello--------------->" + strconv.Itoa(i)))
			e2 := c2.Info([]byte("world--------------->" + strconv.Itoa(i)))
			e3 := c3.Warn([]byte("hello raftx--------------->" + strconv.Itoa(i)))
			if e1 != nil || e2 != nil || e3 != nil {
				t.Log(e1, e2, e3)
			}
		}()
	}
	time.Sleep(30 * time.Second)
	TestFileSync(t) //检查各个节点生成的日志文件是否相同
}

func newlog1() *Logx {
	return NewLogx("log1.log", ":20001", []string{"127.0.0.1:20001", "127.0.0.1:20002", "127.0.0.1:20003"})
}
func newlog2() *Logx {
	return NewLogx("log2.log", ":20002", []string{"127.0.0.1:20001", "127.0.0.1:20002", "127.0.0.1:20003"})
}
func newlog3() *Logx {
	return NewLogx("log3.log", ":20003", []string{"127.0.0.1:20001", "127.0.0.1:20002", "127.0.0.1:20003"})
}
```

- **压测调用**：可以通过 `Benchmark_logx` 函数来进行压力测试，评估系统的性能。
- **并发测试**：`Test_Parallel` 函数用于测试在高并发情况下的表现，并验证各个节点生成的日志文件是否相同。

#### 注意事项

- 如果测试过程中过早退出，可能会导致某些节点的数据未完全同步，出现日志缺失的情况。
- 对于长时间断开后重新连接的节点，由于 `raftx` 分布式易失性数据的特性，可能无法恢复所有历史数据。因此，建议定期进行快照备份，并在节点重启时从其他健康节点同步最新状态。