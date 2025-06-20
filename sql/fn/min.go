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
	"math"
)

// Min is an aggregator that calculates the minimum of values.
type Min struct {
	*aggrImpl
}

// MinOption is a function that configures the Min aggregator.
type MinOption = aggrOption

// WithMinArguments sets the arguments for the Min aggregator.
func WithMinArguments(args []string) MinOption {
	return withAggrArguments(args)
}

// WithMinGroupBys sets the group by column for the Min aggregator.
func WithMinGroupBys(groups ...GroupBy) MinOption {
	return withAggrGroupBys(groups...)
}

// NewMin creates a new Min aggregator with the given options.
func NewMin(opts ...MinOption) (*Min, error) {
	aggr := &Min{
		aggrImpl: newAggr(),
	}

	opts = append(opts,
		withAggrName("MIN"),
		withAggrResetFunc(
			func(aggr *aggrImpl) (float64, error) {
				return math.MaxFloat64, nil
			},
		),
		withAggrAggreateFunc(
			func(aggr *aggrImpl, accumulatedValue float64, inputValue float64) (float64, error) {
				if inputValue < accumulatedValue {
					return inputValue, nil
				}
				return accumulatedValue, nil
			},
		),
		withAggrFinalizeFunc(
			func(aggr *aggrImpl, accumulatedValue float64, accumulatedCount int) (any, error) {
				if accumulatedCount == 0 {
					return nil, nil
				}
				return accumulatedValue, nil
			},
		),
	)

	for _, opt := range opts {
		if err := opt(aggr.aggrImpl); err != nil {
			return nil, err
		}
	}

	if err := aggr.Reset(); err != nil {
		return nil, err
	}

	return aggr, nil
}
