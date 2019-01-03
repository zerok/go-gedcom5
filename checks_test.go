package gedcom5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type byteCheckerTest struct {
	input    byte
	expected bool
}

func TestIsEither(t *testing.T) {
	m := IsEither(IsAt, IsDigit)
	tests := []byteCheckerTest{
		{
			input:    '@',
			expected: true,
		},
		{
			input:    '1',
			expected: true,
		},
		{
			input:    'a',
			expected: false,
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expected, m(test.input))
	}
}

func TestIsAt(t *testing.T) {
	tests := []byteCheckerTest{
		{
			input:    '@',
			expected: true,
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expected, IsAt(test.input))
	}
}

func TestIsDigit(t *testing.T) {
	tests := []byteCheckerTest{
		{
			input:    '8',
			expected: true,
		},
		{
			input:    '@',
			expected: false,
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expected, IsDigit(test.input))
	}
}
