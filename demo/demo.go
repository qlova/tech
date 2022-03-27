package main

import (
	"context"
	"fmt"

	"qlova.tech/web/data"
	"qlova.tech/web/page"
	"qlova.tech/web/site"
	"qlova.tech/web/tree"

	"qlova.tech/use/html/button"
	"qlova.tech/use/html/input"
	"qlova.tech/use/html/paragraph"
)

type Home struct {
	Name string
	Text string
}

func (home *Home) Post(ctx context.Context) error {
	fmt.Println(home.Name)
	return nil
}

func (home *Home) RenderTree(seed tree.Seed) tree.Node {
	about := data.Import[About](seed)

	return tree.New(
		paragraph.New(
			data.Echo("Hello %v", &home.Name),
		),

		input.New(
			data.Sync(&home.Name),
		),

		button.New("click me!", data.Post(home),
			data.When(data.Zero(&home.Name),
				button.Disabled,
			),
		),
		button.New("About", data.Get(about), page.Goto(about)),
	)
}

type About struct {
	Authors []string
}

func (about *About) Get(ctx context.Context) error {
	about.Authors = append(about.Authors, "Qlova", "Splizard")
	return nil
}

func (about *About) RenderTree(seed tree.Seed) tree.Node {
	home := data.Import[Home](seed)

	return tree.New(
		data.Feed(&about.Authors,
			paragraph.New(
				data.Echo("Author %v: %v", data.Index, data.Value),
			),
		),

		button.New("Home", page.Goto(home)),
	)
}

type Demo struct{}

func (demo *Demo) RenderTree(seed tree.Seed) tree.Node {
	home := data.Import[Home](seed)

	return tree.New(
		page.New(
			home,
		),
	)
}

func main() {
	site.Open(new(Demo))
}
