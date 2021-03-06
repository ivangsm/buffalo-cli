package i18n

import (
	"io"
	"os"
)

type namedWriter struct {
	name string
	w    io.Writer
	err  error
}

func (namedWriter) PluginName() string {
	return "namedWriter"
}

var _ NamedWriter = &namedWriter{}

func (f *namedWriter) NamedWriter(filename string) (io.Writer, error) {
	f.name = filename
	if f.w == nil {
		return os.Stdout, f.err
	}
	return f.w, f.err
}
