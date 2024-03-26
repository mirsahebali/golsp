package lsp

type TextDocumentItem struct {
	// Text document's uri
	URI string `json:"uri"`
	//  language ID
	LanguageID string `json:"languageId"`

	Version int `json:"version"`
	// Text Content of the file
	Text string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
	Position     Position         `json:"position"`
}
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
