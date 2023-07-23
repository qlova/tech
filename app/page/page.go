package page

import "qlova.tech/app/show"

type Renderer interface {
	RenderPage() View
}

type isView = show.Layout

type View struct {
	isView
}

func New(nodes ...show.Node) View {
	return View{
		isView: show.Layout{
			Nodes: nodes,
		},
	}
}
