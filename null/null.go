// Package null implements a dummy audio driver that doesn't play any audio.
package null

// Engine is a null audio engine.
type Engine struct{}

// Playable is a null music or sound effect.
type Playable struct{}

// New creates a null engine.
func New() *Engine {
	return &Engine{}
}

// Setup the null engine (do nothing).
func (e *Engine) Setup() error {
	return nil
}

// Teardown the null engine (do nothing).
func (e *Engine) Teardown() error {
	return nil
}

// Playing returns false.
func (e *Engine) Playing() bool {
	return false
}

// PlayingMusic returns false.
func (e *Engine) PlayingMusic() bool {
	return false
}

// PlayingSound returns false.
func (e *Engine) PlayingSound() bool {
	return false
}

// StopAll does nothing.
func (e *Engine) StopAll() {}

// StopMusic does nothing.
func (e *Engine) StopMusic() {}

// StopSounds does nothing.
func (e *Engine) StopSounds() {}

// LoadMusic loads nothing.
func (e *Engine) LoadMusic(filename string) (Playable, error) {
	return Playable{}, nil
}

// LoadMusicBin loads nothing.
func (e *Engine) LoadMusicBin(data []byte) (Playable, error) {
	return Playable{}, nil
}

// LoadSound loads nothing.
func (e *Engine) LoadSound(filename string) (Playable, error) {
	return Playable{}, nil
}

// LoadSoundBin loads nothing.
func (e *Engine) LoadSoundBin(data []byte) (Playable, error) {
	return Playable{}, nil
}

// Play nothing.
func (p Playable) Play(loops int) error {
	return nil
}

// Pause nothing
func (p Playable) Pause() error {
	return nil
}

// Stop nothing.
func (p Playable) Stop() error {
	return nil
}

// Destroy nothing.
func (p Playable) Destroy() error {
	return nil
}
