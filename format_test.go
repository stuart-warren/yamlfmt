package yamlfmt_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stuart-warren/yamlfmt"
)

func assertGotExpected(t *testing.T, got, exp []byte) {
	if !bytes.Equal(got, exp) {
		t.Fatalf("Got:\n%s\nexpected:\n%s\n", got, exp)
	}
	t.Logf("got:\n%v\n", got)
	t.Logf("expected:\n%v\n", exp)
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}
}

func assertExpectedError(t *testing.T, err error, exp string) {
	if err == nil {
		t.Fatalf("Expected an error: %s\n", err)
	}
	if !strings.Contains(err.Error(), exp) {
		t.Fatalf("Got Error, but error unexpected: %s\n", err)
	}
}

func doTest(t *testing.T, in, exp string) {
	expected := []byte(exp)
	out, err := yamlfmt.Format(bytes.NewReader([]byte(in)))
	assertNoError(t, err)
	assertGotExpected(t, out, expected)
}

func TestYamlSort(t *testing.T) {
	var in = `---
k:
  v:
    c: i
  c: l
`
	var expected = `---
k:
  c: l
  v:
    c: i
`
	doTest(t, in, expected)
}

func TestYamlComplexSort(t *testing.T) {
	var in = `---
k:
  c: l
  v:
    # comment
    c: i
  i:
    i: k
  t:
  - c
  - a
  - b
`
	var expected = `---
k:
  c: l
  i:
    i: k
  t:
  - c
  - a
  - b
  v:
    # comment
    c: i
`
	doTest(t, in, expected)
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
	doTest(t, in, expected)
}

func TestYamlMultiSort(t *testing.T) {
	var in = `---
k:
  c: l
  i:
    i: k
  v:
    # comment
    c: i
---
k:
  c: l
  i:
    i: k
  v:
    # comment
    c: i
`
	var expected = `---
k:
  c: l
  i:
    i: k
  v:
    # comment
    c: i
---
k:
  c: l
  i:
    i: k
  v:
    # comment
    c: i
`
	doTest(t, in, expected)
}

func TestYamlLiteralBlockScalarSort(t *testing.T) {
	var in = `---
b: |
  some content
  other content
a: value
`
	var expected = `---
a: value
b: |
  some content
  other content
`
	doTest(t, in, expected)
}

func TestYamlFoldedBlockScalarSort(t *testing.T) {
	var in = `---
b: >
  some content
  other content

a: value
`
	var expected = `---
a: value
b: >
  some content
  other content
`
	doTest(t, in, expected)
}

func TestBadInput(t *testing.T) {
	var in = `,,,`
	_, err := yamlfmt.Format(bytes.NewReader([]byte(in)))
	assertExpectedError(t, err, "did not find expected node content")
}
