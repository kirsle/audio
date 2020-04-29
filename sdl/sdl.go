// Package sdl implements an audio engine using libSDL2.
package sdl

import "github.com/veandco/go-sdl2/mix"

// Engine is the SDL2 audio engine.
type Engine struct {
	initFlags int
}

// New initializes an SDL2 Mixer for the audio engine.
//
// Pass the SDL2 Mixer flags for its initialization. The flags are an OR'd
// value made up of:
//    mix.INIT_MP3
//    mix.INIT_OGG
//    mix.INIT_FLAC
//    mix.INIT_MOD
func New(flags int) (*Engine, error) {
	return &Engine{
		initFlags: flags,
	}, nil
}

// Setup initializes SDL2 Mixer.
func (e *Engine) Setup() error {
	// Initialize SDL2 mixer.
	if err := mix.Init(e.initFlags); err != nil {
		return err
	}

	// Open the audio mixer.
	// the '2' is stereo (two channels), '1' would be mono.
	// 4096 is the chunk size.
	// https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_11.html
	return mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, 2, 4096)
}

// Playing returns if either music or sounds are currently playing.
func (e *Engine) Playing() bool {
	if mix.PlayingMusic() {
		return true
	}
	return mix.Playing(-1) > 0
}

// PlayingMusic returns if music is currently playing.
func (e *Engine) PlayingMusic() bool {
	return mix.PlayingMusic()
}

// PlayingSound returns if sounds are playing.
func (e *Engine) PlayingSound() bool {
	return mix.Playing(-1) > 0
}

// StopAll stops all music and sounds.
func (e *Engine) StopAll() {
	e.StopMusic()
	e.StopSounds()
}

// StopMusic stops all music.
func (e *Engine) StopMusic() {
	mix.HaltMusic()
}

// StopSounds stops all sounds.
func (e *Engine) StopSounds() {
	mix.HaltChannel(-1)
}

// Teardown closes the SDL mixer.
func (e *Engine) Teardown() error {
	mix.CloseAudio()
	mix.Quit()
	return nil
}
