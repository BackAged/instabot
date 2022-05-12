# Instagram Messaging API GO SDK

[![Build Status](https://github.com/BackAged/instabot/actions/workflows/go.yaml/badge.svg?branch=master)](https://github.com/BackAged/instabot/actions/workflows/go.yaml)
[![codecov](https://codecov.io/gh/BackAged/instabot/branch/master/graph/badge.svg)](https://codecov.io/gh/BackAged/instabot)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/BackAged/instabot)
[![Go Report Card](https://goreportcard.com/badge/github.com/BackAged/instabot)](https://goreportcard.com/report/github.com/BackAged/instabot)
[![CodeQL](https://github.com/BackAged/instabot/actions/workflows/codeql-analysis.yaml/badge.svg?branch=master)](https://github.com/BackAged/instabot/actions/workflows/codeql-analysis.yaml)

[![Run on Repl.it](https://repl.it/badge/github/BackAged/instabot)](https://repl.it/github/BackAged/instabot)

## Introduction
Instabot, Instagram Messaging API GO SDK makes it easy to work with instagram messaging API.
It uses Instagram messaging API latest version - `v11.0`

## Requirements

Instabot requires Go 1.13 or later.

## Installation ##

```sh
$ go get -u github.com/mohammadVatandoost/instabot
```

## Instabot configuration

```go
import (
	"context"
	"fmt"
	"log"

	"github.com/mohammadVatandoost/instabot"
)

func main() {
    // instantiating instabot.
	bot, err := instabot.New("your_instagram_business_account_page_access_token")
    ...

    // instantiating with http.Client
    bot, err := instabot.New(
        "your_instagram_business_account_page_access_token",
        instabot.WithHTTPClient(yourHttpClient)
    )
    ...

    // instantiating with mock api server
    bot, err := instabot.New(
        "your_instagram_business_account_page_access_token",
        instabot.APIEndpointBase("http://your_mock_api_server.com")
    )
    ...

}
```

## Example

```go
import (
	"context"
	"fmt"
	"log"

	"github.com/mohammadVatandoost/instabot"
)

func main() {
    // See examples directory for more example.
    
    // instantiating instabot.
	bot, err := instabot.New("your_instagram_business_account_page_access_token")
    ...

    
    // Send text message.
	_, err = bot.SendMessage(
		context.Background(),
		"instagram_user_id_you_want_to_send_message_to",
		instabot.NewTextMessage("hello"),
	)
    ...

    // Set icebreakers
    _, err = bot.SetIceBreakers(
		context.Background(),
		[]*instabot.IceBreaker{
			instabot.NewIceBreaker("frequently asked question 1", "user payload"),
			instabot.NewIceBreaker("frequently asked question 2", "user payload"),
			instabot.NewIceBreaker("frequently asked question 3", "user payload"),
			instabot.NewIceBreaker("frequently asked question 4", "user payload"),
		},
	)
    ...

    // Get user profile.
	profile, err := bot.GetUserProfile(
		context.Background(),
		"instagram_user_id_you_want_to_get_profile",
	)
    ...

	// work with webhook events.
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
	...
}
```