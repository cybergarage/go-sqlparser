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

package fn

// Aggregator is an interface for aggregating data.
type Aggregator interface {
	// Name returns the name of the aggregator.
	Name() string
	// Type returns the type of the aggregator.
	Type() FunctionType
	// Argmuents returns the arguments of the aggregator.
	Arguments() []string
	// GroupBys returns the group by column names and whether it is a group by column.
	GroupBys() ([]GroupBy, bool)
	// Reset resets the aggregator to its initial state.
	Reset(opts ...any) error
	// Aggregate aggregates a map or an array. The map represents a row of data, and the array is a
	// list of rows. If grouping is enabled, the array row must have a group value as the first element.
	Aggregate(v any) error
	// Finalize finalizes the aggregation and returns the result.
	Finalize() (ResultSet, error)
}
