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

func (n *Note) SetChannel(channel uint8) {
	n.channel = channel
}

func (n *Note) Key() uint8 {
	return n.key
}

func (n *Note) SetKey(key uint8) {
	n.key = key
}

func (n *Note) Velocity() uint8 {
	return n.velocity
}

func (n *Note) SetVelocity(velocity uint8) {
	n.velocity = velocity
}

func (n *Note) StartTime() uint32 {
	return n.startTime
}

func (n *Note) SetStartTime(startTime uint32) {
	n.startTime = startTime
}

func (n *Note) EndTime() uint32 {
	return n.endTime
}

func (n *Note) SetEndTime(endTime uint32) {
	n.endTime = endTime
}

