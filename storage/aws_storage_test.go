package storage_test

import (
	"testing"
	"time"

	"github.com/leometzger/mmonitoring-runner/storage"
	"github.com/stretchr/testify/assert"
)

func TestS3KeyPath(t *testing.T) {
	cases := []struct {
		url      string
		datetime time.Time
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
		assert.Equal(t, test.expected, storage.GetS3KeyFromUrl(test.url))
	}
}
