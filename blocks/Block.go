package blocks

type Block interface {
	TypeId()     string
	Comment()    string
	Arguments()  []Argument
}

type block struct {
	typeId     string
	comment    string
	arguments  []Argument
}


func NewBlock(typeId string, comment string, arguments []Argument) Block {
	return &block {
		typeId:     typeId,
		comment:    comment,
		arguments:  arguments,
	}
}


func (b *block) TypeId() string {
	return b.typeId
}

func (b *block) Comment() string {
	return b.comment
}

func (b *block) Arguments() []Argument {
	return b.arguments
}

