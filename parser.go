package gedcom5

import (
	"bytes"
	"io"
)

type FileParser struct {
	lp *LineParser
}

func NewFileParser(s *Scanner) *FileParser {
	return &FileParser{
		lp: NewLineParser(s),
	}
}

func (fp *FileParser) ParseFile() (*File, error) {
	f := File{
		Lines: make([]Line, 0, 10),
	}
	for {
		l, err := fp.lp.ParseLine()
		if err == io.EOF {
			break
		}
		f.Lines = append(f.Lines, *l)
	}
	return &f, nil
}

type LineParser struct {
	s *Scanner
}

func NewLineParser(s *Scanner) *LineParser {
	return &LineParser{s: s}
}

// ParseLine tries to parse a single line from the underlying input.
func (lp *LineParser) ParseLine() (*Line, error) {
	line := &Line{}
	_, err := lp.s.PeekByte()
	if err != nil {
		if err == io.EOF {
			return nil, io.EOF
		}
	}
	num, err := lp.s.ReadNumber()
	if err != nil {
		return nil, err
	}
	line.Level = num

	if err := lp.s.ReadByte(0x20); err != nil {
		return nil, err
	}

	if m, _ := lp.s.PeekBytesEquals([]byte{0x40}); m {
		xrefID, err := lp.s.ReadXRefID()
		if err != nil {
			return nil, err
		}
		line.XRefID = xrefID
		if err := lp.s.ReadByte(0x20); err != nil {
			return nil, err
		}
	}

	// Read the tag
	tag, err := lp.s.ReadAlphaNum()
	if err != nil {
		return nil, err
	}
	line.Tag = tag

	if m, _ := lp.s.PeekBytesEquals([]byte{0x20}); m {
		// We are now dealing with a value and can continue reading
		// until the first terminator.
		lp.s.ConsumeByte()
		var val bytes.Buffer
		if err := lp.s.ConsumeUntilInto(&val, IsEither(IsExact(0xA), IsExact(0xD))); err != nil {
			return nil, err
		}
		line.Value = val.String()
	}

	b, err := lp.s.PeekByte()
	if b == 0xD {
		lp.s.ConsumeByte()
		if n, _ := lp.s.PeekByte(); n == 0xA {
			lp.s.ConsumeByte()
		}
		return line, nil
	} else if b == 0xA {
		lp.s.ConsumeByte()
		if n, _ := lp.s.PeekByte(); n == 0xD {
			lp.s.ConsumeByte()
		}
		return line, nil
	}
	return nil, ErrUnexpectedToken
}

func OwnParseLine(in io.Reader) ([]Line, error) {
	s := NewScanner(in)
	p := &LineParser{s: s}
	result := make([]Line, 0, 10)
	for {
		l, err := p.ParseLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		result = append(result, *l)
	}
	return result, nil
}
