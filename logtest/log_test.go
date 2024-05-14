package logtest

import (
	"log"
	_ "net/http/pprof"
	"os"
	"sync"
	"testing"
	"time"

	logger "github.com/donnie4w/simplelog/logging"
)

func init() {
	// go http.ListenAndServe(":9000", nil)
	// <-time.After(5 * time.Second)
}

func Test_Log(t *testing.T) {
	go Zap()
	go Logger()
	go LoggerNoFORMAT()
	go LoggerWRITE()
	go NativeGoLog()
	<-time.After(300 * time.Second)
}

func Zap() {
	log := initzap()
	var wg sync.WaitGroup
	for c := 0; c < 10; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			}
		}()
	}
	wg.Wait()
}

func Logger() {
	log, _ := logger.NewLogger().SetConsole(false).SetRollingFile(`D:\cfoldTest\`, `golog.txt`, 1, logger.GB)
	var wg sync.WaitGroup
	for c := 0; c < 10; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			}
		}()
	}
	wg.Wait()
}

func LoggerNoFORMAT() {
	log, _ := logger.NewLogger().SetConsole(false).SetFormat(logger.FORMAT_NANO).SetRollingFile(`D:\cfoldTest\`, `gologFormat.txt`, 1, logger.GB)
	var wg sync.WaitGroup
	for c := 0; c < 10; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				log.Debug("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:46: >>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			}
		}()
	}
	wg.Wait()
}

func LoggerWRITE() {
	log, _ := logger.NewLogger().SetRollingFile(`D:\cfoldTest\`, `gologWrite.txt`, 1, logger.GB)
	var wg sync.WaitGroup
	for c := 0; c < 10; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				log.Write([]byte("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:55:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n"))
			}
		}()
	}
	wg.Wait()
}

func NativeGoLog() {
	out, _ := os.OpenFile(`D:\cfoldTest\nativelog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	log := log.New(out, "[debug]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	var wg sync.WaitGroup
	for c := 0; c < 10; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				log.Println(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			}
		}()
	}
	wg.Wait()

}
