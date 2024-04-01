package go_anthropic_api

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Parallel()
	client := NewClient("your-api-key")
	require.Equal(t, client.apikey, "your-api-key")
	require.Equal(t, client.apiUrl, apiUrlV1)
	require.Equal(t, client.apiVersion, defaultApiVersion)
	require.NotNil(t, client.httpClient)
}

func TestSetProxy(t *testing.T) {
	t.Parallel()

	t.Run("Set proxy", func(t *testing.T) {
		client := NewClient("your-api-key")
		err := client.SetProxy("http://localhost:8080")
		require.NoError(t, err)
		require.NotNil(t, client.httpClient.Transport.(*http.Transport).Proxy)
	})

	t.Run("Unset proxy", func(t *testing.T) {
		client := NewClient("your-api-key")
		err := client.SetProxy("http://localhost:8080")
		require.NoError(t, err)
		require.NotNil(t, client.httpClient.Transport.(*http.Transport).Proxy)

		err = client.SetProxy("")
		require.NoError(t, err)
		require.Nil(t, client.httpClient.Transport.(*http.Transport).Proxy)
	})
}
