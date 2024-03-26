package analysis

import (
	"fmt"
	"strings"

	"github.com/mirsahebali/golsp/lsp"
)

type State struct {
	Document map[string]string
}

func NewState() State {
	return State{
		Document: map[string]string{},
	}
}

func (s *State) OpenDocument(uri, text string) {
	s.Document[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Document[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Document[uri]
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf(
				"File: %s, line: %d, char: %d, characters: %d",
				uri,
				position.Line,
				position.Character,
				len(document)),
		},
	}
}
func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 2,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) TextDocCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	text := s.Document[uri]
	actions := []lsp.CodeAction{}

	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "NeoVim",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Censor the inferior editor",
				Edit: &lsp.WorkspaceEdit{
					Changes: replaceChange,
				},
				Command: &lsp.Command{},
			})
		}
	}

	return lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
