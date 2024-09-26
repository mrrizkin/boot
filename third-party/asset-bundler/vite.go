package assetbundler

import (
	goviteparser "github.com/mrrizkin/go-vite-parser"
)

type ViteManifest struct {
	manifest *goviteparser.ViteManifestInfo
}

func newVite(config *goviteparser.Config) *ViteManifest {
	manifest := goviteparser.Parse(*config)
	return &ViteManifest{
		manifest: &manifest,
	}
}

func (v *ViteManifest) Entry(entries ...string) string {
	if v.manifest.IsDev() {
		return v.manifest.RenderDevEntriesTag(entries...)
	}

	return v.manifest.RenderEntriesTag(entries...)
}

func (v *ViteManifest) ReactRefresh() string {
	if v.manifest.IsDev() {
		return v.manifest.RenderReactRefreshTag()
	}

	return ""
}
