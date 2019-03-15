package blocks

import (
	"testing"
)

/*
type mockSource struct { }
func (s *mockSource) Piece() *core.Piece {
	return nil
}
*/
func mockSourceFactory(config []string) (Source, error) {
	return nil, nil
}


func TestInvalidSourceId(t *testing.T) {
	src, err := CreateSource("id does not exist", []string{})
	if src != nil || err == nil {
		t.Fail()
	}
}

func TestListSources(t *testing.T) {
	details := Details{ identifier: "test", description: "Test", arguments: []Argument{} }
	RegisterSource(details, mockSourceFactory)
	sources := Sources()
	if len(sources) != 1 {
		t.FailNow()
	}
	if sources[0].identifier != "test" {
		t.Fail()
	}
	if sources[0].description != "Test" {
		t.Fail()
	}
}

