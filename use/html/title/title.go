/*
	Package head provides the HTML <title> element.

	The <title> HTML element defines the document's title that is shown
	in a browser's title bar or a page's tab. It only contains text.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package title

import "qlova.tech/use/html"

type element string

func (s element) RenderHTML() []byte {
	return []byte(`<title>` + string(s) + `</title>`)
}

// Tag is the html <title> tag.
const Tag = html.Tag("title")

// Set returns an html <title> element with the given title.
func Set(title string) html.Renderer {
	return element(title)
}
