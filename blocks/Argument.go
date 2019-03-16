package blocks

type Argument struct {
	name   string
	value  string
}


func NewArgument(name string, value string) Argument {
	return Argument {
		name:   name,
		value:  value,
	}
}


func (arg *Argument) Name() string {
	return arg.name
}

func (arg *Argument) Value() string {
	return arg.value
}

