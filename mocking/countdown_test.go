package mocking

import "bytes"
import "testing"
import "time"

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
		Countdown(buffer, &CountdownOperationSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// t.Run("sleep before every print", func(t *testing.T) {
	//   spySleepPrinter := &CountdownOperationSpy{}
	//   Countdown(spySleepPrinter, spySleepPrinter)
	//
	//   want := []string{
	//     sleep,
	//     write,
	//     sleep,
	//     write,
	//     sleep,
	//     write,
	//     sleep,
	//     write,
	//   }
	//
	//   if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
	//     t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
	//   }
	// })
}

func TestConfigurationSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
