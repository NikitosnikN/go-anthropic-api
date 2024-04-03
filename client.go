package go_anthropic_api

import "C"
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type Client struct {
	apikey     string
	apiUrl     string
	apiVersion string
	httpClient *http.Client
	mu         sync.Mutex
}

const (
	apiUrlV1            = "https://api.anthropic.com"
	apiAuthHeaderKey    = "x-api-key"
	apiVersionHeaderKey = "anthropic-version"
	defaultApiVersion   = "2023-06-01"
)

func NewClient(apiKey string) *Client {
	return &Client{
		apikey:     apiKey,
		apiUrl:     apiUrlV1,
		apiVersion: defaultApiVersion,
		httpClient: &http.Client{},
		mu:         sync.Mutex{},
	}
}

func (c *Client) SetProxy(proxyUrl string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if proxyUrl != "" {
		proxyURL, err := url.Parse(proxyUrl)
		if err != nil {
			return err
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		c.httpClient.Transport = transport
	} else {
		c.httpClient.Transport = &http.Transport{}
	}
	return nil
}

func (c *Client) SetApiVersion(apiVersion string) {
	c.apiVersion = apiVersion
}

func (c *Client) SetApiUrl(apiUrl string) {
	c.apiUrl = apiUrl
}

func (c *Client) makeRequest(ctx context.Context, path string, method string, body io.Reader) (*http.Request, error) {
	fullUrl := fmt.Sprintf("%s%s", c.apiUrl, path)

	request, err := http.NewRequestWithContext(ctx, method, fullUrl, body)

	if err != nil {
		return nil, err
	}

	request.Header.Add(apiAuthHeaderKey, c.apikey)
	request.Header.Add(apiVersionHeaderKey, c.apiVersion)
	request.Header.Add("content-type", "application/json")
	return request, err
}

func (c *Client) sendRequest(request *http.Request, responsePayload interface{}) error {
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError APIError

		if err = json.Unmarshal(payload, &apiError); err == nil {
			return &apiError
		}

		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return json.Unmarshal(payload, &responsePayload)
}

func (c *Client) sendRequestStream(request *http.Request) (*StreamReader, error) {
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()

		payload, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var apiError APIError

		if err = json.Unmarshal(payload, &apiError); err == nil {
			return nil, &apiError
		}

		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return NewStreamReader(resp.Body), nil
}

// CreateMessageRequest - API call to create message
func (c *Client) CreateMessageRequest(ctx context.Context, request MessagesRequest) (*MessageResponse, error) {
	request.Stream = false

	response := MessageResponse{}

	rawRequest, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	httpRequest, err := c.makeRequest(ctx, "/v1/messages", "POST", bytes.NewReader(rawRequest))

	if err != nil {
		return nil, err
	}

	err = c.sendRequest(httpRequest, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateMessageRequestStream(ctx context.Context, request MessagesRequest) (*StreamReader, error) {
	request.Stream = true

	rawRequest, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	httpRequest, err := c.makeRequest(ctx, "/v1/messages", "POST", bytes.NewReader(rawRequest))

	if err != nil {
		return nil, err
	}

	return c.sendRequestStream(httpRequest)
}
