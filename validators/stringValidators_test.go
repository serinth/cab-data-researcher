package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsOnlyAlphanumeric(t *testing.T) {
	var medallionIdTests = []struct {
		name         string
		medallionIds []string
		expected     bool
	}{
		{"single good case", []string{"801C69A08B51470871A8110F8B0505EE"}, true},
		{"multiple good case", []string{"5455D5FF2BD94D10B304A15D4B7F2735", "abcd123"}, true},
		{"good and bad", []string{"5455D5FF2BD94D10B304A15D4B7F2735", "$houldFa!l"}, false},
		{"-, potential sql injection", []string{"-",}, false},
		{"(, potential sql injection", []string{"("}, false},
		{"), potential sql injection", []string{")"}, false},
		{";, potential sql injection", []string{";"}, false},
		{"good and potential sql injection", []string{"5455D5FF2BD94D10B304A15D4B7F2735", "--"}, false},
	}

	for _, tt := range medallionIdTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ContainsOnlyAlphanumeric(tt.medallionIds))
		})
	}
}
