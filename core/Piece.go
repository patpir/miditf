package core

type Piece struct {
	tracks  []*Track
}


func NewPiece() *Piece {
	return new(Piece)
}

func (p *Piece) AddTrack(track *Track) {
	p.tracks = append(p.tracks, track)
}

func (p *Piece) Tracks() []*Track {
	return p.tracks
}

