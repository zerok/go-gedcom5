package gedcom5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadXRefID(t *testing.T) {
	buf := bytes.NewBufferString("@a1c@")
	s := NewScanner(buf)
	id, err := s.ReadXRefID()
	require.NoError(t, err)
	require.Equal(t, "@a1c@", id)
}

func TestReadNumber(t *testing.T) {
	buf := bytes.NewBufferString("123\n456")
	s := NewScanner(buf)
	num, err := s.ReadNumber()
	require.NoError(t, err)
	require.Equal(t, 123, num)
	require.Equal(t, 3, s.Pos())
	s.ReadByte(0xA)
	other, _ := s.ReadNumber()
	require.NoError(t, s.Error())
	require.Equal(t, 456, other)
}

func TestReadAlphaNum(t *testing.T) {
	buf := bytes.NewBufferString("1a3\n456")
	s := NewScanner(buf)
	res, err := s.ReadAlphaNum()
	require.NoError(t, err)
	require.Equal(t, "1a3", res)
}
