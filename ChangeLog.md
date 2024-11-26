# Changelog

## v2.0.0 (2023-xx-xx)
- Update
  - SELECT
    - Support subqueries
  - Add new ANTLR parser instead of SQLite3 parser

# v1.4.2 (2024-11-xx)
- Updated
  - ResultSet interface to return error in Row() method

# v1.4.1 (2024-11-24)
- Updated
  - Executor interface to remove ErrorHandler

# v1.4.0 (2024-11-16)
- Added
  - Query executor interface
  - net package to handle network connections

# v1.3.6 (2024-11-08)
- Supported
  - DATETIME

# v1.3.5 (2024-11-05)
- Supported
  - USE statement

# v1.3.4 (2024-10/27)
- Update lookup functions

# v1.3.3 (2024-10-13)
- Add Executor interface

# v1.3.2 (2024-10-12)
- Add Query interface

# v1.3.1 (2024-10-11)
- Add Conn interface

# v1.3.0 (2024-10-06)
- Add standard SQL error classes and codes

# v1.2.7 (2023-12-03)
- Updated Parser::Parse() to return empty query errors

# v1.2.6 (2023-09-29)
- Updated
  - ALTER TABLE
    - Supported ALTER ADD INDEX
- Improved
  - Add utility functions for StetementType

# v1.2.5 (2023-09-28)
- Improved
  - Add utility functions for Schema

# v1.2.4 (2023-09-24)
- Updated
  - SELECT
    - Supported schema name
  - COPY
    - Supported column and format options

# v1.2.3 (2023-09-19)
- Updated
  - UPDATE
    - Supported arithmetic operations
      - ＋, －, ×, / and %

# v1.2.2 (2023-09-18)
- Updated
  - VACUUM
    - Supported ANALYZE clause
  - Delete
    - Add option functions

 # v1.2.1 (2023-09-17)
- Updated
  - SELECT
    - Supported multiple order by

## v1.2.0 (2023-09-16)
- Supported
  - Transaction statements
    - BEGIN, COMMIT and ROLLBACK
  - SELECT
    - GROUP BY
  - Added function executors
    - Math functions 
      - ABS, FLOOR and CEIL
    - Aggregate functions
      - COUNT, MAX, MIN, AVG and SUM
- Updated
  - Selector
    - Added utility functions 

## v1.1.0 (2023-09-09)
- Support
  - ALTER DATABASE
  - ALTER TABLE 
    - ADD, RENAME and DROP COLUMN
    - ADD PRIMARY KEY (PostgreSQL)
- Update
  - SELECT
    - Support functions

## v1.0.1 (2023-09-05)
- Improvement
  -  Parser
    -  Fixed parser to enable parsing of signed numbers
    - Testing
    - Added new test queries using PICT data types

## v1.0.0 (2023-08-27)
- Fixed parser and query interfaces
- Supported
  - SELECT
    - LIMIT, ORDER BY

## v0.9.2 (2023-08-18)
- Supported
  - New statements
    - COMMIT
    - COPY
    - TRUNCATE
    - VACCUM

## v0.9.1 (2023-07-28)
- Supported
  - Bind parameters
- Updated
  - Qquery interfaces

## v0.9.0 (2023-07-04)
- Initial public release  
- Supported statements
  - CREATE DATABASE
  - CREATE TABLE
  - CREATE INDEX
  - DROP DATABASE
  - DROP TABLE
  - INSERT
  - SELECT (without subqueries)
  - UPDATE
  - DELETE
