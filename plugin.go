package logrus

import "io"

var (
	GlobalFormatterHook = func(f Formatter) Formatter { return f }
	GlobalWriterHook    = func(w io.Writer) io.Writer { return w }
)
