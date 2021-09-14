package instabot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookEventType(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      WebhookEventType
		afterEach func(t *testing.T, event *WebhookEvent)
	}{
		{
			name: "text message event",
			args: `{
				"object": "instagram",
				"entry": [
				  {
					"id": "<IGID>",
					"time": 1569262486134,
					"messaging": [
					  {
						"sender": {
						  "id": "<IGSID>"
						},
						"recipient": {
						  "id": "<IGID>"
						},
						"timestamp": 1569262485349,
						"message": {
						  "mid": "<MESSAGE_ID>",
						  "text": "<MESSAGE_CONTENT>"
						}
					  }
					]
				  }
				]
			}`,
			want: WebhookEventTypeTextMessage,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				assert.Nil(t, messaging.PostBack)
				message := messaging.Message
				assert.Equal(t, "<MESSAGE_ID>", message.MID)
				assert.Equal(t, "<MESSAGE_CONTENT>", message.Text)
				assert.Nil(t, message.ReplyTo)
				assert.Nil(t, message.QuickReply)
			},
		},
		{
			name: "quick reply event",
			args: `{
				"object": "instagram",
				"entry": [
				  {
					"id": "<IGSID>",
					"time": 1502905976963,
					"messaging": [
					  {
						"sender": {
						  "id": "<IGSID>"
						},
						"recipient": {
						  "id": "<IGID>"
						},
						"timestamp": 1502905976377,
						"message": {
						  "quick_reply": {
							"payload": "<PAYLOAD>"
						  },
						  "mid": "<MID>",
						  "text": "<SOME_TEXT>"
						}
					  }
					]
				  }
				]
			}`,
			want: WebhookEventTypeQuickReply,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				assert.Nil(t, messaging.PostBack)
				message := messaging.Message
				assert.Equal(t, "<MID>", message.MID)
				assert.Equal(t, "<SOME_TEXT>", message.Text)
				assert.Nil(t, message.ReplyTo)
				assert.Equal(t, message.QuickReply.Payload, "<PAYLOAD>")
			},
		},
		{
			name: "echo event",
			args: `{
				"object": "instagram",
				"entry": [
				  {
					"id": "<IGSID>",
					"time": 1502905976963,
					"messaging": [
					  {
						"sender": {
						  "id": "<IGSID>"
						},
						"recipient": {
						  "id": "<IGID>"
						},
						"timestamp": 1502905976377,
						"message": {
						  "mid": "<MID>",
						  "text": "<SOME_TEXT>",
						  "is_echo":true
						}
					  }
					]
				  }
				]
			}`,
			want: WebhookEventTypeEcho,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				assert.Nil(t, messaging.PostBack)
				message := messaging.Message
				assert.Equal(t, "<MID>", message.MID)
				assert.Equal(t, "<SOME_TEXT>", message.Text)
				assert.Nil(t, message.ReplyTo)
			},
		},
		{
			name: "delete event",
			args: `{
				"object": "instagram",
				"entry": [
				  {
					"id": "<IGSID>",
					"time": 1502905976963,
					"messaging": [
					  {
						"sender": {
						  "id": "<IGSID>"
						},
						"recipient": {
						  "id": "<IGID>"
						},
						"timestamp": 1502905976377,
						"message": {
						  "mid": "<MID>",
						  "text": "<SOME_TEXT>",
						  "is_deleted":true
						}
					  }
					]
				  }
				]
			}`,
			want: WebhookEventTypeDeleted,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				assert.Nil(t, messaging.PostBack)
				message := messaging.Message
				assert.Equal(t, "<MID>", message.MID)
				assert.Equal(t, "<SOME_TEXT>", message.Text)
				assert.Nil(t, message.ReplyTo)
			},
		},
		{
			name: "message seen event",
			args: `{
				"object":"instagram",
				"entry":[
				   {
					  "id":"<IGID>",
					  "time":1569262486134,
					  "message":[
						 {
							"sender":{
							   "id":"<IGSID>"
							},
							"recipient":{
							   "id":"<IGID>"
							},
							"timestamp":1569262485349,
							"read":{
							   "mid":"<LAST_MESSAGE_ID_READ>"
							}
						 }
					  ]
				   }
				]
			}`,
			want: WebhookEventTypeMessageSeen,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				assert.Nil(t, messaging.PostBack)
				assert.Nil(t, messaging.Message)
				assert.Equal(t, "<LAST_MESSAGE_ID_READ>", messaging.Read.MID)
			},
		},
		{
			name: "postback (icebreaker) event",
			args: `{
				"object": "instagram",
				"entry": [
				  {
					"id": "<IGSID>",
					"time": 1502905976963,
					"messaging": [
					  {
						"sender": {
						  "id": "<IGSID>"
						},
						"recipient": {
						  "id": "<IGID>"
						},
						"timestamp": 1502905976377,
						"postback": {
						  "mid":"<MESSAGE_ID>",
						  "title": "<SELECTED_ICEBREAKER_QUESTION>",
						  "payload": "<USER_DEFINED_PAYLOAD>"
						}
					  }
					]
				  }
				]
			}`,
			want: WebhookEventTypePostBack,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				assert.Nil(t, messaging.Message)
				assert.Nil(t, messaging.Reaction)
				assert.Equal(t, "<MESSAGE_ID>", messaging.PostBack.MID)
				assert.Equal(t, "<USER_DEFINED_PAYLOAD>", messaging.PostBack.Payload)
			},
		},
		{
			name: "postback (generic template) event",
			args: `{
				"object": "instagram",
				"entry": [
				  {
					"id": "<IGSID>",
					"time": 1502905976963,
					"messaging": [
					  {
						"sender": {
						  "id": "<IGSID>"
						},
						"recipient": {
						  "id": "<IGID>"
						},
						"timestamp": 1502905976377,
						"postback": {
						  "mid":"<MESSAGE_ID>",
						  "title": "<TITLE_FOR_THE_CTA>",
						  "payload": "<USER_DEFINED_PAYLOAD>"
						}
					  }
					]
				  }
				]
			  }
			`,
			want: WebhookEventTypePostBack,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				assert.Nil(t, messaging.Message)
				assert.Nil(t, messaging.Reaction)
				assert.Equal(t, "<MESSAGE_ID>", messaging.PostBack.MID)
				assert.Equal(t, "<USER_DEFINED_PAYLOAD>", messaging.PostBack.Payload)
			},
		},
		{
			name: "media share event",
			args: `{
				"object":"instagram",
				"entry":[
				   {
					  "id":"<IGID>",
					  "time":1569262486134,
					  "messaging":[
						 {
							"sender":{
							   "id":"<IGSID>"
							},
							"recipient":{
							   "id":"<IGID>"
							},
							"timestamp":1569262485349,
							"message":{
							   "mid":"<MESSAGE_ID>",
							   "attachments":[
								  {
									 "type":"image",
									 "payload":{
										"url":"<CDN_LINK>"
									 }
								  }
							   ]
							}
						 }
					  ]
				   }
				]
			}`,
			want: WebhookEventTypeImageMessage,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				message := messaging.Message
				assert.Nil(t, messaging.PostBack)
				assert.Nil(t, messaging.Reaction)
				assert.Equal(t, "<MESSAGE_ID>", messaging.Message.MID)
				assert.Equal(t, "<CDN_LINK>", message.Attachments[0].Payload.URL)
			},
		},
		{
			name: "story mention event",
			args: `{
				"object":"instagram",
				"entry":[
				   {
					  "id":"<IGID>",
					  "time":1569262486134,
					  "messaging":[
						 {
							"sender":{
							   "id":"<IGSID>"
							},
							"recipient":{
							   "id":"<IGID>"
							},
							"timestamp":1569262485349,
							"message":{
							   "mid":"<MESSAGE_ID>",
							   "attachments":[
								  {
									 "type":"story_mention",
									 "payload":{
										"url":"<CDN_URL>"
									 }
								  }
							   ]
							}
						 }
					  ]
				   }
				]
			}`,
			want: WebhookEventTypeStoryMention,
			afterEach: func(t *testing.T, event *WebhookEvent) {
				messaging := event.Entries[0].Messaging[0]
				message := messaging.Message
				assert.Nil(t, messaging.PostBack)
				assert.Nil(t, messaging.Reaction)
				assert.Equal(t, "<MESSAGE_ID>", messaging.Message.MID)
				assert.Equal(t, "<CDN_URL>", message.Attachments[0].Payload.URL)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].Type)

			if tc.afterEach != nil {
				tc.afterEach(t, e)
			}
		})
	}
}
