package instabot

// IceBreaker defines Ice Breaker.
// frequently asked question.
// https://developers.facebook.com/docs/messenger-platform/instagram/features/ice-breakers
type IceBreaker struct {
	Question string `json:"question"`
	Payload  string `json:"payload"`
}

// NewIceBreaker returns a ice breaker.
func NewIceBreaker(question string, payload string) *IceBreaker {
	return &IceBreaker{
		Question: question,
		Payload:  payload,
	}
}
