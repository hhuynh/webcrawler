package crawler

import (
	"net/url"
	"testing"
)

func TestConstructFilePath(t *testing.T) {

	var filePathTable = []struct {
		url      string // input
		expected string // expected result
	}{
		{"http://www.amazon.com", "/tmp/www.amazon.com/index.html"},
		{"http://www.amazon.com/", "/tmp/www.amazon.com/index.html"},
		{"http://www.amazon.com/index.html", "/tmp/www.amazon.com/index.html"},
		{"http://www.amazon.com/us/", "/tmp/www.amazon.com/us/index.html"},
		{"http://www.amazon.com/us", "/tmp/www.amazon.com/us/index.html"},
		{"http://www.amazon.com/us/west/hello.asp", "/tmp/www.amazon.com/us/west/hello.asp"},
		{"http://www.amazon.com/us/west.coast/hello.asp", "/tmp/www.amazon.com/us/west.coast/hello.asp"},
	}

	for _, testData := range filePathTable {
		testUrl, err := url.Parse(testData.url)
		if err != nil {
			t.Errorf("Failed to parse: %s with %v", testData.url, err)
		}
		actual := constructFilePath(testUrl, "/tmp")
		if actual != testData.expected {
			t.Errorf("Expected: %s, but got: %s", testData.expected, actual)
		}
	}
}
