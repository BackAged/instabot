package instabot

import (
	"encoding/json"
)

// WebhookEventType defines webhook event type.
type WebhookEventType string

// all webhook event types.
const (
	WebhookEventTypeTextMessage  WebhookEventType = WebhookEventType("text")
	WebhookEventTypeImageMessage WebhookEventType = WebhookEventType("image")
	WebhookEventTypeAudioMessage WebhookEventType = WebhookEventType("audio")
	WebhookEventTypeVideoMessage WebhookEventType = WebhookEventType("video")
	WebhookEventTypeFileMessage  WebhookEventType = WebhookEventType("file")
	WebhookEventTypeShare        WebhookEventType = WebhookEventType("share")
	WebhookEventTypeMessageReply WebhookEventType = WebhookEventType("message_reply")
	WebhookEventTypeStoryMention WebhookEventType = WebhookEventType("story_mention")
	WebhookEventTypeStoryReply   WebhookEventType = WebhookEventType("story_reply")
	WebhookEventTypeQuickReply   WebhookEventType = WebhookEventType("quick_reply")
	WebhookEventTypeReaction     WebhookEventType = WebhookEventType("reaction")
	WebhookEventTypeMessageSeen  WebhookEventType = WebhookEventType("message_seen")
	WebhookEventTypePostBack     WebhookEventType = WebhookEventType("postback")
	WebhookEventTypeEcho         WebhookEventType = WebhookEventType("echo")
	WebhookEventTypeDeleted      WebhookEventType = WebhookEventType("deleted")
	WebhookEventTypeUnsupported  WebhookEventType = WebhookEventType("unsupported")
)

// ReplyToStory defines story details.
type ReplyToStory struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

// ReplyTo defines either reply to message or story.
type ReplyTo struct {
	MID   string        `json:"mid"`
	Story *ReplyToStory `json:"story"`
}

// AttachmentPayload defines different attachment payload.
type AttachmentPayload struct {
	URL string `json:"url"`
}

// Attachment defines attachment for different type like audio, video, media share, etc.
type Attachment struct {
	Type    string            `json:"type"`
	Payload AttachmentPayload `json:"payload"`
}

// ReferralProduct defines fb/instagram shop product.
type ReferralProduct struct {
	ID string `json:"id"`
}

// Referral holds product referral.
// TODO: integrate later
type Referral struct {
	Product ReferralProduct `json:"product"`
}

// Reaction defines reaction.
type Reaction struct {
	MID      string `json:"mid"`
	Action   string `json:"action"`
	Reaction string `json:"reaction"`
	Emoji    string `json:"emoji"`
}

// Read defines seen message details.
type Read struct {
	MID string `json:"mid"`
}

// WebhookQuickReply defines webhook quickreply.
type WebhookQuickReply struct {
	Payload string `json:"payload"`
}

// Postback defines postback.
type Postback struct {
	MID     string `json:"mid"`
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

// WebhookMessage defines different message event type details.
type WebhookMessage struct {
	MID           string             `json:"mid"`
	Text          string             `json:"text"`
	QuickReply    *WebhookQuickReply `json:"quick_reply"`
	Attachments   []*Attachment      `json:"attachments"`
	ReplyTo       *ReplyTo           `json:"reply_to"`
	IsEcho        bool               `json:"is_echo"`
	IsUnsupported bool               `json:"is_unsupported"`
	IsDeleted     bool               `json:"is_deleted"`
}

func (m *WebhookMessage) isQuickReply() bool {
	return m.QuickReply != nil
}

func (m *WebhookMessage) isStoryReply() bool {
	return m.ReplyTo != nil && m.ReplyTo.Story != nil
}

func (m *WebhookMessage) isMessageReply() bool {
	return m.ReplyTo != nil && m.ReplyTo.MID != ""
}

func (m *WebhookMessage) isStoryMention() bool {
	if len(m.Attachments) > 0 &&
		m.Attachments[0].Type == string(WebhookEventTypeStoryMention) {
		return true
	}

	return false
}

func (m *WebhookMessage) isShare() bool {
	if len(m.Attachments) > 0 &&
		m.Attachments[0].Type == string(WebhookEventTypeShare) {
		return true
	}

	return false
}

func (m *WebhookMessage) isImageMessage() bool {
	if len(m.Attachments) > 0 &&
		m.Attachments[0].Type == string(WebhookEventTypeImageMessage) {
		return true
	}

	return false
}

func (m *WebhookMessage) isAudioMessage() bool {
	if len(m.Attachments) > 0 &&
		m.Attachments[0].Type == string(WebhookEventTypeAudioMessage) {
		return true
	}

	return false
}

func (m *WebhookMessage) isVideoMessage() bool {
	if len(m.Attachments) > 0 &&
		m.Attachments[0].Type == string(WebhookEventTypeVideoMessage) {
		return true
	}

	return false
}

func (m *WebhookMessage) isFileMessage() bool {
	if len(m.Attachments) > 0 &&
		m.Attachments[0].Type == string(WebhookEventTypeFileMessage) {
		return true
	}

	return false
}

// Messaging defines events.
type Messaging struct {
	Type      WebhookEventType
	Sender    *Sender         `json:"sender"`
	Recipient *Recipient      `json:"recipient"`
	Timestamp int64           `json:"timestamp"`
	Message   *WebhookMessage `json:"message"`
	Read      *Read           `json:"read"`
	Reaction  *Reaction       `json:"reaction"`
	Referral  *Referral       `json:"referral"`
	PostBack  *Postback       `json:"postback"`
}

func (m *Messaging) isMessageEvent() bool {
	return m.Message != nil
}

func (m *Messaging) isMessageSeenEvent() bool {
	return m.Read != nil
}

func (m *Messaging) isReactionEvent() bool {
	return m.Reaction != nil
}

func (m *Messaging) isReferralEvent() bool {
	return m.Referral != nil
}

func (m *Messaging) isPostBackEvent() bool {
	return m.PostBack != nil
}

// Entry defines entry.
type Entry struct {
	ID        string       `json:"id"`
	Time      int64        `json:"time"`
	Messaging []*Messaging `json:"messaging"`
}

// WebhookEvent defines instagram webhook event payload.
type WebhookEvent struct {
	Object  string   `json:"object"`
	Entries []*Entry `json:"entry"`
}

func (e *WebhookEvent) setType() {
	for _, entry := range e.Entries {
		for _, event := range entry.Messaging {
			switch {
			case event.isMessageEvent():
				switch {
				case event.Message.IsEcho:
					event.Type = WebhookEventTypeEcho
				case event.Message.IsDeleted:
					event.Type = WebhookEventTypeDeleted
				case event.Message.IsUnsupported:
					event.Type = WebhookEventTypeUnsupported
				case event.Message.isAudioMessage():
					event.Type = WebhookEventTypeAudioMessage
				case event.Message.isFileMessage():
					event.Type = WebhookEventTypeFileMessage
				case event.Message.isImageMessage():
					event.Type = WebhookEventTypeImageMessage
				case event.Message.isVideoMessage():
					event.Type = WebhookEventTypeVideoMessage
				case event.Message.isMessageReply():
					event.Type = WebhookEventTypeMessageReply
				case event.Message.isQuickReply():
					event.Type = WebhookEventTypeQuickReply
				case event.Message.isShare():
					event.Type = WebhookEventTypeShare
				case event.Message.isStoryMention():
					event.Type = WebhookEventTypeStoryMention
				case event.Message.isStoryReply():
					event.Type = WebhookEventTypeStoryReply
				case event.Message.Text != "":
					event.Type = WebhookEventTypeTextMessage
				}
			case event.isMessageSeenEvent():
				event.Type = WebhookEventTypeMessageSeen
			case event.isReactionEvent():
				event.Type = WebhookEventTypeReaction
			case event.isPostBackEvent():
				event.Type = WebhookEventTypePostBack
			}
		}
	}
}

// UnmarshalJSON unmarshal json webhook events.
func (e *WebhookEvent) UnmarshalJSON(buffer []byte) error {
	type rawEntry struct {
		ID        string       `json:"id"`
		Time      int64        `json:"time"`
		Messaging []*Messaging `json:"messaging"`
		Message   []*Messaging `json:"message"`
	}

	// TODO: use this wrapper to parse referral & other
	// inconsistent structure later
	type rawWebhookEvent struct {
		Object  string      `json:"object"`
		Entries []*rawEntry `json:"entry"`
	}

	re := rawWebhookEvent{}

	if err := json.Unmarshal(buffer, &re); err != nil {
		return err
	}

	e.Object = re.Object

	for _, rEntry := range re.Entries {
		eEntry := Entry{
			ID:        rEntry.ID,
			Time:      rEntry.Time,
			Messaging: rEntry.Messaging,
		}

		if len(rEntry.Message) > 0 && len(rEntry.Messaging) == 0 {
			eEntry.Messaging = rEntry.Message
		}

		e.Entries = append(e.Entries, &eEntry)
	}

	e.setType()

	return nil
}
