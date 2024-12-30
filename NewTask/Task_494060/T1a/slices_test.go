package sliceops

import (
	"reflect"
	"testing"
)

func TestAddToSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	newElement := 4
	expected := []int{1, 2, 3, 4}

	result := AddToSlice(slice, newElement)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AddToSlice failed: expected %v, got %v", expected, result)
	}
}

func TestRemoveFromSlice(t *testing.T) {
	slice := []int{1, 2, 3, 2, 4}
	elementToRemove := 2
	expected := []int{1, 3, 2, 4}

	result := RemoveFromSlice(slice, elementToRemove)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveFromSlice failed: expected %v, got %v", expected, result)
	}
}

func TestIsElementInSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	element := 3
	notInSlice := 5

	if !IsElementInSlice(slice, element) {
		t.Errorf("IsElementInSlice failed: element %d not found in slice %v", element, slice)
	}

	if IsElementInSlice(slice, notInSlice) {
		t.Errorf("IsElementInSlice failed: element %d found in slice %v", notInSlice, slice)
	}
}

func TestIterateSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	expectedResult := []int{2, 4, 6}
	var result []int

	funcToApply := func(x int) {
		result = append(result, x*2)
	}

	IterateSlice(slice, funcToApply)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("IterateSlice failed: expected %v, got %v", expectedResult, result)
	}
}