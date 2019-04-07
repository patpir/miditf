package blocks

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
)


func TestPipelineResult(t *testing.T) {
	sourceBlock := NewBlock("source-id", "", make(map[string]interface{}))
	transformBlock := NewBlock("transform-id", "", make(map[string]interface{}))
	visualizationBlock := NewBlock("visu-id", "", make(map[string]interface{}))
	pr := PipelineResult{ sourceBlock, []Block{transformBlock}, visualizationBlock, "Output", errors.New("Error") }

	assert.Equal(t, sourceBlock, pr.Source)
	assert.Equal(t, 1, len(pr.Transformations))
	assert.Equal(t, transformBlock, pr.Transformations[0])
	assert.Equal(t, visualizationBlock, pr.Visualization)
	assert.Equal(t, "Output", pr.Output)
	assert.NotNil(t, pr.Err)
	assert.Equal(t, "Error", pr.Err.Error())
}

