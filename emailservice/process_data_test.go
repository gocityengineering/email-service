package emailservice

import (
	"testing"
)

func TestProcessData(t *testing.T) {
	var tests = []struct {
		description string
		bytes       []byte
		expected    bool
	}{
		{"MustRejectHtml", []byte(`<html><head/><body></body></html>`), false},
		{"MustRejectYaml", []byte(`foo: bar`), false},
		{"MustAcceptValidJsonInput", []byte(`{"recipients":["someone@gocity.com","someoneelse@gocity.com"],"subject":"somesobject","markdownBody":"Some *markdown*"}`), true},
		{"MustRejectInvalidJsonInput", []byte(`{"recipients":"someone@gocity.com","subject":"somesobject","markdownBody":"Some *markdown*"}`), false},
		{"MustRejectMalformedJsonInput", []byte(`{"recipients":"someone@gocity.com"[,"subject":"somesobject","markdownBody":"Some *markdown*"}`), false},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := processData(test.bytes, true)
			if err != nil && test.expected || err == nil && test.expected == false {
				t.Errorf("Unexpected result %t", !test.expected)
			}
		})
	}
}
