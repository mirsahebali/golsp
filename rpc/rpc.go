package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Encode the message taken from a struct or interface
func EncodeMessage(message any) string {
	content, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

// Base message
type BaseMessage struct {
	Method string `json:"method"`
}

// Decode message taken from encoder rpc from bytes to method, contents and return err if the format was invalid or not found
func DecodeMessage(msg []byte) (method string, contents []byte, err error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find separator")
	}
	// Content-Length: <num>
	contentLengthBytes := header[len("Content-Length: "):]
	// Check and return if the Content-Length: <num>  is actually a 'num' or 'int' and return error if not
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err = json.Unmarshal(content, &baseMessage); err != nil {
		return "", nil, err
	}
	return baseMessage.Method, content[:contentLength], nil

}

// Split function for bufio.NewScanner.Split() that satisfies the SplitFunc interface
// to tell the scanner to advance a certain int
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	// Check for some amount of bytes until Content-Length \r\n\r\n
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	// Don't return an error as you're waiting for the msg to arrive
	if !found {
		return 0, nil, nil
	}
	// the 'num' part from Content-Length: <num>
	contentLengthBytes := header[len("Content-Length: "):]
	// Check and return if the Content-Length: <num>  is actually a 'num' or 'int' and return error if not
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}
	// Don't return an error yet as we're waiting for more content to arrive
	if len(content) < contentLength {
		return 0, nil, nil
	}

	// totalLength for split scanner to advance
	// +4 for \r\n\r\n character
	totalLength := len(header) + 4 + contentLength

	return totalLength, data[:totalLength], nil
}
