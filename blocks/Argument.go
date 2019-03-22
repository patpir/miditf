package blocks

type Argument struct {
	name   string
	value  interface{}
}


func NewArgument(name string, value interface{}) Argument {
	return Argument {
		name:   name,
		value:  value,
	}
}


func (arg *Argument) Name() string {
	return arg.name
}

func (arg *Argument) Value() interface{} {
	return arg.value
}

