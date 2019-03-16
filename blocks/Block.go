package blocks

type Block struct {
	typeId     string
	comment    string
	arguments  []Argument
}


func NewBlock(typeId string, comment string, arguments []Argument) Block {
	return Block {
		typeId:     typeId,
		comment:    comment,
		arguments:  arguments,
	}
}


func (b *Block) TypeId() string {
	return b.typeId
}

func (b *Block) Comment() string {
	return b.comment
}

func (b *Block) Arguments() []Argument {
	return b.arguments
}

