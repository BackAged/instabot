package instabot

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
)

func encodePrivateReplyJSON(w io.Writer, commentID string, message Message) error {
	enc := json.NewEncoder(w)

	return enc.Encode(&struct {
		Recipient *PrivateReplyRecipient `json:"recipient"`
		Message   Message                `json:"message"`
	}{
		Recipient: &PrivateReplyRecipient{
			CommentID: commentID,
		},
		Message: message,
	})
}

// SendPrivateReply sends private reply by calling instagram api.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/private-replies
func (c *Client) SendPrivateReply(ctx context.Context, commentID string, message Message) (*SendMessageResponse, error) {
	var buf bytes.Buffer
	if err := encodePrivateReplyJSON(&buf, commentID, message); err != nil {
		return nil, err
	}

	res, err := c.post(ctx, APIEndpointSendMessage, &buf)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return decodeToSendMessageResponse(res)
}
