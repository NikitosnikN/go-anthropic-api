package go_anthropic_api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
)

const (
	contentBlockStart = "content_block_start"
	contentBlockDelta = "content_block_delta"
	contentBlockStop  = "content_block_stop"
)

type streamMessageContentBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type streamMessage struct {
	Type         string                     `json:"type"`
	Index        int                        `json:"index"`
	ContentBlock *streamMessageContentBlock `json:"content_block"`
	Delta        *streamMessageContentBlock `json:"delta"`
}

type StreamReader struct {
	reader      *bufio.Reader
	accumulator string
}

func NewStreamReader(reader io.Reader) *StreamReader {
	return &StreamReader{
		reader:      bufio.NewReader(reader),
		accumulator: "",
	}
}

// ReadMessage reads a single line from the stream.
// Skips any lines that do not start with "data:"
// When it receives message with type content_block_stop, it stops reading the stream.
func (s *StreamReader) ReadMessage(accumulateResponse bool) (string, error) {
	for {
		buff, err := s.reader.ReadBytes('\n')

		if err != nil {
			return "", err
		}

		buff, found := bytes.CutPrefix(buff, []byte("data:"))

		if found {
			var message streamMessage

			err := json.Unmarshal(buff, &message)

			if err != nil {
				return "", err
			}

			if message.Type == contentBlockStop {
				return "", io.EOF
			} else if message.Type == contentBlockDelta && message.Delta != nil {
				if accumulateResponse {
					s.accumulator = s.accumulator + message.Delta.Text
					return s.accumulator, nil
				} else {
					return message.Delta.Text, nil
				}
			}
		}
	}
}
