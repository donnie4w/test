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
	tcf := thrift.NewTCompactProtocolFactoryConf(tconf)
	tp := tcf.GetProtocol(buf)
	ts.Write(context.Background(), tp)
	_r = buf.Bytes()
	return
}

func TDecode[T thrift.TStruct](bs []byte, ts T) (_r T, err error) {
	buf := &thrift.TMemoryBuffer{Buffer: bytes.NewBuffer(bs)}
	tcf := thrift.NewTCompactProtocolFactoryConf(tconf)
	tp := tcf.GetProtocol(buf)
	err = ts.Read(context.Background(), tp)
	return ts, err
}

func PEncode(m proto.Message) ([]byte, error) {
	return proto.Marshal(m)
}

func PDecode(bs []byte, m proto.Message) {
	proto.Unmarshal(bs, m)
}
