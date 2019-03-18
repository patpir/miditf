package blocks

type PipelineResult struct {
	source           Block
	transformations  []Block
	visualization    Block
	output           string
	err              error
}


func (pr *PipelineResult) Source() Block {
	return pr.source
}

func (pr *PipelineResult) Transformations() []Block {
	return pr.transformations
}

func (pr *PipelineResult) Visualization() Block {
	return pr.visualization
}

func (pr *PipelineResult) Output() string {
	return pr.output
}

func (pr *PipelineResult) Err() error {
	return pr.err
}

