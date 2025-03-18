package gravitycombinations

func GetAffinityCombinations(e [][]interface{}, affinities map[interface{}][]interface{}) [][]interface{} {
	var result [][]interface{}
	for _, combination := range e {
		if isValidAffinityCombination(combination, affinities) {
			result = append(result, combination)
		}
	}
	return result
}

func isValidAffinityCombination(combination []interface{}, affinities map[interface{}][]interface{}) bool {
	for element, requiredElements := range affinities {
		if contains(combination, element) {
			if !containsAll(combination, requiredElements) {
				return false
			}
		}
	}
	return true
}

func containsAll(slice []interface{}, items []interface{}) bool {
	for _, item := range items {
		if !contains(slice, item) {
			return false
		}
	}
	return true
}

func contains(slice []interface{}, item interface{}) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
