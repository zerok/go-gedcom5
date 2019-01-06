package gedcom5

import "io"

// LineWriter is used as output interface for serializing lines.
type LineWriter interface {
	WriteLine(Line) error
	IsEmpty() bool
}

type simpleLineWriter struct {
	out      io.Writer
	notEmpty bool
}

func newSimpleLineWriter(out io.Writer) *simpleLineWriter {
	return &simpleLineWriter{
		out: out,
	}
}

func (lw *simpleLineWriter) WriteLine(l Line) error {
	if l.Value != "" {
		lw.notEmpty = true
	}
	if _, err := lw.out.Write([]byte(l.String())); err != nil {
		return err
	}
	if _, err := lw.out.Write([]byte{'\n'}); err != nil {
		return err
	}
	return nil
}

func (lw *simpleLineWriter) IsEmpty() bool {
	return !lw.notEmpty
}

type bufferedLineWriter struct {
	lines []Line
}

func newBufferedLineWriter() *bufferedLineWriter {
	return &bufferedLineWriter{
		lines: make([]Line, 0, 10),
	}
}

func (lw *bufferedLineWriter) WriteLine(l Line) error {
	lw.lines = append(lw.lines, l)
	return nil
}

func (lw *bufferedLineWriter) IsEmpty() bool {
	for _, l := range lw.lines {
		if l.Value != "" {
			return false
		}
	}
	return true
}

func (lw *bufferedLineWriter) WriteTo(out LineWriter) error {
	for _, l := range lw.lines {
		if err := out.WriteLine(l); err != nil {
			return err
		}
	}
	return nil
}
