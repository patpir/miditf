package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockVisualizationFactory(config []Argument) (Visualization, error) {
	return nil, nil
}


func TestInvalidVisualizationId(t *testing.T) {
	visu, err := CreateVisualization("visualization does not exist", []Argument{})
	assert.NotNil(t, err)
	assert.Nil(t, visu)
}

func TestListVisualizations(t *testing.T) {
	info := BlockInfo{ identifier: "test-visu", description: "Test Visualization", argumentInfos: []ArgumentInfo{} }
	RegisterVisualization(info, mockVisualizationFactory)
	visualizations := Visualizations()

	assert.Equal(t, 1, len(visualizations))
	assert.Equal(t, "test-visu", visualizations[0].Identifier())
	assert.Equal(t, "Test Visualization", visualizations[0].Description())
}

