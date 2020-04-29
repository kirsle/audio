package audio_test

import (
	"testing"

	"git.kirsle.net/go/audio/null"
)

func TestNullEngine(t *testing.T) {
	null := null.New()
	null.Setup()
	null.Teardown()
}
