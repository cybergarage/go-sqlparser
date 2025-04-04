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

// IfNotExistsOpt represents a IF NOT EXISTS option.
type IfNotExistsOpt struct {
	enabled bool
}

// NewIfNotExistsWith returns a new IfNotExists option instance.
func NewIfNotExistsWith(v bool) *IfNotExistsOpt {
	return &IfNotExistsOpt{
		enabled: v,
	}
}

// IfNotExists returns the IF NOT EXISTS option.
func (opt *IfNotExistsOpt) IfNotExists() bool {
	return opt.enabled
}

// String returns the string representation.
func (opt *IfNotExistsOpt) String() string {
	if opt.enabled {
		return "IF NOT EXISTS"
	}
	return ""
}

// IfExistsOpt represents a IF EXISTS option.
type IfExistsOpt struct {
	enabled bool
}

// NewIfExistsWith returns a new IfExists option instance.
func NewIfExistsWith(v bool) *IfExistsOpt {
	return &IfExistsOpt{
		enabled: v,
	}
}

// IfExists returns the IF EXISTS option.
func (opt *IfExistsOpt) IfExists() bool {
	return opt.enabled
}

// String returns the string representation.
func (opt *IfExistsOpt) String() string {
	if opt.enabled {
		return "IF EXISTS"
	}
	return ""
}
