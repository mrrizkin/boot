package tag

import (
	"fmt"
	"io"

	"github.com/nikolalohinski/gonja/v2/exec"
	"github.com/nikolalohinski/gonja/v2/nodes"
	"github.com/nikolalohinski/gonja/v2/parser"
	"github.com/nikolalohinski/gonja/v2/tokens"
)

type ViteReactRefreshControlStructure struct {
	position *tokens.Token
}

func (controlStructure *ViteReactRefreshControlStructure) Position() *tokens.Token {
	return controlStructure.position
}
func (controlStructure *ViteReactRefreshControlStructure) String() string {
	t := controlStructure.Position()
	return fmt.Sprintf("ViteReactRefreshControlStructure(Line=%d Col=%d)", t.Line, t.Col)
}

func (controlStructure *ViteReactRefreshControlStructure) Execute(
	r *exec.Renderer,
	tag *nodes.ControlStructureBlock,
) error {
	_, err := io.WriteString(r.Output, vite.ReactRefresh())
	return err
}

func viteReactRefreshParser(p *parser.Parser, args *parser.Parser) (nodes.ControlStructure, error) {
	controlStructure := &ViteReactRefreshControlStructure{
		position: p.Current(),
	}

	if !args.Stream().End() {
		return nil, args.Error("Malformed vite react refresh controlStructure args.", nil)
	}

	return controlStructure, nil
}
