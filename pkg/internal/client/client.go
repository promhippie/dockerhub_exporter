package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jackspirou/syscerts"
	"github.com/promhippie/dockerhub_exporter/pkg/version"
)

const (
	baseURL = "https://hub.docker.com"
)

// New initializes a new client instance.
func New(opts ...Option) (*Client, error) {
	options := newOptions(opts...)

	client := &Client{
		Username: options.Username,
		Password: options.Password,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					RootCAs: syscerts.SystemRootsPool(),
				},
			},
		},
	}

	if client.Username != "" && client.Password != "" {
		if err := client.Login(
			context.Background(),
		); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// Client implements all required api methods.
type Client struct {
	Username   string
	Password   string
	Token      string
	HTTPClient *http.Client
}

// NewRequest prepares a new HTTP request.
func (c *Client) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "http") {
		path = baseURL + path
	}

	req, err := http.NewRequest(method, path, body)

	if err != nil {
		return nil, err
	}

	if c.Token != "" {
		req.Header.Set(
			"Authorization",
			fmt.Sprintf(
				"JWT %s",
				c.Token,
			),
		)
	}

	if body != nil {
		req.Header.Set(
			"Content-Type",
			"application/json",
		)
	}

	req.Header.Set(
		"Accept",
		"application/json",
	)

	req.Header.Set(
		"User-Agent",
		fmt.Sprintf(
			"dockerhub_exporter/%s",
			version.String,
		),
	)

	return req.WithContext(ctx), nil
}

// Do handles the concrete HTTP request.
func (c *Client) Do(r *http.Request, v interface{}) (*Response, error) {
	resp, err := c.HTTPClient.Do(r)

	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp.StatusCode == http.StatusForbidden {
			return &Response{
				Response: resp,
			}, &ForbiddenError{}
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			parsed := make(map[string]string)

			if err := json.Unmarshal(body, &parsed); err == nil {
				for _, k := range []string{"message", "detail"} {
					if msg, ok := parsed[k]; ok {
						return &Response{
								Response: resp,
							}, &GeneralError{
								message: msg,
							}
					}
				}
			}
		}

		return &Response{
			Response: resp,
		}, &UnknownError{}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &Response{
			Response: resp,
		}, err
	}

	resp.Body = ioutil.NopCloser(
		bytes.NewReader(
			body,
		),
	)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(
				w,
				bytes.NewReader(
					body,
				),
			)
		} else {
			err = json.Unmarshal(
				body,
				v,
			)
		}
	}

	return &Response{
		Response: resp,
	}, err
}

// Response wraps the HTTP response object.
type Response struct {
	*http.Response
}
