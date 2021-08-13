package instabot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIceBreakerJSON(t *testing.T) {
	testCases := []struct {
		name string
		args *IceBreaker
		want string
	}{
		{
			name: "ice breaker",
			args: NewIceBreaker("test", "test"),
			want: `{
				"question":"test",
				"payload":"test"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			j, err := json.Marshal(tc.args)
			assert.NoError(t, err)

			assert.JSONEq(t, tc.want, string(j))
		})
	}
}
