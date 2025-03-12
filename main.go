package functime

import (
	"fmt"
	"os"
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
	}
}

func (f *FuncTime) Stop() {
	stop := time.Since(f.start)

	fmt.Fprintf(os.Stderr, "\n\n function call %q took %d seconds \n\n",
		f.function, stop.Seconds())
}
