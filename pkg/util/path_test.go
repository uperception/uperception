package utils_test

import (
	"regexp"
	"testing"

	utils "github.com/leometzger/mmonitoring/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestKeyPathing(t *testing.T) {
	cases := []struct {
		url      string
		expected string
	}{
		{
			url:      "https://google.com",
			expected: "reports\\/google.com\\/20[0-9]{2}\\/(0|1)[0-9]\\/[0-3][0-9]\\/[0-2][0-9]\\/[a-z0-9]{40}\\.json$",
		},
		{
			url:      "https://google.com/search?q=breaking+bad",
			expected: "reports\\/google.com\\/20[0-9]{2}\\/(0|1)[0-9]\\/[0-3][0-9]\\/[0-2][0-9]\\/[a-z0-9]{40}\\.json$",
		},
		{
			url:      "https://facebook.com",
			expected: "reports\\/facebook.com\\/20[0-9]{2}\\/(0|1)[0-9]\\/[0-3][0-9]\\/[0-2][0-9]\\/[a-z0-9]{40}\\.json$",
		},
	}

	for _, test := range cases {
		assert.Regexp(t, regexp.MustCompile(test.expected), utils.GetPathFromUrl(test.url))
	}
}
