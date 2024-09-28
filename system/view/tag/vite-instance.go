package tag

import (
	assetbundler "github.com/mrrizkin/boot/third-party/asset-bundler"
	goviteparser "github.com/mrrizkin/go-vite-parser"
)

var vite = assetbundler.Vite(&goviteparser.Config{
	OutDir:       "/build/",
	ManifestPath: "public/build/manifest.json",
	HotFilePath:  "public/hot",
})
