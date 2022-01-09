package instabot

import (
	"context"
	"net/url"
)

// GetUserProfile fetches user profile by instagram user id.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/user-profile#user-profile-api
func (c *Client) GetUserProfile(ctx context.Context, instagramUserID string) (*GetUserProfileResponse, error) {
	query := url.Values{}
	query.Add(
		"fields",
		"name,profile_pic,is_verified_user,follower_count,is_user_follow_business,is_business_follow_user",
	)

	res, err := c.get(ctx, GetAPIEndpointUserProfile(instagramUserID), query)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return decodeToGetUserProfileResponse(res)
}
