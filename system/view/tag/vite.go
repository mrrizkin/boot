package tag

import (
	"fmt"
	"io"

	"github.com/nikolalohinski/gonja/v2/exec"
	"github.com/nikolalohinski/gonja/v2/nodes"
	"github.com/nikolalohinski/gonja/v2/parser"
	"github.com/nikolalohinski/gonja/v2/tokens"
)

type ViteControlStructure struct {
	entry    string
	position *tokens.Token
}

func (controlStructure *ViteControlStructure) Position() *tokens.Token {
	return controlStructure.position
}
func (controlStructure *ViteControlStructure) String() string {
	t := controlStructure.Position()
	return fmt.Sprintf("ViteControlStructure(Line=%d Col=%d)", t.Line, t.Col)
}

func (controlStructure *ViteControlStructure) Execute(
	r *exec.Renderer,
	tag *nodes.ControlStructureBlock,
) error {
	_, err := io.WriteString(r.Output, vite.Entry(controlStructure.entry))
	return err
}

func viteParser(p *parser.Parser, args *parser.Parser) (nodes.ControlStructure, error) {
	controlStructure := &ViteControlStructure{
		position: p.Current(),
	}

	modeToken := args.Match(tokens.String)
	if modeToken == nil {
		return nil, args.Error("A vite name is required for vite controlStructure.", nil)
	}

	controlStructure.entry = modeToken.Val

	if !args.Stream().End() {
		return nil, args.Error("Malformed vite controlStructure args.", nil)
	}

	return controlStructure, nil
}
