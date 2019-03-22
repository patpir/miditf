package visualize

import (
	"bytes"
	"sort"

	"github.com/gomidi/midi/midimessage/channel"
	"github.com/gomidi/midi/midimessage/meta"
	"github.com/gomidi/midi/smf"
	"github.com/gomidi/midi/smf/smfwriter"

	"github.com/patpir/miditf/core"
	"github.com/patpir/miditf/blocks"
)

type midiFileWriter struct {
}


func newMidiFileWriter(arguments map[string]interface{}) (blocks.Visualization, error) {
	writer := midiFileWriter{}
	return &writer, nil
}


func (w *midiFileWriter) Visualize(piece *core.Piece) (string, error) {

	resolution := smf.MetricTicks(96)

	tracks := piece.Tracks()
	trackCount := len(tracks)

	var bf bytes.Buffer

	wr := smfwriter.New(&bf, smfwriter.NumTracks(uint16(trackCount)), smfwriter.TimeFormat(resolution))
	wr.Write(meta.TimeSig{
		Numerator:                4,
		Denominator:              4,
		ClocksPerClick:           24,
		DemiSemiQuaverPerQuarter: 8,
	})
	wr.Write(meta.BPM(120))

	sortByStartTimeAsc := func(notes []*core.Note) []*core.Note {
		sort.Slice(notes, func(i, j int) bool {
			return notes[i].StartTime() < notes[j].StartTime()
		})
		return notes
	}
	sortByEndTimeAsc := func(notes []*core.Note) []*core.Note {
		sort.Slice(notes, func(i, j int) bool {
			return notes[i].EndTime() < notes[j].EndTime()
		})
		return notes
	}

	for _, track := range tracks {
		noteStartTimes := sortByStartTimeAsc(track.Notes())
		noteEndTimes := sortByEndTimeAsc(track.Notes())

		lastTime := uint32(0)
		startIndex := 0
		endIndex := 0

		channels := []channel.Channel{
			channel.Channel0,
		}

		for startIndex < len(noteStartTimes) && endIndex < len(noteEndTimes) {
			if noteEndTimes[endIndex].EndTime() < noteStartTimes[startIndex].StartTime() {
				note := noteEndTimes[endIndex]
				if lastTime != note.EndTime() {
					delta := note.EndTime() - lastTime
					wr.SetDelta(resolution.Ticks64th() * delta)
					lastTime = note.EndTime()
				}
				wr.Write(channels[note.Channel()].NoteOff(note.Key()))
				endIndex += 1
			} else {
				note := noteStartTimes[startIndex]
				if lastTime != note.StartTime() {
					delta := note.StartTime() - lastTime
					wr.SetDelta(resolution.Ticks64th() * delta)
					lastTime = note.StartTime()
				}
				wr.Write(channels[note.Channel()].NoteOn(note.Key(), note.Velocity()))
				startIndex += 1
			}
		}
		// at this point there can only be unfinished notes -> in the noteEndTimes list
		for endIndex < len(noteEndTimes) {
			note := noteEndTimes[endIndex]
			if lastTime != note.EndTime() {
				delta := note.EndTime() - lastTime
				wr.SetDelta(resolution.Ticks64th() * delta)
				lastTime = note.EndTime()
			}
			wr.Write(channels[note.Channel()].NoteOff(note.Key()))
			endIndex += 1
		}

		wr.Write(meta.EndOfTrack)
	}

	return bf.String(), nil
}


func init() {
	arguments := []blocks.ArgumentInfo{}
	info := blocks.NewBlockInfo("midi-file", "Converts notes to MIDI file format", arguments)
	blocks.DefaultRegistrator().RegisterVisualization(info, newMidiFileWriter)
}

