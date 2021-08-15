package main

import (
	"context"
	"log"

	"github.com/BackAged/instabot"
)

func main() {
	// https://developers.facebook.com/docs/messenger-platform/instagram/features/send-message#send-api

	// instantiating instabot.
	bot, err := instabot.New("your_instagram_business_account_page_access_token")
	if err != nil {
		log.Fatal(err)
	}

	// Sending text message.
	_, err = bot.SendMessage(
		context.Background(),
		"instagram_user_id_you_want_to_send_message_to",
		instabot.NewTextMessage("hello"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Sending image message.
	_, err = bot.SendMessage(
		context.Background(),
		"instagram_user_id_you_want_to_send_message_to",
		instabot.NewImageMessage("image_url"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Sending sticker message.
	_, err = bot.SendMessage(
		context.Background(),
		"instagram_user_id_you_want_to_send_message_to",
		instabot.NewStickerMessage(instabot.StickerTypeHeart),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Sending media share.
	_, err = bot.SendMessage(
		context.Background(),
		"instagram_user_id_you_want_to_send_message_to",
		instabot.NewMediaShareMessage("media_id"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Sending generic template message.
	_, err = bot.SendMessage(
		context.Background(),
		"instagram_user_id_you_want_to_send_message_to",
		instabot.NewGenericTemplateMessage(
			[]*instabot.GenericTemplateElement{
				instabot.NewGenericTemplateElement(
					"Welcome!",
					instabot.WithTemplateImageURL("https://petersfancybrownhats.com/company_image.png"),
					instabot.WithTemplateSubtitle("We have the right hat for everyone."),
					instabot.WithTemplateDefaultAction("https://petersfancybrownhats.com/view?item=103"),
					instabot.WithTemplateButtons(
						[]instabot.Button{
							instabot.NewURLButton("View Website", "https://petersfancybrownhats.com"),
							instabot.NewPostBackButton("Start Chatting", "DEVELOPER_DEFINED_PAYLOAD"),
						},
					),
				),
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Sending product template message.
	_, err = bot.SendMessage(
		context.Background(),
		"instagram_user_id_you_want_to_send_message_to",
		instabot.NewProductTemplateMessage(
			[]*instabot.ProductTemplateElement{
				instabot.NewProductTemplateElement(
					"your_facebook_shop_product_id",
				),
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}

}
