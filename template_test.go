package instabot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultActionJSON(t *testing.T) {
	testCases := []struct {
		name string
		args *TemplateDefaultAction
		want string
	}{
		{
			name: "default action",
			args: NewTemplateDefaultAction("<url>"),
			want: `{
				"type":"web_url",
				"url":"<url>"
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

func TestNewGenericTemplateElement(t *testing.T) {
	type args struct {
		title   string
		options []GenericTemplateElementOption
	}

	type test struct {
		args      args
		afterEach func(gte *GenericTemplateElement)
	}

	testCases := map[string]func(t *testing.T) test{
		"it should have title, when no options are given": func(t *testing.T) test {
			return test{
				args: args{
					title: "title",
				},
				afterEach: func(gte *GenericTemplateElement) {
					assert.Equal(t, "title", gte.Title)
					assert.Equal(t, "", gte.Subtitle)
					assert.Equal(t, "", gte.ImageURL)
					assert.Nil(t, gte.DefaultAction)
					assert.Empty(t, gte.Buttons)
				},
			}
		},
		"it should have subtitle, when subtitle is given": func(t *testing.T) test {
			return test{
				args: args{
					title: "title",
					options: []GenericTemplateElementOption{
						WithTemplateSubtitle("subtitle"),
					},
				},
				afterEach: func(gte *GenericTemplateElement) {
					assert.Equal(t, "title", gte.Title)
					assert.Equal(t, "subtitle", gte.Subtitle)
					assert.Equal(t, "", gte.ImageURL)
					assert.Nil(t, gte.DefaultAction)
					assert.Empty(t, gte.Buttons)
				},
			}
		},
		"it should have image url, when image url is given": func(t *testing.T) test {
			return test{
				args: args{
					title: "title",
					options: []GenericTemplateElementOption{
						WithTemplateImageURL("image url"),
					},
				},
				afterEach: func(gte *GenericTemplateElement) {
					assert.Equal(t, "title", gte.Title)
					assert.Equal(t, "", gte.Subtitle)
					assert.Equal(t, "image url", gte.ImageURL)
					assert.Nil(t, gte.DefaultAction)
					assert.Empty(t, gte.Buttons)
				},
			}
		},
		"it should have default action, when default action is given": func(t *testing.T) test {
			return test{
				args: args{
					title: "title",
					options: []GenericTemplateElementOption{
						WithTemplateDefaultAction("url"),
					},
				},
				afterEach: func(gte *GenericTemplateElement) {
					assert.Equal(t, "title", gte.Title)
					assert.Equal(t, "", gte.Subtitle)
					assert.Equal(t, "", gte.ImageURL)
					assert.Equal(t, "url", gte.DefaultAction.URL)
					assert.Empty(t, gte.Buttons)
				},
			}
		},
		"it should have buttons, when buttons are given given": func(t *testing.T) test {
			buttons := []Button{
				NewURLButton("title", "url"),
			}
			return test{
				args: args{
					title: "title",
					options: []GenericTemplateElementOption{
						WithTemplateButtons(buttons),
					},
				},
				afterEach: func(gte *GenericTemplateElement) {
					assert.Equal(t, "title", gte.Title)
					assert.Equal(t, "", gte.Subtitle)
					assert.Equal(t, "", gte.ImageURL)
					assert.Nil(t, gte.DefaultAction)
					assert.Equal(t, buttons, gte.Buttons)
				},
			}
		},
	}

	for name, fn := range testCases {
		tt := fn(t)

		t.Run(name, func(t *testing.T) {
			gte := NewGenericTemplateElement(tt.args.title, tt.args.options...)

			if tt.afterEach != nil {
				tt.afterEach(gte)
			}
		})
	}
}

func TestGenericTemplateElementJSON(t *testing.T) {
	testCases := []struct {
		name string
		args *GenericTemplateElement
		want string
	}{
		{
			name: "generic template element with subtitle, image, default action, buttons",
			args: NewGenericTemplateElement(
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
			want: `{
				"title": "Welcome!",
				"image_url": "https://petersfancybrownhats.com/company_image.png",
				"subtitle": "We have the right hat for everyone.",
				"default_action": {
					"type": "web_url",
					"url": "https://petersfancybrownhats.com/view?item=103"
				},
				"buttons":[
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
			}`,
		},
		{
			name: "generic template element without button",
			args: NewGenericTemplateElement(
				"Welcome!",
				WithTemplateImageURL("https://petersfancybrownhats.com/company_image.png"),
				WithTemplateSubtitle("We have the right hat for everyone."),
				WithTemplateDefaultAction("https://petersfancybrownhats.com/view?item=103"),
			),
			want: `{
				"title": "Welcome!",
				"image_url": "https://petersfancybrownhats.com/company_image.png",
				"subtitle": "We have the right hat for everyone.",
				"default_action": {
					"type": "web_url",
					"url": "https://petersfancybrownhats.com/view?item=103"
				}
			}`,
		},
		{
			name: "generic template element without subtitle,button",
			args: NewGenericTemplateElement(
				"Welcome!",
				WithTemplateImageURL("https://petersfancybrownhats.com/company_image.png"),
				WithTemplateDefaultAction("https://petersfancybrownhats.com/view?item=103"),
			),
			want: `{
				"title": "Welcome!",
				"image_url": "https://petersfancybrownhats.com/company_image.png",
				"default_action": {
					"type": "web_url",
					"url": "https://petersfancybrownhats.com/view?item=103"
				}
			}`,
		},
		{
			name: "generic template element without subtitle,button, image",
			args: NewGenericTemplateElement(
				"Welcome!",
				WithTemplateDefaultAction("https://petersfancybrownhats.com/view?item=103"),
			),
			want: `{
				"title": "Welcome!",
				"default_action": {
					"type": "web_url",
					"url": "https://petersfancybrownhats.com/view?item=103"
				}
			}`,
		},
		{
			name: "generic template element without subtitle,button, image, default action",
			args: NewGenericTemplateElement(
				"Welcome!",
			),
			want: `{
				"title": "Welcome!"
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


func TestProductTemplateElementJSON(t *testing.T) {
	testCases := []struct {
		name string
		args *ProductTemplateElement
		want string
	}{
		{
			name: "product template element",
			args: NewProductTemplateElement(
				"<PRODUCT_ID>",
			),
			want: `{
				"id": "<PRODUCT_ID>"
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