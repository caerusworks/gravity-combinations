package gravitycombinations

func GetCombinations(e []interface{}) [][]interface{} {
	var result [][]interface{}
	var helper func([]interface{}, int)
	helper = func(current []interface{}, index int) {
		if index == len(e) {
			if len(current) > 0 {
				result = append(result, append([]interface{}(nil), current...))
			}
			return
		}
		helper(current, index+1)
		helper(append(current, e[index]), index+1)
	}
	helper([]interface{}{}, 0)
	return result
}

func GetIntersectingCombinations(e []any, affinities map[interface{}][]interface{}, antiaffinities map[interface{}][]interface{}) [][]interface{} {
	result := GetAntiAffinityCombinations(GetAffinityCombinations(GetCombinations(e), affinities), antiaffinities)
	return result
}
