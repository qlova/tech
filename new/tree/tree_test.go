package tree_test

import (
	"fmt"
	"testing"

	"qlova.tech/use/html"
	"qlova.tech/use/html/elem/body"
	"qlova.tech/use/html/elem/head"
	"qlova.tech/use/html/elem/meta"
)

func TestTree(t *testing.T) {
	fmt.Println(html.Render(html.New(
		head.New(
			meta.ApplicationName("My Application"),
		),
		body.New("Hello, World!"),
	)))
}
