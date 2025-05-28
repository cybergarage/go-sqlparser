# Information Schema

## Columns

<table style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Column</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">MariaDB</th>
<th style="text-align: left;">SQL Server</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>TABLE_CATALOG</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>Always def</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>TABLE_SCHEMA</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>Name of the schema (database)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>TABLE_NAME</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>Name of the table containing the column</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>COLUMN_NAME</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>Name of the column</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DATA_TYPE</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>Data type of the column</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>IS_NULLABLE</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>Indicates if the column can contain NULL values</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>COLUMN_DEFAULT</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>varchar</p></td>
<td style="text-align: left;"><p>Default value of the column</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>CHARACTER_MAXIMUM_LENGTH</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>Maximum length of the column</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>NUMERIC_PRECISION</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>Precision of numeric columns</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>NUMERIC_SCALE</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>Scale of numeric columns</p></td>
</tr>
</tbody>
</table>

- [Information Schema - PostgreSQL](https://www.postgresql.org/docs/current/information-schema.html)

- [The INFORMATION\_SCHEMA TABLES Table - MySQL](https://dev.mysql.com/doc/refman/8.0/en/information-schema.html)

- [Information Schema TABLES Table - MariaDB Knowledge Base](https://mariadb.com/kb/en/information-schema-tables-table/)

- [System Information Schema Views (Transact-SQL) - SQL Server | Microsoft Learn](https://learn.microsoft.com/en-us/sql/relational-databases/system-information-schema-views/system-information-schema-views-transact-sql?view=sql-server-ver16)

- [Introduction to INFORMATION\_SCHEMA | BigQuery | Google Cloud](https://cloud.google.com/bigquery/docs/information-schema-intro)
