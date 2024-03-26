package rpc_test

import (
	"fmt"
	"testing"

	"github.com/mirsahebali/golsp/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {

	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	fmt.Println(expected)
	fmt.Println(actual)
	if expected != actual {
		t.Fatalf("Expected:\n %s\nReceived:\n %s", expected, actual)
	}
}
func TestDecode(t *testing.T) {

	test_string := "Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"

	baseMessage, content, err := rpc.DecodeMessage([]byte(test_string))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if 15 != contentLength {
		t.Fatalf("Expected:\n %d\nReceived:\n %d", 16, content)
	}
	if baseMessage != "hi" {
		t.Fatalf("Recieved: %s\n expected: %s", baseMessage, "hi")
	}
}
