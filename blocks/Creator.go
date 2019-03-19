package blocks

type Creator interface {
	CreateSource(typeId string, args []Argument) (Source, error)
	CreateTransformation(typeId string, args []Argument) (Transformation, error)
	CreateVisualization(typeId string, args []Argument) (Visualization, error)
}

