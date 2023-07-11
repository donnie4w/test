package gopool

import (
	"fmt"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"testing"
	"time"

	gol "github.com/donnie4w/gofer/pool/gopool"
	"github.com/donnie4w/simplelog/logging"
)

//gopool vs go  :compare memory consumption with time consuming

var log1 = logging.NewLogger().SetGzipOn(true).SetConsole(false)
var log2 = logging.NewLogger().SetGzipOn(true).SetConsole(false)

func Test_pool(t *testing.T) {
	log1.SetRollingFileLoop("", "pool1.txt", 300, logging.MB, 5)
	log2.SetRollingFileLoop("", "pool2.txt", 300, logging.MB, 5)
	testPool()
	testgo()
}

func testPool() {
	pool := gol.NewPool(50, 100)
	var wg sync.WaitGroup
	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		pool.Go(func() {
			s := ""
			for i := 0; i < 10; i++ {
				s = fmt.Sprint(s, i)
			}
			log1.Debug(s)
			wg.Done()
		})
	}
	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	fmt.Println(getSysinfo())
}

func testgo() {
	var wg sync.WaitGroup
	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			s := ""
			for i := 0; i < 10; i++ {
				s = fmt.Sprint(s, i)
			}
			log2.Debug(s)
			wg.Done()
		}()
	}
	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	fmt.Println(getSysinfo())
}

func getSysinfo() (alloc, totalalloc uint64, numgc uint32) {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	alloc = rtm.Alloc
	totalalloc = rtm.TotalAlloc
	numgc = rtm.NumGC
	return
}
