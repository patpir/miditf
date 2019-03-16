package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockSourceFactory(config []Argument) (Source, error) {
	return nil, nil
}


func TestInvalidSourceId(t *testing.T) {
	src, err := CreateSource("id does not exist", []Argument{})
	assert.NotNil(t, err)
	assert.Nil(t, src)
}

func TestListSources(t *testing.T) {
	info := BlockInfo{ identifier: "test-source", description: "Test Source", argumentInfos: []ArgumentInfo{} }
	RegisterSource(info, mockSourceFactory)
	sources := Sources()

	assert.Equal(t, 1, len(sources))
	assert.Equal(t, "test-source", sources[0].Identifier())
	assert.Equal(t, "Test Source", sources[0].Description())
}

