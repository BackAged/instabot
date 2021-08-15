package main

import (
	"context"
	"fmt"
	"log"

	"github.com/BackAged/instabot"
)

func main() {
	// https://developers.facebook.com/docs/messenger-platform/instagram/features/user-profile#user-profile-api

	// instantiating instabot.
	bot, err := instabot.New("your_instagram_business_account_page_access_token")
	if err != nil {
		log.Fatal(err)
	}

	// Getting user profile.
	profile, err := bot.GetUserProfile(
		context.Background(),
		"instagram_user_id_you_want_to_get_profile",
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profile)
}
