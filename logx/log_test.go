package logx

import (
	"bytes"
	"github.com/donnie4w/gofer/util"
	"strconv"
	"testing"
	"time"
)

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
	for i := range 1 << 15 {
		go func() {
			e1 := c1.Debug([]byte("hello--------------->" + strconv.Itoa(i)))
			e2 := c2.Info([]byte("world--------------->" + strconv.Itoa(i)))
			e3 := c3.Warn([]byte("hello raftx--------------->" + strconv.Itoa(i)))
			if e1 != nil || e2 != nil || e3 != nil {
				t.Log(e1, e2, e3)
			}
		}()
	}
	time.Sleep(10 * time.Second)
	TestFileSync(t)
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

func TestFileSync(t *testing.T) {
	t.Log("fileByteEq 1&2:", fileByteEq1())
	t.Log("fileByteEq 1&3:", fileByteEq2())
}

func fileByteEq1() bool {
	bs1, _ := util.ReadFile("log1.log")
	bs2, _ := util.ReadFile("log2.log")
	return bytes.Equal(bs1, bs2)
}

func fileByteEq2() bool {
	bs1, _ := util.ReadFile("log1.log")
	bs2, _ := util.ReadFile("log3.log")
	return bytes.Equal(bs1, bs2)
}
