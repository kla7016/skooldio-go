package class_room

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	input := "Apple Banana Apple Banana apple"
	wantOutput := map[string]int{
		"Apple":  2,
		"Banana": 2,
		"apple":  1,
	}

	get := wordCount(input)
	if !reflect.DeepEqual(get, wantOutput) {
		t.Errorf("input %s want %v but %v", input, wantOutput, get)
	}
}
