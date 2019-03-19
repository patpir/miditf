package blocks

import (
	"errors"
)

type Registrator interface {
	RegisterSource(info BlockInfo, factory SourceFactory)
	RegisterTransformation(info BlockInfo, factory TransformationFactory)
	RegisterVisualization(info BlockInfo, factory VisualizationFactory)
	Sources()          []BlockInfo
	Transformations()  []BlockInfo
	Visualizations()   []BlockInfo
	Creator
}


type registrator struct {
	sources              map[string]sourceType
	transformations      map[string]transformationType
	visualizations       map[string]visualizationType
	sourceInfos          []BlockInfo
	transformationInfos  []BlockInfo
	visualizationInfos   []BlockInfo
}

func (r *registrator) RegisterSource(info BlockInfo, factory SourceFactory) {
	r.sources[info.Identifier()] = sourceType{
		info:    info,
		factory: factory,
	}
	r.sourceInfos = append(r.sourceInfos, info)
}

func (r *registrator) RegisterTransformation(info BlockInfo, factory TransformationFactory) {
	r.transformations[info.Identifier()] = transformationType{
		info:    info,
		factory: factory,
	}
	r.transformationInfos = append(r.transformationInfos, info)
}

func (r *registrator) RegisterVisualization(info BlockInfo, factory VisualizationFactory) {
	r.visualizations[info.Identifier()] = visualizationType{
		info:    info,
		factory: factory,
	}
	r.visualizationInfos = append(r.visualizationInfos, info)
}

func (r *registrator) CreateSource(identifier string, argValues []Argument) (Source, error) {
	st, ok := r.sources[identifier]
	if ok {
		return st.factory(argValues)
	}
	return nil, errors.New("Source type does not exist")
}

func (r *registrator) CreateTransformation(identifier string, argValues []Argument) (Transformation, error) {
	tt, ok := r.transformations[identifier]
	if ok {
		return tt.factory(argValues)
	}
	return nil, errors.New("Transformation type does not exist")
}


func (r *registrator) CreateVisualization(identifier string, argValues []Argument) (Visualization, error) {
	vt, ok := r.visualizations[identifier]
	if ok {
		return vt.factory(argValues)
	}
	return nil, errors.New("Visualization type does not exist")
}

func (r *registrator) Sources() []BlockInfo {
	return r.sourceInfos
}

func (r *registrator) Transformations() []BlockInfo {
	return r.transformationInfos
}

func (r *registrator) Visualizations() []BlockInfo {
	return r.visualizationInfos
}


var defaultRegistrator *registrator = &registrator{
	sources: map[string]sourceType{},
	transformations: map[string]transformationType{},
	visualizations: map[string]visualizationType{},
}

func DefaultRegistrator() Registrator {
	return defaultRegistrator
}

func NewRegistrator() Registrator {
	r := new(registrator)
	r.sources = map[string]sourceType{}
	r.transformations = map[string]transformationType{}
	r.visualizations = map[string]visualizationType{}
	return r
}

