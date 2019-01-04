package gedcom5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFamRecord(t *testing.T) {
	data := bytes.NewBufferString(`0 HEAD
0 @i1@ INDI
1 NAME John /Doe/
0 @i2@ INDI
1 NAME Jane /Doe/
0 @i3@ INDI
1 NAME Jack /Doe/
0 @i4@ INDI
1 NAME Joanna /Doe/
0 @fam@ FAM
1 HUSB @i1@
1 WIFE @i2@
1 CHIL @i3@
1 CHIL @i4@
0 TRLR
`)
	var f File
	err := NewDecoder(data).Decode(&f)
	require.NoError(t, err)
	require.NotNil(t, f)
}
