package blocks

type Argument struct {
	name         string
	description  string
}


func NewArgument(name string, description string) Argument {
	return Argument {
		name:         name,
		description:  description,
	}
}


func (arg *Argument) Name() string {
	return arg.name
}

func (arg *Argument) Description() string {
	return arg.description
}

