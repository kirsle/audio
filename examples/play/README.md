# Example: `play`

This example uses the SDL2 Mixer engine to implement a _simple_ command line
program that plays music and sound files.

## Usage

```
play [options] path/to/file.ogg
```

Supports file types `.ogg`, `.mp3` and `.wav`.

Default behavior calls LoadSound() or LoadMusic() using the filename given.
Use the `-binary` option to go through LoadSoundBin() or LoadMusicBin() to
test initializing it by bytes array instead.

### Options

```
-binary
  Opens the file first as a bytes array, and feeds it to the audio engine
  as binary data instead of by passing a filename on disk.

-loop <int>
  Loop the audio file, default 0 will only play it once. -1 for music will
  play forever.
```
