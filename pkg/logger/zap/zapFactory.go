package zap

import (
	"github.com/abdukhashimov/go_api/pkg/logger"
	"github.com/abdukhashimov/go_api/pkg/logger/options"
)

// Factory is the receiver for zap factory
type Factory struct{}

// Build zap logger
func (_ *Factory) Build(cfg *options.Logging) (logger.Logger, error) {
	l, err := RegisterLog(cfg)
	if err != nil {
		return nil, err
	}

	return l, nil
}
