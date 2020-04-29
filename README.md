# audio: Simple Audio Engine for Go

Package `audio` is a simple audio engine for Go that can play some music and
sound files. It currently supports an SDL2 (Mixer) driver suitable for
use on desktop systems like Linux, Mac OS and Windows, with support to load
and play music files (.ogg and .mp3 format, depending on your system libraries)
and sound effects (.wav).

## Example

See the `examples/play/main.go` for a simple command-line media player sample
that uses the SDL2 engine.

```go
package main

import (
    "time"
    "git.kirsle.net/go/audio/sdl"
)

func main() {
    sfx, err := sdl.New(mix.INIT_MP3 | mix.INIT_OGG)
    if err != nil {
        panic(err)
    }

    // Call once at program startup.
    sfx.Setup()
    defer sfx.Teardown()

    // Load a file from disk.
    music, err := sfx.LoadMusic("filename.mp3")
    if err != nil {
        panic(err)
    }

    // Play it.
    music.Play(0)

    // Wait until done.
    for sfx.Playing() {
        time.Sleep(1 * time.Second)
    }
}
```

## License

MIT.
