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
// limitations under the License..

package stmt

import (
	"strconv"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

type bindParam struct {
	v any
}

// NewBindParam creates a new bind parameter.
func NewBindParam(v any) BindParam {
	return &bindParam{v: v}
}

// String returns the string representation of the bind parameter.
func (p *bindParam) String() (string, error) {
	switch v := p.v.(type) {
	case string:
		return "'" + v + "'", nil
	case time.Time:
		if v.Nanosecond() == 0 {
			return "'" + v.Format("2006-01-02 15:04:05") + "'", nil
		}
		return "'" + v.Format("2006-01-02 15:04:05.000000") + "'", nil
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32), nil
	}

	var to string
	return to, safecast.ToString(p.v, &to)
}
