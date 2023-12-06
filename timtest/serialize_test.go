package timtest

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"

	. "github.com/donnie4w/gofer/util"
	"github.com/vmihailenco/msgpack"
)

func init() {
	go http.ListenAndServe(":9001", nil)
	<-time.After(5 * time.Second)
}

var p = &TimMessageBean{
	MsType:    1,                           //1
	OdType:    1,                           //1
	Id:        1 << 60,                     //8
	Mid:       1 << 61,                     //8
	BnType:    1 << 30,                     //4
	FromTid:   &TidBean{Node: "tom"},       //3
	ToTid:     &TidBean{Node: "jerry"},     //5
	Body:      []byte("A good programmer"), //17
	IsOffline: true,                        //1
	Timestamp: 1 << 60,                     //8    ->    56
	// Extend:    map[string]string{"ex1": "A good programmer", "ex2": "A good programmer"}, //40   ->    96
}

func newTimTFMessage() (tm *TimTFMessage) {
	tm = &TimTFMessage{}
	tm.MsType = p.MsType
	tm.OdType = p.OdType
	tm.ID = &p.Id
	tm.Mid = &p.Mid
	tm.BnType = &p.BnType
	tm.FromTid = &TidTF{Node: p.FromTid.Node}
	tm.ToTid = &TidTF{Node: p.ToTid.Node}
	tm.Body = p.Body
	tm.IsOffline = &p.IsOffline
	tm.Timestamp = &p.Timestamp
	// tm.Extend = p.Extend
	return
}

func newTimPbMessage() (tm *TimPbMessage) {
	i1, i2, i3 := int32(1), int32(2), int32(3)
	tm = &TimPbMessage{}
	tm.MsType = &i1
	tm.OdType = &i2
	tm.Id = &p.Id
	tm.Mid = &p.Mid
	tm.BnType = &i3
	tm.FromTid = &TidPb{Node: &p.FromTid.Node}
	tm.ToTid = &TidPb{Node: &p.ToTid.Node}
	tm.Body = p.Body
	tm.IsOffline = &p.IsOffline
	tm.Timestamp = &p.Timestamp
	// tm.Extend = p.Extend
	return
}

func BenchmarkSerialMsgpackEncode(b *testing.B) {
	b.ResetTimer()
	// var bs []byte
	for i := 0; i < b.N; i++ {
		// bs, _ = msgpack.Marshal(p)
		msgpack.Marshal(p)
	}
	// fmt.Println(len(bs)) //254
}

func BenchmarkSerialMsgpackDecode(b *testing.B) {
	bs, _ := msgpack.Marshal(p)
	var p2 TimMessageBean
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(bs, &p2)
	}
}

func BenchmarkSerialGojsonEncode(b *testing.B) {
	// var bs []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// bs, _ = json.Marshal(p)
		json.Marshal(p)
	}
	// fmt.Println(len(bs)) //366
}

func BenchmarkSerialGoJsonDecode(b *testing.B) {
	bs, _ := json.Marshal(p)
	b.ResetTimer()
	var p2 TimMessageBean
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bs, &p2)
	}
}

func TestTThrift(t *testing.T) {
	tm := newTimTFMessage()
	for i := 0; i < 1000000; i++ {
		TDecode(TEncode(tm), &TimTFMessage{})
	}
	<-time.After(100 * time.Second)
}

func BenchmarkSerialThriftEncode(b *testing.B) {
	tm := newTimTFMessage()
	b.ResetTimer()
	// var bs []byte
	for i := 0; i < b.N; i++ {
		// bs = TEncode(tm)
		TEncode(tm)
	}
	// fmt.Println(len(bs)) //124
}

func BenchmarkSerialThriftDecode(b *testing.B) {
	tm := newTimTFMessage()
	bs := TEncode(tm)
	b.ResetTimer()
	// var tm2 *TimTFMessage
	for i := 0; i < b.N; i++ {
		TDecode(bs, &TimTFMessage{})
		// tm2, _ = TDecode(bs, &TimTFMessage{})
	}
	// fmt.Println(tm2)
}

func BenchmarkSerialProtoBufEncode(b *testing.B) {
	tm := newTimPbMessage()
	// var bs []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// bs, _ = PEncode(tm)
		PEncode(tm)
	}
	// fmt.Println(len(bs)) //125
}

func BenchmarkSerialProtoBufDecode(b *testing.B) {
	tm := newTimPbMessage()
	bs, _ := PEncode(tm)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PDecode(bs, &TimPbMessage{})
	}
}

func BenchmarkParallelMsgpackEncode(b *testing.B) {
	// var bs []byte
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			msgpack.Marshal(p)
		}
	})
	// fmt.Println(string(bs))
}

func BenchmarkParallelMsgpackDecode(b *testing.B) {
	bs, _ := msgpack.Marshal(p)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var p2 TimMessageBean
			msgpack.Unmarshal(bs, &p2)
		}
	})
	// fmt.Println(p2)
}

func BenchmarkParallelGojsonEncode(b *testing.B) {
	// var bs []byte
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			json.Marshal(p)
		}
	})
	// fmt.Println(string(bs))
}

func BenchmarkParallelGoJsonDecode(b *testing.B) {
	bs, _ := json.Marshal(p)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var p2 TimMessageBean
			json.Unmarshal(bs, &p2)
		}
	})
	// fmt.Println(p2)
}

func BenchmarkParallelThriftEncode(b *testing.B) {
	tm := newTimTFMessage()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TEncode(tm)
		}
	})
}

func BenchmarkParallelThriftDecode(b *testing.B) {
	tm := newTimTFMessage()
	bs := TEncode(tm)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if b, err := TDecode(bs, &TimTFMessage{}); err != nil {
				fmt.Println(err)
				fmt.Println(b)
				break
			}
		}
	})
}

func BenchmarkParallelProtoBufEncode(b *testing.B) {
	tm := newTimPbMessage()
	// var bs []byte
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			PEncode(tm)
		}
	})
	// fmt.Println(len(bs))
}

func BenchmarkParallelProtoBufDecode(b *testing.B) {
	tm := newTimPbMessage()
	bs, _ := PEncode(tm)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			PDecode(bs, &TimPbMessage{})
		}
	})
}
