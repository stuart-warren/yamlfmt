package yamlfmt_test

import (
	"bytes"
	"testing"

	"github.com/stuart-warren/yamlfmt"
)

var in = `---
foo:
  baz: fuzz
  bar: foo
  buz:
    foo: bar
    baz: boo
    biz:
    - 1
    - 2
    - 3
---
bar:
  foo: baz
`
var expected = `---
foo:
  bar: foo
  baz: fuzz
  buz:
    baz: boo
    biz:
    - 1
    - 2
    - 3
    foo: bar
---
bar:
  foo: baz
`

func TestYamlIn(t *testing.T) {
	exp := []byte(expected)
	out, err := yamlfmt.Format(bytes.NewBuffer([]byte(in)))
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}
	if !bytes.Equal(out, exp) {
		t.Fatalf("Got:\n%q\nexpected:\n%q\n", out, exp)
	}
	t.Logf("got:\n%v\n", out)
	t.Logf("expected:\n%v\n", exp)
}
