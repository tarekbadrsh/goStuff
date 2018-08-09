//time ./goroutineonlyex -t foo > cpu.pprof
//go tool pprof cpu.pprof
//top
//top -cum
//list hashzeros.GetBestZerosGoroutine.func1

//time ./goroutineonlyex -t foo > goroutineonlyex.trace
//go tool trace goroutineonlyex.trace

package main

import (
	"flag"
	"hashzeros"
	"os"
	"runtime/pprof"
)

var tex = flag.String("t", " ", "text need to hash")
var coun = flag.Int("c", 1000000, "count time hash")

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	// trace.Start(os.Stdout)
	// defer trace.Stop()

	flag.Parse()
	text := *tex
	count := *coun

	_ = hashzeros.GetBestZerosGoroutine(text, count, 1000)
	//fmt.Println(pkgResult)
}
