package instabot

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendPrivateReply(t *testing.T) {
	pageAccessToken := "page_access_token"
	privateReplyRecipient := "comment_id"

	type args struct {
		ctx     context.Context
		message Message
	}

	type fields struct {
		wantRequestBody    string
		returnResponse     string
		returnResponseCode int
	}

	type test struct {
		args    args
		fields  fields
		wantErr error
		want    *SendMessageResponse
	}

	tests := map[string]func(t *testing.T) test{
		"text message": func(t *testing.T) test {
			text := "hello"
			mID := "test_message_id"
			args := args{
				ctx:     context.Background(),
				message: NewTextMessage(text),
			}

			fields := fields{
				wantRequestBody: fmt.Sprintf(`{
					"recipient": {
						"comment_id":"%s"
					},
					"message": {
						"text":"%s"
					}
				}`, privateReplyRecipient, text),
				returnResponse: fmt.Sprintf(`{
					"recipient_id":"%s",
					"message_id":"%s"
				}`, privateReplyRecipient, mID),
				returnResponseCode: 200,
			}

			want := &SendMessageResponse{
				RecipientID: privateReplyRecipient,
				MessageID:   mID,
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"text message with quick replies": func(t *testing.T) test {
			text := "hello"
			qTitle := "title"
			qPayload := "payload"
			mID := "test_message_id"
			args := args{
				ctx: context.Background(),
				message: NewTextMessage(
					text,
					WithQuickReplies([]*QuickReply{
						NewTextQuickReply(qTitle, qPayload),
					})),
			}

			fields := fields{
				wantRequestBody: fmt.Sprintf(`{
					"recipient": {
						"comment_id":"%s"
					},
					"message": {
						"text":"%s",
						"quick_replies":[
							{
								"content_type":"text",
								"title":"%s",
								"payload":"%s"
							}
						]
					}
				}`, privateReplyRecipient, text, qTitle, qPayload),
				returnResponse: fmt.Sprintf(`{
					"recipient_id":"%s",
					"message_id":"%s"
				}`, privateReplyRecipient, mID),
				returnResponseCode: 200,
			}

			want := &SendMessageResponse{
				RecipientID: privateReplyRecipient,
				MessageID:   mID,
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"image message": func(t *testing.T) test {
			url := "url"
			mID := "test_message_id"
			args := args{
				ctx:     context.Background(),
				message: NewImageMessage(url),
			}

			fields := fields{
				wantRequestBody: fmt.Sprintf(`{
					"recipient": {
						"comment_id":"%s"
					},
					"message": {
						"attachment":{
							"type":"image", 
							"payload":{
								"url":"%s"
							}
						}
					}
				}`, privateReplyRecipient, url),
				returnResponse: fmt.Sprintf(`{
					"recipient_id":"%s",
					"message_id":"%s"
				}`, privateReplyRecipient, mID),
				returnResponseCode: 200,
			}

			want := &SendMessageResponse{
				RecipientID: privateReplyRecipient,
				MessageID:   mID,
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"sticker message": func(t *testing.T) test {
			sticker := StickerTypeHeart
			mID := "test_message_id"
			args := args{
				ctx:     context.Background(),
				message: NewStickerMessage(sticker),
			}

			fields := fields{
				wantRequestBody: fmt.Sprintf(`{
					"recipient": {
						"comment_id":"%s"
					},
					"message": {
						"attachment":{
							"type":"%s"
						}
					}
				}`, privateReplyRecipient, sticker),
				returnResponse: fmt.Sprintf(`{
					"recipient_id":"%s",
					"message_id":"%s"
				}`, privateReplyRecipient, mID),
				returnResponseCode: 200,
			}

			want := &SendMessageResponse{
				RecipientID: privateReplyRecipient,
				MessageID:   mID,
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"media share message": func(t *testing.T) test {
			mediaID := "test_media_id"
			mID := "test_message_id"
			args := args{
				ctx:     context.Background(),
				message: NewMediaShareMessage(mediaID),
			}

			fields := fields{
				wantRequestBody: fmt.Sprintf(`{
					"recipient": {
						"comment_id":"%s"
					},
					"message": {
						"attachment":{
							"type":"media_share",
							"payload":{
								"id":"%s"
							}
						}
					}
				}`, privateReplyRecipient, mediaID),
				returnResponse: fmt.Sprintf(`{
					"recipient_id":"%s",
					"message_id":"%s"
				}`, privateReplyRecipient, mID),
				returnResponseCode: 200,
			}

			want := &SendMessageResponse{
				RecipientID: privateReplyRecipient,
				MessageID:   mID,
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"generic template message": func(t *testing.T) test {
			mID := "test_message_id"
			args := args{
				ctx: context.Background(),
				message: NewGenericTemplateMessage(
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
			}

			fields := fields{
				wantRequestBody: fmt.Sprintf(`{
					"recipient": {
						"comment_id":"%s"
					},
					"message": {
						"attachment":{
							"type":"template",
							"payload":{
								"template_type":"generic",
								"elements":[
									{
										"title":"Welcome!",
										"image_url":"https://petersfancybrownhats.com/company_image.png",
										"subtitle":"We have the right hat for everyone.",
										"default_action": {
											"type": "web_url",
											"url": "https://petersfancybrownhats.com/view?item=103"
										},
										"buttons":[
											{
												"type":"web_url",
												"url":"https://petersfancybrownhats.com",
												"title":"View Website"
											},
											{
												"type":"postback",
												"title":"Start Chatting",
												"payload":"DEVELOPER_DEFINED_PAYLOAD"
											}              
										]      
									}
								]
							}
						}
					}
				}`, privateReplyRecipient),
				returnResponse: fmt.Sprintf(`{
					"recipient_id":"%s",
					"message_id":"%s"
				}`, privateReplyRecipient, mID),
				returnResponseCode: 200,
			}

			want := &SendMessageResponse{
				RecipientID: privateReplyRecipient,
				MessageID:   mID,
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"text message with empty text- error": func(t *testing.T) test {
			text := ""
			args := args{
				ctx:     context.Background(),
				message: NewTextMessage(text),
			}

			fields := fields{
				wantRequestBody: fmt.Sprintf(`{
					"recipient": {
						"comment_id":"%s"
					},
					"message": {
						"text":"%s"
					}
				}`, privateReplyRecipient, text),
				returnResponse: `{
					"error": {
						"message":"invalid text",
						"type":"invalid message",
						"code":100,
						"error_subcode":23434,
						"fbtrace_id":"fbtrace_id"
					}
				}`,
				returnResponseCode: 400,
			}

			return test{
				args:   args,
				fields: fields,
				want:   nil,
				wantErr: &ErrorResponse{
					StatusCode: 400,
					APIError: APIError{
						Message:   "invalid text",
						Type:      "invalid message",
						Code:      100,
						SubCode:   23434,
						FbTraceID: "fbtrace_id",
					},
				},
			}
		},
	}

	var currentTest string
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc := tests[currentTest](t)

		assert.Equal(t, http.MethodPost, r.Method)

		assert.Equal(t, APIEndpointSendMessage, r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		assert.JSONEq(t, string(tc.fields.wantRequestBody), string(body))

		w.WriteHeader(tc.fields.returnResponseCode)
		w.Write([]byte(tc.fields.returnResponse))
	}))
	defer mockServer.Close()

	for name, fn := range tests {
		currentTest = name
		tt := fn(t)

		t.Run(name, func(t *testing.T) {
			client, err := New(pageAccessToken, WithEndpointBase(mockServer.URL))
			assert.NoError(t, err)

			res, err := client.SendPrivateReply(tt.args.ctx, privateReplyRecipient, tt.args.message)
			if tt.wantErr != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}