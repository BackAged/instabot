package instabot

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/url"
)

func encodeSetIceBreakersJSON(w io.Writer, iceBreakers []*IceBreaker) error {
	enc := json.NewEncoder(w)

	return enc.Encode(&struct {
		Platform    string        `json:"platform"`
		IceBreakers []*IceBreaker `json:"ice_breakers"`
	}{
		Platform:    Platform,
		IceBreakers: iceBreakers,
	})
}

// SetIceBreakers sets a instagram account ice breakers.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers#setting-ice-breakers
func (c *Client) SetIceBreakers(ctx context.Context, iceBreakers []*IceBreaker) (*SetIceBreakersResponse, error) {
	var buf bytes.Buffer
	if err := encodeSetIceBreakersJSON(&buf, iceBreakers); err != nil {
		return nil, err
	}

	res, err := c.post(ctx, APIEndpointMessengerProfile, &buf)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return decodeToSetIceBreakersResponse(res)
}

// GetIceBreakers fetches ice breakers.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers#getting-ice-breakers
func (c *Client) GetIceBreakers(ctx context.Context) (*GetIceBreakersResponse, error) {
	query := url.Values{}
	query.Add("fields", "ice_breakers")
	query.Add("platform", Platform)

	res, err := c.get(ctx, APIEndpointMessengerProfile, query)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return decodeToGetIceBreakersResponse(res)
}

func encodeDeletetIceBreakersJSON(w io.Writer) error {
	enc := json.NewEncoder(w)

	return enc.Encode(&struct {
		Fields []string `json:"fields"`
	}{
		Fields: []string{"ice_breakers"},
	})
}

// DeleteIceBreakers deletes ice breakers.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers#deleting-icebreakers
func (c *Client) DeleteIceBreakers(ctx context.Context) (*DeleteIceBreakersResponse, error) {
	var buf bytes.Buffer
	if err := encodeDeletetIceBreakersJSON(&buf); err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Add("platform", Platform)

	res, err := c.delete(ctx, APIEndpointMessengerProfile, &buf, query)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return decodeToDeleteIceBreakersResponse(res)
}
