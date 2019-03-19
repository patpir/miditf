package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockVisualizationFactory(config []Argument) (Visualization, error) {
	return nil, nil
}


func TestInvalidVisualizationId(t *testing.T) {
	r := NewRegistrator()
	visu, err := r.CreateVisualization("visualization does not exist", []Argument{})
	assert.NotNil(t, err)
	assert.Nil(t, visu)
}

func TestListVisualizations(t *testing.T) {
	r := NewRegistrator()
	info := BlockInfo{ identifier: "test-visu", description: "Test Visualization", argumentInfos: []ArgumentInfo{} }
	r.RegisterVisualization(info, mockVisualizationFactory)
	visualizations := r.Visualizations()

	assert.Equal(t, 1, len(visualizations))
	assert.Equal(t, "test-visu", visualizations[0].Identifier())
	assert.Equal(t, "Test Visualization", visualizations[0].Description())
}

