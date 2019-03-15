package core

type Track struct {
	notes  []*Note
}

func NewTrack() *Track {
	return new(Track)
}

func (t *Track) AddNote(note *Note) {
	t.notes = append(t.notes, note)
}

func (t *Track) Notes() []*Note {
	return t.notes
}

