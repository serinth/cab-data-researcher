package cabService

//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"time"
//)
//
//func TestParseDateForDayOnlyQuery(t *testing.T) {
//
//	var dateTests = []struct {
//		in string
//		expected *betweenDateTuple
//	}{
//		{"2017-01-20", &betweenDateTuple{start: "2017-01-20", end: "2017-01-21"}},
//		{"2017-01-31", &betweenDateTuple{start: "2017-01-31", end: "2017-02-01"}},
//	}
//
//	for _, tt := range dateTests {
//		t.Run(tt.in, func(t *testing.T) {
//			parsedTime, _ := time.Parse(ISO8601Layout, tt.in)
//			actual := parseDateForDayOnlyQuery(parsedTime)
//
//			assert.Equal(t, tt.expected.start, actual.start)
//			assert.Equal(t, tt.expected.end, actual.end)
//		})
//	}
//}