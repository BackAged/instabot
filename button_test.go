package instabot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButtonType(t *testing.T) {
	testCases := []struct {
		name string
		args Button
		want ButtonType
	}{
		{
			name: "web url button",
			args: NewURLButton("title", "link"),
			want: ButtonTypeURL,
		},
		{
			name: "post back button",
			args: NewPostBackButton("title", "payload"),
			want: ButtonTypePostBack,
		},
		{
			name: "call button",
			args: NewCallButton("title", "phone number"),
			want: ButtonTypeCall,
		},
		{
			name: "log in button",
			args: NewLogInButton("<YOUR_LOGIN_URL>"),
			want: ButtonTypeLogin,
		},
		{
			name: "log out button",
			args: NewLogOutButton(),
			want: ButtonTypeLogOut,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.args.Type())
		})
	}
}

func TestButtonJSON(t *testing.T) {
	testCases := []struct {
		name string
		args Button
		want string
	}{
		{
			name: "web url button",
			args: NewURLButton("title", "link"),
			want: `{
				"type": "web_url",
				"url": "link",
				"title": "title"
			}`,
		},
		{
			name: "post back button",
			args: NewPostBackButton("title", "payload"),
			want: `{
				"type": "postback",
				"payload": "payload",
				"title": "title"
			}`,
		},
		{
			name: "call button",
			args: NewCallButton("title", "phone number"),
			want: `{
				"type": "phone_number",
				"payload": "phone number",
				"title": "title"
			}`,
		},
		{
			name: "log in button",
			args: NewLogInButton("<YOUR_LOGIN_URL>"),
			want: `{
				"type": "account_link",
				"url": "<YOUR_LOGIN_URL>"
			}`,
		},
		{
			name: "log out button",
			args: NewLogOutButton(),
			want: `{
				"type": "account_unlink"
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
