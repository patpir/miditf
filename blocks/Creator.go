package blocks

type Creator interface {
	CreateSource(typeId string, args map[string]interface{}) (Source, error)
	CreateTransformation(typeId string, args map[string]interface{}) (Transformation, error)
	CreateVisualization(typeId string, args map[string]interface{}) (Visualization, error)
}

