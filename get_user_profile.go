package instabot

import (
	"context"
)

// GetUserProfile fetches user profile by instagram user id.
func (c *Client) GetUserProfile(ctx context.Context, instagramUserID string) (*GetUserProfileResponse, error) {
	res, err := c.get(ctx, GetAPIEndpointUserProfile(instagramUserID), nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return decodeToGetUserProfileResponse(res)
}
