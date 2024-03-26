package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/mirsahebali/golsp/analysis"
	"github.com/mirsahebali/golsp/lsp"
	"github.com/mirsahebali/golsp/rpc"
)

func main() {
	logger := getLogger("/home/mirsahebali/projects/lsp/log.txt")
	// Prints to the log file
	logger.Println("Lesss Go!!!")
	// I/O scanner which is buffered for better reading in performance
	scanner := bufio.NewScanner(os.Stdin)
	// Split function to pass the scanner on how to split and read from the buffered I/O
	scanner.Split(rpc.Split)
	state := analysis.NewState()
	writer := os.Stdout
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("We got this error: %s", err)
		}

		handleMessage(logger, writer, state, method, contents)
	}

}

// handles the message recieved from encoded rpc and does smth with it
func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Recieved msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("We couldn't parse this at initialize: %s", err)
		}
		logger.Printf("Connected to %s Version: %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		WriteResponse(writer, msg)
		logger.Print("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentParamsNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("ERROR: textDocument/didOpen : %s", err)
			return
		}
		logger.Printf("Opened file at %s:", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			log.Printf("ERROR: textDocument/didChange: %s", err)
			return
		}
		logger.Printf("Changed: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			log.Printf("ERROR: textDocument/hover : %s", err)
			return
		}

		// Create a response
		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		// Write back that response
		WriteResponse(writer, response)

	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			log.Printf("ERROR: textDocument/definition: %s", err)
			return
		}

		// Create a response
		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		// Write back that response
		WriteResponse(writer, response)
	}
}

func WriteResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

// Logger function to put the logs in a 'filename' log file in the specified dir
func getLogger(filename string) *log.Logger {
	// Open the log file to create, truncate and read and write only and could be opened by anyone later
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("Thas not a good file")
	}
	// Logs in our logfile the lsp name i.e, golsp and data and time of the event
	return log.New(logfile, "[golsp]", log.Ldate|log.Ltime)
}
