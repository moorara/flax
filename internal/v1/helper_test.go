package v1

import (
	"hash/fnv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashBool(t *testing.T) {
	tests := []struct {
		first  []bool
		second []bool
		equal  bool
	}{
		{
			first:  []bool{false},
			second: []bool{false},
			equal:  true,
		},
		{
			first:  []bool{true},
			second: []bool{true},
			equal:  true,
		},
		{
			first:  []bool{false},
			second: []bool{true},
			equal:  false,
		},
		{
			first:  []bool{false, true},
			second: []bool{false, true},
			equal:  true,
		},
		{
			first:  []bool{false, true},
			second: []bool{true, false},
			equal:  false,
		},
	}

	for _, tc := range tests {
		h1 := fnv.New64a()
		h2 := fnv.New64a()

		hashBool(h1, tc.first...)
		hashBool(h2, tc.second...)

		if tc.equal {
			assert.Equal(t, h1.Sum64(), h2.Sum64())
		} else {
			assert.NotEqual(t, h1.Sum64(), h2.Sum64())
		}
	}
}

func TestHashString(t *testing.T) {
	tests := []struct {
		first  []string
		second []string
		equal  bool
	}{
		{
			first:  []string{"alice"},
			second: []string{"alice"},
			equal:  true,
		},
		{
			first:  []string{"bob"},
			second: []string{"bob"},
			equal:  true,
		},
		{
			first:  []string{"alice"},
			second: []string{"bob"},
			equal:  false,
		},
		{
			first:  []string{"alice", "bob"},
			second: []string{"alice", "bob"},
			equal:  true,
		},
		{
			first:  []string{"alice", "bob"},
			second: []string{"bob", "alice"},
			equal:  false,
		},
	}

	for _, tc := range tests {
		h1 := fnv.New64a()
		h2 := fnv.New64a()

		hashString(h1, tc.first...)
		hashString(h2, tc.second...)

		if tc.equal {
			assert.Equal(t, h1.Sum64(), h2.Sum64())
		} else {
			assert.NotEqual(t, h1.Sum64(), h2.Sum64())
		}
	}
}

func TestHashArray(t *testing.T) {
	tests := []struct {
		canonical bool
		first     []string
		second    []string
		equal     bool
	}{
		{
			canonical: true,
			first:     []string{"GET"},
			second:    []string{"GET"},
			equal:     true,
		},
		{
			canonical: true,
			first:     []string{"GET", "POST"},
			second:    []string{"GET", "POST"},
			equal:     true,
		},
		{
			canonical: true,
			first:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			second:    []string{"DELETE", "PATCH", "PUT", "POST", "GET"},
			equal:     true,
		},
		{
			canonical: true,
			first:     []string{"GET", "POST", "PUT", "DELETE"},
			second:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			equal:     false,
		},
	}

	for _, tc := range tests {
		h1 := fnv.New64a()
		h2 := fnv.New64a()

		hashArray(h1, tc.canonical, tc.first)
		hashArray(h2, tc.canonical, tc.second)

		if tc.equal {
			assert.Equal(t, h1.Sum64(), h2.Sum64())
		} else {
			assert.NotEqual(t, h1.Sum64(), h2.Sum64())
		}
	}
}

func TestHashMap(t *testing.T) {
	tests := []struct {
		canonical bool
		first     map[string]string
		second    map[string]string
		equal     bool
	}{
		{
			canonical: true,
			first: map[string]string{
				"name": "John Doe",
			},
			second: map[string]string{
				"name": "John Doe",
			},
			equal: true,
		},
		{
			canonical: true,
			first: map[string]string{
				"name":  "John Doe",
				"email": "john.doe@example.com",
			},
			second: map[string]string{
				"email": "john.doe@example.com",
				"name":  "John Doe",
			},
			equal: true,
		},
		{
			canonical: true,
			first: map[string]string{
				"name":  "John Doe",
				"email": "john.doe@example.com",
			},
			second: map[string]string{
				"email": "john.doe@example.com",
			},
			equal: false,
		},
	}

	for _, tc := range tests {
		h1 := fnv.New64a()
		h2 := fnv.New64a()

		hashMap(h1, tc.canonical, tc.first)
		hashMap(h2, tc.canonical, tc.second)

		if tc.equal {
			assert.Equal(t, h1.Sum64(), h2.Sum64())
		} else {
			assert.NotEqual(t, h1.Sum64(), h2.Sum64())
		}
	}
}

func TestFindID(t *testing.T) {
	tests := []struct {
		idProp        string
		object        JSON
		expectedError string
		expectedValue interface{}
	}{
		{
			idProp: "id",
			object: JSON{
				"id": "11111111-1111-1111-1111-111111111111",
			},
			expectedValue: "11111111-1111-1111-1111-111111111111",
		},
		{
			idProp: "id",
			object: JSON{
				"_id": "11111111-1111-1111-1111-111111111111",
			},
			expectedError: "identifier id does not exist",
		},
		{
			object: JSON{
				"id": "11111111-1111-1111-1111-111111111111",
			},
			expectedValue: "11111111-1111-1111-1111-111111111111",
		},
		{
			object: JSON{
				"_id": "11111111-1111-1111-1111-111111111111",
			},
			expectedValue: "11111111-1111-1111-1111-111111111111",
		},
		{
			object: JSON{
				"Id": "11111111-1111-1111-1111-111111111111",
			},
			expectedValue: "11111111-1111-1111-1111-111111111111",
		},
		{
			object: JSON{
				"ID": "11111111-1111-1111-1111-111111111111",
			},
			expectedValue: "11111111-1111-1111-1111-111111111111",
		},
		{
			object: JSON{
				"key": "11111111-1111-1111-1111-111111111111",
			},
			expectedError: "cannot find an identifier",
		},
	}

	for _, tc := range tests {
		val, err := findID(tc.idProp, tc.object)

		if tc.expectedError != "" {
			assert.Contains(t, err.Error(), tc.expectedError)
			assert.Nil(t, val)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedValue, val)
		}
	}
}
