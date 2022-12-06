package processor

import (
	"context"
	"encoding/json"

	"github.com/abdukhashimov/go_api/internal/pkg/logger"
)

type ExecutorWithResp[K any, Z any] func(ctx context.Context, payload K) (Z, error)
type Executor[K any] func(ctx context.Context, payload K) error

func ExecuteWithResp[T any, U any, K any](ctx context.Context, input T, dbPayload K, f ExecutorWithResp[K, U]) (U, error) {
	var (
		resp U
	)

	logger.Log.Debug("Execute with Resp: ", input)

	rawBytes, err := json.Marshal(input)
	if err != nil {
		logger.Log.Error("failed to convert input to json", err)
		return resp, err
	}

	err = json.Unmarshal(rawBytes, &dbPayload)
	if err != nil {
		logger.Log.Error("failed to convert json to struct", err)
		return resp, err
	}

	resp, err = f(ctx, dbPayload)
	if err != nil {
		logger.Log.Error("failed to execute database operation", err)
		return resp, err
	}

	logger.Log.Debug("Request succeeded")

	return resp, nil
}

func ExecuteManyWithResp[T any, U any, Z any](ctx context.Context, input T, dbPayload U, f ExecutorWithResp[U, Z]) (Z, error) {
	var (
		resp Z
		err  error
	)

	logger.Log.Debug("ExecuteMany with Resp: ", input)

	rawBytes, err := json.Marshal(input)
	if err != nil {
		logger.Log.Error("failed to convert input to json", err)
		return resp, err
	}

	err = json.Unmarshal(rawBytes, &dbPayload)
	if err != nil {
		logger.Log.Error("failed to convert json to struct", err)
		return resp, err
	}

	resp, err = f(ctx, dbPayload)
	if err != nil {
		logger.Log.Error("failed to execute database operation", err)
		return resp, err
	}

	logger.Log.Debug("Request succeeded")

	return resp, err
}

func Execute[T any, U any](ctx context.Context, input T, dbPayload U, f Executor[U]) error {
	logger.Log.Debug("Execute with Resp: ", input)

	rawBytes, err := json.Marshal(input)
	if err != nil {
		logger.Log.Error("failed to convert input to json", err)
		return err
	}

	err = json.Unmarshal(rawBytes, &dbPayload)
	if err != nil {
		logger.Log.Error("failed to convert json to struct", err)
		return err
	}

	logger.Log.Debug("Request succeeded")

	return f(ctx, dbPayload)
}
