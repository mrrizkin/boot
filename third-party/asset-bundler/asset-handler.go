package assetbundler

import (
	goviteparser "github.com/mrrizkin/go-vite-parser"
)

func Vite(config *goviteparser.Config) *ViteManifest {
	return newVite(config)
}
