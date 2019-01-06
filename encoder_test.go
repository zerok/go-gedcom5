package gedcom5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncoder(t *testing.T) {
	f := File{}
	f.Records = []Record{
		&IndividualRecord{
			PersonalNames: []PersonalName{
				{Name: "John Doe 1"},
				{Name: "John Doe 2"},
			},
			Sex: "M",
			Residence: Event{
				Address: Address{
					Val: "Addressline 1",
				},
			},
		},
	}
	out := bytes.Buffer{}
	err := NewEncoder(&out).Encode(&f)
	require.NoError(t, err)
	require.Equal(t, `0 HEAD
0 INDI
1 NAME John Doe 1
1 NAME John Doe 2
1 SEX M
1 RESI
2 ADDR Addressline 1
0 TRLR
`, out.String())
}
