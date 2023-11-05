package query

import (
	"context"
)

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, qry Q) (R, error)
}
