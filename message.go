package instabot

import (
	"encoding/json"
)

// Message type interface.
type Message interface {
	Type() MessageType
}

// MessageType defines instagram message type.
type MessageType string

// all message type.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/send-message#supported-messages
const (
	MessageTypeText       MessageType = MessageType("text")
	MessageTypeImage      MessageType = MessageType("image")
	MessageTypeSticker    MessageType = MessageType("sticker")
	MessageTypeMediaShare MessageType = MessageType("media_share")
	MessageTypeReacton    MessageType = MessageType("reaction")
	MessageTypeTemplate   MessageType = MessageType("template")
)

// TextMessage defines text message.
type TextMessage struct {
	messageType     MessageType
	Text            string
	quickReplyItems []*QuickReply
}

// TextMessageOption defines optional argument for new text message construction.
type TextMessageOption func(*TextMessage)

// WithQuickReplies adds quick reply items to text message.
// Quick reply is only supported with text message.
func WithQuickReplies(quickReplyItems []*QuickReply) TextMessageOption {
	return func(m *TextMessage) {
		m.quickReplyItems = quickReplyItems
	}
}

// NewTextMessage returns a new text message.
func NewTextMessage(text string, options ...TextMessageOption) *TextMessage {
	m := &TextMessage{
		messageType: MessageTypeText,
		Text:        text,
	}

	for _, option := range options {
		option(m)
	}

	return m
}

// AttachQuickReplies attach quick reply items to text message
func (m *TextMessage) AttachQuickReplies(quickReplies []*QuickReply) {
	m.quickReplyItems = quickReplies
}

// Type returns message type.
func (m *TextMessage) Type() MessageType {
	return m.messageType
}

// MarshalJSON returns json of the message.
func (m *TextMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Text            string        `json:"text"`
		QuickReplyItems []*QuickReply `json:"quick_replies,omitempty"`
	}{
		Text:            m.Text,
		QuickReplyItems: m.quickReplyItems,
	})
}

// ImageMessage defines image message.
type ImageMessage struct {
	messageType MessageType
	ImageURL    string
}

// NewImageMessage returns a new image message.
func NewImageMessage(imageURL string) *ImageMessage {
	return &ImageMessage{
		messageType: MessageTypeImage,
		ImageURL:    imageURL,
	}
}

// Type returns message type.
func (m *ImageMessage) Type() MessageType {
	return m.messageType
}

// MarshalJSON returns json of the message.
func (m *ImageMessage) MarshalJSON() ([]byte, error) {
	type Payload struct {
		ImageURL string `json:"url"`
	}

	type Attachment struct {
		Type    string   `json:"type"`
		Payload *Payload `json:"payload"`
	}

	return json.Marshal(&struct {
		Attachment *Attachment `json:"attachment"`
	}{
		Attachment: &Attachment{
			Type: string(m.messageType),
			Payload: &Payload{
				ImageURL: m.ImageURL,
			},
		},
	})
}

// StickerType defines sticker type.
type StickerType string

// all sticker type.
const (
	StickerTypeHeart StickerType = StickerType("like_heart")
)

// StickerMessage defines sticker message.
type StickerMessage struct {
	messageType MessageType
	Sticker     StickerType
}

// NewStickerMessage returns sticker message.
func NewStickerMessage(sticker StickerType) *StickerMessage {
	return &StickerMessage{
		messageType: MessageTypeSticker,
		Sticker:     sticker,
	}
}

// Type returns message type.
func (m *StickerMessage) Type() MessageType {
	return m.messageType
}

// MarshalJSON returns json of the message.
func (m *StickerMessage) MarshalJSON() ([]byte, error) {
	type Attachment struct {
		Type string `json:"type"`
	}

	return json.Marshal(&struct {
		Attachment *Attachment `json:"attachment"`
	}{
		Attachment: &Attachment{
			Type: string(m.Sticker),
		},
	})
}

// MediaShareMessage defines media share message.
type MediaShareMessage struct {
	messageType MessageType
	MediaID     string
}

// NewMediaShareMessage returns new media share message.
func NewMediaShareMessage(mediaID string) *MediaShareMessage {
	return &MediaShareMessage{
		messageType: MessageTypeMediaShare,
		MediaID:     mediaID,
	}
}

// Type returns message type.
func (m *MediaShareMessage) Type() MessageType {
	return m.messageType
}

// MarshalJSON returns json of the message.
func (m *MediaShareMessage) MarshalJSON() ([]byte, error) {
	type Payload struct {
		MediaID string `json:"id"`
	}

	type Attachment struct {
		Type    string   `json:"type"`
		Payload *Payload `json:"payload"`
	}

	return json.Marshal(&struct {
		Attachment *Attachment `json:"attachment"`
	}{
		Attachment: &Attachment{
			Type: string(m.messageType),
			Payload: &Payload{
				MediaID: m.MediaID,
			},
		},
	})
}

// GenericTemplateMessage defines generic template message.
// A generic template message could have maximum of 10 elements.
type GenericTemplateMessage struct {
	messageType  MessageType
	templateType TemplateType
	Elements     []*GenericTemplateElement
}

// NewGenericTemplateMessage returns generic template element.
// A generic template message could have maximum of 10 elements.
func NewGenericTemplateMessage(elements []*GenericTemplateElement) *GenericTemplateMessage {
	return &GenericTemplateMessage{
		messageType:  MessageTypeTemplate,
		templateType: TemplateTypeGeneric,
		Elements:     elements,
	}
}

// Type returns message type.
func (m *GenericTemplateMessage) Type() MessageType {
	return m.messageType
}

// TemplateType returns template type.
func (m *GenericTemplateMessage) TemplateType() TemplateType {
	return m.templateType
}

// MarshalJSON returns json of the message.
func (m *GenericTemplateMessage) MarshalJSON() ([]byte, error) {
	type Payload struct {
		TemplateType string                    `json:"template_type"`
		Elements     []*GenericTemplateElement `json:"elements"`
	}

	type Attachment struct {
		Type    string   `json:"type"`
		Payload *Payload `json:"payload"`
	}

	return json.Marshal(&struct {
		Attachment *Attachment `json:"attachment"`
	}{
		Attachment: &Attachment{
			Type: string(m.messageType),
			Payload: &Payload{
				TemplateType: string(m.templateType),
				Elements:     m.Elements,
			},
		},
	})
}

// ProductTemplateMessage defines product template message.
// A product template message could have maximum of 10 elements.
type ProductTemplateMessage struct {
	messageType  MessageType
	templateType TemplateType
	Elements     []*ProductTemplateElement
}

// NewProductTemplateMessage returns product template element.
// A product template message could have maximum of 10 elements.
func NewProductTemplateMessage(elements []*ProductTemplateElement) *ProductTemplateMessage {
	return &ProductTemplateMessage{
		messageType:  MessageTypeTemplate,
		templateType: TemplateTypeProduct,
		Elements:     elements,
	}
}

// Type returns message type.
func (m *ProductTemplateMessage) Type() MessageType {
	return m.messageType
}

// TemplateType returns template type.
func (m *ProductTemplateMessage) TemplateType() TemplateType {
	return m.templateType
}

// MarshalJSON returns json of the message.
func (m *ProductTemplateMessage) MarshalJSON() ([]byte, error) {
	type Payload struct {
		TemplateType string                    `json:"template_type"`
		Elements     []*ProductTemplateElement `json:"elements"`
	}

	type Attachment struct {
		Type    string   `json:"type"`
		Payload *Payload `json:"payload"`
	}

	return json.Marshal(&struct {
		Attachment *Attachment `json:"attachment"`
	}{
		Attachment: &Attachment{
			Type: string(m.messageType),
			Payload: &Payload{
				TemplateType: string(m.templateType),
				Elements:     m.Elements,
			},
		},
	})
}
