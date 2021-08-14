package instabot

import "encoding/json"

// Template defines template.
type Template interface {
	TemplateType() TemplateType
}

// TemplateType defines available template type.
type TemplateType string

// all template type.
// generic, product template are available on instagram.
const (
	TemplateTypeProduct TemplateType = TemplateType("product")
	TemplateTypeGeneric TemplateType = TemplateType("generic")
)

// TemplateDefaultAction template default action.
type TemplateDefaultAction struct {
	buttonType ButtonType
	URL        string
}

// NewTemplateDefaultAction returns new TemplateDefaultAction.
func NewTemplateDefaultAction(URL string) *TemplateDefaultAction {
	return &TemplateDefaultAction{
		buttonType: ButtonTypeURL,
		URL:        URL,
	}
}

// MarshalJSON returns json of the button.
func (d *TemplateDefaultAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	}{
		Type: string(d.buttonType),
		URL:  d.URL,
	})
}

// GenericTemplateElement defines generic template element.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/generic-template#elements
type GenericTemplateElement struct {
	Title         string
	Subtitle      string
	ImageURL      string
	DefaultAction *TemplateDefaultAction
	Buttons       []Button
}

// GenericTemplateElementOption defines new GenericTemplateElement
// instantiation optional argument.
type GenericTemplateElementOption func(*GenericTemplateElement)

// WithTemplateSubtitle sets subtitle of a GenericTemplateElement.
func WithTemplateSubtitle(subtitle string) GenericTemplateElementOption {
	return func(gte *GenericTemplateElement) {
		gte.Subtitle = subtitle
	}
}

// WithTemplateImageURL sets image url of a GenericTemplateElement.
func WithTemplateImageURL(imageURL string) GenericTemplateElementOption {
	return func(gte *GenericTemplateElement) {
		gte.ImageURL = imageURL
	}
}

// WithTemplateDefaultAction sets default action of a GenericTemplateElement.
func WithTemplateDefaultAction(URL string) GenericTemplateElementOption {
	return func(gte *GenericTemplateElement) {
		gte.DefaultAction = &TemplateDefaultAction{
			buttonType: ButtonTypeURL,
			URL:        URL,
		}
	}
}

// WithTemplateButtons sets buttons of a GenericTemplateElement,
// a maximum of 3 buttons per element is supported.
func WithTemplateButtons(buttons []Button) GenericTemplateElementOption {
	return func(gte *GenericTemplateElement) {
		gte.Buttons = buttons
	}
}

// NewGenericTemplateElement returns new GenericTemplateElement.
func NewGenericTemplateElement(title string, opts ...GenericTemplateElementOption) *GenericTemplateElement {
	gte := &GenericTemplateElement{
		Title: title,
	}

	for _, opt := range opts {
		opt(gte)
	}

	return gte
}

// MarshalJSON marshal generic template element to JSON.
func (e *GenericTemplateElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Title         string                 `json:"title"`
		Subtitle      string                 `json:"subtitle,omitempty"`
		ImageURL      string                 `json:"image_url,omitempty"`
		DefaultAction *TemplateDefaultAction `json:"default_action,omitempty"`
		Buttons       []Button               `json:"buttons,omitempty"`
	}{
		Title:         e.Title,
		Subtitle:      e.Subtitle,
		ImageURL:      e.ImageURL,
		DefaultAction: e.DefaultAction,
		Buttons:       e.Buttons,
	})
}

// ProductTemplateElement defines product template element.
// https://developers.facebook.com/docs/messenger-platform/send-messages/template/product
type ProductTemplateElement struct {
	ProductID string
}

// NewProductTemplateElement returns a new product template element.
func NewProductTemplateElement(productID string) *ProductTemplateElement {
	return &ProductTemplateElement{
		ProductID: productID,
	}
}

// MarshalJSON marshal product template element to JSON.
func (e *ProductTemplateElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ProductID string `json:"id"`
	}{
		ProductID: e.ProductID,
	})
}
