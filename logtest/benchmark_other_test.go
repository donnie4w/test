package logtest

import (
	logger "github.com/donnie4w/simplelog/logging"
	"testing"
)

// go-logger FORMAT_NANO模式
func BenchmarkSerialGoLoggerNoFORMAT(b *testing.B) {
	b.StopTimer()
	log, _ := logger.NewLogger().SetConsole(false).SetFormat(logger.FORMAT_NANO).SetRollingFile(``, `gologFormat.txt`, 1, logger.GB)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:46:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

// go-logger write 方法
func BenchmarkSerialGoLoggerWrite(b *testing.B) {
	b.StopTimer()
	log, _ := logger.NewLogger().SetRollingFile(`D:\cfoldTest\`, `gologWrite.txt`, 1, logger.GB)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Write([]byte("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:55:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n"))
	}
}

// go-logger FORMAT_NANO模式
func BenchmarkParallelGoLoggerNoFORMAT(b *testing.B) {
	log, _ := logger.NewLogger().SetConsole(false).SetFormat(logger.FORMAT_NANO).SetRollingFile(``, `gologFormat.txt`, 1, logger.GB)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:46:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
	})
}

// go-logger write 方法
func BenchmarkParallelGoLoggerWrite(b *testing.B) {
	log, _ := logger.NewLogger().SetRollingFile(``, `gologWrite.txt`, 1, logger.GB)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Write([]byte("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:55:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n"))
		}
	})
}
