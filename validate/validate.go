package validate

import (
	"context"
	"middleware"
)

type Validator interface {
	Validate() error
}

func Middleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if v, ok := req.(Validator); ok {
				err := v.Validate()
				if err != nil {
					return nil, err
				}
			}
			return handler(ctx, req)
		}
	}
}
