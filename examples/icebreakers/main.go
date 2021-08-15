package main

import (
	"context"
	"fmt"
	"log"

	"github.com/BackAged/instabot"
)

func main() {
	// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers#ice-breakers

	// instantiating instabot.
	bot, err := instabot.New("your_instagram_business_account_page_access_token")
	if err != nil {
		log.Fatal(err)
	}

	// Setting icebreaker of a instagram business account id.
	// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers#setting-ice-breakers
	_, err = bot.SetIceBreakers(
		context.Background(),
		[]*instabot.IceBreaker{
			instabot.NewIceBreaker("frequently asked question 1", "user payload"),
			instabot.NewIceBreaker("frequently asked question 2", "user payload"),
			instabot.NewIceBreaker("frequently asked question 3", "user payload"),
			instabot.NewIceBreaker("frequently asked question 4", "user payload"),
		},
	)

	// Get icebreaker of a instagram business account id.
	// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers#getting-ice-breakers
	icebreakers, err := bot.GetIceBreakers(
		context.Background(),
	)

	fmt.Println(icebreakers)

	// Delete icebreaker of a instagram business account id.
	// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers#deleting-icebreakers
	_, err = bot.DeleteIceBreakers(
		context.Background(),
	)
}
