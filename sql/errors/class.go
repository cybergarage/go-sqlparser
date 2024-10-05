// Copyright (C) 2024 The go-sqlparser Authors. All rights reserved.
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

package errors

// PostgreSQL: Appendix A. PostgreSQL Error Codes
// https://www.postgresql.org/docs/current/errcodes-appendix.html

// Class represents a SQL error class.
type Class string

const (
	// Class 00 — Successful Completion.
	Class00 Class = "00"
	// Class 01 — Warning.
	Class01 Class = "01"
	// Class 02 — No Data.
	Class02 Class = "02"
	// Class 03 — SQL Statement Not Yet Complete.
	Class03 Class = "03"
	// Class 08 — Connection Exception.
	Class08 Class = "08"
	// Class 09 — Triggered Action Exception.
	Class09 Class = "09"
	// Class 0A — Feature Not Supported.
	Class0A Class = "0A"
	// Class 0B — Invalid Transaction Initiation.
	Class0B Class = "0B"
	// Class 0F — Locator Exception.
	Class0F Class = "0F"
	// Class 0L — Invalid Grantor.
	Class0L Class = "0L"
	// Class 0P — Invalid Role Specification.
	Class0P Class = "0P"
	// Class 0Z — Diagnostics Exception.
	Class0Z Class = "0Z"
	// Class 20 — Case Not Found.
	Class20 Class = "20"
	// Class 21 — Cardinality Violation.
	Class21 Class = "21"
	// Class 22 — Data Exception.
	Class22 Class = "22"
	// Class 23 — Integrity Constraint Violation.
	Class23 Class = "23"
	// Class 24 — Invalid Cursor State.
	Class24 Class = "24"
	// Class 25 — Invalid Transaction State.
	Class25 Class = "25"
	// Class 26 — Invalid SQL Statement Name.
	Class26 Class = "26"
	// Class 27 — Triggered Data Change Violation.
	Class27 Class = "27"
	// Class 28 — Invalid Authorization Specification.
	Class28 Class = "28"
	// Class 2B — Dependent Privilege Descriptors Still Exist.
	Class2B Class = "2B"
	// Class 2D — Invalid Transaction Termination.
	Class2D Class = "2D"
	// Class 2F — SQL Routine Exception.
	Class2F Class = "2F"
	// Class 34 — Invalid Cursor Name.
	Class34 Class = "34"
	// Class 38 — External Routine Exception.
	Class38 Class = "38"
	// Class 39 — External Routine Invocation Exception.
	Class39 Class = "39"
	// Class 3B — Savepoint Exception.
	Class3B Class = "3B"
	// Class 3D — Invalid Catalog Name.
	Class3D Class = "3D"
	// Class 3F — Invalid Schema Name.
	Class3F Class = "3F"
	// Class 40 — Transaction Rollback.
	Class40 Class = "40"
	// Class 42 — Syntax Error or Access Rule Violation.
	Class42 Class = "42"
	// Class 44 — WITH CHECK OPTION Violation.
	Class44 Class = "44"
	// Class 53 — Insufficient Resources.
	Class53 Class = "53"
	// Class 54 — Program Limit Exceeded.
	Class54 Class = "54"
	// Class 55 — Object Not In Prerequisite State.
	Class55 Class = "55"
	// Class 57 — Operator Intervention.
	Class57 Class = "57"
	// Class 58 — System Error.
	Class58 Class = "58"
	// Class F0 — Configuration File Error.
	ClassF0 Class = "F0"
)

// String returns the class as a string
func (c Class) String() string {
	switch c {
	case Class00:
		return "Successful Completion"
	case Class01:
		return "Warning"
	case Class02:
		return "No Data"
	case Class03:
		return "SQL Statement Not Yet Complete"
	case Class08:
		return "Connection Exception"
	case Class09:
		return "Triggered Action Exception"
	case Class0A:
		return "Feature Not Supported"
	case Class0B:
		return "Invalid Transaction Initiation"
	case Class0F:
		return "Locator Exception"
	case Class0L:
		return "Invalid Grantor"
	case Class0P:
		return "Invalid Role Specification"
	case Class0Z:
		return "Diagnostics Exception"
	case Class20:
		return "Case Not Found"
	case Class21:
		return "Cardinality Violation"
	case Class22:
		return "Data Exception"
	case Class23:
		return "Integrity Constraint Violation"
	case Class24:
		return "Invalid Cursor State"
	case Class25:
		return "Invalid Transaction State"
	case Class26:
		return "Invalid SQL Statement Name"
	case Class27:
		return "Triggered Data Change Violation"
	case Class28:
		return "Invalid Authorization Specification"
	case Class2B:
		return "Dependent Privilege Descriptors Still Exist"
	case Class2D:
		return "Invalid Transaction Termination"
	case Class2F:
		return "SQL Routine Exception"
	case Class34:
		return "Invalid Cursor Name"
	case Class38:
		return "External Routine Exception"
	case Class39:
		return "External Routine Invocation Exception"
	case Class3B:
		return "Savepoint Exception"
	case Class3D:
		return "Invalid Catalog Name"
	case Class3F:
		return "Invalid Schema Name"
	case Class40:
		return "Transaction Rollback"
	case Class42:
		return "Syntax Error or Access Rule Violation"
	case Class44:
		return "WITH CHECK OPTION Violation"
	case Class53:
		return "Insufficient Resources"
	case Class54:
		return "Program Limit Exceeded"
	case Class55:
		return "Object Not In Prerequisite State"
	case Class57:
		return "Operator Intervention"
	case Class58:
		return "System Error"
	case ClassF0:
		return "Configuration File Error"
	}
	return ""
}
