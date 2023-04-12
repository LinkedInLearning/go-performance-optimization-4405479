package validate

import (
	"testing"
	"time"
)

var (
	events []Event
)

func init() {
	const size = 10_000
	events = make([]Event, size)
	now := time.Now()
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			events[i].Time = now
			events[i].Action = "write"
			events[i].Login = "elliot"
			events[i].Path = "/etc/passwd"
		}
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out := Validate(events)
		if size := len(out); size != len(events) {
			b.Fatal(size)
		}
	}
}
