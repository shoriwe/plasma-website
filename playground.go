package main

import "github.com/shoriwe/static/pkg/engine"

func HandlePlayground(e1 *engine.Engine) error {
	return e1.HandlePath("/playground.html",
		func(e *engine.Engine) ([]byte, error) {
			playgroundCssURL, playgroundCssGetError := e.AssetURL("css/playground.css")
			if playgroundCssGetError != nil {
				return nil, playgroundCssGetError
			}
			markdownCssURL, markdownCssGetError := e.AssetURL("vendor/css/github-markdown.css")
			if markdownCssGetError != nil {
				return nil, markdownCssGetError
			}
			prismCssURL, prismCssGetError := e.AssetURL("vendor/prism/prism.css")
			if prismCssGetError != nil {
				return nil, prismCssGetError
			}
			prismScriptURL, prismScriptGetError := e.AssetURL("vendor/prism/prism.js")
			if prismScriptGetError != nil {
				return nil, prismScriptGetError
			}
			prismLiveCssURL, prismLiveCssGetError := e.AssetURL("vendor/prism/prism-live.css")
			if prismLiveCssGetError != nil {
				return nil, prismLiveCssGetError
			}
			prismLiveScriptURL, prismLiveURlGetError := e.AssetURL("vendor/prism/prism-live.js")
			if prismLiveURlGetError != nil {
				return nil, prismLiveURlGetError
			}
			playgroundScriptURL, playgroundScriptGetError := e.ScriptURL("playground")
			if playgroundScriptGetError != nil {
				return nil, playgroundScriptGetError
			}
			playground, playgroundCompileError := e.RenderTemplate("playground.html",
				map[string]string{
					"PLAYGROUND_STYLE":  playgroundCssURL,
					"MARKDOWN_STYLE":    markdownCssURL,
					"PRISM_STYLE":       prismCssURL,
					"PRISM_LIVE_STYLE":  prismLiveCssURL,
					"PRISM_SCRIPT":      prismScriptURL,
					"PRISM_LIVE_SCRIPT": prismLiveScriptURL,
					"PLAYGROUND_SCRIPT": playgroundScriptURL,
				},
			)
			if playgroundCompileError != nil {
				return nil, playgroundCompileError
			}
			navigationCssURL, navigationCssGetError := e.AssetURL("css/page.css")
			if navigationCssGetError != nil {
				return nil, navigationCssGetError
			}
			renderPlayground, renderError := e.RenderTemplate("page.html",
				map[string]string{
					"NAV_STYLE": navigationCssURL,
					"BODY":      string(playground),
				},
			)
			if renderError != nil {
				return nil, renderError
			}
			return e.MinifyHTML(renderPlayground)
		},
	)
}