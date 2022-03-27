package node

import (
	"strings"

	"qlova.tech/use/html"
	"qlova.tech/use/js"
)

func OnClick(script ...js.Renderer) html.Attribute {
	var s strings.Builder
	for _, r := range script {
		s.Write(r.RenderJS())
	}
	return html.Attr("onclick", s.String())
}
