package gravitycombinations

import (
	"fmt"
	"reflect"
)

type GravitatedInterface interface {
	IsGravitated() bool
}

type CombinedInterface struct {
	EntityInterface GravitatedInterface
	Affinities      []GravitatedInterface
	Antiaffinities  []GravitatedInterface
}

type InterfaceSet struct {
	Interfaces []CombinedInterface
}

func findIndex(slice []interface{}, item reflect.Type) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// Convert Structures to
func TranslateSSetToIndexedSets(StructureSet InterfaceSet) ([]interface{}, []interface{}, map[interface{}][]interface{}, map[interface{}][]interface{}) {
	// Convert Underlying Structures, Their Affinities and Antiaffinities to integers representing the index of the Structure in the StructureSet
	// Index Provided Set Of Structures First
	indexes := make([]interface{}, len(StructureSet.Interfaces))
	interfaceindexes := make([]interface{}, len(StructureSet.Interfaces))
	for i := range StructureSet.Interfaces {
		indexes[i] = reflect.TypeOf(StructureSet.Interfaces[i].EntityInterface)
		interfaceindexes[i] = i
	}
	// Index Provided Set Of Affinities
	affinities := map[interface{}][]interface{}{}
	for i, structure := range StructureSet.Interfaces {
		affinities[i] = make([]interface{}, len(structure.Affinities))
		for j, affinity := range structure.Affinities {
			idx := findIndex(indexes, reflect.TypeOf(affinity))
			if idx != -1 {
				affinities[i][j] = idx
			}
		}
	}

	// Index Provided Set Of Antiaffinities
	antiaffinities := map[interface{}][]interface{}{}
	for i, structure := range StructureSet.Interfaces {
		antiaffinities[i] = make([]interface{}, len(structure.Antiaffinities))
		for j, antiaffinity := range structure.Antiaffinities {
			idx := findIndex(indexes, reflect.TypeOf(antiaffinity))
			if idx != -1 {
				antiaffinities[i][j] = idx
			}
		}
	}
	fmt.Println(indexes, affinities, antiaffinities)
	return indexes, interfaceindexes, affinities, antiaffinities
}

func GetStructuresCombinations(iSet *InterfaceSet) [][]reflect.Type {
	// Convert Underlying Interfaces, Affinities and Antiaffinities to integers representing the index of the interface in the InterfaceSet
	interfaces, interfaceindexes, affinities, antiaffinities := TranslateSSetToIndexedSets(*iSet)
	// Get intersecting combinations
	intersectingCombinations := GetIntersectingCombinations(interfaceindexes, affinities, antiaffinities)
	// Convert back to original interfaces
	returnCombinations := make([][]reflect.Type, len(intersectingCombinations))
	for i, combination := range intersectingCombinations {
		returnCombinations[i] = make([]reflect.Type, len(combination))
		for j := range combination {
			e := combination[j]
			t := interfaceindexes[e.(int)]
			returnCombinations[i][j] = interfaces[t.(int)].(reflect.Type)
		}
	}
	return returnCombinations
}
