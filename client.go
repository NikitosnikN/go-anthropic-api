package go_anthropic_api

import "C"
import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	apikey     string
	apiUrl     string
	apiVersion string
	proxyUrl   string
}

const (
	apiUrlV1            = "https://api.anthropic.com/v1"
	apiAuthHeaderKey    = "x-api-key"
	apiVersionHeaderKey = "anthropic-version"
	defaultApiVersion   = "2023-06-01"
)

func NewClient(apiKey string) *Client {
	return &Client{
		apikey:     apiKey,
		apiUrl:     apiUrlV1,
		apiVersion: defaultApiVersion,
		proxyUrl:   "",
	}
}

func (c *Client) SetProxy(proxyUrl string) {
	c.proxyUrl = proxyUrl
}

func (c *Client) SetApiVersion(apiVersion string) {
	c.apiVersion = apiVersion
}

func (c *Client) SetApiUrl(apiUrl string) {
	c.apiUrl = apiUrl
}

func (c *Client) makeRequest(ctx context.Context, path string, method string, body io.Reader) (*http.Request, error) {
	fullUrl := fmt.Sprintf("%s/%s", c.apiUrl, path)

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
	transport := &http.Transport{}

	if c.proxyUrl != "" {
		proxyURL, err := url.Parse(c.proxyUrl)
		if err != nil {
			return err
		}

		transport.Proxy = http.ProxyURL(proxyURL)
	}

	httpClient := http.Client{
		Transport: transport,
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(payload, &responsePayload)
}

func (c *Client) sendRequestStream(request *http.Request) (*bufio.Reader, error) {
	transport := &http.Transport{}

	if c.proxyUrl != "" {
		proxyURL, err := url.Parse(c.proxyUrl)
		if err != nil {
			return nil, err
		}

		transport.Proxy = http.ProxyURL(proxyURL)
	}

	httpClient := http.Client{
		Transport: transport,
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	reader := bufio.NewReader(resp.Body)

	return reader, nil
}
