package WordsFilter

import (
	"sync"
)

type WordsFilter struct {
	placeholder string
	node        *Node
	mutex       sync.Mutex
}

// New creates a filter.
func NewWordsFilter(placeholder string) *WordsFilter {
	return &WordsFilter{
		placeholder: placeholder,
		node:        NewNode(make(map[string]*Node), ""),
	}
}

// Convert sensitive text lists into sensitive word tree nodes
func (wf *WordsFilter) Generate(texts []string) map[string]*Node {
	root := make(map[string]*Node)
	for _, text := range texts {
		wf.Add(text, root)
	}
	return root
}

// Add sensitive words to specified sensitive words Map.
func (wf *WordsFilter) Add(text string, root map[string]*Node) {
	wf.mutex.Lock()
	defer wf.mutex.Unlock()
	wf.node.Add(text, root, wf.placeholder)
}

// Replace sensitive words in strings and return new strings.
func (wf *WordsFilter) Replace(text string, root map[string]*Node) string {
	return wf.node.Replace(text, root)
}

// Whether the string contains sensitive words.
func (wf *WordsFilter) Contains(text string, root map[string]*Node) bool {
	return wf.node.Contains(text, root)
}

// Remove specified sensitive words from sensitive word map.
func (wf *WordsFilter) Remove(text string, root map[string]*Node) {
	wf.mutex.Lock()
	defer wf.mutex.Unlock()
	wf.node.Remove(text, root)
}
