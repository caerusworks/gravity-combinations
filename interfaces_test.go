package gravitycombinations

import (
	"reflect"
	"testing"
)

const expstr = "expected %v, got %v"

type EntityOne struct {
	Items []string // Field
}

func (e EntityOne) IsGravitated() bool {
	return true
}

func (e EntityOne) GetItemOne(i int) string {
	return e.Items[i]
}

func (e EntityOne) FillItems(items []string) {
	e.Items = items
}

type EntityTwo struct {
	Items []string // Field
}

func (e EntityTwo) IsGravitated() bool {
	return true
}

func (e EntityTwo) GetItemTwo(i int) string {
	return e.Items[i]
}

func (e EntityTwo) FillItems(items []string) {
	e.Items = items
}

type EntityThree struct {
	Items []string // Field
}

func (e EntityThree) IsGravitated() bool {
	return true
}

func (e EntityThree) GetItemThree(i int) string {
	return e.Items[i]
}

func (e EntityThree) FillItems(items []string) {
	e.Items = items
}

type EntityFour struct {
	Items []string // Field
}

func (e EntityFour) IsGravitated() bool {
	return true
}

func (e EntityFour) GetItemFour(i int) string {
	return e.Items[i]
}

func (e EntityFour) FillItems(items []string) {
	e.Items = items
}

func r(i any) reflect.Type {
	return reflect.TypeOf(i)
}

func TestClassesCombinations(t *testing.T) {
	// Test Case 1
	InterfacesEnumeration1 := &InterfaceSet{
		Interfaces: []CombinedInterface{
			{
				EntityInterface: &EntityOne{},
				Affinities:      []GravitatedInterface{},
				Antiaffinities:  []GravitatedInterface{},
			},
			{
				EntityInterface: &EntityTwo{},
				Affinities:      []GravitatedInterface{},
				Antiaffinities:  []GravitatedInterface{&EntityThree{}, &EntityOne{}},
			},
			{
				EntityInterface: &EntityThree{},
				Affinities:      []GravitatedInterface{},
				Antiaffinities:  []GravitatedInterface{&EntityTwo{}, &EntityOne{}},
			},
		},
	}
	expected1 := [][]reflect.Type{
		{r(&EntityOne{})},
		{r(&EntityTwo{})},
		{r(&EntityThree{})},
	}
	result1 := GetStructuresCombinations(InterfacesEnumeration1)
	if !compareUnorderedTypes(result1, expected1) {
		t.Errorf(expstr, expected1, result1)
	}

	// Test Case 2
	InterfacesEnumeration2 := &InterfaceSet{
		Interfaces: []CombinedInterface{
			{
				EntityInterface: &EntityOne{},
				Affinities:      []GravitatedInterface{&EntityTwo{}},
				Antiaffinities:  []GravitatedInterface{},
			},
			{
				EntityInterface: &EntityTwo{},
				Affinities:      []GravitatedInterface{&EntityOne{}},
				Antiaffinities:  []GravitatedInterface{},
			},
			{
				EntityInterface: &EntityThree{},
				Affinities:      []GravitatedInterface{},
				Antiaffinities:  []GravitatedInterface{&EntityFour{}},
			},
			{
				EntityInterface: &EntityFour{},
				Affinities:      []GravitatedInterface{},
				Antiaffinities:  []GravitatedInterface{&EntityThree{}},
			},
		},
	}
	expected2 := [][]reflect.Type{
		{r(&EntityOne{}), r(&EntityTwo{})},
		{r(&EntityOne{}), r(&EntityTwo{}), r(&EntityFour{})},
		{r(&EntityOne{}), r(&EntityTwo{}), r(&EntityThree{})},
		{r(&EntityThree{})},
		{r(&EntityFour{})},
	}
	result2 := GetStructuresCombinations(InterfacesEnumeration2)
	if !compareUnorderedTypes(result2, expected2) {
		t.Errorf(expstr, expected2, result2)
	}

	// Test Case 3
	InterfacesEnumeration3 := &InterfaceSet{
		Interfaces: []CombinedInterface{
			{
				EntityInterface: &EntityOne{},
				Affinities:      []GravitatedInterface{&EntityTwo{}, &EntityThree{}},
				Antiaffinities:  []GravitatedInterface{},
			},
			{
				EntityInterface: &EntityTwo{},
				Affinities:      []GravitatedInterface{&EntityOne{}},
				Antiaffinities:  []GravitatedInterface{&EntityFour{}},
			},
			{
				EntityInterface: &EntityThree{},
				Affinities:      []GravitatedInterface{&EntityOne{}},
				Antiaffinities:  []GravitatedInterface{},
			},
			{
				EntityInterface: &EntityFour{},
				Affinities:      []GravitatedInterface{},
				Antiaffinities:  []GravitatedInterface{&EntityTwo{}},
			},
		},
	}
	expected3 := [][]reflect.Type{
		{r(&EntityOne{}), r(&EntityTwo{}), r(&EntityThree{})},
		{r(&EntityFour{})},
	}
	result3 := GetStructuresCombinations(InterfacesEnumeration3)
	if !compareUnorderedTypes(result3, expected3) {
		t.Errorf(expstr, expected3, result3)
	}
}

func compareUnorderedTypes(a, b [][]reflect.Type) bool {
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
