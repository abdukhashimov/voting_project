package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ExecutorWithResp[K any, Z any] func(ctx context.Context, payload K) (Z, error)
type Executor[K any] func(ctx context.Context, payload K) error
type Delete func(ctx context.Context, key string) error

type QuerierWithSingleResp[K any] func(ctx context.Context, key string) (K, error)

func QueryWithSingleResp[T any](c *gin.Context, key string, f QuerierWithSingleResp[T]) (T, error) {
	ctxWithTimeout, cancel := makeContextWithTimeout()
	defer cancel()

	valueOfKey := c.Param(key)

	return f(ctxWithTimeout, valueOfKey)
}

func QueryWithResp[T any, D any](c *gin.Context, servicePayload T, f ExecutorWithResp[T, D]) (D, error) {
	ctxWithTimeout, cancel := makeContextWithTimeout()
	defer cancel()

	return f(ctxWithTimeout, servicePayload)
}

func ExecuteWithResp[T any, D any](c *gin.Context, servicePayload T, f ExecutorWithResp[T, D]) (D, error) {
	var (
		resp D
		err  error
	)

	ctxWithTimeout, cancel := makeContextWithTimeout()
	defer cancel()

	err = c.ShouldBindJSON(&servicePayload)
	if err != nil {
		return resp, err
	}

	resp, err = f(ctxWithTimeout, servicePayload)

	return resp, err
}

func Execute[T any](c *gin.Context, servicePayload T, f Executor[T]) error {
	ctxWithTimeout, cancel := makeContextWithTimeout()
	defer cancel()

	err := c.ShouldBindJSON(&servicePayload)
	if err != nil {
		return err
	}

	return f(ctxWithTimeout, servicePayload)
}

func ExecuteDelete(c *gin.Context, key string, f Delete) error {
	ctxWithTimeout, cancel := makeContextWithTimeout()
	defer cancel()

	valueOfKey := c.Param(key)

	return f(ctxWithTimeout, valueOfKey)
}
