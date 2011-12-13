package lisp

import (
	"strings"
	"testing"
)

type readTokenTest struct{
	str string
	tok token
}

var readTokenTests = []readTokenTest{
	{"1\n", "1"},
	{"32.5a\n", "32.5a"},
	{"32.5 a\n", "32.5"},
	{"\"b\"\n", "\"b\""},
	{" \t5;6", "5"},
	{"(", "("},
	{")(", ")"},
}

func TestReadToken(t *testing.T) {
	for _, test := range readTokenTests {
		r := strings.NewReader(test.str)
		tok, err := readToken(r)
		if err != nil {
			t.Errorf("ERROR: %s", err.Error())
		}
		if tok != test.tok {
			t.Errorf("%s != %s", tok, test.tok)
		}
	}
}