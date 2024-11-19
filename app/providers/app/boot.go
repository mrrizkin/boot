package app

import (
	"github.com/nikolalohinski/gonja/v2/exec"
	"github.com/nikolalohinski/gonja/v2/parser"

	"github.com/mrrizkin/boot/app/providers/asset"
	"github.com/mrrizkin/boot/app/providers/view"
	"github.com/mrrizkin/boot/config"
)

func Boot(
	asset *asset.Asset,
	view *view.View,
	appCfg *config.App,
) {
	view.AddContext(exec.NewContext(map[string]interface{}{
		"appConfig": appCfg,
	}))

	view.AddControlStructure(exec.NewControlStructureSet(map[string]parser.ControlStructureParser{
		"vite":         viteParser(asset),
		"reactRefresh": viteReactRefreshParser(asset),
	}))

	view.Compile()
}
