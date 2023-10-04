package api

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegexp(t *testing.T) {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	testCases := []struct {
		date  string
		valid bool
	}{
		{
			date:  "234234-2342-234",
			valid: false,
		},
		{
			date:  "2023-10-04",
			valid: true,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.valid, re.MatchString(tc.date))
	}
}
