package instabot

import (
	"encoding/json"
)

// Button defines button type.
type Button interface {
	Type() ButtonType
}

// ButtonType defines button type.
type ButtonType string

// all button type.
// https://developers.facebook.com/docs/messenger-platform/send-messages/buttons
const (
	ButtonTypeURL      ButtonType = ButtonType("web_url")
	ButtonTypePostBack ButtonType = ButtonType("postback")
	ButtonTypeCall     ButtonType = ButtonType("phone_number")
	ButtonTypeLogin    ButtonType = ButtonType("account_link")
	ButtonTypeLogOut   ButtonType = ButtonType("account_unlink")
	ButtonTypeGamePlay ButtonType = ButtonType("game_play")
)

// URLButton defines URL button.
// https://developers.facebook.com/docs/messenger-platform/send-messages/buttons#url
type URLButton struct {
	buttonType ButtonType
	Title      string
	URL        string
}

// NewURLButton returns a new url button.
func NewURLButton(title string, URL string) *URLButton {
	return &URLButton{
		buttonType: ButtonTypeURL,
		Title:      title,
		URL:        URL,
	}
}

// Type returns button type.
func (b *URLButton) Type() ButtonType {
	return b.buttonType
}

// MarshalJSON returns json of the button.
func (b *URLButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type  string `json:"type"`
		Title string `json:"title"`
		URL   string `json:"url"`
	}{
		Type:  string(b.buttonType),
		Title: b.Title,
		URL:   b.URL,
	})
}

// PostBackButton defines post back button.
// https://developers.facebook.com/docs/messenger-platform/send-messages/buttons#postback
type PostBackButton struct {
	buttonType ButtonType
	Title      string
	Payload    string
}

// NewPostBackButton returns a new post back button.
func NewPostBackButton(title string, payload string) *PostBackButton {
	return &PostBackButton{
		buttonType: ButtonTypePostBack,
		Title:      title,
		Payload:    payload,
	}
}

// Type returns button type.
func (b *PostBackButton) Type() ButtonType {
	return b.buttonType
}

// MarshalJSON returns json of the button.
func (b *PostBackButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    string `json:"type"`
		Title   string `json:"title"`
		Payload string `json:"payload"`
	}{
		Type:    string(b.buttonType),
		Title:   b.Title,
		Payload: b.Payload,
	})
}

// CallButton defines call button.
// https://developers.facebook.com/docs/messenger-platform/send-messages/buttons#call
type CallButton struct {
	buttonType  ButtonType
	Title       string
	PhoneNumber string
}

// NewCallButton returns a new call button.
func NewCallButton(title string, PhoneNumber string) *CallButton {
	return &CallButton{
		buttonType:  ButtonTypeCall,
		Title:       title,
		PhoneNumber: PhoneNumber,
	}
}

// Type returns button type.
func (b *CallButton) Type() ButtonType {
	return b.buttonType
}

// MarshalJSON returns json of the button.
func (b *CallButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    string `json:"type"`
		Title   string `json:"title"`
		Payload string `json:"payload"`
	}{
		Type:    string(b.buttonType),
		Title:   b.Title,
		Payload: b.PhoneNumber,
	})
}

// LogInButton defines log in button.
// https://developers.facebook.com/docs/messenger-platform/send-messages/buttons#login
type LogInButton struct {
	buttonType ButtonType
	URL        string
}

// NewLogInButton returns a new log in button.
func NewLogInButton(URL string) *LogInButton {
	return &LogInButton{
		buttonType: ButtonTypeLogin,
		URL:        URL,
	}
}

// Type returns button type.
func (b *LogInButton) Type() ButtonType {
	return b.buttonType
}

// MarshalJSON returns json of the button.
func (b *LogInButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	}{
		Type: string(b.buttonType),
		URL:  b.URL,
	})
}

// LogOutButton defines log out button.
// https://developers.facebook.com/docs/messenger-platform/send-messages/buttons#logout
type LogOutButton struct {
	buttonType ButtonType
}

// NewLogOutButton returns a new log out button.
func NewLogOutButton() *LogOutButton {
	return &LogOutButton{
		buttonType: ButtonTypeLogOut,
	}
}

// Type returns button type.
func (b *LogOutButton) Type() ButtonType {
	return b.buttonType
}

// MarshalJSON returns json of the button.
func (b *LogOutButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
	}{
		Type: string(b.buttonType),
	})
}

// TODO:- define GamePlayButton.
