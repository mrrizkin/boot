package tag

import "github.com/nikolalohinski/gonja/v2/parser"

var All = map[string]parser.ControlStructureParser{
	"vite":         viteParser,
	"reactRefresh": viteReactRefreshParser,
}
