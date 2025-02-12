# go-functime

It's not a profiler, just dumb function measurements. You may find it handy when
you work in an environment with no observability infrastructure yet and want to
debug something locally.

## Example

```go
package main

import (
	"github.com/tyuio9/go-functime"
)

type Foo struct {
	v int
}

func (f *Foo) Process() {
	// A call to Stop() will emit a log at DEBUG level.
	//
	defer functime.Start("Foo.Process").Stop()

	/* magic */
}
```
