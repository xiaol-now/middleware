package middleware

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestChain(t *testing.T) {
	m := Chain(middleware1, middleware2, middleware3)
	handler := m(func(ctx context.Context, req any) (any, error) {
		fmt.Printf("%s\n", req)
		return req, nil
	})

	res, err := handler(context.Background(), "hello")
	if err != nil {
		t.Errorf("expect %v, got %v", nil, err)
	}
	if !reflect.DeepEqual(res, "hello") {
		t.Errorf("expect %v, got %v", "hello", res)
	}

	// Output:
	//    before middleware3
	//	  before middleware2
	//    before middleware1
	//    hello
	//    after middleware1
	//    after middleware2
	//    after middleware3
}

func middleware1(next Handler) Handler {
	return func(ctx context.Context, req any) (any, error) {
		println("before middleware1")
		defer println("after middleware1")
		return next(ctx, req)
	}
}

func middleware2(next Handler) Handler {
	return func(ctx context.Context, req any) (any, error) {
		println("before middleware2")
		defer println("after middleware2")
		return next(ctx, req)
	}
}

func middleware3(next Handler) Handler {
	return func(ctx context.Context, req any) (any, error) {
		println("before middleware3")
		defer println("after middleware3")
		return next(ctx, req)
	}
}
