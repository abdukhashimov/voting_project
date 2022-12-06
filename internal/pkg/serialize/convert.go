package serialize

import (
	"context"
	"encoding/json"
)

func MarshalUnMarshal(ctx context.Context, input any, output any) error {
	rawBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawBytes, output)
}
