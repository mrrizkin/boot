package vite

import (
	goVite "github.com/mrrizkin/go-vite-parser"
)

var viteInstance goVite.ViteManifestInfo

func init() {
	viteInstance = goVite.Parse(goVite.Config{
		OutDir:       "/build/",
		ManifestPath: "public/build/manifest.json",
		HotFilePath:  "public/hot",
	})
}
