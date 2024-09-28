package tag

import (
	"fmt"

	"github.com/nikolalohinski/gonja/v2/exec"
	"github.com/nikolalohinski/gonja/v2/nodes"
	"github.com/nikolalohinski/gonja/v2/parser"
	"github.com/nikolalohinski/gonja/v2/tokens"
)

type OnceControlStructure struct {
	wrapper *nodes.Wrapper
}

func (controlStructure *OnceControlStructure) Position() *tokens.Token {
	return controlStructure.wrapper.Position()
}
func (controlStructure *OnceControlStructure) String() string {
	t := controlStructure.Position()
	return fmt.Sprintf("OnceControlStructure(Line=%d Col=%d)", t.Line, t.Col)
}

func (controlStructure *OnceControlStructure) Execute(
	r *exec.Renderer,
	tag *nodes.ControlStructureBlock,
) error {
	id, ok := r.Environment.Context.Get("gonja-tag-state-id")
	if !ok {
		return nil
	}

	if State.ShouldRender(id.(string), controlStructure.wrapper.String()) {
		sub := r.Inherit()
		err := sub.ExecuteWrapper(controlStructure.wrapper)
		if err != nil {
			return err
		}
	}

	return nil
}

func onceParser(p *parser.Parser, args *parser.Parser) (nodes.ControlStructure, error) {
	wrapper, _, err := p.WrapUntil("endonce")
	if err != nil {
		return nil, err
	}

	if !args.Stream().End() {
		return nil, args.Error("Malformed once controlStructure args.", nil)
	}

	return &OnceControlStructure{
		wrapper: wrapper,
	}, nil
}
