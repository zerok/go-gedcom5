package gedcom5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSourceRecord(t *testing.T) {
	data := bytes.NewBufferString(`0 HEAD
0 @i1@ SOUR
1 TITL title
0 TRLR
`)
	var f File
	err := NewDecoder(data).Decode(&f)
	require.NoError(t, err)
	require.NotNil(t, f)
	require.Len(t, f.Records, 1)
	sour := f.Records[0].(*SourceRecord)
	require.Equal(t, "title", sour.Title)
}
