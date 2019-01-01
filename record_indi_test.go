package gedcom5

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndividualRecord(t *testing.T) {
	tests := []struct {
		name      string
		lines     []Line
		expectErr bool
		result    *IndividualRecord
	}{
		{
			name: "with one name",
			lines: []Line{
				Line{Level: 1, Tag: "NAME"},
				Line{Level: 2, Tag: "SURN", Value: "surname"},
				Line{Level: 1, Tag: "BIRT"},
				Line{Level: 2, Tag: "DATE", Value: "1. 2. 1903"},
			},
			expectErr: false,
			result: &IndividualRecord{
				lvl: 0,
				PersonalNames: []PersonalName{
					{
						lvl: 1,
						lines: []Line{
							Line{Level: 2, Tag: "SURN", Value: "surname"},
						},
						Surname: "surname",
					},
				},
				Birth: Birth{
					lvl:  1,
					Date: "1. 2. 1903",
					lines: []Line{
						Line{Level: 2, Tag: "DATE", Value: "1. 2. 1903"},
					},
				},
			},
		},
		{
			name: "with two names",
			lines: []Line{
				Line{Level: 1, Tag: "NAME"},
				Line{Level: 2, Tag: "SURN", Value: "surname"},
				Line{Level: 1, Tag: "NAME", Value: "name2"},
				Line{Level: 1, Tag: "BIRT"},
				Line{Level: 2, Tag: "DATE", Value: "1. 2. 1903"},
			},
			expectErr: false,
			result: &IndividualRecord{
				lvl: 0,
				PersonalNames: []PersonalName{
					{
						lvl: 1,
						lines: []Line{
							Line{Level: 2, Tag: "SURN", Value: "surname"},
						},
						Surname: "surname",
					},
					{
						lvl:   1,
						lines: []Line{},
						Name:  "name2",
					},
				},
				Birth: Birth{
					lvl:  1,
					Date: "1. 2. 1903",
					lines: []Line{
						Line{Level: 2, Tag: "DATE", Value: "1. 2. 1903"},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewIndividualRecord()
			r.SetLines(test.lines)
			test.result.SetLines(test.lines)
			err := r.Decode(context.Background())
			if test.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			if err == nil {
				require.Equal(t, test.result, r)
			}
		})
	}
}

func TestPersonalName(t *testing.T) {
	tests := []struct {
		name      string
		lines     []Line
		expectErr bool
		result    *PersonalName
	}{
		{
			name: "all fields set",
			lines: []Line{
				{
					Level: 1,
					Tag:   "NPFX",
					Value: "prefix",
				},
				{
					Level: 1,
					Tag:   "NSFX",
					Value: "suffix",
				},
				{
					Level: 1,
					Tag:   "SURN",
					Value: "surname",
				},
				{
					Level: 1,
					Tag:   "GIVN",
					Value: "given",
				},
				{
					Level: 1,
					Tag:   "SPFX",
					Value: "surname-prefix",
				},
			},
			expectErr: false,
			result: &PersonalName{
				lvl:           0,
				Prefix:        "prefix",
				Suffix:        "suffix",
				Given:         "given",
				SurnamePrefix: "surname-prefix",
				Surname:       "surname",
			},
		},
	}

	for _, test := range tests {
		p := PersonalName{}
		p.SetLevel(0)
		p.SetLines(test.lines)
		test.result.SetLines(test.lines)
		err := p.Decode(context.Background())
		if test.expectErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		if err == nil {
			require.Equal(t, test.result, &p)
		}
	}
}
