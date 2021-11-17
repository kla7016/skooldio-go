package class_room

import "testing"

func TestDoublePointer(t *testing.T) {
	tempInput := 2
	input := 2
	want := 4
	doublePointer(&input)
	if input != want {
		t.Errorf("input %d want %d but %d \n", tempInput, want, input)
	}
}
