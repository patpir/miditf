package core

import "testing"

func TestCreateNote(t *testing.T) {
	note := NewNote()
	if note == nil {
		t.Fail()
	}
}

