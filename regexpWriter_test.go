package revregexp

import (
	"regexp"
	"regexp/syntax"
	"strings"
	"testing"
)

func TestReciprocity(t *testing.T) {
	rules := []string{
		`nice`,
		`\d`,
		`\d+`,
		`\d{0,6}`,
		`\d{0,}`,
		`[A-Z]`,
		`[A-Z]+`,
		`[A-Z]*`,
		`\dfoo`,
		`\d+foo`,
		`\d{0,6}foo`,
		`\d{0,}foo`,
		`[A-Z]foo`,
		`[A-Z]+foo`,
		`[A-Z]*foo`,
	}

	for _, r := range rules {
		pat, _ := syntax.Parse(r, syntax.Perl)
		rexp := regexp.MustCompile(r)

		t.Run(r, func(t *testing.T) {
			t.Parallel()

			wr := &strings.Builder{}

			for i := 0; i < 100; i++ {
				wr.Reset()

				err := Default.Write(pat, wr)
				if err != nil {
					t.Fatalf("failed to reverse: %s", err)
				}

				if !rexp.MatchString(wr.String()) {
					t.Fatal("generated string does not match pattern")
				}
			}
		})
	}
}
