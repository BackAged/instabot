package instabot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickReplyJSON(t *testing.T) {
	testCases := []struct {
		name      string
		args      *QuickReply
		want      string
		afterEach func(t *testing.T)
	}{
		{
			name: "quick reply",
			args: NewTextQuickReply("<TITLE_1>", "<POSTBACK_PAYLOAD_1>"),
			want: `{
				"content_type":"text",
				"title":"<TITLE_1>",
				"payload":"<POSTBACK_PAYLOAD_1>"
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
