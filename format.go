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
	var docs []interface{}
	d := yaml.NewDecoder(r)
	for {
		var doc interface{}
		err := d.Decode(&doc)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed decoding: %s", err)
		}
		docs = append(docs, doc)
	}

	out := bytes.NewBuffer(nil)
	for _, doc := range docs {
		e := yaml.NewEncoder(out)
		e.SetIndent(indent)
		defer e.Close()
		out.WriteString("---\n")
		if err := e.Encode(&doc); err != nil {
			return nil, fmt.Errorf("failed encoding: %s", err)
		}
	}
	return out.Bytes(), nil
}
