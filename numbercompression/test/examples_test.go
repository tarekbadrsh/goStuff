package numbercompression_test

import (
	"fmt"

	"github.com/tarekbadrshalaan/goStuff/numbercompression"
)

func ExampleCompresNumberDefault() {
	var num int64 = 123456789123456789
	res := numbercompression.CompresNumberDefault(num)
	fmt.Printf("Compres Number %d is %v", num, res)
	// Output:
	// Compres Number 123456789123456789 is V8F0su0m2G
}

func ExampleUncompresNumberDefault() {
	input := "V8F0su0m2G"
	res := numbercompression.UncompresNumberDefault(input)
	fmt.Printf("Uncompress string to number %v is %d", input, res)
	// Output:
	// Uncompress string to number V8F0su0m2G is 123456789123456789
}
