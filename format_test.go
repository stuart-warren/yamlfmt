package yamlfmt_test

import (
	"bytes"
	"testing"

	"github.com/stuart-warren/yamlfmt"
)

func TestYamlIn(t *testing.T) {
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

func TestCommentedYamlIn(t *testing.T) {
	var in = `---
bar:
  foo: baz # comment
`
	var expected = `---
bar:
  foo: baz # comment
`
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
