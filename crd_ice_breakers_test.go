package instabot

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetIceBreakers(t *testing.T) {
	pageAccessToken := "page_access_token"

	type args struct {
		ctx         context.Context
		iceBreakers []*IceBreaker
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
		want    *SetIceBreakersResponse
	}

	tests := map[string]func(t *testing.T) test{
		"set icebreaker success": func(t *testing.T) test {
			args := args{
				ctx: context.Background(),
				iceBreakers: []*IceBreaker{
					NewIceBreaker("test?", "test"),
				},
			}

			fields := fields{
				wantRequestBody: `{
					"platform": "instagram",
					"ice_breakers": [
					   {
						  "question": "test?",
						  "payload": "test"
					   }
					]
				}`,
				returnResponse: fmt.Sprintf(`{
					"result": "success"
				}`),
				returnResponseCode: 200,
			}

			want := &SetIceBreakersResponse{
				Result: "success",
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"set icebreaker error": func(t *testing.T) test {
			args := args{
				ctx: context.Background(),
				iceBreakers: []*IceBreaker{
					NewIceBreaker("test?", "test"),
					NewIceBreaker("test?", "test"),
					NewIceBreaker("test?", "test"),
					NewIceBreaker("test?", "test"),
					NewIceBreaker("test?", "test"),
				},
			}

			fields := fields{
				wantRequestBody: `{
					"platform": "instagram",
					"ice_breakers": [
					   {
						  "question": "test?",
						  "payload": "test"
					   },
					   {
						"question": "test?",
						"payload": "test"
					   },
					   {
						"question": "test?",
						"payload": "test"
					   },
					   {
						"question": "test?",
						"payload": "test"
					   },
					   {
						"question": "test?",
						"payload": "test"
					   }
					]
				}`,
				returnResponse: fmt.Sprintf(`{
					"error": {
						"message":"error",
						"type":"invalid message",
						"code":100,
						"error_subcode":23434,
						"fbtrace_id":"fbtrace_id"
					}
				}`),
				returnResponseCode: 400,
			}

			return test{
				args:   args,
				fields: fields,
				want:   nil,
				wantErr: &ErrorResponse{
					StatusCode: 400,
					APIError: APIError{
						Message:   "error",
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

		assert.Equal(t, APIEndpointMessengerProfile, r.URL.Path)

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

			res, err := client.SetIceBreakers(tt.args.ctx, tt.args.iceBreakers)
			if tt.wantErr != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}

func TestGetIceBreakers(t *testing.T) {
	pageAccessToken := "page_access_token"

	type args struct {
		ctx context.Context
	}

	type fields struct {
		wantRequestQuery   url.Values
		returnResponse     string
		returnResponseCode int
	}

	type test struct {
		args    args
		fields  fields
		wantErr error
		want    *GetIceBreakersResponse
	}

	tests := map[string]func(t *testing.T) test{
		"get icebreakers success": func(t *testing.T) test {
			args := args{
				ctx: context.Background(),
			}

			q := url.Values{}
			q.Add("platform", Platform)
			q.Add("fields", "ice_breakers")
			q.Add("access_token", pageAccessToken)

			fields := fields{
				wantRequestQuery: q,
				returnResponse: fmt.Sprintf(`{
					"data": [
						{
						  "ice_breakers": [
								{
									"question": "<QUESTION>",
									"payload": "<PAYLOAD>"
								}
							]
					  	}
				   ]
				}`),
				returnResponseCode: 200,
			}

			want := &GetIceBreakersResponse{
				Data: []IceBreakers{
					{
						IceBreakers: []IceBreaker{
							*NewIceBreaker("<QUESTION>", "<PAYLOAD>"),
						},
					},
				},
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"get icebreakers error": func(t *testing.T) test {
			args := args{
				ctx: context.Background(),
			}

			q := url.Values{}
			q.Add("platform", Platform)
			q.Add("fields", "ice_breakers")
			q.Add("access_token", pageAccessToken)

			fields := fields{
				wantRequestQuery: q,
				returnResponse: fmt.Sprintf(`{
					"error": {
						"message": "error",
						"type": "invalid message",
						"code": 100,
						"error_subcode": 23434,
						"fbtrace_id": "fbtrace_id"
					}
				}`),
				returnResponseCode: 400,
			}

			return test{
				args:   args,
				fields: fields,
				want:   nil,
				wantErr: &ErrorResponse{
					StatusCode: 400,
					APIError: APIError{
						Message:   "error",
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

		assert.Equal(t, http.MethodGet, r.Method)

		assert.Equal(t, APIEndpointMessengerProfile, r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, tc.fields.wantRequestQuery, query)

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

			res, err := client.GetIceBreakers(tt.args.ctx)
			if tt.wantErr != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}

func TestDeleteIceBreakers(t *testing.T) {
	pageAccessToken := "page_access_token"

	type args struct {
		ctx context.Context
	}

	type fields struct {
		wantRequestQuery   url.Values
		returnResponse     string
		returnResponseCode int
	}

	type test struct {
		args    args
		fields  fields
		wantErr error
		want    *DeleteIceBreakersResponse
	}

	tests := map[string]func(t *testing.T) test{
		"delete icebreakers success": func(t *testing.T) test {
			args := args{
				ctx: context.Background(),
			}

			q := url.Values{}
			q.Add("platform", Platform)
			q.Add("access_token", pageAccessToken)

			fields := fields{
				wantRequestQuery: q,
				returnResponse: fmt.Sprintf(`{
					"result": "success"
				}`),
				returnResponseCode: 200,
			}

			want := &DeleteIceBreakersResponse{
				Result: "success",
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"delete icebreakers error": func(t *testing.T) test {
			args := args{
				ctx: context.Background(),
			}

			q := url.Values{}
			q.Add("platform", Platform)
			q.Add("access_token", pageAccessToken)

			fields := fields{
				wantRequestQuery: q,
				returnResponse: fmt.Sprintf(`{
					"error": {
						"message": "error",
						"type": "invalid message",
						"code": 100,
						"error_subcode": 23434,
						"fbtrace_id": "fbtrace_id"
					}
				}`),
				returnResponseCode: 400,
			}

			return test{
				args:   args,
				fields: fields,
				want:   nil,
				wantErr: &ErrorResponse{
					StatusCode: 400,
					APIError: APIError{
						Message:   "error",
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

		assert.Equal(t, http.MethodDelete, r.Method)

		assert.Equal(t, APIEndpointMessengerProfile, r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, tc.fields.wantRequestQuery, query)

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

			res, err := client.DeleteIceBreakers(tt.args.ctx)
			if tt.wantErr != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}
