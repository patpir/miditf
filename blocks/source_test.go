package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockSourceFactory(config []string) (Source, error) {
	return nil, nil
}


func TestInvalidSourceId(t *testing.T) {
	src, err := CreateSource("id does not exist", []string{})
	assert.NotNil(t, err)
	assert.Nil(t, src)
}

func TestListSources(t *testing.T) {
	details := Details{ identifier: "test-source", description: "Test Source", arguments: []Argument{} }
	RegisterSource(details, mockSourceFactory)
	sources := Sources()

	assert.Equal(t, 1, len(sources))
	assert.Equal(t, "test-source", sources[0].Identifier())
	assert.Equal(t, "Test Source", sources[0].Description())
}

