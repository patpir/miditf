package sources

import (
	"errors"
	"os"

	"github.com/gomidi/midi/midimessage/channel"
	"github.com/gomidi/midi/smf"
	"github.com/gomidi/midi/smf/smfreader"

	"github.com/patpir/miditf/core"
	"github.com/patpir/miditf/blocks"
)

type midiFileReader struct {
	filename string
}


func newMidiFileReader(arguments []blocks.Argument) (blocks.Source, error) {
	if len(arguments) != 1 {
		return nil, errors.New("Invalid number of arguments")
	}

	ms := midiFileReader{}
	ms.filename = arguments[0].Value()
	return &ms, nil
}


func (mfr *midiFileReader) Piece() (*core.Piece, error) {
	result := core.NewPiece()

	// 16 MIDI channels and 128 MIDI keys
	var channelKeyNotes [16][128]*core.Note

	currentTime := uint32(0)
	trackId := int16(-1)
	var currentTrack *core.Track = nil

	f, err := os.Open(mfr.filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rd := smfreader.New(f)

	err = rd.ReadHeader()
	if err != nil {
		return nil, err
	}

	header := rd.Header()

	for {
		msg, err := rd.Read()
		if err != nil {
			break
		}

		currentTime += rd.Delta()

		if rd.Track() != trackId {
			currentTrack = core.NewTrack()
			result.AddTrack(currentTrack)
			trackId = rd.Track()
		}

		switch m := msg.(type) {
		case channel.NoteOn:
			c, k := m.Channel(), m.Key()
			startTime := header.TimeFormat.(smf.MetricTicks).In64ths(currentTime)
			channelKeyNotes[c][k] = core.NewNote(
				c,
				k,
				m.Velocity(),
				startTime,
				uint32(0),
			)
		case channel.NoteOff:
			c, k := m.Channel(), m.Key()
			endTime := header.TimeFormat.(smf.MetricTicks).In64ths(currentTime)
			currentTrack.AddNote(core.NewNote(
				c,
				k,
				channelKeyNotes[c][k].Velocity(),
				channelKeyNotes[c][k].StartTime(),
				endTime,
			))
			channelKeyNotes[c][k] = nil
		}
	}

	return result, nil
}


func init() {
	arguments := []blocks.ArgumentInfo{
		blocks.NewArgumentInfo("filename", "Path to a SMF/MIDI file", false),
	}
	info := blocks.NewBlockInfo("midi-file", "Reads a single MIDI file", arguments)
	blocks.DefaultRegistrator().RegisterSource(info, newMidiFileReader)
}

