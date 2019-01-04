package gedcom5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultimediaRecord(t *testing.T) {
	data := bytes.NewBufferString(`0 HEAD
0 @i1@ OBJE
0 TRLR
`)
	var f File
	err := NewDecoder(data).Decode(&f)
	require.NoError(t, err)
	require.NotNil(t, f)
}
