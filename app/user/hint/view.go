package hint

import "qlova.tech/app/show"

type View interface {
	HintView(*show.View)
}

var Row row

type row struct{}

func (row) HintView(view *show.View) {
	view.Hints.Row = true
}
