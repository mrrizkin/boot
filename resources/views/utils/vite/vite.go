package vite

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
)

type (
	entryInfo struct {
		File    string   `json:"file"`
		CSS     []string `json:"css"`
		Imports []string `json:"imports"`
	}

	manifestType map[string]entryInfo

	Config struct {
		OutputDirectory string
		PublicPath      string
		ManifestFile    string
		HotFile         string
	}

	Vite struct {
		manifest      manifestType
		viteOrigin    string
		isDevelopment bool
		config        Config
	}

	CodeBlock struct {
		Preload string
		CSS     string
		JS      string
	}
)

func ViteEntries(vite *Vite, entryType string, entries ...string) string {
	css := ""
	js := ""
	preload := ""

	for _, entry := range entries {
		vite, err := vite.Invoke(entry, entryType)
		if err != nil {
			continue
		}

		css += vite.CSS
		js += vite.JS
		preload += vite.Preload
	}

	return css + js + preload
}

func ViteReactRefresh(vite *Vite) string {
	return fmt.Sprintf(`<script type="module">
	import RefreshRuntime from '%s/@react-refresh'
	RefreshRuntime.injectIntoGlobalHook(window)
	window.$RefreshReg$ = () => {}
	window.$RefreshSig$ = () => (type) => type
	window.__vite_plugin_react_preamble_installed__ = true
	</script>`, vite.Origin())
}

func New(config Config) *Vite {
	hotFilePath := path.Clean(path.Join(config.OutputDirectory, config.HotFile))
	_, err := os.Stat(hotFilePath)
	isDevelopment := err == nil

	// create random string
	randomBytes := make([]byte, 16)
	_, _ = rand.Read(randomBytes)

	manifest := make(manifestType)
	if !isDevelopment {
		manifestPath := path.Join(config.OutputDirectory, config.ManifestFile)
		content, err := os.ReadFile(manifestPath)
		if err == nil {
			_ = json.Unmarshal(content, &manifest)
		}
	}

	viteOrigin := ""
	if isDevelopment {
		content, err := os.ReadFile(hotFilePath)
		if err == nil {
			viteOrigin = string(content)
		}
	}

	return &Vite{
		isDevelopment: isDevelopment,
		manifest:      manifest,
		viteOrigin:    viteOrigin,
		config:        config,
	}
}

func (v *Vite) Invoke(input, inputType string) (*CodeBlock, error) {
	if v.isDevelopment {
		return v.serveDevelopmentCodeBlock(input, inputType)
	}

	return v.serveCodeBlock(input, inputType)
}

func (v *Vite) serveCodeBlock(input, inputType string) (*CodeBlock, error) {
	entryPath := path.Clean(input)
	entryInfo, ok := v.manifest[entryPath]
	if !ok {
		return nil, fmt.Errorf("entry not found in %s", v.config.ManifestFile)
	}

	preloads := make([]string, 0)
	preloads = append(preloads, entryPath)
	preloads = append(preloads, entryInfo.Imports...)
	styleTag := ""
	preloadTag := ""
	for _, pre := range preloads {
		_, ok := v.manifest[pre]
		if ok && v.manifest[pre].File != "" {
			preloadTag += createPreloadTag(v.config.PublicPath + v.manifest[pre].File)
		}

		if ok && len(v.manifest[pre].CSS) > 0 {
			for _, cssPath := range v.manifest[pre].CSS {
				if cssPath != "" {
					styleTag += createStyleTag(v.config.PublicPath + cssPath)
				}
			}
		}
	}

	legacyEntryInfo, legacyOk := v.manifest[fmt.Sprintf("%s-legacy", input)]
	legacyPolyfillsInfo, legacyPolyfillsOk := v.manifest["vite/legacy-polyfills-legacy"]
	if !legacyOk || !legacyPolyfillsOk {
		scriptTag := ""
		switch inputType {
		case "script":
			scriptTag = createScriptTag(v.config.PublicPath + entryInfo.File)
		case "style":
			styleTag += createStyleTag(v.config.PublicPath + entryInfo.File)
		case "preload":
			preloadTag += createPreloadTag(v.config.PublicPath + entryInfo.File)
		}

		return &CodeBlock{
			Preload: preloadTag,
			CSS:     styleTag,
			JS:      scriptTag,
		}, nil
	}

	return &CodeBlock{
		Preload: preloadTag,
		CSS:     styleTag,
		JS: fmt.Sprintf(
			`<script type="module">const script=document.createElement("script");try {if(!"noModule" in HTMLScriptElement.prototype) throw "";import.meta.url;import("_").catch(()=>1);(async function*(){})().next();script.type = "module";script.src="%s%s";document.body.appendChild(script);window._isRunManifestJs = true;}catch(error) {}</script><script>window.onload=function(){if(window._isRunManifestJs) return;const script = document.createElement("script");script.src = "%s%s",script.onload=function(){System.import("%s%s");},document.body.appendChild(script);}</script>`,
			v.config.PublicPath,
			entryInfo.File,
			v.config.PublicPath,
			legacyPolyfillsInfo.File,
			v.config.PublicPath,
			legacyEntryInfo.File,
		),
	}, nil
}

func (v *Vite) serveDevelopmentCodeBlock(input string, inputType string) (*CodeBlock, error) {
	viteClientLink, err := v.createViteLink("/@vite/client")
	if err != nil {
		return nil, err
	}
	viteClient := createScriptTag(viteClientLink)

	script := ""
	style := ""
	preload := ""
	switch inputType {
	case "script":
		scriptLink, err := v.createViteLink(input)
		if err != nil {
			return nil, err
		}
		script = viteClient + createScriptTag(scriptLink)
	case "style":
		styleLink, err := v.createViteLink(input)
		if err != nil {
			return nil, err
		}
		style = viteClient + createStyleTag(styleLink)
	case "preload":
		preloadLink, err := v.createViteLink(input)
		if err != nil {
			return nil, err
		}
		preload = viteClient + createPreloadTag(preloadLink)
	}

	codeBlock := CodeBlock{
		JS:      script,
		CSS:     style,
		Preload: preload,
	}

	return &codeBlock, nil
}

func (v *Vite) createViteLink(input string) (string, error) {
	return url.JoinPath(v.viteOrigin, input)
}

func (v *Vite) Origin() string {
	return v.viteOrigin
}

func createPreloadTag(path string) string {
	return fmt.Sprintf(`<link rel="modulepreload" href="%s" />`, path)
}

func createStyleTag(path string) string {
	return fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, path)
}

func createScriptTag(path string) string {
	return fmt.Sprintf(`<script type="module" src="%s"></script>`, path)
}
