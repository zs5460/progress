package progress

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	count := 5460
	bar := New(count)
	//bar.SetGraph("▫")

	for i := 0; i <= count; i++ {
		time.Sleep(1 * time.Millisecond)
		bar.Incr()
	}
}
