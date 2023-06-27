// Copyright (C) 2022 The go-sqlparser Authors All rights reserved.
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

package query

// IfNotExists represents a IF NOT EXISTS option.
type IfNotExists struct {
	enabled bool
}

// NewIfNotExistsWith returns a new IfNotExists option instance.
func NewIfNotExistsWith(v bool) *IfNotExists {
	return &IfNotExists{
		enabled: v,
	}
}

// IfNotExists returns the IF NOT EXISTS option.
func (opt *IfNotExists) IfNotExists() bool {
	return opt.enabled
}

// Enabled returns true whether the option is enabled.
func (opt *IfNotExists) Enabled() bool {
	return opt.enabled
}

// String returns the string representation.
func (opt *IfNotExists) String() string {
	if opt.enabled {
		return "IF NOT EXISTS"
	}
	return ""
}

// IfExists represents a IF EXISTS option.
type IfExists struct {
	enabled bool
}

// NewIfExistsWith returns a new IfExists option instance.
func NewIfExistsWith(v bool) *IfExists {
	return &IfExists{
		enabled: v,
	}
}

// Enabled returns true whether the option is enabled.
func (opt *IfExists) Enabled() bool {
	return opt.enabled
}

// IfExists returns the IF EXISTS option.
func (opt *IfExists) IfExists() bool {
	return opt.enabled
}

// String returns the string representation.
func (opt *IfExists) String() string {
	if opt.enabled {
		return "IF EXISTS"
	}
	return ""
}
