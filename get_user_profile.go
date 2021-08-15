package instabot

import (
	"context"
)

// GetUserProfile fetches user profile by instagram user id.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/user-profile#user-profile-api
func (c *Client) GetUserProfile(ctx context.Context, instagramUserID string) (*GetUserProfileResponse, error) {
	res, err := c.get(ctx, GetAPIEndpointUserProfile(instagramUserID), nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return decodeToGetUserProfileResponse(res)
}
