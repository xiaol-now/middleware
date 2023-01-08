package middleware

import "context"

type Handler func(ctx context.Context, req any) (any, error)

type Middleware func(Handler) Handler

func Chain(m ...Middleware) Middleware {
	return func(next Handler) Handler {
		for _, middleware := range m {
			next = middleware(next)
		}
		return next
	}
}
