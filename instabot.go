package instabot

import "context"

// InstaBot defines InstaBot client interface.
type InstaBot interface {
	SendMessage(ctx context.Context, recipient string, message Message) (*SendMessageResponse, error)
	SetIceBreakers(ctx context.Context, iceBreakers []*IceBreaker) (*SetIceBreakersResponse, error)
	GetIceBreakers(ctx context.Context) (*GetIceBreakersResponse, error)
	DeleteIceBreakers(ctx context.Context) (*DeleteIceBreakersResponse, error)
	GetUserProfile(ctx context.Context, instagramUserID string) (*GetUserProfileResponse, error)
}

// compile time interface implementation check.
var _ InstaBot = (*Client)(nil)
