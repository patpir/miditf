package core

import "testing"

func TestCreatePiece(t *testing.T) {
	piece := NewPiece()
	if piece == nil {
		t.Fail()
	}
}

