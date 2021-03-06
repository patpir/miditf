package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockSourceFactory(arguments map[string]interface{}) (Source, error) {
	return nil, nil
}


func TestInvalidSourceId(t *testing.T) {
	r := NewRegistrator()
	src, err := r.CreateSource("id does not exist", make(map[string]interface{}))
	assert.NotNil(t, err)
	assert.Nil(t, src)
}

func TestListSources(t *testing.T) {
	r := NewRegistrator()
	info := BlockInfo{ identifier: "test-source", description: "Test Source", argumentInfos: []ArgumentInfo{} }
	r.RegisterSource(info, mockSourceFactory)
	sources := r.Sources()

	assert.Equal(t, 1, len(sources))
	assert.Equal(t, "test-source", sources[0].Identifier())
	assert.Equal(t, "Test Source", sources[0].Description())
}

