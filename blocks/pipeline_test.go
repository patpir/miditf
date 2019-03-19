package blocks

import (
	"testing"
	"fmt"
	"strconv"

	"github.com/stretchr/testify/assert"

	"github.com/patpir/miditf/core"
)


func TestSources(t *testing.T) {
	source := NewBlock("my-source", "Some source", []Argument{})
	pipeline := NewPipeline()
	pipeline.AddSource(source)

	sources := pipeline.Sources()
	assert.Equal(t, 1, len(sources))
	assert.Equal(t, source, sources[0])
}

func TestTransformations(t *testing.T) {
	transform := NewBlock("my-transform", "Some transformation", []Argument{})
	pipeline := NewPipeline()
	pipeline.AddTransformation(transform)

	transforms := pipeline.Transformations()
	assert.Equal(t, 1, len(transforms))
	assert.Equal(t, transform, transforms[0])
}

func TestVisualizations(t *testing.T) {
	visu := NewBlock("my-visu", "Some visualization", []Argument{})
	pipeline := NewPipeline()
	pipeline.AddVisualization(visu)

	visualizations := pipeline.Visualizations()
	assert.Equal(t, 1, len(visualizations))
	assert.Equal(t, visu, visualizations[0])
}


type emptyTrackSource struct { }
func (s *emptyTrackSource) Piece() *core.Piece {
	p := core.NewPiece()
	p.AddTrack(core.NewTrack())
	return p
}

type noteAppender struct {
	key uint8
}
func (appender *noteAppender) Transform(piece *core.Piece) *core.Piece {
	notes := piece.Tracks()[0].Notes()
	var startTime uint32 = 0
	if len(notes) > 0 {
		startTime = notes[len(notes) - 1].EndTime() + 1
	}
	note := core.NewNote(0, appender.key, 80, startTime, startTime + 15)
	piece.Tracks()[0].AddNote(note)
	return piece
}

type noteLister struct { }
func (lister *noteLister) Visualize(piece *core.Piece) string {
	result := ""
	tracks := piece.Tracks()
	for i, track := range tracks {
		result = fmt.Sprintf("%sTrack %2d:", result, i)
		notes := track.Notes()
		for _, note := range notes {
			result = fmt.Sprintf("%s %d", result, note.Key())
		}
		result += "\n"
	}
	return result
}

func mockRegistrator() Registrator {
	r := NewRegistrator()
	r.RegisterSource(NewBlockInfo("empty", "Empty source", []ArgumentInfo{}), func(arguments []Argument) (Source, error) {
		return &emptyTrackSource{}, nil
	})
	r.RegisterTransformation(NewBlockInfo("append-note", "Append a note", []ArgumentInfo{}), func(arguments []Argument) (Transformation, error) {
		key, err := strconv.Atoi(arguments[0].Value())
		return &noteAppender{ key: uint8(key) }, err
	})
	r.RegisterVisualization(NewBlockInfo("list-notes", "List all notes", []ArgumentInfo{}), func(arguments []Argument) (Visualization, error) {
		return &noteLister{}, nil
	})
	return r
}

func createPipeline(creator Creator, sourceErrors []bool, transformErrors []bool, visuErrors []bool, keys []int) *Pipeline {
	p := NewPipeline()
	p.creator = creator
	for si, sourceError := range sourceErrors {
		if sourceError {
			p.AddSource(NewBlock("source-error", fmt.Sprintf("start %d", si), []Argument{}))
		} else {
			p.AddSource(NewBlock("empty", fmt.Sprintf("start %d", si), []Argument{}))
		}
	}
	for ti, transformError := range transformErrors {
		if transformError {
			p.AddTransformation(NewBlock("transform-error", fmt.Sprintf("transform %d", ti), []Argument{}))
		} else {
			p.AddTransformation(NewBlock("append-note", fmt.Sprintf("transform %d", ti), []Argument{
				NewArgument("key", strconv.Itoa(keys[ti])),
			}))
		}
	}
	for vi, visuError := range visuErrors {
		if visuError {
			p.AddVisualization(NewBlock("visu-error", fmt.Sprintf("visu %d", vi), []Argument{}))
		} else {
			p.AddVisualization(NewBlock("list-notes", fmt.Sprintf("visu %d", vi), []Argument{}))
		}
	}
	return p
}

func TestPerformSuccessMulti(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]bool{ false, false, },
		[]bool{ false, false, },
		[]bool{ false, false, },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 4)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	for result := range ch {
		count += 1
		assert.NotNil(t, result.source)
		assert.Equal(t, 2, len(result.transformations))
		assert.NotNil(t, result.visualization)
		assert.Equal(t, "Track  0: 60 69\n", result.output)
		assert.Nil(t, result.err)
	}
	assert.Equal(t, 4, count)
}

func TestPerformErrorSingleSource(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]bool{ true, },
		[]bool{ false, },
		[]bool{ false, },
		[]int{ 60, },
	)

	ch := make(chan PipelineResult, 1)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	for result := range ch {
		count += 1
		assert.NotNil(t, result.source)
		assert.NotNil(t, result.err)
		assert.Equal(t, 0, len(result.transformations))
		assert.Nil(t, result.visualization)
		assert.Equal(t, "", result.output)
	}
	assert.Equal(t, 1, count)
}

func TestPerformErrorFirstTransform(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]bool{ false, false, },
		[]bool{ true, false, },
		[]bool{ false, false, },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 2)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	for result := range ch {
		count += 1
		assert.NotNil(t, result.source)
		assert.Equal(t, 1, len(result.transformations))
		assert.NotNil(t, result.err)
		assert.Nil(t, result.visualization)
		assert.Equal(t, "", result.output)
	}
	assert.Equal(t, 2, count)
}

func TestPerformErrorFirstSource(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]bool{ true, false, },
		[]bool{ false, false, },
		[]bool{ false, false, },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 3)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	errorCount := 0
	for result := range ch {
		count += 1
		if result.err != nil {
			errorCount += 1
			assert.NotNil(t, result.source)
			assert.NotNil(t, result.err)
			assert.Equal(t, 0, len(result.transformations))
			assert.Nil(t, result.visualization)
			assert.Equal(t, "", result.output)
		} else {
			assert.NotNil(t, result.source)
			assert.Equal(t, 2, len(result.transformations))
			assert.NotNil(t, result.visualization)
			assert.Equal(t, "Track  0: 60 69\n", result.output)
			assert.Nil(t, result.err)
		}
	}
	assert.Equal(t, 1, errorCount)
	assert.Equal(t, 3, count)
}

func TestPerformErrorFirstVisu(t *testing.T) {
	registrator := mockRegistrator()
	visu, err := registrator.CreateVisualization("visu-error", []Argument{})
	assert.Nil(t, visu)
	assert.NotNil(t, err)
	pipeline := createPipeline(
		registrator,
		[]bool{ false, false, },
		[]bool{ false, false, },
		[]bool{ true, false, },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 4)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	errorCount := 0
	for result := range ch {
		count += 1
		if result.err != nil {
			errorCount += 1
			assert.NotNil(t, result.source)
			assert.Equal(t, 2, len(result.transformations))
			assert.NotNil(t, result.visualization)
			assert.NotNil(t, result.err)
			assert.Equal(t, "", result.output)
		} else {
			assert.NotNil(t, result.source)
			assert.Equal(t, 2, len(result.transformations))
			assert.NotNil(t, result.visualization)
			assert.Equal(t, "Track  0: 60 69\n", result.output)
			assert.Nil(t, result.err)
		}
	}
	assert.Equal(t, 2, errorCount)
	assert.Equal(t, 4, count)
}

