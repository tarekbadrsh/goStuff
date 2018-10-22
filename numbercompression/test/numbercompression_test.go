package numbercompression_test

import (
	"testing"

	"github.com/tarekbadrshalaan/goStuff/numbercompression"
)

//!+test
//go test -v

func TestCompresNubmerUncompresNumberDefault1000000(t *testing.T) {
	for index := int64(0); index < 1000000; index++ {
		encode := numbercompression.CompresNumberDefault(index)
		decode := numbercompression.UncompresNumberDefault(encode)
		if decode != index {
			t.Errorf("TestCompresNubmerUncompresNubmer\nindex:%v\nencode:%v\ndecode%v\n", index, encode, decode)
		}
	}
}

func TestCompresNubmerUncompresNumberDefault1000005000000(t *testing.T) {
	for index := int64(1000000000000); index < 1000005000000; index++ {
		encode := numbercompression.CompresNumberDefault(index)
		decode := numbercompression.UncompresNumberDefault(encode)
		if decode != index {
			t.Errorf("TestCompresNubmerUncompresNubmer\nindex:%v\nencode:%v\ndecode%v\n", index, encode, decode)
		}
	}
}

//!-tests

//!+bench
//go test -v  -bench=.
func BenchmarkCompresNubmerUncompresNubmerDefault(b *testing.B) {
	for index := int64(0); index < int64(b.N); index++ {
		encode := numbercompression.CompresNumberDefault(index)
		decode := numbercompression.UncompresNumberDefault(encode)
		if decode != index {
			b.Errorf("BenchmarkCompresNubmerUncompresNumberDefault\nindex:%v\nencode:%v\ndecode%v\n", index, encode, decode)
		}
	}
}

func BenchmarkCompresNubmerUncompresNumberDefault1000000000(b *testing.B) {
	for index := int64(0); index < int64(b.N); index++ {
		i := index + 1000000000000
		encode := numbercompression.CompresNumberDefault(i)
		decode := numbercompression.UncompresNumberDefault(encode)
		if decode != i {
			b.Errorf("BenchmarkCompresNubmerUncompresNumberDefault\nindex:%v\nencode:%v\ndecode%v\n", i, encode, decode)
		}
	}
}

//!-bench
