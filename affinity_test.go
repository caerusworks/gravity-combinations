package gravitycombinations

import (
	"testing"
)

func TestGetAffinityCombinations(t *testing.T) {
	tests := []struct {
		e          []interface{}
		affinities map[interface{}][]interface{}
		expected   [][]interface{}
	}{
		{
			e: []interface{}{"A", "B", "C"},
			affinities: map[interface{}][]interface{}{
				"B": {"A"},
			},
			expected: [][]interface{}{
				{"A"},
				{"C"},
				{"A", "C"},
				{"A", "B"},
				{"A", "B", "C"},
			},
		},
		{
			e: []interface{}{"A", "B", "C"},
			affinities: map[interface{}][]interface{}{
				"B": {"C"},
			},
			expected: [][]interface{}{
				{"A", "B", "C"},
				{"B", "C"},
				{"C"},
				{"A"},
				{"A", "C"},
			},
		},
		{
			e: []interface{}{"A", "B", "C"},
			affinities: map[interface{}][]interface{}{
				"D": {"A"},
			},
			expected: [][]interface{}{
				{"A"},
				{"B"},
				{"C"},
				{"A", "B"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			e: []interface{}{"A", "B", "C"},
			affinities: map[interface{}][]interface{}{
				"C": {"A", "B"},
			},
			expected: [][]interface{}{
				{"A", "B", "C"},
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			e: []interface{}{"A", "B", "C"},
			affinities: map[interface{}][]interface{}{
				"A": {"B", "C"},
			},
			expected: [][]interface{}{
				{"A", "B", "C"},
				{"B"},
				{"C"},
				{"B", "C"},
			},
		},
		{
			e: []interface{}{"A", "B", "C", "D", "E"},
			affinities: map[interface{}][]interface{}{
				"A": {"B", "C"},
				"E": {"D"},
				"D": {"E"},
				"C": {"A", "B"},
			},
			expected: [][]interface{}{
				{"A", "B", "C"},
				{"B"},
				{"A", "B", "C", "D", "E"},
				{"D", "E"},
				{"B", "D", "E"},
			},
		},
		{
			e: []interface{}{"A", "B", "C"},
			affinities: map[interface{}][]interface{}{
				"A": {"B", "C"},
				"B": {"A", "C"},
				"C": {"A", "B"},
			},
			expected: [][]interface{}{
				{"A", "B", "C"},
			},
		},
	}

	for _, test := range tests {
		intermediate := GetCombinations(test.e)
		result := GetAffinityCombinations(intermediate, test.affinities)
		if !compareUnordered(result, test.expected) {
			t.Errorf("GetAffinityCombinations(%v, %v) = %v; expected %v", test.e, test.affinities, result, test.expected)
		}
	}
}
