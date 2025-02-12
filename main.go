package functime

import (
	"log/slog"
	"time"
)

type FuncTime struct {
	start    time.Time
	function string
}

func Start(function string) *FuncTime {
	return &FuncTime{
		start:    time.Now(),
		function: function,
		log:      log,
	}
}

func (f *FuncTime) Stop() {
	stop := time.Since(f.start)

	slog.Debug("function call",
		"name", f.function, "miliseconds", stop.Milliseconds(), "seconds", stop.Seconds())
}
