package instabot

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	pageAccessToken := "testPageAccessToken"
	httpClient := http.DefaultClient
	apiBaseEndpont := "https://www.testapi.com"

	type args struct {
		pageAccessToken string
		options         []ClientOption
	}

	type test struct {
		args      args
		wantErr   error
		afterEach func(client *Client)
	}

	testCases := map[string]func(t *testing.T) test{
		"it should return error, when page access token is not given": func(t *testing.T) test {
			return test{
				wantErr: ErrMissingPageAccessToken,
			}
		},
		"it should return client, when page access token is given": func(t *testing.T) test {
			return test{
				args: args{
					pageAccessToken: pageAccessToken,
				},
				wantErr: nil,
				afterEach: func(client *Client) {
					assert.Equal(t, pageAccessToken, client.pageAccessToken)
					assert.NotNil(t, client.httpClient)
					assert.NotNil(t, client.endpointBase)
				},
			}
		},
		"it should return client, when optional base endpoint is given": func(t *testing.T) test {
			return test{
				args: args{
					pageAccessToken: pageAccessToken,
					options:         []ClientOption{WithEndpointBase(apiBaseEndpont)},
				},
				wantErr: nil,
				afterEach: func(client *Client) {
					assert.Equal(t, pageAccessToken, client.pageAccessToken)
					assert.NotNil(t, client.httpClient)
					assert.Equal(t, apiBaseEndpont, client.endpointBase.String())
				},
			}
		},
		"it should return client, when optional httpClient is given": func(t *testing.T) test {
			return test{
				args: args{
					pageAccessToken: pageAccessToken,
					options:         []ClientOption{WithHTTPClient(httpClient)},
				},
				wantErr: nil,
				afterEach: func(client *Client) {
					assert.Equal(t, pageAccessToken, client.pageAccessToken)
					assert.NotNil(t, client.endpointBase)
					assert.Equal(t, httpClient, client.httpClient)
				},
			}
		},
	}

	for name, fn := range testCases {
		tt := fn(t)

		t.Run(name, func(t *testing.T) {
			client, err := New(tt.args.pageAccessToken, tt.args.options...)
			if tt.wantErr != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			if tt.afterEach != nil {
				tt.afterEach(client)
			}
		})
	}
}
