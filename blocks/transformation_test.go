package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockTransformationFactory(arguments map[string]interface{}) (Transformation, error) {
	return nil, nil
}


func TestInvalidTransformationId(t *testing.T) {
	r := NewRegistrator()
	tf, err := r.CreateTransformation("transformation does not exist", make(map[string]interface{}))
	assert.NotNil(t, err)
	assert.Nil(t, tf)
}

func TestListTransformations(t *testing.T) {
	r := NewRegistrator()
	info := BlockInfo{ identifier: "test-tf", description: "Test Transformation", argumentInfos: []ArgumentInfo{} }
	r.RegisterTransformation(info, mockTransformationFactory)
	transformations := r.Transformations()

	assert.Equal(t, 1, len(transformations))
	assert.Equal(t, "test-tf", transformations[0].identifier)
	assert.Equal(t, "Test Transformation", transformations[0].description)
}

