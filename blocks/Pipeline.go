package blocks

import (
	"sync"

	"github.com/patpir/miditf/core"
)

type Pipeline struct {
	creator         Creator
	sources         []Block
	transformations []Block
	visualizations  []Block
}

type intermediateResult struct {
	source           Block
	transformations  []Block
	piece            *core.Piece
	err              error
}


func NewPipeline() *Pipeline {
	return new(Pipeline)
}


func (p *Pipeline) Sources() []Block {
	return p.sources
}

func (p *Pipeline) AddSource(block Block) {
	p.sources = append(p.sources, block)
}

func (p *Pipeline) Transformations() []Block {
	return p.transformations
}

func (p *Pipeline) AddTransformation(block Block) {
	p.transformations = append(p.transformations, block)
}

func (p *Pipeline) Visualizations() []Block {
	return p.visualizations
}

func (p *Pipeline) AddVisualization(block Block) {
	p.visualizations = append(p.visualizations, block)
}


func (p *Pipeline) Perform(c chan PipelineResult) {
	var wg sync.WaitGroup
	for _, s := range p.sources {
		wg.Add(1)
		go func(source Block, transformations []Block) {
			defer wg.Done()
			transformed := transformSingleSource(p.creator, source, transformations)

			if transformed.err != nil {
				result := PipelineResult{
					source: transformed.source,
					transformations: transformed.transformations,
					visualization: nil,
					output: "",
					err: transformed.err,
				}
				c <- result
			} else {
				for _, v := range p.visualizations {
					wg.Add(1)
					go func(visualization Block) {
						defer wg.Done()
						c <- visualizeTransformationResult(p.creator, transformed, visualization)
					}(v)
				}
			}
		}(s, p.transformations)
	}
	wg.Wait()
}


func transformSingleSource(creator Creator, source Block, transformations []Block) intermediateResult {
	result := intermediateResult{
		source: source,
		transformations: nil,
		piece: nil,
		err: nil,
	}
	src, err := creator.CreateSource(source.TypeId(), source.Arguments())
	if err != nil {
		result.err = err
		return result
	} else {
		result.piece = src.Piece()
		for _, transformation := range transformations {
			result.transformations = append(result.transformations, transformation)
			transform, err := creator.CreateTransformation(transformation.TypeId(), transformation.Arguments())
			if err != nil {
				result.err = err
				break
			}
			result.piece = transform.Transform(result.piece)
		}
	}
	return result
}

func visualizeTransformationResult(creator Creator, tfResult intermediateResult, visualization Block) PipelineResult {
	result := PipelineResult{
		source: tfResult.source,
		transformations: tfResult.transformations,
		visualization: visualization,
		output: "",
		err: nil,
	}

	visu, err := creator.CreateVisualization(visualization.TypeId(), visualization.Arguments())
	if err != nil {
		result.err = err
	} else {
		result.output = visu.Visualize(tfResult.piece)
	}

	return result
}

