package core

type Note struct {
	Channel    uint8
	Key        uint8
	Velocity   uint8
	StartTime  uint32
	EndTime    uint32
}

func NewNote() *Note {
	return new(Note)
}

