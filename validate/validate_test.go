package validate

import (
	"context"
	"errors"
	"middleware"
	"testing"
)

type testData struct {
	name  string
	count int
	isErr bool
}

func (t *testData) Validate() error {
	if t.count < 0 {
		return errors.New("cannot be less than 0")
	}
	return nil
}

func TestMiddleware(t *testing.T) {
	fn := func(ctx context.Context, req any) (any, error) {
		return req, nil
	}
	tests := []*testData{
		{name: "v1", count: 1, isErr: false},
		{name: "v2", count: 0, isErr: false},
		{name: "v3", count: -1, isErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := middleware.Chain(Middleware())
			_, err := m(fn)(context.Background(), tt)

			if err != nil && !tt.isErr {
				t.Errorf("err: %s", err)
			}
			if err == nil && tt.isErr {
				t.Errorf("unexpected error")
			}
		})
	}
}
