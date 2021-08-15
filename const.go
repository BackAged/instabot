package instabot

import "fmt"

// instabot const values.
const (
	Platform   string = "instagram"
	APIVersion string = "v11.0"
)

// instagram messaging api endpoints.
var (
	APIEndpointBase             = "https://graph.facebook.com"
	APIEndpointSendMessage      = fmt.Sprintf("/%s/me/messages", APIVersion)
	APIEndpointMessengerProfile = fmt.Sprintf("/%s/me/messenger_profile", APIVersion)
	GetAPIEndpointUserProfile   = func(instagramUserID string) string {
		return fmt.Sprintf("/%s/%s", APIVersion, instagramUserID)
	}
)
