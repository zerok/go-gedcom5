package gedcom5

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLineString(t *testing.T) {
	l := Line{Level: 1, Tag: "TAG", XRefID: "@a@", Value: "abc"}
	require.Equal(t, "1 @a@ TAG abc", l.String())
}

func TestMultilineValueDecoding(t *testing.T) {
	tests := []struct {
		name      string
		lines     []Line
		result    string
		expectErr bool
	}{
		{
			name: "Using CONT",
			lines: []Line{
				Line{Level: 1, Tag: "ADDR", Value: "Line 1"},
				Line{Level: 2, Tag: "CONT", Value: "Line 2"},
			},
			result:    "Line 1\nLine 2",
			expectErr: false,
		},
		{
			name: "Using CONC",
			lines: []Line{
				Line{Level: 1, Tag: "ADDR", Value: "Line 1"},
				Line{Level: 2, Tag: "CONC", Value: "Line 2"},
			},
			result:    "Line 1Line 2",
			expectErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := struct {
				Addr string `gedcom5:"ADDR"`
			}{}
			ld := NewLineDecoder(&val, 0)
			err := ld.Decode(context.Background(), test.lines)
			if test.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, test.result, val.Addr)
		})
	}
}

func TestMultilineValueOfValuableDecoding(t *testing.T) {
	tests := []struct {
		name      string
		lines     []Line
		result    string
		expectErr bool
	}{
		{
			name: "Using CONT",
			lines: []Line{
				Line{Level: 1, Tag: "ADDR", Value: "Line 1"},
				Line{Level: 2, Tag: "CONT", Value: "Line 2"},
			},
			result:    "Line 1\nLine 2",
			expectErr: false,
		},
	}

	for _, test := range tests {
		addr := struct {
			Address Address `gedcom5:"ADDR"`
		}{}
		ld := NewLineDecoder(&addr, 0)
		err := ld.Decode(context.Background(), test.lines)
		if test.expectErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, test.result, addr.Address.Value())
	}

}
