package instabot

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetQuickReplyEvent(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      *QuickReplyEvent
		afterEach func(t *testing.T, event *WebhookEvent)
	}{
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
			want: &QuickReplyEvent{
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1502905976377, 0).UTC(),
				MID:       "<MID>",
				Text:      "<SOME_TEXT>",
				Data: &WebhookQuickReply{
					Payload: "<PAYLOAD>",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, WebhookEventTypeQuickReply, e.Entries[0].Messaging[0].Type)
			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].GetQuickReplyEvent())

			if tc.afterEach != nil {
				tc.afterEach(t, e)
			}
		})
	}
}

func TestGetPostbackEvent(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      *PostBackEvent
		afterEach func(t *testing.T, event *WebhookEvent)
	}{
		{
			name: "postback event",
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
			want: &PostBackEvent{
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1502905976377, 0).UTC(),
				Data: &Postback{
					MID:     "<MESSAGE_ID>",
					Title:   "<SELECTED_ICEBREAKER_QUESTION>",
					Payload: "<USER_DEFINED_PAYLOAD>",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, WebhookEventTypePostBack, e.Entries[0].Messaging[0].Type)
			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].GetPostBackEvent())

			if tc.afterEach != nil {
				tc.afterEach(t, e)
			}
		})
	}
}

func TestGetStoryReplyEvent(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      *StoryReplyEvent
		afterEach func(t *testing.T, event *WebhookEvent)
	}{
		{
			name: "story reply event",
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
							   "text":"<MESSAGE_CONTENT>",
							   "reply_to":{
								  "story":{
									 "url":"<CDN_URL>",
									 "id":"story_id"
								  }
							   }
							}
						 }
					  ]
				   }
				]
			}`,
			want: &StoryReplyEvent{
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1569262485349, 0).UTC(),
				MID:       "<MESSAGE_ID>",
				Text:      "<MESSAGE_CONTENT>",
				Story: &ReplyToStory{
					ID:  "story_id",
					URL: "<CDN_URL>",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, WebhookEventTypeStoryReply, e.Entries[0].Messaging[0].Type)
			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].GetStoryReplyEvent())

			if tc.afterEach != nil {
				tc.afterEach(t, e)
			}
		})
	}
}

func TestGetStoryMentionEvent(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      *StoryMentionEvent
		afterEach func(t *testing.T, event *WebhookEvent)
	}{
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
			want: &StoryMentionEvent{
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1569262485349, 0).UTC(),
				MID:       "<MESSAGE_ID>",
				Story: &ReplyToStory{
					URL: "<CDN_URL>",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, WebhookEventTypeStoryMention, e.Entries[0].Messaging[0].Type)
			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].GetStoryMentionEvent())

			if tc.afterEach != nil {
				tc.afterEach(t, e)
			}
		})
	}
}

func TestGetTextMessageEvent(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      *TextMessageEvent
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
			want: &TextMessageEvent{
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1569262485349, 0).UTC(),
				MID:       "<MESSAGE_ID>",
				Text:      "<MESSAGE_CONTENT>",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, WebhookEventTypeTextMessage, e.Entries[0].Messaging[0].Type)
			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].GetTextMessageEvent())

			if tc.afterEach != nil {
				tc.afterEach(t, e)
			}
		})
	}
}

func TestGetMediaMessageEvent(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      *MediaMessageEvent
		afterEach func(t *testing.T, event *MediaMessageEvent)
	}{
		{
			name: "image media message event",
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
			want: &MediaMessageEvent{
				Type: WebhookEventTypeImageMessage,
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1569262485349, 0).UTC(),
				MID:       "<MESSAGE_ID>",
				Media: &AttachmentPayload{
					URL: "<CDN_LINK>",
				},
			},
		},
		{
			name: "audio media message event",
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
									 "type":"audio",
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
			want: &MediaMessageEvent{
				Type: WebhookEventTypeAudioMessage,
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1569262485349, 0).UTC(),
				MID:       "<MESSAGE_ID>",
				Media: &AttachmentPayload{
					URL: "<CDN_LINK>",
				},
			},
		},
		{
			name: "video media message event",
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
									 "type":"video",
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
			want: &MediaMessageEvent{
				Type: WebhookEventTypeVideoMessage,
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1569262485349, 0).UTC(),
				MID:       "<MESSAGE_ID>",
				Media: &AttachmentPayload{
					URL: "<CDN_LINK>",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].GetMediaMessageEvent())

			if tc.afterEach != nil {
				tc.afterEach(t, e.Entries[0].Messaging[0].GetMediaMessageEvent())
			}
		})
	}
}

func TestGetMessageReplyEvent(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      *MessageReplyEvent
		afterEach func(t *testing.T, event *WebhookEvent)
	}{
		{
			name: "message reply event",
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
							   "text":"<MESSAGE_CONTENT>",
							   "reply_to":{
								  "mid":"<MESSAGE_ID>"
							   }
							}
						 }
					  ]
				   }
				]
			}`,
			want: &MessageReplyEvent{
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp:  time.Unix(1569262485349, 0).UTC(),
				MID:        "<MESSAGE_ID>",
				Text:       "<MESSAGE_CONTENT>",
				ReplyToMID: "<MESSAGE_ID>",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, WebhookEventTypeMessageReply, e.Entries[0].Messaging[0].Type)
			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].GetMessageReplyEvent())

			if tc.afterEach != nil {
				tc.afterEach(t, e)
			}
		})
	}
}

func TestGetMessageReactionEvent(t *testing.T) {
	testCases := []struct {
		name      string
		args      string
		want      *MessageReactionEvent
		afterEach func(t *testing.T, event *WebhookEvent)
	}{
		{
			name: "message reaction event",
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
						"reaction" : {
						  "mid" : "<MID>",
						  "action": "react",
						  "reaction": "love",
						  "emoji": "something"
						} 
					  }
					]
				  }
				]
			}`,
			want: &MessageReactionEvent{
				Sender: &Sender{
					ID: "<IGSID>",
				},
				Recipient: &Recipient{
					ID: "<IGID>",
				},
				Timestamp: time.Unix(1569262485349, 0).UTC(),
				Reaction: &Reaction{
					MID:      "<MID>",
					Action:   "react",
					Reaction: "love",
					Emoji:    "something",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := new(WebhookEvent)
			err := json.Unmarshal([]byte(tc.args), e)
			assert.NoError(t, err)

			assert.Equal(t, WebhookEventTypeReaction, e.Entries[0].Messaging[0].Type)
			assert.Equal(t, tc.want, e.Entries[0].Messaging[0].GetMessageReactionEvent())

			if tc.afterEach != nil {
				tc.afterEach(t, e)
			}
		})
	}
}
