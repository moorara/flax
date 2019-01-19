package v1

import (
	"errors"
	"fmt"
	"hash"
	"sort"
	"strconv"
)

var (
	defaultIdentifiers = []string{"id", "Id", "ID", "_id"}
)

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

func hashArray(h hash.Hash, canonical bool, a []string) {
	buff := make([]string, len(a))
	for i := range a {
		buff[i] = a[i]
	}

	if canonical {
		sort.Strings(buff)
	}

	for _, str := range buff {
		h.Write([]byte(str))
	}
}

func hashMap(h hash.Hash, canonical bool, m map[string]string) {
	buff := make([]Pair, 0)
	for name, value := range m {
		buff = append(buff, Pair{
			Name:  name,
			Value: value,
		})
	}

	if canonical {
		sort.Slice(buff, func(i, j int) bool {
			return buff[i].Name < buff[j].Name
		})
	}

	for _, p := range buff {
		h.Write([]byte(p.Name))
		h.Write([]byte(p.Value))
	}
}

func findID(idProp string, object JSON) (interface{}, error) {
	if idProp != "" {
		if val, ok := object[idProp]; ok {
			return val, nil
		}

		return nil, fmt.Errorf("identifier %s does not exist", idProp)
	}

	for _, idProp = range defaultIdentifiers {
		if val, ok := object[idProp]; ok {
			return val, nil
		}
	}

	return nil, errors.New("cannot find an identifier")
}
