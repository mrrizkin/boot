package tag

import (
	"fmt"

	"github.com/nikolalohinski/gonja/v2/exec"
	"github.com/nikolalohinski/gonja/v2/nodes"
	"github.com/nikolalohinski/gonja/v2/parser"
	"github.com/nikolalohinski/gonja/v2/tokens"
)

type StackControlStructure struct {
	name     string
	position *tokens.Token
}

func (controlStructure *StackControlStructure) Position() *tokens.Token {
	return controlStructure.position
}
func (controlStructure *StackControlStructure) String() string {
	t := controlStructure.Position()
	return fmt.Sprintf("AutoescapeControlStructure(Line=%d Col=%d)", t.Line, t.Col)
}

func (controlStructure *StackControlStructure) Execute(
	r *exec.Renderer,
	tag *nodes.ControlStructureBlock,
) error {
	requestID, ok := r.Environment.Context.Get("requestID")
	if !ok {
		return nil
	}

	wrappers := StackStore.Get(requestID.(string), controlStructure.name)
	for _, wrapper := range wrappers {
		sub := r.Inherit()
		err := sub.ExecuteWrapper(wrapper)
		if err != nil {
			return err
		}
	}

	return nil
}

func stackParser(p *parser.Parser, args *parser.Parser) (nodes.ControlStructure, error) {
	controlStructure := &StackControlStructure{
		position: p.Current(),
	}

	modeToken := args.Match(tokens.String)
	if modeToken == nil {
		return nil, args.Error("A stack name is required for stack controlStructure.", nil)
	}

	controlStructure.name = modeToken.Val

	if !args.Stream().End() {
		return nil, args.Error("Malformed stack controlStructure args.", nil)
	}

	return controlStructure, nil
}
