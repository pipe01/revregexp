package revregexp

import (
	"fmt"
	"math/rand"
	"regexp/syntax"
)

type StringWriter interface {
	WriteString(str string) (int, error)
	WriteRune(r rune) (int, error)
}

type RegexpWriter struct {
	MaxQuantifierLength int // The maximum length for open-ended quantifiers (*, ?, etc)
}

// Default is an instance of RegexpWriter with the default configurtion
var Default = &RegexpWriter{
	MaxQuantifierLength: 5,
}

var anyCharNotNewline = []rune{'0', '9', 'A', 'Z', 'a', 'z'}

func (w *RegexpWriter) Write(r *syntax.Regexp, wr StringWriter) error {
	var err error

	switch r.Op {
	case syntax.OpConcat:
		for _, sub := range r.Sub {
			if err := w.Write(sub, wr); err != nil {
				return err
			}
		}

	case syntax.OpLiteral:
		_, err = wr.WriteString(string(r.Rune))

	case syntax.OpCharClass:
		err = w.writeCharClass(r.Rune, wr)

	case syntax.OpPlus:
		err = w.writeTimes(r.Sub[0], 1, -1, wr)

	case syntax.OpStar:
		err = w.writeTimes(r.Sub[0], 0, -1, wr)

	case syntax.OpRepeat:
		err = w.writeTimes(r.Sub[0], r.Min, r.Max, wr)

	case syntax.OpCapture:
		err = w.Write(r.Sub[0], wr)

	case syntax.OpAnyCharNotNL:
		err = w.writeCharClass(anyCharNotNewline, wr)

	case syntax.OpAlternate:
		err = w.writeAny(r.Sub, wr)

	default:
		err = fmt.Errorf("unknown syntax operator %d", r.Op)
	}

	return err
}

func (w *RegexpWriter) writeCharClass(r []rune, wr StringWriter) error {
	paircnt := len(r) / 2
	pairidx := rand.Intn(paircnt)

	return w.writeCharRange(r[pairidx*2], r[pairidx*2+1], wr)
}

func (w *RegexpWriter) writeCharRange(min, max rune, wr StringWriter) error {
	offset := rand.Intn(int(max - min + 1))
	rune := min + rune(offset)

	_, err := wr.WriteRune(rune)
	return err
}

func (w *RegexpWriter) writeAny(sub []*syntax.Regexp, wr StringWriter) error {
	idx := rand.Intn(len(sub))

	return w.Write(sub[idx], wr)
}

func (w *RegexpWriter) writeTimes(r *syntax.Regexp, min, max int, wr StringWriter) error {
	if max == -1 {
		max = w.MaxQuantifierLength
	}

	var n int

	if max == min {
		n = min
	} else {
		n = rand.Intn(max-min) + min
	}

	for i := 0; i < n; i++ {
		if err := w.Write(r, wr); err != nil {
			return err
		}
	}

	return nil
}
