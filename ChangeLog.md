# Changelog

## v2.0.0 (2023-xx-xx)
### Updates
- SELECT
  - Support subqueries
- Add new ANTLR parser instead of SQLite3 parser

## v1.4.3 (2025-01-XX)
### Fixes
- SQLiteParser to support bind parameters
### Updates
- Added test queries in go-sqltest
 
## v1.4.2 (2024-12-11)
### Updates
- ResultSet Interface
  - Return error in `Row()` method
### Additions
- Utility functions for `CREATE INDEX` and `DROP INDEX` to convert to ALTER TABLE statements

## v1.4.1 (2024-11-24)
### Updates
- Executor Interface
  - Remove `ErrorHandler`

## v1.4.0 (2024-11-16)
### Additions
- Query executor interface
- `net` package to handle network connections

## v1.3.6 (2024-11-08)
### Support
- DATETIME

## v1.3.5 (2024-11-05)
### Support
- USE Statement

## v1.3.4 (2024-10-27)
### Updates
- Lookup functions

## v1.3.3 (2024-10-13)
### Additions
- Executor Interface

## v1.3.2 (2024-10-12)
### Additions
- Query Interface

## v1.3.1 (2024-10-11)
### Additions
- Conn Interface

## v1.3.0 (2024-10-06)
### Additions
- Standard SQL error classes and codes

## v1.2.7 (2023-12-03)
### Updates
- Parser::Parse()
  - Return empty query errors

## v1.2.6 (2023-09-29)
### Updates
- ALTER TABLE
  - Support `ALTER ADD INDEX`
### Improvements
- Utility functions for `StatementType`

## v1.2.5 (2023-09-28)
### Improvements
- Utility functions for `Schema`

## v1.2.4 (2023-09-24)
### Updates
- SELECT
  - Support schema name
- COPY
  - Support column and format options

## v1.2.3 (2023-09-19)
### Updates
- UPDATE
  - Support arithmetic operations (`+`, `-`, `*`, `/`, `%`)

## v1.2.2 (2023-09-18)
### Updates
- VACUUM
  - Support `ANALYZE` clause
- DELETE
  - Add option functions

## v1.2.1 (2023-09-17)
### Updates
- SELECT
  - Support multiple `ORDER BY`

## v1.2.0 (2023-09-16)
### Support
- Transaction Statements
  - `BEGIN`, `COMMIT`, `ROLLBACK`
- SELECT
  - `GROUP BY`
### Additions
- Function Executors
  - Math functions (`ABS`, `FLOOR`, `CEIL`)
  - Aggregate functions (`COUNT`, `MAX`, `MIN`, `AVG`, `SUM`)
### Updates
- Selector
  - Utility functions

## v1.1.0 (2023-09-09)
### Support
- ALTER DATABASE
- ALTER TABLE
  - `ADD`, `RENAME`, `DROP COLUMN`
  - `ADD PRIMARY KEY` (PostgreSQL)
### Updates
- SELECT
  - Support functions

## v1.0.1 (2023-09-05)
### Improvements
- Parser
  - Enable parsing of signed numbers
- Testing
  - Add new test queries using PICT data types

## v1.0.0 (2023-08-27)
### Fixes
- Parser and query interfaces
### Support
- SELECT
  - `LIMIT`, `ORDER BY`

## v0.9.2 (2023-08-18)
### Support
- New statements
  - `COMMIT`, `COPY`, `TRUNCATE`, `VACUUM`

## v0.9.1 (2023-07-28)
### Support
- Bind parameters
### Updates
- Query interfaces

## v0.9.0 (2023-07-04)
### Initial Public Release
- Supported statements
  - `CREATE DATABASE`, `CREATE TABLE`, `CREATE INDEX`
  - `DROP DATABASE`, `DROP TABLE`
  - `INSERT`, `SELECT` (without subqueries)
  - `UPDATE`, `DELETE`