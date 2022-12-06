package config

import (
	"fmt"
)

func validateConfig(cfg *Config) error {
	if err := validateGeneralConfig(cfg); err != nil {
		return err
	}

	if err := validateMongoPsqlConfig(cfg); err != nil {
		return err
	}

	if err := validateLoggingConfig(cfg); err != nil {
		return err
	}

	return nil
}

func validateGeneralConfig(cfg *Config) error {
	if cfg.Project.Name == "" {
		return buildError("project name")
	}

	if cfg.Project.Mode == "" {
		return buildError("project mode")
	}

	if cfg.Project.Version == "" {
		return buildError("project version")
	}

	return nil
}

func validateMongoPsqlConfig(cfg *Config) error {
	if cfg.PSQL.URI == "" {
		return buildError("psql uri")
	}

	return nil
}

func validateLoggingConfig(cfg *Config) error {
	if cfg.Logging.Code == "" {
		return buildError("logger type")
	}

	if cfg.Logging.LogLevel == "" {
		return buildError("logger level")
	}

	if cfg.Logging.DateTimeFormat == "" {
		return buildError("logger datetime format")
	}

	return nil
}

func buildError(key string) error {
	return fmt.Errorf("application config. %s is not specified", key)
}
