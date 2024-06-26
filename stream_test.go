package go_anthropic_api

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

var rawStreamResponse = `event: message_start

data: {"type":"message_start","message":{"id":"msg_X","type":"message","role":"assistant","content":[],"model":"claude-3-haiku-20240307","stop_reason":null,"stop_sequence":null,"usage":{"input_tokens":8,"output_tokens":1}}   }



event: content_block_start

data: {"type":"content_block_start","index":0,"content_block":{"type":"text","text":""}  }



event: ping

data: {"type": "ping"}



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":"Hello"}  }



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":"!"} }



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":" How"}}



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":" can"}     }



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":" I"}    }



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":" assist"}    }



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":" you"}       }



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":" today"}      }



event: content_block_delta

data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":"?"}            }



event: content_block_stop

data: {"type":"content_block_stop","index":0  }



event: message_delta

data: {"type":"message_delta","delta":{"stop_reason":"end_turn","stop_sequence":null},"usage":{"output_tokens":12}   }



event: message_stop

data: {"type":"message_stop"          }`

func TestStreamReader(t *testing.T) {
	reader := NewStreamReader(bufio.NewReader(bytes.NewReader([]byte(rawStreamResponse))))

	result := []string{}

	expectedResult := []string{
		"Hello", "!", " How", " can", " I", " assist", " you", " today", "?",
	}

	for {
		part, err := reader.ReadMessage(false)

		if err == io.EOF {
			break
		}

		require.NoError(t, err)

		result = append(result, part)
	}

	require.Equal(t, expectedResult, result)
}
