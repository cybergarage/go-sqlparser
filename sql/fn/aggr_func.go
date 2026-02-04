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

import (
	"fmt"
	"strings"
)

// AggregatorOption is a function that configures the Aggregator.
type AggregatorOption = aggrOption

// WithAggregatorName is an option to set the name of the aggregator.
func WithAggregatorName(name string) AggregatorOption {
	return withAggrName(name)
}

// WithAggregatorArguments is an option to set the arguments for the aggregator.
func WithAggregatorArguments(args []string) AggregatorOption {
	return withAggrArguments(args)
}

// WithAggregatorGroupBys is an option to set the group by clause for the aggregator.
func WithAggregatorGroupBys(groups ...GroupBy) AggregatorOption {
	return withAggrGroupBys(groups...)
}

// NewAggregatorForName creates an Aggregator instance for the specified function name.
func NewAggregatorForName(name string, opts ...AggregatorOption) (Aggregator, error) {
	upperName := strings.ToUpper(name)
	switch {
	case strings.HasPrefix(upperName, SumFunctionName):
		return NewSum(opts...)
	case strings.HasPrefix(upperName, AvgFunctionName):
		return NewAvg(opts...)
	case strings.HasPrefix(upperName, MinFunctionName):
		return NewMin(opts...)
	case strings.HasPrefix(upperName, MaxFunctionName):
		return NewMax(opts...)
	case strings.HasPrefix(upperName, CountFunctionName):
		return NewCount(opts...)
	}
	return nil, fmt.Errorf("%w aggregator: %s", ErrNotSupported, name)
}
