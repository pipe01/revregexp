package revregexp_test

import (
	"fmt"
	"regexp/syntax"

	"github.com/pipe01/revregexp"
)

func ExampleReversePattern() {
	r, _ := syntax.Parse(`\d{5}`, syntax.Perl)

	str, _ := revregexp.ReversePattern(r)

	fmt.Println(str)
	// Outputs: 38572
}

func ExampleReverse() {
	str, _ := revregexp.Reverse(`\d{5}`)

	fmt.Println(str)
	// Outputs: 19347
}
