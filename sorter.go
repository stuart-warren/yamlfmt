package yamlfmt

import (
	"sort"

	"gopkg.in/yaml.v3"
)

type nodes []*yaml.Node

func (i nodes) Len() int { return len(i) / 2 }

func (i nodes) Swap(x, y int) {
	x *= 2
	y *= 2
	i[x], i[y] = i[y], i[x]         // keys
	i[x+1], i[y+1] = i[y+1], i[x+1] // values
}

func (i nodes) Less(x, y int) bool {
	x *= 2
	y *= 2
	return i[x].Value < i[y].Value
}

func sortYAML(node *yaml.Node) *yaml.Node {
	if node.Kind == yaml.DocumentNode {
		for i, n := range node.Content {
			node.Content[i] = sortYAML(n)
		}
	}
	if node.Kind == yaml.SequenceNode {
		for i, n := range node.Content {
			node.Content[i] = sortYAML(n)
		}
	}
	if node.Kind == yaml.MappingNode {
		for i, n := range node.Content {
			node.Content[i] = sortYAML(n)
		}
		sort.Sort(nodes(node.Content))
	}
	return node
}
