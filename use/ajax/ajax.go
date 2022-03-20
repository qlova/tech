package ajax

import (
	"context"

	"qlova.tech/use/data"
)

type Poster interface {
	Post(ctx context.Context) (data.Patch, error)
}

func Post(p Poster) Poster {
	return p
}
