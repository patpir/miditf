package blocks

type BlockInfo struct {
	identifier     string
	description    string
	argumentInfos  []ArgumentInfo
}


func NewBlockInfo(identifier string, description string, argumentInfos []ArgumentInfo) BlockInfo {
	return BlockInfo {
		identifier:     identifier,
		description:    description,
		argumentInfos:  argumentInfos,
	}
}


func (b *BlockInfo) Identifier() string {
	return b.identifier
}

func (b *BlockInfo) Description() string {
	return b.description
}

func (b *BlockInfo) ArgumentInfos() []ArgumentInfo {
	return b.argumentInfos
}

