package tag

import "github.com/nikolalohinski/gonja/v2/parser"

var All = map[string]parser.ControlStructureParser{
	"once":         onceParser,
	"push":         pushParser,
	"stack":        stackParser,
	"vite":         viteParser,
	"reactRefresh": viteReactRefreshParser,
}
