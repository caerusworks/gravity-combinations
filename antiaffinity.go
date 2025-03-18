package gravitycombinations

func GetAntiAffinityCombinations(e [][]any, antiaffinities map[any][]any) [][]any {
	var result [][]any
	for _, combination := range e {
		if isValidCombination(combination, antiaffinities) {
			result = append(result, combination)
		}
	}
	return result
}

func isValidCombination(combination []any, antiaffinities map[any][]any) bool {
	for element, requiredElements := range antiaffinities {
		if contains(combination, element) {
			if hasRequiredElements(combination, requiredElements) {
				return false
			}
		}
	}
	return true
}

func hasRequiredElements(combination []any, requiredElements []any) bool {
	for _, required := range requiredElements {
		if contains(combination, required) {
			return true
		}
	}
	return false
}
