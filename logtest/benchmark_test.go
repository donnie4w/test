package logtest

import (
	"log"
	"os"
	"testing"

	_ "net/http/pprof"

	logger "github.com/donnie4w/simplelog/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initzap() *zap.Logger {
	writeSyncer, _ := os.OpenFile(`D:\cfoldTest\zaplog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666) //日志文件存放目录
	encoderConfig := zap.NewProductionEncoderConfig()                                                 //指定时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)               //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	log := zap.New(core, zap.AddCaller())                             //AddCaller()为显示文件名和行号
	return log
}

// Zap 常规打印
func BenchmarkZap(b *testing.B) {
	b.StopTimer()
	log := initzap()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

// go-logger 常规打印
func BenchmarkLogger(b *testing.B) {
	b.StopTimer()
	log, _ := logger.NewLogger().SetConsole(false).SetRollingFile(`D:\cfoldTest\`, `golog.txt`, 1, logger.GB)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

// go-logger FORMAT_NANO模式
func BenchmarkLoggerNoFORMAT(b *testing.B) {
	b.StopTimer()
	log, _ := logger.NewLogger().SetConsole(false).SetFormat(logger.FORMAT_NANO).SetRollingFile(`D:\cfoldTest\`, `gologFormat.txt`, 1, logger.GB)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:46:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

// go-logger write 方法
func BenchmarkLoggerWrite(b *testing.B) {
	b.StopTimer()
	log, _ := logger.NewLogger().SetRollingFile(`D:\cfoldTest\`, `gologWrite.txt`, 1, logger.GB)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Write([]byte("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:55:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n"))
	}
}

// go自动log
func BenchmarkNativeGoLog(b *testing.B) {
	b.StopTimer()
	out, _ := os.OpenFile(`D:\cfoldTest\nativelog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	log := log.New(out, "[debug]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Println(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

// Zap 常规打印
func BenchmarkParallelZap(b *testing.B) {
	log := initzap()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
	})

}

// go-logger 常规打印
func BenchmarkParallelLogger(b *testing.B) {
	// b.SetParallelism(8)
	log, _ := logger.NewLogger().SetConsole(false).SetRollingFile(`D:\cfoldTest\`, `golog.txt`, 1, logger.GB)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
	})
}

// go-logger FORMAT_NANO模式
func BenchmarkParallelLoggerNoFORMAT(b *testing.B) {
	log, _ := logger.NewLogger().SetConsole(false).SetFormat(logger.FORMAT_NANO).SetRollingFile(`D:\cfoldTest\`, `gologFormat.txt`, 1, logger.GB)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:46:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
	})
}

// go-logger write 方法
func BenchmarkParallelLoggerWrite(b *testing.B) {
	log, _ := logger.NewLogger().SetRollingFile(`D:\cfoldTest\`, `gologWrite.txt`, 1, logger.GB)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Write([]byte("[DEBUG]2023/06/10 01:25:55.028277 log_test.go:55:>>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n"))
		}
	})
}

// go自动log
func BenchmarkParallelNativeGoLog(b *testing.B) {
	out, _ := os.OpenFile(`D:\cfoldTest\nativelog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	log := log.New(out, "[debug]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Println(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
	})
}
