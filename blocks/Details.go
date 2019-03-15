package blocks

type Details struct {
	identifier   string
	description  string
	arguments    []Argument
}


func NewDetails(identifier string, description string, arguments []Argument) Details {
	return Details {
		identifier:   identifier,
		description:  description,
		arguments:    arguments,
	}
}


func (details *Details) Identifier() string {
	return details.identifier
}

func (details *Details) Description() string {
	return details.description
}

func (details *Details) Arguments() []Argument {
	return details.arguments
}

