package hashzeros

import (
	"testing"
)

//!+test

func TestGetBestZeros(t *testing.T) {
	var test = struct {
		text          string
		nonceLimit    int
		oprationLimit int
		want          map[int]bool
	}{
		"test", 1000000, 10000,
		map[int]bool{93721: true, 195035: true,
			286531: true, 291889: true, 323790: true,
			401718: true, 839016: true, 845012: true,
			853118: true, 906372: true, 949642: true}}

	got := GetBestZeros(test.text, test.nonceLimit, test.oprationLimit)
	if _, ok := test.want[got.Nonce]; !ok {
		t.Errorf("GetBestZeros(%v,%v,%v) \nexpected:%v\ngot:%v\n", test.text, test.nonceLimit, test.oprationLimit, test.want, got)
	}
}

//!-tests

//!+bench
func BenchmarkGetBestZeros(b *testing.B) {
	GetBestZeros("test", b.N, 1000)
}

//!-bench
