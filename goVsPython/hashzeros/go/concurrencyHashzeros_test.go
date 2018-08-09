package hashzeros

import (
	"sync"
	"testing"
)

//!+test

func TestConcurrencyGetBestZeros(t *testing.T) {
	type Test struct {
		text          string
		nonceLimit    int
		oprationLimit int
		want          map[int]bool
	}
	var tests = []Test{
		{"test", 1000, 100, map[int]bool{338: true, 304: true, 825: true, 849: true}},
		{"test", 100000, 1000, map[int]bool{93721: true}},
		{"test", 1000000, 10000,
			map[int]bool{93721: true, 195035: true,
				286531: true, 291889: true, 323790: true,
				401718: true, 839016: true, 845012: true,
				853118: true, 906372: true, 949642: true}},

		{"foo", 1000, 100, map[int]bool{78: true, 270: true, 417: true, 930: true}},
		{"foo", 10000000, 10000, map[int]bool{4970608: true}},
	}

	var wg sync.WaitGroup
	for i, test := range tests {
		wg.Add(1)
		go func(i int, test Test, wg *sync.WaitGroup) {
			defer wg.Done()
			got := ConcurrencyGetBestZeros(test.text, test.nonceLimit, test.oprationLimit)

			if _, ok := test.want[got.Nonce]; !ok {
				t.Errorf("index : %d ; GetBestZeros(%v,%v,%v) \nexpected:%v\ngot:%v\n",
					i, test.text, test.nonceLimit, test.oprationLimit, test.want, got)
			}
		}(i, test, &wg)
	}
	wg.Wait()

}

//!-tests

//!+bench

func BenchmarkConcurrencyGetBestZeros(b *testing.B) {
	ConcurrencyGetBestZeros("test", b.N, 1000)
}

//!-bench
