package logx

//这是raftx的使用示例，主要利用raftx支持易失性数据的功能特性进行分布式数据同步
//该示例展示如何简单的开发一个分布式日志同步库
//该示例中，有3个节点，即3个分布于不同服务的写日志相互同步操作
//无论哪个节点写日志，都会通过raftx快速同步至其他两个节点，并且执行顺序完全一致
//最后校验日志文件的数据是否完全一致
//注意，该示例主要目的在于raftx的部分功能展示，并不考虑分布式日志库的功能完整性与严谨性
//易失性数据本身的数据特点可能不适合作为需要严谨数据的同步机制，它更适合用于一些标识状态，或统计等需要不断更新的数据

import (
	"github.com/donnie4w/go-logger/logger"
	"github.com/donnie4w/raftx"
	"github.com/donnie4w/raftx/raft"
)

type Logx struct {
	log   *logger.Logging
	raftx raftx.Raftx
}

func NewLogx(filePath string, listen string, peers []string) (r *Logx) {
	log := logger.NewLogger().SetOption(&logger.Option{FileOption: &logger.FileSizeMode{Filename: filePath, Maxsize: 1 << 30}, Format: logger.FORMAT_DATE | logger.FORMAT_SHORTFILENAME})
	rx := raftx.NewRaftx(&raft.Config{ListenAddr: listen, PeerAddr: peers})
	go rx.Open() //启动raftx服务
	r = &Logx{log: log, raftx: rx}
	go func() {
		rx.WaitRun() //等待raftx集群可正常服务
		r.init()
	}()
	return
}

func (lx *Logx) init() {
	//注册Debug的事件监听
	lx.raftx.MemWatch([]byte{0}, func(key, value []byte, watchType raft.WatchType) {
		lx.log.Debug(string(value))
	}, true, raft.ADD, raft.UPDATE)

	//注册Info的事件监听
	lx.raftx.MemWatch([]byte{1}, func(key, value []byte, watchType raft.WatchType) {
		lx.log.Info(string(value))
	}, true, raft.ADD, raft.UPDATE)

	//注册Warn的事件监听
	lx.raftx.MemWatch([]byte{2}, func(key, value []byte, watchType raft.WatchType) {
		lx.log.Warn(string(value))
	}, true, raft.ADD, raft.UPDATE)

	//注册Error的事件监听
	lx.raftx.MemWatch([]byte{3}, func(key, value []byte, watchType raft.WatchType) {
		lx.log.Error(string(value))
	}, true, raft.ADD, raft.UPDATE)
}

func (lx *Logx) Debug(value []byte) error {
	return lx.raftx.MemCommand([]byte{0}, value, 0, raft.MEM_PUT)
}

func (lx *Logx) Info(value []byte) error {
	return lx.raftx.MemCommand([]byte{1}, value, 0, raft.MEM_PUT)
}

func (lx *Logx) Warn(value []byte) error {
	return lx.raftx.MemCommand([]byte{2}, value, 0, raft.MEM_PUT)
}

func (lx *Logx) Error(value []byte) error {
	return lx.raftx.MemCommand([]byte{3}, value, 0, raft.MEM_PUT)
}
