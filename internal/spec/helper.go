package spec

import (
	"errors"
	"fmt"
	"hash"
	"sort"
	"strconv"
)

// JSON is the type for json objects.
type JSON map[string]interface{}

var defaultIdentifiers = []string{"id", "Id", "ID", "_id"}

func findID(key string, obj JSON) (interface{}, error) {
	if key != "" {
		if val, ok := obj[key]; ok {
			return val, nil
		}
		return nil, fmt.Errorf("identifier %q does not exist", key)
	}

	for _, key = range defaultIdentifiers {
		if val, ok := obj[key]; ok {
			return val, nil
		}
	}

	return nil, errors.New("cannot find an identifier")
}

// Pair is a key-value pair
type Pair struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

func hashBool(h hash.Hash, args ...bool) {
	for _, b := range args {
		h.Write([]byte(strconv.FormatBool(b)))
	}
}

func hashString(h hash.Hash, args ...string) {
	for _, s := range args {
		h.Write([]byte(s))
	}
}

func hashStringSlice(h hash.Hash, canonical bool, s []string) {
	buff := make([]string, len(s))
	for i := range s {
		buff[i] = s[i]
	}

	if canonical {
		sort.Strings(buff)
	}

	for _, str := range buff {
		h.Write([]byte(str))
	}
}

func hashStringMap(h hash.Hash, canonical bool, m map[string]string) {
	buff := make([]Pair, 0)
	for key, value := range m {
		buff = append(buff, Pair{
			Key:   key,
			Value: value,
		})
	}

	if canonical {
		sort.Slice(buff, func(i, j int) bool {
			return buff[i].Key < buff[j].Key
		})
	}

	for _, p := range buff {
		h.Write([]byte(p.Key))
		h.Write([]byte(p.Value))
	}
}
