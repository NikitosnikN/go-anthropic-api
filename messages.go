package go_anthropic_api

import (
	"bytes"
	"context"
	"encoding/json"
)

type MessageRole string

const (
	User      MessageRole = "user"
	Assistant MessageRole = "assistant"
)

type MessageMetadata struct {
	UserId string `json:"user_id,omitempty"`
}

type MessageContentType struct {
	Type      string `json:"type,omitempty"`
	MediaType string `json:"media_type,omitempty"`
	Data      string `json:"data,omitempty"`
}

type MessageContent struct {
	Type   string              `json:"type"`
	Text   string              `json:"text,omitempty"`
	Source *MessageContentType `json:"source,omitempty"`
}

type Message struct {
	Role    string            `json:"role"`
	Content []*MessageContent `json:"content"`
}

type MessagesRequest struct {
	Model         string           `json:"model"`
	Messages      []*Message       `json:"messages"`
	System        string           `json:"system,omitempty"`
	MaxTokens     int              `json:"max_tokens"`
	Metadata      *MessageMetadata `json:"metadata,omitempty"`
	StopSequences []string         `json:"stop_sequences,omitempty"`
	Stream        bool             `json:"stream,omitempty"`
	Temperature   float32          `json:"temperature,omitempty"`
	TopP          float32          `json:"top_p,omitempty"`
	TopK          int32            `json:"top_k,omitempty"`
}

func (m *MessagesRequest) AddTextMessage(role MessageRole, text string) {
	if m.Messages == nil {
		m.Messages = []*Message{}
	}
	message := &Message{
		Role: string(role),
		Content: []*MessageContent{
			&MessageContent{
				Type:   "text",
				Text:   text,
				Source: nil,
			},
		},
	}
	m.Messages = append(m.Messages, message)
}

type MessageResponseUsage struct {
	InputToken  int32 `json:"input_token"`
	OutputToken int32 `json:"output_token"`
}

type MessageResponse struct {
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	Role          string                `json:"role"`
	Model         string                `json:"model"`
	Content       []*MessageContent     `json:"content"`
	StopReason    string                `json:"stop_reason"`
	StopSequences string                `json:"stop_sequences"`
	Usage         *MessageResponseUsage `json:"usage"`
}

// CreateMessageRequest - API call to create message
func (c *Client) CreateMessageRequest(ctx context.Context, request MessagesRequest) (*MessageResponse, error) {
	request.Stream = false

	response := MessageResponse{}

	rawRequest, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	httpRequest, err := c.makeRequest(ctx, "messages", "POST", bytes.NewReader(rawRequest))

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

	httpRequest, err := c.makeRequest(ctx, "messages", "POST", bytes.NewReader(rawRequest))

	if err != nil {
		return nil, err
	}

	return c.sendRequestStream(httpRequest)
}
