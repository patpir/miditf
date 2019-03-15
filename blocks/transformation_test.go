package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockTransformationFactory(config []string) (Transformation, error) {
	return nil, nil
}


func TestInvalidTransformationId(t *testing.T) {
	tf, err := CreateTransformation("transformation does not exist", []string{})
	assert.NotNil(t, err)
	assert.Nil(t, tf)
}

func TestListTransformations(t *testing.T) {
	details := Details{ identifier: "test-tf", description: "Test Transformation", arguments: []Argument{} }
	RegisterTransformation(details, mockTransformationFactory)
	transformations := Transformations()

	assert.Equal(t, 1, len(transformations))
	assert.Equal(t, "test-tf", transformations[0].identifier)
	assert.Equal(t, "Test Transformation", transformations[0].description)
}

