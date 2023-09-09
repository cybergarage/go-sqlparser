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

// FunctionExecutor represents a function executor interface.
type FunctionExecutor interface {
	// Name returns the name of the function.
	Name() string
	// Execute returns the executed value with the specified arguments.
	Execute(...any) (any, error)
}

// CastFunction represents a cast function interface.
type CastFunction interface {
	FunctionExecutor
	// Cast returns the casted value.
	Cast(any) (any, error)
}

// Function represents an aggregator function interface.
type AggregatorFunction interface {
	FunctionExecutor
	// Aggregate returns the latest aggregated value.
	Aggregate(string, any) (any, error)
}
