package logtest

import (
	"log"
	"log/slog"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"testing"

	logger "github.com/donnie4w/go-logger/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initzap() *zap.Logger {
	writeSyncer, _ := os.OpenFile(`zaplog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666) //日志文件存放目录
	encoderConfig := zap.NewProductionEncoderConfig()                                    //指定时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	log := zap.New(core, zap.AddCaller())
	return log
}

// go 自带的log库
func Benchmark_Serial_NativeLog(b *testing.B) {
	out, _ := os.OpenFile(`nativelog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	log := log.New(out, "[debug]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Println(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

// Zap 常规打印
func Benchmark_Serial_Zap(b *testing.B) {
	log := initzap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

// go-logger 常规打印
func Benchmark_Serial_GoLogger(b *testing.B) {
	log := logger.NewLogger().SetOption(&logger.Option{Console: false, FileOption: &logger.FileSizeMode{Filename: "golog.txt", Maxsize: 1 << 30}})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

// go 新增的slog库
func Benchmark_Serial_Slog(b *testing.B) {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	out, _ := os.OpenFile(`slog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	h := slog.NewTextHandler(out, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace})
	log := slog.New(h)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info(">>>aaaaaaaaaaaaaaaaaaaaaa")
	}
}

// slog and go-logger
func Benchmark_Serial_SlogAndGoLogger(b *testing.B) {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	loggingFile, _ := logger.NewLogger().SetRollingFile(``, `slogLogger.txt`, 1, logger.GB)
	h := slog.NewTextHandler(loggingFile, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace})
	log := slog.New(h)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info(">>>aaaaaaaaaaaaaaaaaaaaaa")
	}
}

// go自带log库
func Benchmark_Parallel_NativeLog(b *testing.B) {
	out, _ := os.OpenFile(`nativelog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	log := log.New(out, "[debug]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Println(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
	})
}

// Zap 常规打印
func Benchmark_Parallel_Zap(b *testing.B) {
	log := initzap()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
	})

}

// go-logger 常规打印
func Benchmark_Parallel_GoLogger(b *testing.B) {
	// b.SetParallelism(8)
	log := logger.NewLogger().SetOption(&logger.Option{Console: false, FileOption: &logger.FileSizeMode{Filename: "golog.txt", Maxsize: 1 << 30}})
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug(">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
	})
}

func Benchmark_Parallel_SLog(b *testing.B) {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	out, _ := os.OpenFile(`slog.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	h := slog.NewTextHandler(out, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace})
	log := slog.New(h)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info(">>>aaaaaaaaaaaaaaaaaaaaaa")
		}
	})
}

func Benchmark_Parallel_SLogAndGoLogger(b *testing.B) {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	loggingFile, _ := logger.NewLogger().SetRollingFile(``, `slogLogger.txt`, 1, logger.GB)
	h := slog.NewTextHandler(loggingFile, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace})
	log := slog.New(h)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info(">>>aaaaaaaaaaaaaaaaaaaaaa")
		}
	})
}
