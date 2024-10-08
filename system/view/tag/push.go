package tag

import (
	"fmt"

	"github.com/nikolalohinski/gonja/v2/exec"
	"github.com/nikolalohinski/gonja/v2/nodes"
	"github.com/nikolalohinski/gonja/v2/parser"
	"github.com/nikolalohinski/gonja/v2/tokens"
)

type PushControlStructure struct {
	name    string
	wrapper *nodes.Wrapper
}

func (controlStructure *PushControlStructure) Position() *tokens.Token {
	return controlStructure.wrapper.Position()
}
func (controlStructure *PushControlStructure) String() string {
	t := controlStructure.Position()
	return fmt.Sprintf("PushControlStructure(Line=%d Col=%d)", t.Line, t.Col)
}

func (controlStructure *PushControlStructure) Execute(
	r *exec.Renderer,
	tag *nodes.ControlStructureBlock,
) error {
	id, ok := r.Environment.Context.Get("gonja-tag-state-id")
	if !ok {
		return nil
	}

	State.Push(id.(string), controlStructure.name, controlStructure.wrapper)
	return nil
}

func pushParser(p *parser.Parser, args *parser.Parser) (nodes.ControlStructure, error) {
	controlStructure := new(PushControlStructure)

	wrapper, _, err := p.WrapUntil("endpush")
	if err != nil {
		return nil, err
	}
	controlStructure.wrapper = wrapper

	modeToken := args.Match(tokens.String)
	if modeToken == nil {
		return nil, args.Error("A stack name is required for push controlStructure.", nil)
	}

	controlStructure.name = modeToken.Val

	if !args.Stream().End() {
		return nil, args.Error("Malformed push controlStructure args.", nil)
	}

	return controlStructure, nil
}
