package blocks

import (
	"testing"
)

func TestInvalidSourceId(t *testing.T) {
	src, err := CreateSource("id does not exist", []string{})
	if src != nil || err == nil {
		t.Fail()
	}
}

