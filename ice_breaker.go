package instabot

import "encoding/json"

// IceBreaker defines Ice Breaker.
// frequently asked question.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers
type IceBreaker struct {
	Question string
	Payload  string
}

// NewIceBreaker returns a ice breaker.
func NewIceBreaker(question string, payload string) *IceBreaker {
	return &IceBreaker{
		Question: question,
		Payload:  payload,
	}
}

// MarshalJSON returns json of the ice breaker.
func (i *IceBreaker) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Question string `json:"question"`
		Payload  string `json:"payload"`
	}{
		Question: i.Question,
		Payload:  i.Payload,
	})
}
