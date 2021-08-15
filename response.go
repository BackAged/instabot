package instabot

import (
	"encoding/json"
	"io"
	"net/http"
)

// APIError defines error received from the api.
type APIError struct {
	Message   string `json:"message"`
	Type      string `json:"type"`
	Code      int32  `json:"code"`
	SubCode   int32  `json:"error_subcode"`
	FbTraceID string `json:"fbtrace_id"`
}

// ErrorResponse defines error response received from instagram api.
type ErrorResponse struct {
	StatusCode int      `json:"status_code,omitempty"`
	APIError   APIError `json:"error"`
}

// Error return error messsage.
func (e ErrorResponse) Error() string {
	// TODO: improve error message format
	j, _ := json.Marshal(e)

	return string(j)
}

func checkErrorResponse(res *http.Response) error {
	if res.StatusCode/100 == 2 {
		return nil
	}

	decoder := json.NewDecoder(res.Body)

	response := ErrorResponse{
		StatusCode: res.StatusCode,
	}

	if err := decoder.Decode(&response); err != nil {
		// TODO: debug log, bug??
		return &response
	}

	return &response
}

// SendMessageResponse send message api success response.
type SendMessageResponse struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}

func decodeToSendMessageResponse(res *http.Response) (*SendMessageResponse, error) {
	if err := checkErrorResponse(res); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)

	response := SendMessageResponse{}

	if err := decoder.Decode(&response); err != nil {
		if err == io.EOF {
			return &response, nil
		}

		return nil, err
	}

	return &response, nil
}

// SetIceBreakersResponse defines set ice breaker api success response.
type SetIceBreakersResponse struct {
	Result string `json:"result"`
}

func decodeToSetIceBreakersResponse(res *http.Response) (*SetIceBreakersResponse, error) {
	if err := checkErrorResponse(res); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)

	response := SetIceBreakersResponse{}

	if err := decoder.Decode(&response); err != nil {
		if err == io.EOF {
			return &response, nil
		}

		return nil, err
	}

	return &response, nil
}

// IceBreakers holds list of ice breakers.
type IceBreakers struct {
	IceBreakers []IceBreaker `json:"ice_breakers"`
}

// GetIceBreakersResponse defines get ice breaker api success response.
type GetIceBreakersResponse struct {
	Data []IceBreakers `json:"data"`
}

func decodeToGetIceBreakersResponse(res *http.Response) (*GetIceBreakersResponse, error) {
	if err := checkErrorResponse(res); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)

	response := GetIceBreakersResponse{}

	if err := decoder.Decode(&response); err != nil {
		if err == io.EOF {
			return &response, nil
		}

		return nil, err
	}

	return &response, nil
}

// DeleteIceBreakersResponse defines delete ice breaker api success response.
type DeleteIceBreakersResponse struct {
	Result string `json:"result"`
}

func decodeToDeleteIceBreakersResponse(res *http.Response) (*DeleteIceBreakersResponse, error) {
	if err := checkErrorResponse(res); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)

	response := DeleteIceBreakersResponse{}

	if err := decoder.Decode(&response); err != nil {
		if err == io.EOF {
			return &response, nil
		}

		return nil, err
	}

	return &response, nil
}

// GetUserProfileResponse defines instagram get user profile response.
type GetUserProfileResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProfilePic string `json:"profile_pic"`
}

func decodeToGetUserProfileResponse(res *http.Response) (*GetUserProfileResponse, error) {
	if err := checkErrorResponse(res); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)

	response := GetUserProfileResponse{}

	if err := decoder.Decode(&response); err != nil {
		if err == io.EOF {
			return &response, nil
		}

		return nil, err
	}

	return &response, nil
}
