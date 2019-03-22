package blocks

type Block interface {
	TypeId()     string
	Comment()    string
	Arguments()  map[string]interface{}
}

type block struct {
	typeId     string
	comment    string
	arguments  map[string]interface{}
}


func NewBlock(typeId string, comment string, arguments map[string]interface{}) Block {
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

func (b *block) Arguments() map[string]interface{} {
	return b.arguments
}

