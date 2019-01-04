package gedcom5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepositoryRecord(t *testing.T) {
	data := bytes.NewBufferString(`0 HEAD
0 @i1@ REPO
0 TRLR
`)
	var f File
	err := NewDecoder(data).Decode(&f)
	require.NoError(t, err)
	require.NotNil(t, f)
}
