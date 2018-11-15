package model

import (
	"hash/fnv"
	"path"
	"sort"
	"strings"
)

type (
	// Expectation is the type for a request expectation
	Expectation interface {
		GetHash() string
	}

	// Pair is a name-value pair
	Pair struct {
		Name  string `json:"name" yaml:"name"`
		Value string `json:"value" yaml:"value"`
	}

	// HTTPExpectation is the model for an http request expectation
	HTTPExpectation struct {
		Methods []string `json:"methods" yaml:"methods"`
		Path    string   `json:"path" yaml:"path"`
		Queries []Pair   `json:"queries" yaml:"queries"`
		Headers []Pair   `json:"headers" yaml:"headers"`
	}
)

// Canonical returns the canonical form of an http expectation
func (e HTTPExpectation) canonical() HTTPExpectation {
	if e.Methods == nil {
		e.Methods = []string{}
	}

	sort.Strings(e.Methods)
	for i := range e.Methods {
		e.Methods[i] = strings.ToUpper(e.Methods[i])
	}

	e.Path = "/" + e.Path
	e.Path = path.Clean(e.Path)

	if e.Queries == nil {
		e.Queries = []Pair{}
	}

	sort.Slice(e.Queries, func(i, j int) bool {
		return e.Queries[i].Name < e.Queries[j].Name
	})

	if e.Headers == nil {
		e.Headers = []Pair{}
	}

	sort.Slice(e.Headers, func(i, j int) bool {
		return e.Headers[i].Name < e.Headers[j].Name
	})

	return e
}

// GetHash returns a hash for an http expectation
func (e HTTPExpectation) GetHash() uint64 {
	e = e.canonical()
	hash := fnv.New64a()

	for _, m := range e.Methods {
		hash.Write([]byte(m))
	}

	hash.Write([]byte(e.Path))

	for _, q := range e.Queries {
		hash.Write([]byte(q.Name))
		hash.Write([]byte(q.Value))
	}

	for _, h := range e.Headers {
		hash.Write([]byte(h.Name))
		hash.Write([]byte(h.Value))
	}

	return hash.Sum64()
}
