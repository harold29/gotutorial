package mocking

import "bytes"
import "testing"

func TestCountdown(t *testing.T) {
	// t.Run("first test", func(t *testing.T) {
	//   buffer := &bytes.Buffer{}
	//
	//   Countdown(buffer)
	//
	//   got := buffer.String()
	//   want := "3"
	//
	//   if got != want {
	//     t.Errorf("got %q want %q", got, want)
	//   }
	// })

	t.Run("second test", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		}
	})
}
