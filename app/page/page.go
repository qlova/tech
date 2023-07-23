package page

import "qlova.tech/app/show"

type Renderer interface {
	RenderPage() View
}

type isView = show.View

type View struct {
	isView
}

func New(nodes ...show.Node) View {
	return View{
		isView: show.View{
			Nodes: nodes,
		},
	}
}
