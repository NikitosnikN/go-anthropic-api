# go-anthropic-ai

This is a Golang library for interacting with Anthropic's Claude AI API.
It provides a simple and easy-to-use interface for sending requests to the API and receiving responses.

## Installation

To install the library, use the following command:

```
go get github.com/NikitosnikN/go-anthropic-api
```

## Usage

Here are some examples how to use library:

### Basic text request

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

### Streaming response without accumulating response text

```go
package main

import (
	"context"
	"fmt"
	claude "github.com/NikitosnikN/go-anthropic-api"
	"io"
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
	stream, err := client.CreateMessageRequestStream(context.Background(), message)

	if err != nil {
		panic(err)
	}

	// Read response
	for {
		part, err := stream.ReadMessage(false)

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(part)
		// Hello
		// !
		// 	How
		//  can
		//  I
		//  assist
		//  you
		//  today
		// ?
	}
}
```

### Streaming response with response text accumulation

```go
package main

import (
	"context"
	"fmt"
	claude "github.com/NikitosnikN/go-anthropic-api"
	"io"
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
	stream, err := client.CreateMessageRequestStream(context.Background(), message)

	if err != nil {
		panic(err)
	}

	// Read response
	for {
		part, err := stream.ReadMessage(true)

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(part)
		// Hello
		// Hello!
		// Hello! How
		// Hello! How can
		// Hello! How can I
		// Hello! How can I assist
		// Hello! How can I assist you
		// Hello! How can I assist you today
		// Hello! How can I assist you today?
	}
}
```


### Sending image with caption

```go
package main

import (
	"context"
	"fmt"
	claude "github.com/NikitosnikN/go-anthropic-api"
	"io"
	"os"
)

func main() {
	// Create a new Claude client
	client := claude.NewClient("your-api-key")

	// Read image file
	file, err := os.Open("giraffe.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Create a message request instance
	message := claude.MessagesRequest{
		Model:     "claude-3-haiku-20240307",
		MaxTokens: 1024,
	}

	message.AddImageMessage("user", content, "image/jpeg", "who is it?")

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
- [X] Add support for Messages API streaming responses
- [X] Add image support for Messages API
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