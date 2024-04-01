package go_anthropic_api

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMessageRequest(t *testing.T) {
	t.Parallel()
	request := NewMessageRequest("claude-3-haiku", 1024)

	require.Equal(t, request.Model, "claude-3-haiku")
	require.Equal(t, request.MaxTokens, 1024)
	require.Equal(t, len(request.Messages), 0)
}

func TestAddTextMessage(t *testing.T) {
	t.Parallel()

	t.Run("add single message", func(t *testing.T) {
		request := NewMessageRequest("claude-3-haiku", 1024)
		request.AddTextMessage("user", "hello")
		require.Equal(t, len(request.Messages), 1)
		require.Equal(t, request.Messages[0].Role, "user")
		require.Equal(t, request.Messages[0].Content[0].Type, "text")
		require.Equal(t, request.Messages[0].Content[0].Text, "hello")
	})
	t.Run("add multiple messages", func(t *testing.T) {
		request := NewMessageRequest("claude-3-haiku", 1024)
		request.AddTextMessage("user", "hello")
		request.AddTextMessage("assistant", "world")
		require.Equal(t, len(request.Messages), 2)
	})
}

func TestAddImage(t *testing.T) {
	t.Parallel()

	image := []byte(`test`)
	imageInBase64 := "dGVzdA=="
	request := NewMessageRequest("claude-3-haiku", 1024)

	request.AddImageMessage("user", image, "image/jpeg", "hello")
	require.Equal(t, len(request.Messages), 1)
	require.Equal(t, request.Messages[0].Content[0].Type, "image")
	require.Equal(t, request.Messages[0].Content[0].Source.Type, "base64")
	require.Equal(t, request.Messages[0].Content[0].Source.MediaType, "image/jpeg")
	require.Equal(t, request.Messages[0].Content[0].Source.Data, imageInBase64)
}

func TestClearMessages(t *testing.T) {
	t.Parallel()

	request := NewMessageRequest("claude-3-haiku", 1024)
	request.AddTextMessage("user", "hello")

	require.Equal(t, len(request.Messages), 1)

	request.ClearMessages()
	require.Equal(t, len(request.Messages), 0)
}
