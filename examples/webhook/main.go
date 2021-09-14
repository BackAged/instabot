package main

import (
	"encoding/json"
	"log"

	"github.com/BackAged/instabot"
)

func main() {
	// https://developers.facebook.com/docs/messenger-platform/instagram/features/webhook

	payload := []byte(`{
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
		],
	}`)

	webhookEvent := new(instabot.WebhookEvent)

	if err := json.Unmarshal(payload, webhookEvent); err != nil {
		log.Fatal("failed to unmarshal event")
	}

	for _, entry := range webhookEvent.Entries {
		for _, event := range entry.Messaging {
			switch event.Type {
			case instabot.WebhookEventTypeEcho:
			case instabot.WebhookEventTypeDeleted:
				log.Println(event.GetMessageDeleteEvent())
			case instabot.WebhookEventTypeUnsupported:
			case instabot.WebhookEventTypeMessageSeen:
				log.Println(event.GetMessageSeenEvent())
			case instabot.WebhookEventTypeMessageReply:
				log.Println(event.GetMessageReplyEvent())
			case instabot.WebhookEventTypeShare:
				log.Println(event.GetMessageShareEvent())
			case instabot.WebhookEventTypeReaction:
				log.Println(event.GetMessageReactionEvent())
			case instabot.WebhookEventTypeTextMessage:
				log.Println(event.GetTextMessageEvent())
			case instabot.WebhookEventTypeImageMessage:
				log.Println(event.GetMediaMessageEvent())
			case instabot.WebhookEventTypeAudioMessage:
				log.Println(event.GetMediaMessageEvent())
			case instabot.WebhookEventTypeVideoMessage:
				log.Println(event.GetMediaMessageEvent())
			case instabot.WebhookEventTypeFileMessage:
				log.Println(event.GetMediaMessageEvent())
			case instabot.WebhookEventTypeQuickReply:
				log.Println(event.GetQuickReplyEvent())
			case instabot.WebhookEventTypePostBack:
				log.Println(event.GetPostBackEvent())
			case instabot.WebhookEventTypeStoryMention:
				log.Println(event.GetStoryMentionEvent())
			case instabot.WebhookEventTypeStoryReply:
				log.Println(event.GetStoryReplyEvent())
			default:
				log.Println("unexpected event")
			}
		}
	}
}
