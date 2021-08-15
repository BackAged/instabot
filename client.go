package instabot

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"path"
)

// Client defines instabot.
type Client struct {
	pageAccessToken string
	endpointBase    *url.URL
	httpClient      *http.Client
}

// ClientOption defines optional argument for new client construction.
type ClientOption func(*Client) error

// WithHTTPClient sets client http client.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c

		return nil
	}
}

// WithEndpointBase sets client base endpoint.
func WithEndpointBase(endpointBase string) ClientOption {
	return func(client *Client) error {
		url, err := url.ParseRequestURI(endpointBase)
		if err != nil {
			return err
		}

		client.endpointBase = url

		return nil
	}
}

// New returns a new bot client instance.
func New(pageAccessToken string, options ...ClientOption) (*Client, error) {
	if pageAccessToken == "" {
		return nil, ErrMissingPageAccessToken
	}

	client := &Client{
		pageAccessToken: pageAccessToken,
	}

	for _, option := range options {
		err := option(client)
		if err != nil {
			return nil, err
		}
	}

	if client.endpointBase == nil {
		u, err := url.ParseRequestURI(APIEndpointBase)
		if err != nil {
			return nil, err
		}

		client.endpointBase = u
	}

	if client.httpClient == nil {
		client.httpClient = http.DefaultClient
	}

	return client, nil
}

func (client *Client) url(base *url.URL, endpoint string) string {
	u := *base
	u.Path = path.Join(u.Path, endpoint)

	// pass page_access_token to query string
	query := u.Query()

	query.Add("access_token", client.pageAccessToken)

	u.RawQuery = query.Encode()

	return u.String()
}

// abstraction for later usage like to set some global property of request, ex- header etc.
func (client *Client) do(req *http.Request) (*http.Response, error) {
	return client.httpClient.Do(req)
}

func (client *Client) get(ctx context.Context, endpoint string, query url.Values) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		client.url(client.endpointBase, endpoint),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if query != nil {
		query.Set("access_token", req.URL.Query().Get("access_token"))

		req.URL.RawQuery = query.Encode()
	}

	return client.do(req)
}

func (client *Client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		client.url(client.endpointBase, endpoint),
		body,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return client.do(req)
}

func (client *Client) put(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		client.url(client.endpointBase, endpoint),
		body,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return client.do(req)
}

func (client *Client) delete(ctx context.Context, endpoint string, body io.Reader, query url.Values) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		client.url(client.endpointBase, endpoint),
		body,
	)
	if err != nil {
		return nil, err
	}

	if query != nil {
		query.Set("access_token", req.URL.Query().Get("access_token"))

		req.URL.RawQuery = query.Encode()
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return client.do(req)
}
