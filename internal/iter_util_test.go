package internal

import (
	"reflect"
	"testing"
)

func TestFilterList(t *testing.T) {
	numbers := make([]uint8, 10)
	for i := range cap(numbers) {
		numbers[i] = uint8(i) + 1
	}
	filteredNumbers, newCount := FilterList(&numbers, func(n uint8) bool {
		return n%2 > 0
	})

	if newCount != 5 {
		t.Errorf("Expected new length of %v but got %v", 5, newCount)
	}
	expectedNewSlice := []uint8{1, 3, 5, 7, 9}
	if !reflect.DeepEqual(filteredNumbers, expectedNewSlice) {
		t.Errorf("Expected filtered numbers to equal %v but instead got %v", expectedNewSlice, filteredNumbers)
	}
}

func TestFilterListWithEmptySlice(t *testing.T) {
	strings := []string{}
	_, newCount := FilterList(&strings, func(_ string) bool { return true })
	if newCount != 0 {
		t.Errorf("Expected new count to be 0 but got %v", newCount)
	}
}
