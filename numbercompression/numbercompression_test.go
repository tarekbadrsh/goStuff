package numbercompression

import "testing"

//!+test
//go test -v

func TestCompresNubmerUncompresNubmerDefault1000000(t *testing.T) {
	for index := int64(0); index < 1000000; index++ {
		encode := CompresNubmerDefault(index)
		decode := UncompresNubmerDefault(encode)
		if decode != index {
			t.Errorf("TestCompresNubmerUncompresNubmer\nindex:%v\nencode:%v\ndecode%v\n", index, encode, decode)
		}
	}
}

func TestCompresNubmerUncompresNubmerDefault1000005000000(t *testing.T) {
	for index := int64(1000000000000); index < 1000005000000; index++ {
		encode := CompresNubmerDefault(index)
		decode := UncompresNubmerDefault(encode)
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
		encode := CompresNubmerDefault(index)
		decode := UncompresNubmerDefault(encode)
		if decode != index {
			b.Errorf("BenchmarkCompresNubmerUncompresNubmerDefault\nindex:%v\nencode:%v\ndecode%v\n", index, encode, decode)
		}
	}
}

func BenchmarkCompresNubmerUncompresNubmerDefault1000000000(b *testing.B) {
	for index := int64(0); index < int64(b.N); index++ {
		i := index + 1000000000000
		encode := CompresNubmerDefault(i)
		decode := UncompresNubmerDefault(encode)
		if decode != i {
			b.Errorf("BenchmarkCompresNubmerUncompresNubmerDefault\nindex:%v\nencode:%v\ndecode%v\n", i, encode, decode)
		}
	}
}

//!-bench
