package blocks

import (
	"testing"
	"errors"
	"fmt"

	"github.com/stretchr/testify/assert"

	"github.com/patpir/miditf/core"
)


func TestSources(t *testing.T) {
	source := NewBlock("my-source", "Some source", make(map[string]interface{}))
	pipeline := NewPipeline()
	pipeline.AddSource(source)

	sources := pipeline.Sources()
	assert.Equal(t, 1, len(sources))
	assert.Equal(t, source, sources[0])
}

func TestTransformations(t *testing.T) {
	transform := NewBlock("my-transform", "Some transformation", make(map[string]interface{}))
	pipeline := NewPipeline()
	pipeline.AddTransformation(transform)

	transforms := pipeline.Transformations()
	assert.Equal(t, 1, len(transforms))
	assert.Equal(t, transform, transforms[0])
}

func TestVisualizations(t *testing.T) {
	visu := NewBlock("my-visu", "Some visualization", make(map[string]interface{}))
	pipeline := NewPipeline()
	pipeline.AddVisualization(visu)

	visualizations := pipeline.Visualizations()
	assert.Equal(t, 1, len(visualizations))
	assert.Equal(t, visu, visualizations[0])
}


type emptyTrackSource struct { }
func (s *emptyTrackSource) Piece() (*core.Piece, error) {
	p := core.NewPiece()
	p.AddTrack(core.NewTrack())
	return p, nil
}

type noteAppender struct {
	key uint8
}
func (appender *noteAppender) Transform(piece *core.Piece) (*core.Piece, error) {
	notes := piece.Tracks()[0].Notes()
	var startTime uint32 = 0
	if len(notes) > 0 {
		startTime = notes[len(notes) - 1].EndTime() + 1
	}
	note := core.NewNote(0, appender.key, 80, startTime, startTime + 15)
	piece.Tracks()[0].AddNote(note)
	return piece, nil
}

type noteLister struct { }
func (lister *noteLister) Visualize(piece *core.Piece) (string, error) {
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
	return result, nil
}

type produceError struct { }
func (e *produceError) Piece() (*core.Piece, error) {
	return nil, errors.New("Source Error")
}
func (e *produceError) Transform(piece *core.Piece) (*core.Piece, error) {
	return nil, errors.New("Transform Error")
}
func (e *produceError) Visualize(piece *core.Piece) (string, error) {
	return "", errors.New("Visu Error")
}

func mockRegistrator() Registrator {
	r := NewRegistrator()

	r.RegisterSource(NewBlockInfo("empty", "Empty source", []ArgumentInfo{}), func(arguments map[string]interface{}) (Source, error) {
		return &emptyTrackSource{}, nil
	})
	r.RegisterTransformation(NewBlockInfo("append-note", "Append a note", []ArgumentInfo{}), func(arguments map[string]interface{}) (Transformation, error) {
		key, ok := arguments["key"].(int)
		if ok {
			return &noteAppender{ key: uint8(key) }, nil
		} else {
			return nil, InvalidArgumentTypeError
		}
	})
	r.RegisterVisualization(NewBlockInfo("list-notes", "List all notes", []ArgumentInfo{}), func(arguments map[string]interface{}) (Visualization, error) {
		return &noteLister{}, nil
	})

	r.RegisterSource(NewBlockInfo("error", "Error", []ArgumentInfo{}), func(arguments map[string]interface{}) (Source, error) {
		return &produceError{}, nil
	})
	r.RegisterTransformation(NewBlockInfo("error", "Error", []ArgumentInfo{}), func(arguments map[string]interface{}) (Transformation, error) {
		return &produceError{}, nil
	})
	r.RegisterVisualization(NewBlockInfo("error", "Error", []ArgumentInfo{}), func(arguments map[string]interface{}) (Visualization, error) {
		return &produceError{}, nil
	})

	return r
}

func createPipeline(creator Creator, sources []string, transformations []string, visualizations []string, keys []int) *Pipeline {
	p := NewPipeline()
	p.creator = creator
	for si, source := range sources {
		p.AddSource(NewBlock(source, fmt.Sprintf("start %d", si), make(map[string]interface{})))
	}
	for ti, transformation := range transformations {
		p.AddTransformation(NewBlock(transformation, fmt.Sprintf("transform %d", ti), map[string]interface{}{
			"key": keys[ti],
		}))
	}
	for vi, visualization := range visualizations {
		p.AddVisualization(NewBlock(visualization, fmt.Sprintf("visu %d", vi), make(map[string]interface{})))
	}
	return p
}

func TestPerformSuccessMulti(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]string{ "empty", "empty", },
		[]string{ "append-note", "append-note", },
		[]string{ "list-notes", "list-notes", },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 4)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	for result := range ch {
		count += 1
		assert.NotNil(t, result.Source)
		assert.Equal(t, 2, len(result.Transformations))
		assert.NotNil(t, result.Visualization)
		assert.Equal(t, "Track  0: 60 69\n", result.Output)
		assert.Nil(t, result.Err)
	}
	assert.Equal(t, 4, count)
}

func TestPerformWithoutTransform(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]string{ "empty", },
		[]string{},
		[]string{ "list-notes", },
		[]int{},
	)

	ch := make(chan PipelineResult, 1)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	for result := range ch {
		count += 1
		assert.NotNil(t, result.Source)
		assert.Equal(t, 0, len(result.Transformations))
		assert.NotNil(t, result.Visualization)
		assert.Equal(t, "Track  0:\n", result.Output)
		assert.Nil(t, result.Err)
	}
	assert.Equal(t, 1, count)
}

func TestPerformUnknownSingleSource(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]string{ "unknown", },
		[]string{ "append-note", },
		[]string{ "list-notes", },
		[]int{ 60, },
	)

	ch := make(chan PipelineResult, 1)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	for result := range ch {
		count += 1
		assert.NotNil(t, result.Source)
		assert.NotNil(t, result.Err)
		assert.Equal(t, 0, len(result.Transformations))
		assert.Nil(t, result.Visualization)
		assert.Equal(t, "", result.Output)
	}
	assert.Equal(t, 1, count)
}

func TestPerformUnknownFirstSource(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]string{ "unknown", "empty", },
		[]string{ "append-note", "append-note", },
		[]string{ "list-notes", "list-notes", },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 3)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	errorCount := 0
	for result := range ch {
		count += 1
		if result.Err != nil {
			errorCount += 1
			assert.NotNil(t, result.Source)
			assert.NotNil(t, result.Err)
			assert.Equal(t, 0, len(result.Transformations))
			assert.Nil(t, result.Visualization)
			assert.Equal(t, "", result.Output)
		} else {
			assert.NotNil(t, result.Source)
			assert.Equal(t, 2, len(result.Transformations))
			assert.NotNil(t, result.Visualization)
			assert.Equal(t, "Track  0: 60 69\n", result.Output)
			assert.Nil(t, result.Err)
		}
	}
	assert.Equal(t, 1, errorCount)
	assert.Equal(t, 3, count)
}

func TestPerformUnknownFirstTransform(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]string{ "empty", "empty", },
		[]string{ "unknown", "append-note", },
		[]string{ "list-notes", "list-notes", },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 2)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	for result := range ch {
		count += 1
		assert.NotNil(t, result.Source)
		assert.Equal(t, 1, len(result.Transformations))
		assert.NotNil(t, result.Err)
		assert.Nil(t, result.Visualization)
		assert.Equal(t, "", result.Output)
	}
	assert.Equal(t, 2, count)
}

func TestPerformUnknownFirstVisu(t *testing.T) {
	registrator := mockRegistrator()
	visu, err := registrator.CreateVisualization("visu-error", make(map[string]interface{}))
	assert.Nil(t, visu)
	assert.NotNil(t, err)
	pipeline := createPipeline(
		registrator,
		[]string{ "empty", "empty", },
		[]string{ "append-note", "append-note", },
		[]string{ "unknown", "list-notes", },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 4)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	errorCount := 0
	for result := range ch {
		count += 1
		if result.Err != nil {
			errorCount += 1
			assert.NotNil(t, result.Source)
			assert.Equal(t, 2, len(result.Transformations))
			assert.NotNil(t, result.Visualization)
			assert.NotNil(t, result.Err)
			assert.Equal(t, "", result.Output)
		} else {
			assert.NotNil(t, result.Source)
			assert.Equal(t, 2, len(result.Transformations))
			assert.NotNil(t, result.Visualization)
			assert.Equal(t, "Track  0: 60 69\n", result.Output)
			assert.Nil(t, result.Err)
		}
	}
	assert.Equal(t, 2, errorCount)
	assert.Equal(t, 4, count)
}


func TestPerformErrorFirstSource(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]string{ "error", "empty", },
		[]string{ "append-note", "append-note", },
		[]string{ "list-notes", "list-notes", },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 3)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	errorCount := 0
	for result := range ch {
		count += 1
		if result.Err != nil {
			errorCount += 1
			assert.NotNil(t, result.Source)
			assert.NotNil(t, result.Err)
			assert.Equal(t, 0, len(result.Transformations))
			assert.Nil(t, result.Visualization)
			assert.Equal(t, "", result.Output)
		} else {
			assert.NotNil(t, result.Source)
			assert.Equal(t, 2, len(result.Transformations))
			assert.NotNil(t, result.Visualization)
			assert.Equal(t, "Track  0: 60 69\n", result.Output)
			assert.Nil(t, result.Err)
		}
	}
	assert.Equal(t, 1, errorCount)
	assert.Equal(t, 3, count)
}

func TestPerformErrorFirstTransformation(t *testing.T) {
	registrator := mockRegistrator()
	pipeline := createPipeline(
		registrator,
		[]string{ "empty", "empty", },
		[]string{ "error", "append-note", },
		[]string{ "list-notes", "list-notes", },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 2)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	for result := range ch {
		count += 1
		assert.NotNil(t, result.Source)
		assert.Equal(t, 1, len(result.Transformations))
		assert.NotNil(t, result.Err)
		assert.Nil(t, result.Visualization)
		assert.Equal(t, "", result.Output)
	}
	assert.Equal(t, 2, count)
}

func TestPerformErrorFirstVisualization(t *testing.T) {
	registrator := mockRegistrator()
	visu, err := registrator.CreateVisualization("visu-error", make(map[string]interface{}))
	assert.Nil(t, visu)
	assert.NotNil(t, err)
	pipeline := createPipeline(
		registrator,
		[]string{ "empty", "empty", },
		[]string{ "append-note", "append-note", },
		[]string{ "error", "list-notes", },
		[]int{ 60, 69, },
	)

	ch := make(chan PipelineResult, 4)
	pipeline.Perform(ch)
	close(ch)
	count := 0
	errorCount := 0
	for result := range ch {
		count += 1
		if result.Err != nil {
			errorCount += 1
			assert.NotNil(t, result.Source)
			assert.Equal(t, 2, len(result.Transformations))
			assert.NotNil(t, result.Visualization)
			assert.NotNil(t, result.Err)
			assert.Equal(t, "", result.Output)
		} else {
			assert.NotNil(t, result.Source)
			assert.Equal(t, 2, len(result.Transformations))
			assert.NotNil(t, result.Visualization)
			assert.Equal(t, "Track  0: 60 69\n", result.Output)
			assert.Nil(t, result.Err)
		}
	}
	assert.Equal(t, 2, errorCount)
	assert.Equal(t, 4, count)
}

