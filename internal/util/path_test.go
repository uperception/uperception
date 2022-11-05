package utils_test

import (
	"testing"

	"github.com/leometzger/mmonitoring-runner/utils"
	"github.com/stretchr/testify/assert"
)

func TestKeyPathing(t *testing.T) {
	cases := []struct {
		url      string
		expected string
	}{
		{
			url:      "https://google.com",
			expected: "reports/google.com/2022/10/10/11/30/id.json",
		},
		{
			url:      "https://google.com/search?q=breaking+bad",
			expected: "reports/google.com/2022/02/10/23/30/id.json",
		},
		{
			url:      "https://facebook.com",
			expected: "reports/facebook.com/2022/10/10/11/30/id.json",
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.expected, utils.GetPathFromUrl(test.url))
	}
}
