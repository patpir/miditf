package blocks

type ArgumentInfo struct {
	name         string
	description  string
}


func NewArgumentInfo(name string, description string) ArgumentInfo {
	return ArgumentInfo {
		name:         name,
		description:  description,
	}
}


func (arg *ArgumentInfo) Name() string {
	return arg.name
}

func (arg *ArgumentInfo) Description() string {
	return arg.description
}

