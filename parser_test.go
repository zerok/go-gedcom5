package gedcom5

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	fp := NewFileParser(NewScanner(bytes.NewBufferString("0 HEAD\n0 @a@ INDI\n0 TRLR\n")))
	f, err := fp.ParseFile(context.Background())
	require.NoError(t, err)
	require.NotNil(t, f)
	require.Len(t, f.Records, 1)
	require.Equal(t, "@a@", f.Records[0].(Identifyable).ID())
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		input     string
		expectErr bool
		line      *Line
		name      string
	}{
		{
			name:      "Level and tag with LF",
			input:     "123 TAG\n",
			expectErr: false,
			line:      &Line{Level: 123, Tag: "TAG"},
		},
		{
			name:      "Level and tag with CRLF",
			input:     "123 TAG\r\n",
			expectErr: false,
			line:      &Line{Level: 123, Tag: "TAG"},
		},
		{
			name:      "Level and tag with CR",
			input:     "123 TAG\r",
			expectErr: false,
			line:      &Line{Level: 123, Tag: "TAG"},
		},
		{
			name:      "Level and tag with LFCR",
			input:     "123 TAG\n\r",
			expectErr: false,
			line:      &Line{Level: 123, Tag: "TAG"},
		},
		{
			name:      "Level and tag with LFCR and xref_id",
			input:     "123 @a@ TAG\n\r",
			expectErr: false,
			line:      &Line{Level: 123, Tag: "TAG", XRefID: "@a@"},
		},
		{
			name:      "Level and tag with LFCR and xref_id and value",
			input:     "123 @a@ TAG value\n\r",
			expectErr: false,
			line:      &Line{Level: 123, Tag: "TAG", XRefID: "@a@", Value: "value"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lines, err := OwnParseLine(bytes.NewBufferString(test.input))
			if test.expectErr && err == nil {
				t.Fatalf("error expected for %s", test.name)
			}
			if err != nil && !test.expectErr {
				t.Fatalf("unexpected error for %s: %s", test.name, err.Error())
			}
			if err == nil {
				if test.line != nil {
					line := lines[0]
					require.Equal(t, test.line.Level, line.Level)
					require.Equal(t, test.line.Tag, line.Tag)
				}
			}
		})
	}
}
