package gedcom5

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

type ScanError struct {
	pos int
	msg string
}

var ErrUnexpectedToken = errors.New("unexpected token")

func (e *ScanError) Error() string {
	return fmt.Sprintf("scan failed at %d: %s", e.pos, e.msg)
}

type Token int

type Scanner struct {
	r   *bufio.Reader
	err error
	pos int
}

func NewScanner(in io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(in),
	}
}

func (s *Scanner) Pos() int {
	return s.pos
}

func (s *Scanner) Failed() bool {
	return s.err != nil
}

func (s *Scanner) Error() error {
	return s.err
}

// PeekBytesEquals peeks at the next len(needle) bytes and returns
// true if these next bytes match the needle.
//
// This does not update the failed status!
func (s *Scanner) PeekBytesEquals(needle []byte) (bool, error) {
	if s.Failed() {
		return false, s.err
	}
	actual, err := s.r.Peek(len(needle))
	if err != nil {
		return false, err
	}
	for i := 0; i < len(needle); i++ {
		if needle[i] != actual[i] {
			return false, nil
		}
	}
	return true, nil
}

func (s *Scanner) PeekByte() (byte, error) {
	if s.Failed() {
		return 0, s.Error()
	}
	next, err := s.r.Peek(1)
	if err != nil {
		return 0, err
	}
	return next[0], nil
}

func (s *Scanner) ConsumeByte() (byte, error) {
	if s.Failed() {
		return 0, s.Error()
	}
	s.pos += 1
	b, err := s.r.ReadByte()
	if err != nil {
		s.fail(err)
	}
	return b, err
}

func (s *Scanner) ConsumeByteInto(w io.Writer) error {
	if s.Failed() {
		return s.Error()
	}
	b, err := s.ConsumeByte()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte{b})
	return err
}

func (s *Scanner) ConsumeUntilInto(w io.Writer, delimCheck ByteTypeChecker) error {
	if s.Failed() {
		return s.Error()
	}
	for {
		next, err := s.PeekByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if delimCheck(next) {
			return nil
		}
		if _, err = w.Write([]byte{next}); err != nil {
			return err
		}
		s.pos += 1
		if _, err := s.ConsumeByte(); err != nil {
			return err
		}
	}
	return nil
}

type ByteTypeChecker func(byte) bool

func (s *Scanner) PeekByteWithType(check ...ByteTypeChecker) (byte, error) {
	if s.Failed() {
		return 0, s.Error()
	}
	b, err := s.PeekByte()
	if err != nil {
		return 0, err
	}
	for _, ch := range check {
		if ch(b) {
			return b, nil
		}
	}
	return 0, ErrUnexpectedToken
}

func (s *Scanner) fail(err error) {
	if s.err == nil {
		s.err = err
	}
}

func (s *Scanner) ReadXRefID() (string, error) {
	var res bytes.Buffer
	_, err := s.PeekByteWithType(IsAt)
	if err != nil {
		return "", err
	}
	s.ConsumeByteInto(&res)

	s.ConsumeUntilInto(&res, IsAt)

	_, err = s.PeekByteWithType(IsAt)
	if err != nil {
		return "", err
	}
	s.ConsumeByteInto(&res)
	return res.String(), nil
}

func (s *Scanner) ReadAlphaNum() (string, error) {
	var res bytes.Buffer
	for {
		b, err := s.PeekByteWithType(IsAlphaNum)
		if err != nil {
			if err == ErrUnexpectedToken {
				break
			}
			return "", err
		}
		if _, err := s.ConsumeByte(); err != nil {
			return "", err
		}
		res.WriteByte(b)
	}
	return res.String(), nil
}

// ReadNumber consumes a series of digits and returns an error
func (s *Scanner) ReadNumber() (int, error) {
	if s.Failed() {
		return 0, s.Error()
	}
	next, err := s.PeekByte()
	if err != nil {
		return -1, err
	}
	if !IsDigit(next) {
		return -1, s.scanError("digit expected")
	}
	var res bytes.Buffer
	res.WriteByte(next)
	if _, err := s.ConsumeByte(); err != nil {
		return -1, err
	}
	for {
		b, err := s.PeekByteWithType(IsDigit)
		if isTokenEnd(err) {
			break
		}
		if err != nil {
			return 0, err
		}
		res.WriteByte(b)
		s.ConsumeByte()
	}
	return strconv.Atoi(res.String())

}
func isTokenEnd(err error) bool {
	if err != nil {
		if err == ErrUnexpectedToken || err == io.EOF {
			return true
		}
	}
	return false
}

func (s *Scanner) ReadByte(b byte) error {
	if s.Failed() {
		return s.Error()
	}
	found, err := s.PeekByte()
	if err != nil {
		return err
	}
	if found != b {
		return s.scanErrorf("expected `%v`, found `%v`", b, found)
	}
	_, err = s.ConsumeByte()
	return err
}

func (s *Scanner) scanError(msg string) error {
	e := &ScanError{
		pos: s.pos,
		msg: msg,
	}
	return e
}

func (s *Scanner) scanErrorf(format string, args ...interface{}) error {
	return s.scanError(fmt.Sprintf(format, args...))
}
