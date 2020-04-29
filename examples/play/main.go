// Example program using the SDL2 Mix engine.
//
// Plays a music or sound file from disk.
// Usage: `play <filepath>`
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"git.kirsle.net/go/audio/sdl"
	"github.com/veandco/go-sdl2/mix"
)

// CLI flags.
var (
	flagBinary bool
	flagLoop   int
)

func init() {
	flag.BoolVar(&flagBinary, "binary", false,
		"Feed the file as bytes instead of by filename on disk.")
	flag.IntVar(&flagLoop, "loop", 0, "Number of times to loop the audio.")
}

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: play <path/to/file.mp3>")
		fmt.Println("Supports .ogg, .mp3, .wav")
		os.Exit(1)
	}

	sfx, err := sdl.New(mix.INIT_MP3 | mix.INIT_OGG)
	if err != nil {
		panic(err)
	}

	sfx.Setup()

	var (
		filename = os.Args[len(os.Args)-1]
		sound    sdl.Track
	)
	if flagBinary {
		fmt.Printf("Opening '%s' and feeding to engine as binary\n", filename)
		sound = loadBinary(sfx, filename)
	} else {
		fmt.Printf("Playing '%s' from filesystem\n", filename)
		sound = loadFilename(sfx, filename)
	}

	fmt.Printf("Begin playback")
	sound.Play(flagLoop)

	for sfx.Playing() {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}

	fmt.Println("Done.")
}

func loadFilename(sfx *sdl.Engine, filename string) sdl.Track {
	var (
		sound sdl.Track
		err   error
	)

	switch filepath.Ext(filename) {
	case ".ogg", ".mp3":
		sound, err = sfx.LoadMusic(filename)
		if err != nil {
			panic(err)
		}
	case ".wav":
		sound, err = sfx.LoadSound(filename)
		if err != nil {
			panic(err)
		}
	default:
		panic("Unsupported file type")
	}
	return sound
}

// Loads the audio file by sending the byte stream into the audio engine
// instead of having the audio engine open it from filesystem itself.
func loadBinary(sfx *sdl.Engine, filename string) sdl.Track {
	bin, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var sound sdl.Track

	switch filepath.Ext(filename) {
	case ".ogg", ".mp3":
		sound, err = sfx.LoadMusicBin(bin)
		if err != nil {
			panic(err)
		}
	case ".wav":
		sound, err = sfx.LoadSoundBin(bin)
		if err != nil {
			panic(err)
		}
	default:
		panic("Unsupported file type: " + filepath.Ext(filename))
	}
	return sound
}
