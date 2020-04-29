package audio

// Engine is a music and sound effects driver.
type Engine interface {
	// Setup runs initialization tasks for the audio engine.
	Setup() error

	// Teardown runs closing tasks for the audio engine to shut down gracefully.
	Teardown() error

	// Playing returns a bool if something is actively playing.
	// PlayingMusic and PlayingSound check specifically if music or sound
	// effects are currently playing.
	Playing() bool
	PlayingMusic() bool
	PlayingSound() bool

	// StopAll stops all music and sound effects.
	// StopMusic and StopSounds to selectively stop either the music or all
	// sound effects, respectively.
	StopAll()
	StopMusic()
	StopSounds()

	// LoadMusic opens a music file from disk and loads it into memory.
	// LoadMusicBin to load file by bytes in memory instead.
	LoadMusic(filename string) (Playable, error)
	LoadMusicBin(data []byte) (Playable, error)

	// LoadSound opens a sound effect file.
	// LoadSoundBin to load file by bytes in memory instead.
	LoadSound(filename string) (Playable, error)
	LoadSoundBin(data []byte) (Playable, error)
}

// Playable is a music or sound effect object that can be played and managed.
type Playable interface {
	Play(loops int) error
	Pause() error
	Stop() error

	// Destroy deallocates and frees the memory.
	Destroy() error
}
