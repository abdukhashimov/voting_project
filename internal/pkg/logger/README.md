## Application Logger

This package contains Application Logger
To use App Logger it needs to create any logger and set the logger for the application

Initializing Example:
```go
package main

import (
	"github.com/abdukhashimov/go_api/internal/logger"

	logConfig "github.com/abdukhashimov/go_api/pkg/logger/config"
	logFactory "github.com/abdukhashimov/go_api/pkg/logger/factory"
)

func main() {
	lc := &logConfig.Logging{
		ProjectName:    "HTTP and gRPC Template",
		Code:           logConfig.ZAP,
		LogLevel:       logConfig.DEBUG,
		DateTimeFormat: "2006-01-02 15:04:05",
		DateFormat:     "2006-01-02",
		Encoding:       "json",
		DevMode:        true,
	}

	log, err := logFactory.Build(lc)

	if err != nil {
		panic(err)
	}

	logger.SetLogger(log)
```

Usage Example:
```go
package app

import (
	"github.com/abdukhashimov/go_api/internal/logger"
)

func foo() {
	logger.Log.Debug("debug")
	logger.Log.Info("info")
	logger.Log.Warn("warn")
	logger.Log.Error("error")
}
```
