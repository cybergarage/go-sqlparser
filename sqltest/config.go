// Copyright (C) 2025 The go-sqlparser Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqltest

// ConfigOption represents a configuration option.
type ConfigOption func(*Config)

// ValidationMode represents a validation mode.
type ValidationMode int

const (
	// DefaultValidation represents the default validation mode.
	DefaultValidation ValidationMode = iota
	// StrictValidation represents the strict validation mode.
	StrictValidation
	// ParseOnlyValidation represents the parsing validation mode.
	ParseOnlyValidation
)

// Config represents a configuration.
type Config struct {
	skipErrors bool
	mode       ValidationMode
}

// NewConfig returns a new Config instance.
func NewConfig(opts ...ConfigOption) *Config {
	return NewDefaultConfig(opts...)
}

// WithConfigSkipErrors sets the skip errors option.
func WithConfigSkipErrors(skipErrors bool) ConfigOption {
	return func(c *Config) {
		c.skipErrors = skipErrors
	}
}

// WithConfigValidationMode sets the validation mode.
func WithConfigValidationMode(mode ValidationMode) ConfigOption {
	return func(c *Config) {
		c.mode = mode
	}
}

// NewDefaultConfig returns a new Config instance with default options.
func NewDefaultConfig(opts ...ConfigOption) *Config {
	cfg := &Config{
		skipErrors: false,
		mode:       DefaultValidation,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// SkipErrors returns the skip errors option.
func (c *Config) SkipErrors() bool {
	return c.skipErrors
}

// ValidationMode returns the validation mode.
func (c *Config) ValidationMode() ValidationMode {
	return c.mode
}
