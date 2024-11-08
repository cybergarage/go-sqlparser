// Copyright (C) 2019 The go-sqlparser Authors. All rights reserved.
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

type BindParam interface {
	Name() string
}

// bindParam represents a bind param.
type bindParam struct {
	name string
}

// NewBindParam returns a bind param instance.
func NewBindParamWith(name string) *bindParam {
	param := &bindParam{
		name: name,
	}
	return param
}

// Name returns the bind param name.
func (param *bindParam) Name() string {
	return param.name
}

// String returns the string representation.
func (param *bindParam) String() string {
	return param.name
}

// BindParams represens a bind param array.
type BindParams []BindParam

// NewBindParams returns a bind param array instance.
func NewBindParams() BindParams {
	return make(BindParams, 0)
}
