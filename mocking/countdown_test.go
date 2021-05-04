package mocking

import "bytes"
import "testing"

func TestCountdown(t *testing.T) {
	t.Run("first test", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer)

		got := buffer.String()
		want := "3"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
