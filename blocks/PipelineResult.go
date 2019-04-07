package blocks

type PipelineResult struct {
	Source           Block
	Transformations  []Block
	Visualization    Block
	Output           string
	Err              error
}

