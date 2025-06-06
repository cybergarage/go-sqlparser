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

import (
	"testing"
)

func TestNewDataDefFromStrings(t *testing.T) {
	tests := []struct {
		dataTypeStrings  []string
		expectedDataType DataType
		expectedSize     int
	}{
		{[]string{"TEXT"}, TextType, -1},
		{[]string{"VARCHAR", "(", "10", ")"}, VarCharType, 10},
		{[]string{"INT"}, IntType, -1},
		{[]string{"CHAR", "(", "5", ")"}, CharType, 5},
		{[]string{"FLOAT"}, FloatType, -1},
		{[]string{"DECIMAL", "(", "10", ")"}, DecimalType, 10},
		{[]string{"BOOLEAN"}, BooleanType, -1},
		{[]string{"DATE"}, DateType, -1},
	}

	for _, test := range tests {
		dataDef, err := NewDataDefFromStrings(test.dataTypeStrings)
		if err != nil {
			t.Errorf("%v: %v", err, test.dataTypeStrings)
			continue
		}
		if dataDef.DataType() != test.expectedDataType {
			t.Errorf("actual (%v) != expected (%v)", dataDef.DataType(), test.expectedDataType)
		}
		if dataDef.DataTypeSize() != test.expectedSize {
			t.Errorf("actual (%v) != expected (%v)", dataDef.DataTypeSize(), test.expectedSize)
		}
	}
}
