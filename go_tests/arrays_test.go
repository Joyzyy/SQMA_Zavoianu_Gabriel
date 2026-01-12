package gotests_test

import (
	gotests "ism/sqma"
	"testing"
)

func TestAddElement(t *testing.T) {
	arr := []int{1, 2, 3}
	element := 4
	expected := []int{1, 2, 3, 4}

	result := gotests.AddElement(arr, element)
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("at index %d: expected %d, got %d", i, v, result[i])
		}
	}
}

func TestRemoveElement(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	index := 2
	expected := []int{1, 2, 4}

	result := gotests.RemoveElement(arr, index)
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("at index %d: expected %d, got %d", i, v, result[i])
		}
	}
}

func TestGetElement(t *testing.T) {
	arr := []int{10, 20, 30}
	index := 1
	expected := 20

	result, ok := gotests.GetElement(arr, index)
	if !ok {
		t.Fatalf("expected to get element at index %d", index)
	}
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
