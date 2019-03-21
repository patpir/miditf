package blocks

type ArgumentInfo struct {
	name         string
	description  string
	optional     bool
}


func NewArgumentInfo(name string, description string, optional bool) ArgumentInfo {
	return ArgumentInfo {
		name:         name,
		description:  description,
		optional:     optional,
	}
}


func (arg *ArgumentInfo) Name() string {
	return arg.name
}

func (arg *ArgumentInfo) Description() string {
	return arg.description
}

func (arg *ArgumentInfo) IsOptional() bool {
	return arg.optional
}

