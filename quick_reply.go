package instabot

import "encoding/json"

// QuickReplyType defines quick reply content type.
type QuickReplyType string

// all quick reply content type.
// Only text type quick reply is supported for instagram.
const (
	QuickReplyTypeText QuickReplyType = QuickReplyType("text")
)

// QuickReply defines quick reply.
// quick reply is only supported for text message on instagram platform.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/quick-replies
type QuickReply struct {
	quickReplyType QuickReplyType
	Title          string
	Payload        string
}

// NewTextQuickReply returns Text type quick reply.
func NewTextQuickReply(title string, payload string) *QuickReply {
	return &QuickReply{
		quickReplyType: QuickReplyTypeText,
		Title:          title,
		Payload:        payload,
	}
}

// MarshalJSON returns json of the quick reply item.
func (q *QuickReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ContentType QuickReplyType `json:"content_type"`
		Title       string         `json:"title"`
		Payload     string         `json:"payload"`
	}{
		ContentType: q.quickReplyType,
		Title:       q.Title,
		Payload:     q.Payload,
	})
}
