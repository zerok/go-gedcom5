package gedcom5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecoder(t *testing.T) {
	data := bytes.NewBufferString("0 HEAD\n0 @a@ INDI\n1 NAME Jane /Doe/\n0 TRLR\n")
	f := &File{}
	err := NewDecoder(data).Decode(f)
	require.NoError(t, err)
	require.NotNil(t, f)
}
