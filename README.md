# go-anthropic-ai

This is a Golang library for interacting with Anthropic's Claude AI API.
It provides a simple and easy-to-use interface for sending requests to the API and receiving responses.

## Installation

To install the library, use the following command:

```
go get github.com/NikitosnikN/go-anthropic-api
```

## Usage

Here's a basic example of how to use the library:

```go
package main

import (
	"fmt"
	"context"
	claude "github.com/NikitosnikN/go-anthropic-api"
)

func main() {
	// Create a new Claude client
	client := claude.NewClient("your-api-key")

	// Create a message request instance
	message := claude.MessagesRequest{
		Model:     "claude-3-haiku-20240307",
		MaxTokens: 1024,
	}
	message.AddTextMessage("user", "hello world")

	// Send request 
	response, _ := client.CreateMessageRequest(context.Background(), message)

	// Print the response
	fmt.Println(response)
}
```

## Features

- Simple and intuitive API
- Supports sending requests and receiving responses
- Handles authentication with API keys
- Provides error handling and logging
- Well-documented code and examples

## Roadmap

Here are some planned features and improvements for future releases:

- [X] Add support for Messages API
- [ ] Add support for Messages API streaming responses
- [ ] Implement rate limiting and retry mechanisms
- [ ] Provide more configuration options for the client
- [ ] Improve error handling and logging
- [ ] Add unit tests and integration tests
- [ ] Create a CLI tool for interacting with the API

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a
pull request.

## License

This library is licensed under the [MIT License](LICENCE).

## Contact

If you have any questions or feedback, please contact the maintainer at me@nikitayugov.com