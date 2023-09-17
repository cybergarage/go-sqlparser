# Changelog

## v2.0.0 (2023-xx-xx)
- Update
  - SELECT
    - Support subqueries
  - Add new ANTLR parser instead of SQLite3 parser

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
