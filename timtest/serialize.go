package timtest

import (
	"google.golang.org/protobuf/proto"
)

func PEncode(m proto.Message) ([]byte, error) {
	return proto.Marshal(m)
}

func PDecode(bs []byte, m proto.Message) {
	proto.Unmarshal(bs, m)
}
