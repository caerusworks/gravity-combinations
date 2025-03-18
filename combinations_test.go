package gravitycombinations

import (
	"reflect"
	"testing"
)

func TestGetCombinations(t *testing.T) {
	tests := []struct {
		input    []any
		expected [][]any
	}{
		{
			input:    []any{},
			expected: [][]any{},
		},
		{
			input:    []any{1},
			expected: [][]any{[]interface{}{1}},
		},
		{
			input:    []any{1, 2},
			expected: [][]any{[]interface{}{1}, []interface{}{2}, []interface{}{1, 2}},
		},
		{
			input:    []interface{}{1, 2, 3},
			expected: [][]any{[]interface{}{1}, []interface{}{2}, []interface{}{1, 2}, []interface{}{3}, []interface{}{1, 3}, []interface{}{2, 3}, []interface{}{1, 2, 3}},
		},
		{
			input:    []interface{}{"a", "b"},
			expected: [][]interface{}{[]any{"a"}, []any{"b"}, []any{"a", "b"}},
		},
	}

	for _, test := range tests {
		result := GetCombinations(test.input)
		if !compareUnordered(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestGetIntersectingCombinations(t *testing.T) {
	tests := []struct {
		input          []any
		affinities     map[any][]any
		antiaffinities map[any][]any
		expected       [][]any
	}{
		{
			input:          []interface{}{},
			affinities:     map[interface{}][]interface{}{},
			antiaffinities: map[interface{}][]interface{}{},
			expected:       [][]interface{}{},
		},
		{
			input:          []any{1},
			affinities:     map[any][]any{1: {2}},
			antiaffinities: map[any][]any{1: {3}},
			expected:       [][]any{},
		},
		{
			input:          []interface{}{1, 2},
			affinities:     map[interface{}][]interface{}{1: {2}},
			antiaffinities: map[interface{}][]interface{}{1: {3}},
			expected:       [][]interface{}{[]any{1, 2}, {2}},
		},
		{
			input:          []interface{}{1, 2, 3},
			affinities:     map[interface{}][]interface{}{1: {2}},
			antiaffinities: map[interface{}][]interface{}{1: {3}},
			expected:       [][]interface{}{[]any{1, 2}, {2, 3}, {2}, {3}},
		},
		{
			input:          []interface{}{"a", "b"},
			affinities:     map[interface{}][]interface{}{"a": {"b"}},
			antiaffinities: map[interface{}][]interface{}{"a": {"c"}},
			expected: [][]interface{}{
				[]any{"a", "b"}, {"b"},
			},
		},
	}

	for _, test := range tests {
		result := GetIntersectingCombinations(test.input, test.affinities, test.antiaffinities)
		if !compareUnordered(result, test.expected) {
			t.Errorf("For input %v with affinities %v and antiaffinities %v, expected %v, but got %v", test.input, test.affinities, test.antiaffinities, test.expected, result)
		}
	}
}

func compareUnordered(a, b [][]any) bool {
	if len(a) != len(b) {
		return false
	}

	visited := make([]bool, len(b))
	for _, subA := range a {
		found := false
		for j, subB := range b {
			if !visited[j] && reflect.DeepEqual(subA, subB) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
