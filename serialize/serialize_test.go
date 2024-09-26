package serialize

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/donnie4w/gofer/util"
	"github.com/vmihailenco/msgpack"
)

type TestBean struct {
	Name    string
	Int64   int64
	Int32   int32
	Int16   int16
	Int8    int8
	ListVal [][]byte
	MapVal  map[string][]byte
}

func BenchmarkMsgpackEncode(b *testing.B) {
	p1 := TestBean{
		Name:    "i哈哈aaaaaa12345",
		Int64:   1 << 60,
		Int32:   1 << 30,
		Int16:   1 << 14,
		Int8:    1 << 6,
		ListVal: [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:  map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	// var bs []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(p1)
	}
	// fmt.Println(string(bs))
}

func BenchmarkMsgpackDecode(b *testing.B) {
	p1 := TestBean{
		Name:    "i哈哈aaaaaa12345",
		Int64:   1 << 60,
		Int32:   1 << 30,
		Int16:   1 << 14,
		Int8:    1 << 6,
		ListVal: [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:  map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	bs, _ := msgpack.Marshal(p1)
	var p2 TestBean
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(bs, &p2)
	}
	// fmt.Println(p2)
}

func BenchmarkGojsonEncode(b *testing.B) {
	p1 := TestBean{
		Name:    "i哈哈aaaaaa12345",
		Int64:   1 << 60,
		Int32:   1 << 30,
		Int16:   1 << 14,
		Int8:    1 << 6,
		ListVal: [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:  map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	// var bs []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(p1)
	}
	// fmt.Println(string(bs))
}

func BenchmarkGoJsonDecode(b *testing.B) {
	p1 := TestBean{
		Name:    "i哈哈aaaaaa12345",
		Int64:   1 << 60,
		Int32:   1 << 30,
		Int16:   1 << 14,
		Int8:    1 << 6,
		ListVal: [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:  map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	bs, _ := json.Marshal(p1)
	b.ResetTimer()
	var p2 TestBean
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bs, &p2)
	}
	// fmt.Println(p2)
}

func BenchmarkThriftEncode(b *testing.B) {
	bean := &Bean{Name: "i哈哈aaaaaa12345", LongVal: 1 << 60, IntVal: 1 << 30, Int16Val: 1 << 14, Int8Val: 1 << 6}
	bean.ListVal = [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")}
	bean.MapVal = map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")}
	// var bs []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TEncode(bean)
	}
	// fmt.Println(len(bs))
}

func BenchmarkThriftDecode(b *testing.B) {
	bean := &Bean{Name: "i哈哈aaaaaa12345", LongVal: 1 << 60, IntVal: 1 << 30, Int16Val: 1 << 14, Int8Val: 1 << 6}
	bean.ListVal = [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")}
	bean.MapVal = map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")}
	bs := TEncode(bean)
	b.ResetTimer()
	// var _tm *Bean
	for i := 0; i < b.N; i++ {
		TDecode(bs, &Bean{})
	}
	// fmt.Println(_tm)
}

func BenchmarkProtoBufEncode(b *testing.B) {
	name := "i哈哈aaaaaa12345"
	i64, i32 := int64(1<<60), int32(1<<30)
	p1 := &PBean{
		Name:     &name,
		Int64Val: &i64,
		Int32Val: &i32,
		ListVal:  [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:   map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	// var bs []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PEncode(p1)
	}
	// fmt.Println(len(bs))
}

func BenchmarkProtoBufDecode(b *testing.B) {
	name := "i哈哈aaaaaa12345"
	i64, i32 := int64(1<<60), int32(1<<30)
	p1 := &PBean{
		Name:     &name,
		Int64Val: &i64,
		Int32Val: &i32,
		ListVal:  [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:   map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	bs, _ := PEncode(p1)
	p := &PBean{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PDecode(bs, p)
	}
}

func BenchmarkParallelMsgpackEncode(b *testing.B) {
	p1 := TestBean{
		Name:    "i哈哈aaaaaa12345",
		Int64:   1 << 60,
		Int32:   1 << 30,
		Int16:   1 << 14,
		Int8:    1 << 6,
		ListVal: [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:  map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	// var bs []byte
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			msgpack.Marshal(p1)
		}
	})
	// fmt.Println(string(bs))
}

func BenchmarkParallelMsgpackDecode(b *testing.B) {
	p1 := TestBean{
		Name:    "i哈哈aaaaaa12345",
		Int64:   1 << 60,
		Int32:   1 << 30,
		Int16:   1 << 14,
		Int8:    1 << 6,
		ListVal: [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:  map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	bs, _ := msgpack.Marshal(p1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var p2 TestBean
			msgpack.Unmarshal(bs, &p2)
		}
	})
	// fmt.Println(p2)
}

func BenchmarkParallelGojsonEncode(b *testing.B) {
	p1 := TestBean{
		Name:    "i哈哈aaaaaa12345",
		Int64:   1 << 60,
		Int32:   1 << 30,
		Int16:   1 << 14,
		Int8:    1 << 6,
		ListVal: [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:  map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	// var bs []byte
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			json.Marshal(p1)
		}
	})
	// fmt.Println(string(bs))
}

func BenchmarkParallelGoJsonDecode(b *testing.B) {
	p1 := TestBean{
		Name:    "i哈哈aaaaaa12345",
		Int64:   1 << 60,
		Int32:   1 << 30,
		Int16:   1 << 14,
		Int8:    1 << 6,
		ListVal: [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:  map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	bs, _ := json.Marshal(p1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var p2 TestBean
			json.Unmarshal(bs, &p2)
		}
	})
	// fmt.Println(p2)
}

func BenchmarkParallelThriftEncode(b *testing.B) {
	bean := &Bean{Name: "i哈哈aaaaaa12345", LongVal: 1 << 60, IntVal: 1 << 30, Int16Val: 1 << 14, Int8Val: 1 << 6}
	bean.ListVal = [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")}
	bean.MapVal = map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TEncode(bean)
		}
	})
}

func BenchmarkParallelThriftDecode(b *testing.B) {
	bean := &Bean{Name: "i哈哈aaaaaa12345", LongVal: 1 << 60, IntVal: 1 << 30, Int16Val: 1 << 14, Int8Val: 1 << 6}
	bean.ListVal = [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")}
	bean.MapVal = map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")}
	bs := TEncode(bean)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if b, err := TDecode(bs, &Bean{}); err != nil {
				fmt.Println(err)
				fmt.Println(b)
				break
			}
		}
	})
}

func BenchmarkParallelProtoBufEncode(b *testing.B) {
	name := "i哈哈aaaaaa12345"
	i64, i32 := int64(1<<60), int32(1<<30)
	p1 := &PBean{
		Name:     &name,
		Int64Val: &i64,
		Int32Val: &i32,
		ListVal:  [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:   map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	// var bs []byte
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			PEncode(p1)
		}
	})
	// fmt.Println(len(bs))
}

func BenchmarkParallelProtoBufDecode(b *testing.B) {
	name := "i哈哈aaaaaa12345"
	i64, i32 := int64(1<<60), int32(1<<30)
	p1 := &PBean{
		Name:     &name,
		Int64Val: &i64,
		Int32Val: &i32,
		ListVal:  [][]byte{[]byte("i哈哈aaaaaa12345"), []byte("i哈哈bbbbbb12345"), []byte("i哈哈ccccccc12345")},
		MapVal:   map[string][]byte{"aaaa": []byte("i哈哈aaaaaa123456"), "bbb": []byte("i哈哈bbbbbb123456"), "ccc": []byte("i哈哈ccccccc123456")},
	}
	bs, _ := PEncode(p1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			PDecode(bs, &PBean{})
		}
	})
}
