package core

type Note struct {
	channel    uint8
	key        uint8
	velocity   uint8
	startTime  uint32
	endTime    uint32
}

func NewNote(channel uint8, key uint8, velocity uint8, startTime uint32, endTime uint32) *Note {
	note := new(Note)
	note.channel = channel
	note.key = key
	note.velocity = velocity
	note.startTime = startTime
	note.endTime = endTime
	return note
}

func (n *Note) Channel() uint8 {
	return n.channel
}

func (n *Note) Key() uint8 {
	return n.key
}

func (n *Note) Velocity() uint8 {
	return n.velocity
}

func (n *Note) StartTime() uint32 {
	return n.startTime
}

func (n *Note) EndTime() uint32 {
	return n.endTime
}

