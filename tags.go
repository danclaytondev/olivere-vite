package vite

import (
	"fmt"
	"html/template"
	"net/url"
)

type Generator struct {
	ViteURL    string
	Entrypoint string
}

func (g Generator) Tags() template.HTML {
	devClientURL, err := url.JoinPath(g.ViteURL, "@vite/client")
	if err != nil {
		panic(err)
	}
	entrypointURL, err := url.JoinPath(g.ViteURL, g.Entrypoint)
	if err != nil {
		panic(err)
	}

	clientTag := fmt.Sprintf(`<script type="module" src="%s"></script>`, devClientURL)
	entrypointTag := fmt.Sprintf(`<script type="module" src="%s"></script>`, entrypointURL)
	return template.HTML(clientTag + "\n" + entrypointTag)
}
