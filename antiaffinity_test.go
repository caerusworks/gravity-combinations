package gravitycombinations

import (
	"testing"
)

func TestGetAntiAffinityCombinations(t *testing.T) {
	tests := []struct {
		e              []interface{}
		antiaffinities map[interface{}][]interface{}
		expected       [][]interface{}
	}{
		{
			e:              []interface{}{1, 2, 3},
			antiaffinities: map[interface{}][]interface{}{1: {2}},
			expected:       [][]interface{}{{1, 3}, {2, 3}, {3}, {1}, {2}},
		},
		{
			e:              []interface{}{1, 2, 3, 4},
			antiaffinities: map[interface{}][]interface{}{1: {2}, 3: {4}},
			expected:       [][]interface{}{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {1}, {2}, {3}, {4}},
		},
		{
			e:              []interface{}{1, 2, 3},
			antiaffinities: map[interface{}][]interface{}{},
			expected:       [][]interface{}{{1, 2, 3}, {1, 2}, {1, 3}, {2, 3}, {1}, {2}, {3}},
		},
		{
			e:              []interface{}{1, 2, 3},
			antiaffinities: map[interface{}][]interface{}{1: {2}, 2: {3}},
			expected:       [][]interface{}{{1, 3}, {1}, {2}, {3}},
		},
		{
			e:              []interface{}{1, 2, 3},
			antiaffinities: map[interface{}][]interface{}{1: {2}, 2: {1, 3}},
			expected:       [][]interface{}{{1, 3}, {1}, {2}, {3}},
		},
	}

	for _, test := range tests {
		intermediate := GetCombinations(test.e)
		result := GetAntiAffinityCombinations(intermediate, test.antiaffinities)
		if !compareUnordered(result, test.expected) {
			t.Errorf("GetAntiAffinityCombinations(%v, %v) = %v; expected %v", test.e, test.antiaffinities, result, test.expected)
		}
	}
}
