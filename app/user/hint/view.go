package hint

import "qlova.tech/app/show"

type Layout interface {
	HintLayout(*show.Layout)
}

var Row row

type row struct{}

func (row) HintLayout(view *show.Layout) {
	view.Hints.Row = true
}
