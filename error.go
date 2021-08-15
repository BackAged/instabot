package instabot

import "errors"

var (
	// ErrMissingPageAccessToken happens when instantiating instabot
	// with empty page access token.
	ErrMissingPageAccessToken = errors.New("missing page access token")
)
