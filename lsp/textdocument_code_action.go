package lsp

type TextDocumentCodeActionRequest struct {
	Request
	Params TextDocumentCodeActionParams `json:"params"`
}

type TextDocumentCodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}
type TextDocumentCodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
	Context      CodeActionContext      `json:"codeActionContext"`
}

type CodeActionContext struct{}

type CodeAction struct {
	Title   string         `json:"title"`
	Edit    *WorkspaceEdit `json:"edit,omitempty"`
	Command *Command       `json:"command,omitempty"`
}

type WorkspaceEdit struct {
	Changes map[string][]TextEdit `json:"changes"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

type Command struct {
	Title   string `json:"title"`
	Command string `json:"command"`
}
