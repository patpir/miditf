package core

import "testing"

func TestCreateTrack(t *testing.T) {
	track := NewTrack()
	if track == nil {
		t.Fail()
	}
}

