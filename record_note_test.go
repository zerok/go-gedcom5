package gedcom5

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestNoteRecord(t *testing.T) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	data := bytes.NewBufferString(`0 HEAD
0 @i1@ NOTE line1
1 CONT line2
0 TRLR
`)
	var f File
	err := NewDecoder(data).WithContext(logger.WithContext(context.Background())).Decode(&f)
	require.NoError(t, err)
	require.NotNil(t, f)
	require.Len(t, f.Records, 1)
	require.Equal(t, "line1\nline2", f.Records[0].(*NoteRecord).Value())
}
