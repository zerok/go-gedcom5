package gedcom5

import (
	"bytes"
	"fmt"
)

type Line struct {
	Level  int
	XRefID string
	Tag    string
	Value  string
}

func (l *Line) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%d", l.Level))
	if l.XRefID != "" {
		out.WriteByte(0x20)
		out.WriteString(l.XRefID)
	}
	if l.Tag != "" {
		out.WriteByte(0x20)
		out.WriteString(l.Tag)
	}
	if l.Value != "" {
		out.WriteByte(0x20)
		out.WriteString(l.Value)
	}
	return out.String()
}
