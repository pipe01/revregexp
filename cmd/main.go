package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp/syntax"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/pipe01/revregexp"
)

func main() {
	count := flag.Int("count", 1, "How many strings to generate, separated by newlines")

	var opts struct {
		Count int `short:"c" long:"count" description:"How many strings to print, separated by newlines"`
	}

	args, err := flags.Parse(&opts)
	if err != nil {
		fmt.Printf("failed to parse flags: %s", err)
		os.Exit(1)
	}

	if len(args) != 1 {
		fmt.Println("missing pattern argument")
		os.Exit(1)
	}

	r, err := syntax.Parse(args[0], syntax.Perl)
	if err != nil {
		fmt.Printf("failed to parse pattern: %s\n", err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	str := &strings.Builder{}
	for i := 0; i < *count; i++ {
		str.Reset()
		revregexp.Default.Write(r, str)

		fmt.Println(str.String())
	}
}
