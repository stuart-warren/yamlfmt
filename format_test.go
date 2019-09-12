package yamlfmt_test

import (
	"bytes"
	"testing"

	"github.com/stuart-warren/yamlfmt"
)

func TestYamlSort(t *testing.T) {
	var in = `---
foo:
  buz: foo
  bar:
    foo: bar
    biz:
    - 1
`
	var expected = `---
foo:
  bar:
    biz:
    - 1
    foo: bar
  buz: foo
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

func TestCommentedYaml(t *testing.T) {
	var in = `---
bar:
  foo: baz # comment
  boo: fizz
`
	var expected = `---
bar:
  boo: fizz
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
