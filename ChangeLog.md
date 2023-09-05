# Changelog

## v2.0.0 (2023-xx-xx)
- Improvements
  - SELECT
    - Support subqueries
  - Add new ANTLR parser instead of SQLite3 parser

## v1.1.0 (2023-xx-xx)
- Support new statements
  - ALTER DATABASE
  - ALTER TABLE 

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
