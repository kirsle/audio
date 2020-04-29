package sdl

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// Track is a music or sound effect file.
type Track struct {
	isMusic bool // false = is sound effect

	// If isMusic
	mus *mix.Music

	// Sound effect
	wav     *mix.Chunk
	channel int
}

// LoadMusic loads a music file from disk.
func (e *Engine) LoadMusic(filename string) (Track, error) {
	mus, err := mix.LoadMUS(filename)
	return Track{
		isMusic: true,
		mus:     mus,
	}, err
}

// LoadMusicBin loads a music file from bytes data.
func (e *Engine) LoadMusicBin(data []byte) (Track, error) {
	// Create an SDL RWOps from the binary.
	rw, err := sdl.RWFromMem(data)
	if err != nil {
		return Track{}, err
	}

	mus, err := mix.LoadMUSRW(rw, 0)
	return Track{
		isMusic: true,
		mus:     mus,
	}, err
}

// LoadSound loads a wave file from disk.
func (e *Engine) LoadSound(filename string) (Track, error) {
	wav, err := mix.LoadWAV(filename)
	return Track{
		isMusic: false,
		wav:     wav,
		channel: -1,
	}, err
}

// LoadSoundBin loads a wave file from bytes data.
func (e *Engine) LoadSoundBin(data []byte) (Track, error) {
	// Create an SDL RWOps from the binary.
	rw, err := sdl.RWFromMem(data)
	if err != nil {
		return Track{}, err
	}

	wav, err := mix.LoadWAVRW(rw, false)
	return Track{
		isMusic: false,
		wav:     wav,
		channel: -1,
	}, err
}

// Play the track.
func (t *Track) Play(loops int) error {
	if t.isMusic {
		return t.mus.Play(loops)
	}

	// Normalize the `loops` value for sound effects to work around
	// a quirk in the SDL mixer between Music and Sound behaviors:
	//
	// For music:
	//    loops=0 plays it one time
	//    loops=1 plays it one time
	//    loops=2 plays it twice
	// For sounds:
	//    loops=0 plays it one time
	//    loops=1 plays it twice!
	//    loops=2 plays it three times!
	//
	// For both, a loops of -1 plays it on an infinite loop. So to make
	// the API consistent on our end, subtract 1 from a Sound loop only
	// when the given value is >= 1 itself.
	if loops > 0 {
		loops--
	}
	channel, err := t.wav.Play(-1, loops)
	t.channel = channel
	return err
}

// Pause the track.
func (t Track) Pause() error {
	if t.isMusic {
		mix.PauseMusic()
		return nil
	}
	mix.Pause(t.channel)
	return nil
}

// Stop the track.
func (t Track) Stop() error {
	if t.isMusic {
		mix.HaltMusic()
		return nil
	}
	mix.HaltChannel(t.channel)
	return nil
}

// Destroy the track.
func (t Track) Destroy() error {
	if t.isMusic {
		t.mus.Free()
		return nil
	}
	t.wav.Free()
	return nil
}
