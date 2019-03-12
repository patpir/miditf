package core

type Track struct {
	notes  []Note
}

func NewTrack() *Track {
	return new(Track)
}

