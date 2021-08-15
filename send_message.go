package instabot

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
)

func encodeSendMessageJSON(w io.Writer, recipeint string, message Message) error {
	enc := json.NewEncoder(w)

	return enc.Encode(&struct {
		Recipient *Recipient `json:"recipient"`
		Message   Message    `json:"message"`
	}{
		Recipient: &Recipient{
			ID: recipeint,
		},
		Message: message,
	})
}

// SendMessage sends message by calling instagram api.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/send-message#send-api
func (c *Client) SendMessage(ctx context.Context, recipient string, message Message) (*SendMessageResponse, error) {
	var buf bytes.Buffer
	if err := encodeSendMessageJSON(&buf, recipient, message); err != nil {
		return nil, err
	}

	res, err := c.post(ctx, APIEndpointSendMessage, &buf)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return decodeToSendMessageResponse(res)
}
