package instabot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageType(t *testing.T) {
	testCases := []struct {
		name      string
		args      Message
		want      MessageType
		afterEach func(t *testing.T)
	}{
		{
			name: "text message",
			args: NewTextMessage("Hello, world"),
			want: MessageTypeText,
		},
		{
			name: "image message",
			args: NewImageMessage("www.image.com"),
			want: MessageTypeImage,
		},
		{
			name: "sticker message",
			args: NewStickerMessage(StickerTypeHeart),
			want: MessageTypeSticker,
		},
		{
			name: "media share message",
			args: NewMediaShareMessage("1000"),
			want: MessageTypeMediaShare,
		},
		{
			name: "generic template message",
			args: NewGenericTemplateMessage(
				[]*GenericTemplateElement{},
			),
			want: MessageTypeTemplate,
		},
		{
			name: "product template message",
			args: NewProductTemplateMessage(
				[]*ProductTemplateElement{},
			),
			want: MessageTypeTemplate,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.args.Type())
		})
	}
}

func TestTemplateMessageType(t *testing.T) {
	testCases := []struct {
		name      string
		args      Template
		want      TemplateType
		afterEach func(t *testing.T)
	}{
		{
			name: "generic template message",
			args: NewGenericTemplateMessage(
				[]*GenericTemplateElement{},
			),
			want: TemplateTypeGeneric,
		},
		{
			name: "product template message",
			args: NewProductTemplateMessage(
				[]*ProductTemplateElement{},
			),
			want: TemplateTypeProduct,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.args.TemplateType())
		})
	}
}

func TestMessageJSON(t *testing.T) {
	testCases := []struct {
		name      string
		args      Message
		want      string
		afterEach func(t *testing.T)
	}{
		{
			name: "text message",
			args: NewTextMessage("<TEXT>"),
			want: `{
				"text": "<TEXT>"
			}`,
		},
		{
			name: "text message with quick reply",
			args: NewTextMessage(
				"<SOME_TEXT>",
				WithQuickReplies(
					[]*QuickReply{
						NewTextQuickReply("<TITLE_1>", "<POSTBACK_PAYLOAD_1>"),
						NewTextQuickReply("<TITLE_2>", "<POSTBACK_PAYLOAD_2>"),
					},
				),
			),
			want: `{
				"text": "<SOME_TEXT>",
				"quick_replies": [
				  {
					"content_type": "text",
					"title": "<TITLE_1>",
					"payload": "<POSTBACK_PAYLOAD_1>"
				  },
				  {
					"content_type": "text",
					"title": "<TITLE_2>",
					"payload": "<POSTBACK_PAYLOAD_2>"
				  }
				]
			}`,
		},
		{
			name: "image message",
			args: NewImageMessage("<ASSET_URL>"),
			want: `{
				"attachment":{
					"type": "image", 
					"payload": {
						"url": "<ASSET_URL>"
					}
			  	}
			}`,
		},
		{
			name: "sticker message",
			args: NewStickerMessage(StickerTypeHeart),
			want: `{
				"attachment": {
					"type": "like_heart"
				}
			}`,
		},
		{
			name: "media share message",
			args: NewMediaShareMessage("<MEDIA_ID>"),
			want: `{
				"attachment": {
					"type": "media_share", 
					"payload": {
					  "id": "<MEDIA_ID>"
					}
				}
			}`,
		},
		{
			name: "generic template message",
			args: NewGenericTemplateMessage(
				[]*GenericTemplateElement{
					NewGenericTemplateElement(
						"Welcome!",
						WithTemplateImageURL("https://petersfancybrownhats.com/company_image.png"),
						WithTemplateSubtitle("We have the right hat for everyone."),
						WithTemplateDefaultAction("https://petersfancybrownhats.com/view?item=103"),
						WithTemplateButtons(
							[]Button{
								NewURLButton("View Website", "https://petersfancybrownhats.com"),
								NewPostBackButton("Start Chatting", "DEVELOPER_DEFINED_PAYLOAD"),
							},
						),
					),
				},
			),
			want: `{
				"attachment": {
					"type": "template",
					"payload": {
						"template_type": "generic",
						"elements": [
							{
								"title": "Welcome!",
								"image_url": "https://petersfancybrownhats.com/company_image.png",
								"subtitle": "We have the right hat for everyone.",
								"default_action": {
									"type": "web_url",
									"url": "https://petersfancybrownhats.com/view?item=103"
								},
								"buttons": [
									{
										"type": "web_url",
										"url": "https://petersfancybrownhats.com",
										"title": "View Website"
									},
									{
										"type": "postback",
										"title": "Start Chatting",
										"payload": "DEVELOPER_DEFINED_PAYLOAD"
									}              
								]      
							}
						]
					}
				}
			  }`,
		},
		{
			name: "generic template message without button",
			args: NewGenericTemplateMessage(
				[]*GenericTemplateElement{
					NewGenericTemplateElement(
						"Welcome!",
						WithTemplateImageURL("https://petersfancybrownhats.com/company_image.png"),
						WithTemplateSubtitle("We have the right hat for everyone."),
						WithTemplateDefaultAction("https://petersfancybrownhats.com/view?item=103"),
					),
				},
			),
			want: `{
				"attachment": {
					"type": "template",
					"payload": {
						"template_type": "generic",
						"elements": [
							{
								"title": "Welcome!",
								"image_url": "https://petersfancybrownhats.com/company_image.png",
								"subtitle": "We have the right hat for everyone.",
								"default_action": {
									"type": "web_url",
									"url": "https://petersfancybrownhats.com/view?item=103"
								}
							}
						]
					}
				}
			  }`,
		},
		{
			name: "generic template message without button & default action",
			args: NewGenericTemplateMessage(
				[]*GenericTemplateElement{
					NewGenericTemplateElement(
						"Welcome!",
						WithTemplateImageURL("https://petersfancybrownhats.com/company_image.png"),
						WithTemplateSubtitle("We have the right hat for everyone."),
					),
				},
			),
			want: `{
				"attachment": {
					"type": "template",
					"payload": {
						"template_type": "generic",
						"elements": [
							{
								"title": "Welcome!",
								"image_url": "https://petersfancybrownhats.com/company_image.png",
								"subtitle": "We have the right hat for everyone."
							}
						]
					}
				}
			  }`,
		},
		{
			name: "product template message",
			args: NewProductTemplateMessage(
				[]*ProductTemplateElement{
					NewProductTemplateElement("<PRODUCT_ID_1>"),
					NewProductTemplateElement("<PRODUCT_ID_2>"),
				},
			),
			want: `{
				"attachment": {
				    "type": "template",
					"payload": {
					   "template_type": "product",
					   "elements": [
							{
						  		"id": "<PRODUCT_ID_1>"
							},
							{
						 		"id": "<PRODUCT_ID_2>"
							}
					 	]
				  	}
				}
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			j, err := json.Marshal(tc.args)
			assert.NoError(t, err)

			assert.JSONEq(t, tc.want, string(j))
		})
	}
}
