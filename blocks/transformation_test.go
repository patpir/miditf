package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockTransformationFactory(config []Argument) (Transformation, error) {
	return nil, nil
}


func TestInvalidTransformationId(t *testing.T) {
	tf, err := CreateTransformation("transformation does not exist", []Argument{})
	assert.NotNil(t, err)
	assert.Nil(t, tf)
}

func TestListTransformations(t *testing.T) {
	info := BlockInfo{ identifier: "test-tf", description: "Test Transformation", argumentInfos: []ArgumentInfo{} }
	RegisterTransformation(info, mockTransformationFactory)
	transformations := Transformations()

	assert.Equal(t, 1, len(transformations))
	assert.Equal(t, "test-tf", transformations[0].identifier)
	assert.Equal(t, "Test Transformation", transformations[0].description)
}

