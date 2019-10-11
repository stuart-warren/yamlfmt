package yamlfmt

import (
	"bytes"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

const indent = 2

// Format reads in a yaml document and outputs the yaml in a standard format.
// Dictionary keys are sorted lexicagraphically
// Indents are set to 2
// Lists are not indented
func Format(r io.Reader) ([]byte, error) {
	dec := yaml.NewDecoder(r)
	out := bytes.NewBuffer(nil)
	for {
		enc := yaml.NewEncoder(out)
		enc.SetIndent(indent)
		defer enc.Close()
		var doc yaml.Node
		err := dec.Decode(&doc)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed decoding: %s", err)
		}
		out.WriteString("---\n")
		if err := enc.Encode(sortYAML(&doc)); err != nil {
			return nil, fmt.Errorf("failed encoding: %s", err)
		}
		enc.Close()
	}
	return out.Bytes(), nil
}
