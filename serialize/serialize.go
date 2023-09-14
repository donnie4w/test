package serialize

import (
	"bytes"
	"context"

	// thrift "github.com/apache/thrift/lib/go/thrift"

	"github.com/donnie4w/gothrift/thrift"
	"google.golang.org/protobuf/proto"
)

var tconf = &thrift.TConfiguration{}

func TEncode(ts thrift.TStruct) (_r []byte) {
	buf := &thrift.TMemoryBuffer{Buffer: bytes.NewBuffer([]byte{})}
	protocol := thrift.NewTCompactProtocolConf(buf, tconf)
	ts.Write(context.Background(), protocol)
	protocol.Flush(context.Background())
	_r = buf.Bytes()
	return
}

func TDecode[T thrift.TStruct](bs []byte, ts T) (_r T, err error) {
	buf := &thrift.TMemoryBuffer{Buffer: bytes.NewBuffer(bs)}
	protocol := thrift.NewTCompactProtocolConf(buf, tconf)
	err = ts.Read(context.Background(), protocol)
	return ts, err
}

func PEncode(m proto.Message) ([]byte, error) {
	return proto.Marshal(m)
}

func PDecode(bs []byte, m proto.Message) {
	proto.Unmarshal(bs, m)
}
