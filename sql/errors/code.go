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

// Code represents a SQL error code.
type Code string

const (
	SuccessfulCompletion                            Code = "00000"
	Warning                                         Code = "01000"
	DynamicResultSetsReturned                       Code = "0100C"
	ImplicitZeroBitPadding                          Code = "01008"
	NullValueEliminatedInSetFunction                Code = "01003"
	PrivilegeNotGranted                             Code = "01007"
	PrivilegeNotRevoked                             Code = "01006"
	StringDataRightTruncationWarning                Code = "01004"
	DeprecatedFeature                               Code = "01P01"
	NoData                                          Code = "02000"
	NoAdditionalDynamicResultSetsReturned           Code = "02001"
	SQLStatementNotYetComplete                      Code = "03000"
	ConnectionException                             Code = "08000"
	ConnectionDoesNotExist                          Code = "08003"
	ConnectionFailure                               Code = "08006"
	SQLClientUnableToEstablishSQLConnection         Code = "08001"
	SQLServerRejectedEstablishmentOfSQLConnection   Code = "08004"
	TransactionResolutionUnknown                    Code = "08007"
	ProtocolViolation                               Code = "08P01"
	TriggeredActionException                        Code = "09000"
	FeatureNotSupported                             Code = "0A000"
	InvalidTransactionInitiation                    Code = "0B000"
	LocatorException                                Code = "0F000"
	InvalidLocatorSpecification                     Code = "0F001"
	InvalidGrantor                                  Code = "0L000"
	InvalidGrantOperation                           Code = "0LP01"
	InvalidRoleSpecification                        Code = "0P000"
	DiagnosticsException                            Code = "0Z000"
	StackedDiagnosticsAccessedWithoutActiveHandler  Code = "0Z002"
	CaseNotFound                                    Code = "20000"
	CardinalityViolation                            Code = "21000"
	DataException                                   Code = "22000"
	ArraySubscriptError                             Code = "2202E"
	CharacterNotInRepertoire                        Code = "22021"
	DatetimeFieldOverflow                           Code = "22008"
	DivisionByZero                                  Code = "22012"
	ErrorInAssignment                               Code = "22005"
	EscapeCharacterConflict                         Code = "2200B"
	IndicatorOverflow                               Code = "22022"
	IntervalFieldOverflow                           Code = "22015"
	InvalidArgumentForLogarithm                     Code = "2201E"
	InvalidArgumentForNtileFunction                 Code = "22014"
	InvalidArgumentForNthValueFunction              Code = "22016"
	InvalidArgumentForPowerFunction                 Code = "2201F"
	InvalidArgumentForWidthBucketFunction           Code = "2201G"
	InvalidCharacterValueForCast                    Code = "22018"
	InvalidDatetimeFormat                           Code = "22007"
	InvalidEscapeCharacter                          Code = "22019"
	InvalidEscapeOctet                              Code = "2200D"
	InvalidEscapeSequence                           Code = "22025"
	NonstandardUseOfEscapeCharacter                 Code = "22P06"
	InvalidIndicatorParameterValue                  Code = "22010"
	InvalidParameterValue                           Code = "22023"
	InvalidPrecedingOrFollowingSize                 Code = "22013"
	InvalidRegularExpression                        Code = "2201B"
	InvalidRowCountInLimitClause                    Code = "2201W"
	InvalidRowCountInResultOffsetClause             Code = "2201X"
	InvalidTablesampleArgument                      Code = "2202H"
	InvalidTablesampleRepeat                        Code = "2202G"
	InvalidTimeZoneDisplacementValue                Code = "22009"
	InvalidUseOfEscapeCharacter                     Code = "2200C"
	MostSpecificTypeMismatch                        Code = "2200G"
	NullDataValueNotAllowed                         Code = "22004"
	NullDataValueNoIndicatorParameter               Code = "22002"
	NumericValueOutOfRange                          Code = "22003"
	SequenceGeneratorLimitExceeded                  Code = "2200H"
	StringDataLengthMismatch                        Code = "22026"
	StringDataRightTruncation                       Code = "22001"
	SubstringError                                  Code = "22011"
	TrimError                                       Code = "22027"
	UnterminatedCString                             Code = "22024"
	ZeroLengthCharacterString                       Code = "2200F"
	FloatingPointException                          Code = "22P01"
	InvalidTextRepresentation                       Code = "22P02"
	InvalidBinaryRepresentation                     Code = "22P03"
	BadCopyFileFormat                               Code = "22P04"
	UntranslatableCharacter                         Code = "22P05"
	NotAnXMLDocument                                Code = "2200L"
	InvalidXMLDocument                              Code = "2200M"
	InvalidXMLContent                               Code = "2200N"
	InvalidXMLComment                               Code = "2200S"
	InvalidXMLProcessingInstruction                 Code = "2200T"
	DuplicateJSONObjectKeyValue                     Code = "22030"
	InvalidArgumentForSQLJSONDatetimeFunction       Code = "22031"
	InvalidJSONText                                 Code = "22032"
	InvalidSQLJSONSubscript                         Code = "22033"
	MoreThanOneSQLJSONItem                          Code = "22034"
	NoSQLJSONItem                                   Code = "22035"
	NonNumericSQLJSONItem                           Code = "22036"
	NonUniqueKeysInAJSONObject                      Code = "22037"
	SingletonSQLJSONItemRequired                    Code = "22038"
	SQLJSONArrayNotFound                            Code = "22039"
	SQLJSONMemberNotFound                           Code = "2203A"
	SQLJSONNumberNotFound                           Code = "2203B"
	SQLJSONObjectNotFound                           Code = "2203C"
	TooManyJSONArrayElements                        Code = "2203D"
	TooManyJSONObjectMembers                        Code = "2203E"
	SQLJSONScalarRequired                           Code = "2203F"
	SQLJSONItemCannotBeCastToTargetType             Code = "2203G"
	IntegrityConstraintViolation                    Code = "23000"
	RestrictViolation                               Code = "23001"
	NotNullViolation                                Code = "23502"
	ForeignKeyViolation                             Code = "23503"
	UniqueViolation                                 Code = "23505"
	CheckViolation                                  Code = "23514"
	ExclusionViolation                              Code = "23P01"
	InvalidCursorState                              Code = "24000"
	InvalidTransactionState                         Code = "25000"
	ActiveSQLTransaction                            Code = "25001"
	BranchTransactionAlreadyActive                  Code = "25002"
	HeldCursorRequiresSameIsolationLevel            Code = "25008"
	InappropriateAccessModeForBranchTransaction     Code = "25003"
	InappropriateIsolationLevelForBranchTransaction Code = "25004"
	NoActiveSQLTransactionForBranchTransaction      Code = "25005"
	ReadOnlySQLTransaction                          Code = "25006"
	SchemaAndDataStatementMixingNotSupported        Code = "25007"
	NoActiveSQLTransaction                          Code = "25P01"
	InFailedSQLTransaction                          Code = "25P02"
	IdleInTransactionSessionTimeout                 Code = "25P03"
	TransactionTimeout                              Code = "25P04"
	InvalidSQLStatementName                         Code = "26000"
	TriggeredDataChangeViolation                    Code = "27000"
	InvalidAuthorizationSpecification               Code = "28000"
	InvalidPassword                                 Code = "28P01"
	DependentPrivilegeDescriptorsStillExist         Code = "2B000"
	DependentObjectsStillExist                      Code = "2BP01"
	InvalidTransactionTermination                   Code = "2D000"
	SQLRoutineException                             Code = "2F000"
	FunctionExecutedNoReturnStatement               Code = "2F005"
	SQLRoutineModifyingSQLDataNotPermitted          Code = "2F002"
	SQLRoutineProhibitedSQLStatementAttempted       Code = "2F003"
	SQLRoutineReadingSQLDataNotPermitted            Code = "2F004"
	InvalidCursorName                               Code = "34000"
	ExternalRoutineException                        Code = "38000"
	ContainingSQLNotPermitted                       Code = "38001"
	ExternalRoutineModifyingSQLDataNotPermitted     Code = "38002"
	ExternalRoutineProhibitedSQLStatementAttempted  Code = "38003"
	ExternalRoutineReadingSQLDataNotPermitted       Code = "38004"
	ExternalRoutineInvocationException              Code = "39000"
	InvalidSQLStateReturned                         Code = "39001"
	NullValueNotAllowed                             Code = "39004"
	TriggerProtocolViolated                         Code = "39P01"
	SRFProtocolViolated                             Code = "39P02"
	EventTriggerProtocolViolated                    Code = "39P03"
	SavepointException                              Code = "3B000"
	InvalidSavepointSpecification                   Code = "3B001"
	InvalidCatalogName                              Code = "3D000"
	InvalidSchemaName                               Code = "3F000"
	TransactionRollback                             Code = "40000"
	TransactionIntegrityConstraintViolation         Code = "40002"
	SerializationFailure                            Code = "40001"
	StatementCompletionUnknown                      Code = "40003"
	DeadlockDetected                                Code = "40P01"
	SyntaxErrorOrAccessRuleViolation                Code = "42000"
	SyntaxError                                     Code = "42601"
	InsufficientPrivilege                           Code = "42501"
	CannotCoerce                                    Code = "42846"
	GroupingError                                   Code = "42803"
	WindowingError                                  Code = "42P20"
	InvalidRecursion                                Code = "42P19"
	InvalidForeignKey                               Code = "42830"
	InvalidName                                     Code = "42602"
	NameTooLong                                     Code = "42622"
	ReservedName                                    Code = "42939"
	DatatypeMismatch                                Code = "42804"
	IndeterminateDatatype                           Code = "42P18"
	CollationMismatch                               Code = "42P21"
	IndeterminateCollation                          Code = "42P22"
	WrongObjectType                                 Code = "42809"
	GeneratedAlways                                 Code = "428C9"
	UndefinedColumn                                 Code = "42703"
	UndefinedFunction                               Code = "42883"
	UndefinedTable                                  Code = "42P01"
	UndefinedParameter                              Code = "42P02"
	UndefinedObject                                 Code = "42704"
	DuplicateColumn                                 Code = "42701"
	DuplicateCursor                                 Code = "42P03"
	DuplicateDatabase                               Code = "42P04"
	DuplicateFunction                               Code = "42723"
	DuplicatePreparedStatement                      Code = "42P05"
	DuplicateSchema                                 Code = "42P06"
	DuplicateTable                                  Code = "42P07"
	DuplicateAlias                                  Code = "42712"
	DuplicateObject                                 Code = "42710"
	AmbiguousColumn                                 Code = "42702"
	AmbiguousFunction                               Code = "42725"
	AmbiguousParameter                              Code = "42P08"
	AmbiguousAlias                                  Code = "42P09"
	InvalidColumnReference                          Code = "42P10"
	InvalidColumnDefinition                         Code = "42611"
	InvalidCursorDefinition                         Code = "42P11"
	InvalidDatabaseDefinition                       Code = "42P12"
	InvalidFunctionDefinition                       Code = "42P13"
	InvalidPreparedStatementDefinition              Code = "42P14"
	InvalidSchemaDefinition                         Code = "42P15"
	InvalidTableDefinition                          Code = "42P16"
	InvalidObjectDefinition                         Code = "42P17"
	WithCheckOptionViolation                        Code = "44000"
	InsufficientResources                           Code = "53000"
	DiskFull                                        Code = "53100"
	OutOfMemory                                     Code = "53200"
	TooManyConnections                              Code = "53300"
	ConfigurationLimitExceeded                      Code = "53400"
	ProgramLimitExceeded                            Code = "54000"
	StatementTooComplex                             Code = "54001"
	TooManyColumns                                  Code = "54011"
	TooManyArguments                                Code = "54023"
	ObjectNotInPrerequisiteState                    Code = "55000"
	ObjectInUse                                     Code = "55006"
	CantChangeRuntimeParam                          Code = "55P02"
	LockNotAvailable                                Code = "55P03"
	UnsafeNewEnumValueUsage                         Code = "55P04"
	OperatorIntervention                            Code = "57000"
	QueryCanceled                                   Code = "57014"
	AdminShutdown                                   Code = "57P01"
	CrashShutdown                                   Code = "57P02"
	CannotConnectNow                                Code = "57P03"
	DatabaseDropped                                 Code = "57P04"
	IdleSessionTimeout                              Code = "57P05"
	SystemError                                     Code = "58000"
	IOError                                         Code = "58030"
	UndefinedFile                                   Code = "58P01"
	DuplicateFile                                   Code = "58P02"
	ConfigFileError                                 Code = "F0000"
	LockFileExists                                  Code = "F0001"
	FDWError                                        Code = "HV000"
	FDWColumnNameNotFound                           Code = "HV005"
	FDWDynamicParameterValueNeeded                  Code = "HV002"
	FDWFunctionSequenceError                        Code = "HV010"
	FDWInconsistentDescriptorInformation            Code = "HV021"
	FDWInvalidAttributeValue                        Code = "HV024"
	FDWInvalidColumnName                            Code = "HV007"
	FDWInvalidColumnNumber                          Code = "HV008"
	FDWInvalidDataType                              Code = "HV004"
	FDWInvalidDataTypeDescriptors                   Code = "HV006"
	FDWInvalidDescriptorFieldIdentifier             Code = "HV091"
	FDWInvalidHandle                                Code = "HV00B"
	FDWInvalidOptionIndex                           Code = "HV00C"
	FDWInvalidOptionName                            Code = "HV00D"
	FDWInvalidStringLength                          Code = "HV090"
	FDWInvalidStringFormat                          Code = "HV00A"
	FDWInvalidUseOfNullPointer                      Code = "HV009"
	FDWTooManyHandles                               Code = "HV014"
	FDWOutOfMemory                                  Code = "HV001"
	FDWNoSchemas                                    Code = "HV00P"
	FDWOptionNameNotFound                           Code = "HV00J"
	FDWReplyHandle                                  Code = "HV00K"
	FDWSchemaNotFound                               Code = "HV00Q"
	FDWTableNotFound                                Code = "HV00R"
	FDWUnableToCreateExecution                      Code = "HV00L"
	FDWUnableToCreateReply                          Code = "HV00M"
	FDWUnableToEstablishConnection                  Code = "HV00N"
	PLPGSQLError                                    Code = "P0000"
	RaiseException                                  Code = "P0001"
	NoDataFound                                     Code = "P0002"
	TooManyRows                                     Code = "P0003"
	AssertFailure                                   Code = "P0004"
	InternalError                                   Code = "XX000"
	DataCorrupted                                   Code = "XX001"
	IndexCorrupted                                  Code = "XX002"
)

// String returns the string representation of the code.
func (c Code) String() string { // nolint: maintidx
	switch c {
	case SuccessfulCompletion:
		return "Successful Completion"
	case Warning:
		return "Warning"
	case DynamicResultSetsReturned:
		return "Dynamic Result Sets Returned"
	case ImplicitZeroBitPadding:
		return "Implicit Zero Bit Padding"
	case NullValueEliminatedInSetFunction:
		return "Null Value Eliminated In Set Function"
	case PrivilegeNotGranted:
		return "Privilege Not Granted"
	case PrivilegeNotRevoked:
		return "Privilege Not Revoked"
	case StringDataRightTruncationWarning:
		return "String Data Right Truncation Warning"
	case DeprecatedFeature:
		return "Deprecated Feature"
	case NoData:
		return "No Data"
	case NoAdditionalDynamicResultSetsReturned:
		return "No Additional Dynamic Result Sets Returned"
	case SQLStatementNotYetComplete:
		return "SQL Statement Not Yet Complete"
	case ConnectionException:
		return "Connection Exception"
	case ConnectionDoesNotExist:
		return "Connection Does Not Exist"
	case ConnectionFailure:
		return "Connection Failure"
	case SQLClientUnableToEstablishSQLConnection:
		return "SQL Client Unable To Establish SQL Connection"
	case SQLServerRejectedEstablishmentOfSQLConnection:
		return "SQL Server Rejected Establishment Of SQL Connection"
	case TransactionResolutionUnknown:
		return "Transaction Resolution Unknown"
	case ProtocolViolation:
		return "Protocol Violation"
	case TriggeredActionException:
		return "Triggered Action Exception"
	case FeatureNotSupported:
		return "Feature Not Supported"
	case InvalidTransactionInitiation:
		return "Invalid Transaction Initiation"
	case LocatorException:
		return "Locator Exception"
	case InvalidLocatorSpecification:
		return "Invalid Locator Specification"
	case InvalidGrantor:
		return "Invalid Grantor"
	case InvalidGrantOperation:
		return "Invalid Grant Operation"
	case InvalidRoleSpecification:
		return "Invalid Role Specification"
	case DiagnosticsException:
		return "Diagnostics Exception"
	case StackedDiagnosticsAccessedWithoutActiveHandler:
		return "Stacked Diagnostics Accessed Without Active Handler"
	case CaseNotFound:
		return "Case Not Found"
	case CardinalityViolation:
		return "Cardinality Violation"
	case DataException:
		return "Data Exception"
	case ArraySubscriptError:
		return "Array Subscript Error"
	case CharacterNotInRepertoire:
		return "Character Not In Repertoire"
	case DatetimeFieldOverflow:
		return "Datetime Field Overflow"
	case DivisionByZero:
		return "Division By Zero"
	case ErrorInAssignment:
		return "Error In Assignment"
	case EscapeCharacterConflict:
		return "Escape Character Conflict"
	case IndicatorOverflow:
		return "Indicator Overflow"
	case IntervalFieldOverflow:
		return "Interval Field Overflow"
	case InvalidArgumentForLogarithm:
		return "Invalid Argument For Logarithm"
	case InvalidArgumentForNtileFunction:
		return "Invalid Argument For Ntile Function"
	case InvalidArgumentForNthValueFunction:
		return "Invalid Argument For Nth Value Function"
	case InvalidArgumentForPowerFunction:
		return "Invalid Argument For Power Function"
	case InvalidArgumentForWidthBucketFunction:
		return "Invalid Argument For Width Bucket Function"
	case InvalidCharacterValueForCast:
		return "Invalid Character Value For Cast"
	case InvalidDatetimeFormat:
		return "Invalid Datetime Format"
	case InvalidEscapeCharacter:
		return "Invalid Escape Character"
	case InvalidEscapeOctet:
		return "Invalid Escape Octet"
	case InvalidEscapeSequence:
		return "Invalid Escape Sequence"
	case NonstandardUseOfEscapeCharacter:
		return "Nonstandard Use Of Escape Character"
	case InvalidIndicatorParameterValue:
		return "Invalid Indicator Parameter Value"
	case InvalidParameterValue:
		return "Invalid Parameter Value"
	case InvalidPrecedingOrFollowingSize:
		return "Invalid Preceding Or Following Size"
	case InvalidRegularExpression:
		return "Invalid Regular Expression"
	case InvalidRowCountInLimitClause:
		return "Invalid Row Count In Limit Clause"
	case InvalidRowCountInResultOffsetClause:
		return "Invalid Row Count In Result Offset Clause"
	case InvalidTablesampleArgument:
		return "Invalid Tablesample Argument"
	case InvalidTablesampleRepeat:
		return "Invalid Tablesample Repeat"
	case InvalidTimeZoneDisplacementValue:
		return "Invalid Time Zone Displacement Value"
	case InvalidUseOfEscapeCharacter:
		return "Invalid Use Of Escape Character"
	case MostSpecificTypeMismatch:
		return "Most Specific Type Mismatch"
	case NullDataValueNotAllowed:
		return "Null Data Value Not Allowed"
	case NullDataValueNoIndicatorParameter:
		return "Null Data Value No Indicator Parameter"
	case NumericValueOutOfRange:
		return "Numeric Value Out Of Range"
	case SequenceGeneratorLimitExceeded:
		return "Sequence Generator Limit Exceeded"
	case StringDataLengthMismatch:
		return "String Data Length Mismatch"
	case StringDataRightTruncation:
		return "String Data Right Truncation"
	case SubstringError:
		return "Substring Error"
	case TrimError:
		return "Trim Error"
	case UnterminatedCString:
		return "Unterminated C String"
	case ZeroLengthCharacterString:
		return "Zero Length Character String"
	case FloatingPointException:
		return "Floating Point Exception"
	case InvalidTextRepresentation:
		return "Invalid Text Representation"
	case InvalidBinaryRepresentation:
		return "Invalid Binary Representation"
	case BadCopyFileFormat:
		return "Bad Copy File Format"
	case UntranslatableCharacter:
		return "Untranslatable Character"
	case NotAnXMLDocument:
		return "Not An XML Document"
	case InvalidXMLDocument:
		return "Invalid XML Document"
	case InvalidXMLContent:
		return "Invalid XML Content"
	case InvalidXMLComment:
		return "Invalid XML Comment"
	case InvalidXMLProcessingInstruction:
		return "Invalid XML Processing Instruction"
	case DuplicateJSONObjectKeyValue:
		return "Duplicate JSON Object Key Value"
	case InvalidArgumentForSQLJSONDatetimeFunction:
		return "Invalid Argument For SQL/JSON Datetime Function"
	case InvalidJSONText:
		return "Invalid JSON Text"
	case InvalidSQLJSONSubscript:
		return "Invalid SQL/JSON Subscript"
	case MoreThanOneSQLJSONItem:
		return "More Than One SQL/JSON Item"
	case NoSQLJSONItem:
		return "No SQL/JSON Item"
	case NonNumericSQLJSONItem:
		return "Non Numeric SQL/JSON Item"
	case NonUniqueKeysInAJSONObject:
		return "Non Unique Keys In A JSON Object"
	case SingletonSQLJSONItemRequired:
		return "Singleton SQL/JSON Item Required"
	case SQLJSONArrayNotFound:
		return "SQL/JSON Array Not Found"
	case SQLJSONMemberNotFound:
		return "SQL/JSON Member Not Found"
	case SQLJSONNumberNotFound:
		return "SQL/JSON Number Not Found"
	case SQLJSONObjectNotFound:
		return "SQL/JSON Object Not Found"
	case TooManyJSONArrayElements:
		return "Too Many JSON Array Elements"
	case TooManyJSONObjectMembers:
		return "Too Many JSON Object Members"
	case SQLJSONScalarRequired:
		return "SQL/JSON Scalar Required"
	case SQLJSONItemCannotBeCastToTargetType:
		return "SQL/JSON Item Cannot Be Cast To Target Type"
	case IntegrityConstraintViolation:
		return "Integrity Constraint Violation"
	case RestrictViolation:
		return "Restrict Violation"
	case NotNullViolation:
		return "Not Null Violation"
	case ForeignKeyViolation:
		return "Foreign Key Violation"
	case UniqueViolation:
		return "Unique Violation"
	case CheckViolation:
		return "Check Violation"
	case ExclusionViolation:
		return "Exclusion Violation"
	case InvalidCursorState:
		return "Invalid Cursor State"
	case InvalidTransactionState:
		return "Invalid Transaction State"
	case ActiveSQLTransaction:
		return "Active SQL Transaction"
	case BranchTransactionAlreadyActive:
		return "Branch Transaction Already Active"
	case HeldCursorRequiresSameIsolationLevel:
		return "Held Cursor Requires Same Isolation Level"
	case InappropriateAccessModeForBranchTransaction:
		return "Inappropriate Access Mode For Branch Transaction"
	case InappropriateIsolationLevelForBranchTransaction:
		return "Inappropriate Isolation Level For Branch Transaction"
	case NoActiveSQLTransactionForBranchTransaction:
		return "No Active SQL Transaction For Branch Transaction"
	case ReadOnlySQLTransaction:
		return "Read Only SQL Transaction"
	case SchemaAndDataStatementMixingNotSupported:
		return "Schema And Data Statement Mixing Not Supported"
	case NoActiveSQLTransaction:
		return "No Active SQL Transaction"
	case InFailedSQLTransaction:
		return "In Failed SQL Transaction"
	case IdleInTransactionSessionTimeout:
		return "Idle In Transaction Session Timeout"
	case TransactionTimeout:
		return "Transaction Timeout"
	case InvalidSQLStatementName:
		return "Invalid SQL Statement Name"
	case TriggeredDataChangeViolation:
		return "Triggered Data Change Violation"
	case InvalidAuthorizationSpecification:
		return "Invalid Authorization Specification"
	case InvalidPassword:
		return "Invalid Password"
	case DependentPrivilegeDescriptorsStillExist:
		return "Dependent Privilege Descriptors Still Exist"
	case DependentObjectsStillExist:
		return "Dependent Objects Still Exist"
	case InvalidTransactionTermination:
		return "Invalid Transaction Termination"
	case SQLRoutineException:
		return "SQL Routine Exception"
	case FunctionExecutedNoReturnStatement:
		return "Function Executed No Return Statement"
	case SQLRoutineModifyingSQLDataNotPermitted:
		return "SQL Routine Modifying SQL Data Not Permitted"
	case SQLRoutineProhibitedSQLStatementAttempted:
		return "SQL Routine Prohibited SQL Statement Attempted"
	case SQLRoutineReadingSQLDataNotPermitted:
		return "SQL Routine Reading SQL Data Not Permitted"
	case InvalidCursorName:
		return "Invalid Cursor Name"
	case ExternalRoutineException:
		return "External Routine Exception"
	case ContainingSQLNotPermitted:
		return "Containing SQL Not Permitted"
	case ExternalRoutineModifyingSQLDataNotPermitted:
		return "External Routine Modifying SQL Data Not Permitted"
	case ExternalRoutineProhibitedSQLStatementAttempted:
		return "External Routine Prohibited SQL Statement Attempted"
	case ExternalRoutineReadingSQLDataNotPermitted:
		return "External Routine Reading SQL Data Not Permitted"
	case ExternalRoutineInvocationException:
		return "External Routine Invocation Exception"
	case InvalidSQLStateReturned:
		return "Invalid SQL State Returned"
	case NullValueNotAllowed:
		return "Null Value Not Allowed"
	case TriggerProtocolViolated:
		return "Trigger Protocol Violated"
	case SRFProtocolViolated:
		return "SRF Protocol Violated"
	case EventTriggerProtocolViolated:
		return "Event Trigger Protocol Violated"
	case SavepointException:
		return "Savepoint Exception"
	case InvalidSavepointSpecification:
		return "Invalid Savepoint Specification"
	case InvalidCatalogName:
		return "Invalid Catalog Name"
	case InvalidSchemaName:
		return "Invalid Schema Name"
	case TransactionRollback:
		return "Transaction Rollback"
	case TransactionIntegrityConstraintViolation:
		return "Transaction Integrity Constraint Violation"
	case SerializationFailure:
		return "Serialization Failure"
	case StatementCompletionUnknown:
		return "Statement Completion Unknown"
	case DeadlockDetected:
		return "Deadlock Detected"
	case SyntaxErrorOrAccessRuleViolation:
		return "Syntax Error Or Access Rule Violation"
	case SyntaxError:
		return "Syntax Error"
	case InsufficientPrivilege:
		return "Insufficient Privilege"
	case CannotCoerce:
		return "Cannot Coerce"
	case GroupingError:
		return "Grouping Error"
	case WindowingError:
		return "Windowing Error"
	case InvalidRecursion:
		return "Invalid Recursion"
	case InvalidForeignKey:
		return "Invalid Foreign Key"
	case InvalidName:
		return "Invalid Name"
	case NameTooLong:
		return "Name Too Long"
	case ReservedName:
		return "Reserved Name"
	case DatatypeMismatch:
		return "Datatype Mismatch"
	case IndeterminateDatatype:
		return "Indeterminate Datatype"
	case CollationMismatch:
		return "Collation Mismatch"
	case IndeterminateCollation:
		return "Indeterminate Collation"
	case WrongObjectType:
		return "Wrong Object Type"
	case GeneratedAlways:
		return "Generated Always"
	case UndefinedColumn:
		return "Undefined Column"
	case UndefinedFunction:
		return "Undefined Function"
	case UndefinedTable:
		return "Undefined Table"
	case UndefinedParameter:
		return "Undefined Parameter"
	case UndefinedObject:
		return "Undefined Object"
	case DuplicateColumn:
		return "Duplicate Column"
	case DuplicateCursor:
		return "Duplicate Cursor"
	case DuplicateDatabase:
		return "Duplicate Database"
	case DuplicateFunction:
		return "Duplicate Function"
	case DuplicatePreparedStatement:
		return "Duplicate Prepared Statement"
	case DuplicateSchema:
		return "Duplicate Schema"
	case DuplicateTable:
		return "Duplicate Table"
	case DuplicateAlias:
		return "Duplicate Alias"
	case DuplicateObject:
		return "Duplicate Object"
	case AmbiguousColumn:
		return "Ambiguous Column"
	case AmbiguousFunction:
		return "Ambiguous Function"
	case AmbiguousParameter:
		return "Ambiguous Parameter"
	case AmbiguousAlias:
		return "Ambiguous Alias"
	case InvalidColumnReference:
		return "Invalid Column Reference"
	case InvalidColumnDefinition:
		return "Invalid Column Definition"
	case InvalidCursorDefinition:
		return "Invalid Cursor Definition"
	case InvalidDatabaseDefinition:
		return "Invalid Database Definition"
	case InvalidFunctionDefinition:
		return "Invalid Function Definition"
	case InvalidPreparedStatementDefinition:
		return "Invalid Prepared Statement Definition"
	case InvalidSchemaDefinition:
		return "Invalid Schema Definition"
	case InvalidTableDefinition:
		return "Invalid Table Definition"
	case InvalidObjectDefinition:
		return "Invalid Object Definition"
	case WithCheckOptionViolation:
		return "With Check Option Violation"
	case InsufficientResources:
		return "Insufficient Resources"
	case DiskFull:
		return "Disk Full"
	case OutOfMemory:
		return "Out Of Memory"
	case TooManyConnections:
		return "Too Many Connections"
	case ConfigurationLimitExceeded:
		return "Configuration Limit Exceeded"
	case ProgramLimitExceeded:
		return "Program Limit Exceeded"
	case StatementTooComplex:
		return "Statement Too Complex"
	case TooManyColumns:
		return "Too Many Columns"
	case TooManyArguments:
		return "Too Many Arguments"
	case ObjectNotInPrerequisiteState:
		return "Object Not In Prerequisite State"
	case ObjectInUse:
		return "Object In Use"
	case CantChangeRuntimeParam:
		return "Can't Change Runtime Param"
	case LockNotAvailable:
		return "Lock Not Available"
	case UnsafeNewEnumValueUsage:
		return "Unsafe New Enum Value Usage"
	case OperatorIntervention:
		return "Operator Intervention"
	case QueryCanceled:
		return "Query Canceled"
	case AdminShutdown:
		return "Admin Shutdown"
	case CrashShutdown:
		return "Crash Shutdown"
	case CannotConnectNow:
		return "Cannot Connect Now"
	case DatabaseDropped:
		return "Database Dropped"
	case IdleSessionTimeout:
		return "Idle Session Timeout"
	case SystemError:
		return "System Error"
	case IOError:
		return "IO Error"
	case UndefinedFile:
		return "Undefined File"
	case DuplicateFile:
		return "Duplicate File"
	case ConfigFileError:
		return "Config File Error"
	case LockFileExists:
		return "Lock File Exists"
	case FDWError:
		return "FDW Error"
	case FDWColumnNameNotFound:
		return "FDW Column Name Not Found"
	case FDWDynamicParameterValueNeeded:
		return "FDW Dynamic Parameter Value Needed"
	case FDWFunctionSequenceError:
		return "FDW Function Sequence Error"
	case FDWInconsistentDescriptorInformation:
		return "FDW Inconsistent Descriptor Information"
	case FDWInvalidAttributeValue:
		return "FDW Invalid Attribute Value"
	case FDWInvalidColumnName:
		return "FDW Invalid Column Name"
	case FDWInvalidColumnNumber:
		return "FDW Invalid Column Number"
	case FDWInvalidDataType:
		return "FDW Invalid Data Type"
	case FDWInvalidDataTypeDescriptors:
		return "FDW Invalid Data Type Descriptors"
	case FDWInvalidDescriptorFieldIdentifier:
		return "FDW Invalid Descriptor Field Identifier"
	case FDWInvalidHandle:
		return "FDW Invalid Handle"
	case FDWInvalidOptionIndex:
		return "FDW Invalid Option Index"
	case FDWInvalidOptionName:
		return "FDW Invalid Option Name"
	case FDWInvalidStringLength:
		return "FDW Invalid String Length"
	case FDWInvalidStringFormat:
		return "FDW Invalid String Format"
	case FDWInvalidUseOfNullPointer:
		return "FDW Invalid Use Of Null Pointer"
	case FDWTooManyHandles:
		return "FDW Too Many Handles"
	case FDWOutOfMemory:
		return "FDW Out Of Memory"
	case FDWNoSchemas:
		return "FDW No Schemas"
	case FDWOptionNameNotFound:
		return "FDW Option Name Not Found"
	case FDWReplyHandle:
		return "FDW Reply Handle"
	case FDWSchemaNotFound:
		return "FDW Schema Not Found"
	case FDWTableNotFound:
		return "FDW Table Not Found"
	case FDWUnableToCreateExecution:
		return "FDW Unable To Create Execution"
	case FDWUnableToCreateReply:
		return "FDW Unable To Create Reply"
	case FDWUnableToEstablishConnection:
		return "FDW Unable To Establish Connection"
	case PLPGSQLError:
		return "PL/PGSQL Error"
	case RaiseException:
		return "Raise Exception"
	case NoDataFound:
		return "No Data Found"
	case TooManyRows:
		return "Too Many Rows"
	case AssertFailure:
		return "Assert Failure"
	case InternalError:
		return "Internal Error"
	case DataCorrupted:
		return "Data Corrupted"
	case IndexCorrupted:
		return "Index Corrupted"
	default:
		return "Unknown Code"
	}
}
