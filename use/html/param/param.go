/*
	Package param provides the HTML <param> element.

	The <param> HTML element defines parameters for an
	<object> element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package param

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

func Set(name, value string) tree.Node {
	return html.New(
		html.Tag("param"),
		html.Void,

		html.Attr("name", name),
		html.Attr("value", value),
	)
}
