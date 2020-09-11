package progress

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	count := 21
	bar := New(count)

	for i := 0; i < count; i++ {
		time.Sleep(10 * time.Millisecond)
		bar.Incr()
	}

	count = 211
	bar = New(count)
	bar.SetGraph("▫")

	for i := 0; i <= count; i++ {
		time.Sleep(1 * time.Millisecond)
		bar.Incr()
	}

}
