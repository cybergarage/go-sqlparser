# Data Types

`go-sqlparser` supports a variety of data types that can be used in SQL queries. Below is a list of the supported data types along with their descriptions.

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
<th style="text-align: left;">go-sqlparser</th>
<th style="text-align: left;">SQL99</th>
<th style="text-align: left;">Oracle</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">SQLite</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>CHAR(n)</p></td>
<td style="text-align: left;"><p>CHAR(n)</p></td>
<td style="text-align: left;"><p>CHAR(n)</p></td>
<td style="text-align: left;"><p>CHAR(n)</p></td>
<td style="text-align: left;"><p>CHAR(n)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>VARCHAR(n)</p></td>
<td style="text-align: left;"><p>VARCHAR2(n)</p></td>
<td style="text-align: left;"><p>VARCHAR(n)</p></td>
<td style="text-align: left;"><p>VARCHAR(n)</p></td>
<td style="text-align: left;"><p>VARCHAR(n)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>STRING</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>NUMBER(n)</p></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>INT</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>NUMBER(n)</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>INT</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>NUMBER(n)</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>FLOAT</p></td>
<td style="text-align: left;"><p>REAL</p></td>
<td style="text-align: left;"><p>BINARY_FLOAT</p></td>
<td style="text-align: left;"><p>REAL</p></td>
<td style="text-align: left;"><p>FLOAT</p></td>
<td style="text-align: left;"><p>REAL</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DOUBLE</p></td>
<td style="text-align: left;"><p>DOUBLE PRECISION</p></td>
<td style="text-align: left;"><p>BINARY_DOUBLE</p></td>
<td style="text-align: left;"><p>DOUBLE PRECISION</p></td>
<td style="text-align: left;"><p>DOUBLE</p></td>
<td style="text-align: left;"><p>REAL</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DATE</p></td>
<td style="text-align: left;"><p>DATE</p></td>
<td style="text-align: left;"><p>DATE</p></td>
<td style="text-align: left;"><p>DATE</p></td>
<td style="text-align: left;"><p>DATE</p></td>
<td style="text-align: left;"><p>DATE</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DATETIME</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>DATETIME</p></td>
<td style="text-align: left;"><p>DATETIME</p></td>
<td style="text-align: left;"><p>DATETIME</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
<td style="text-align: left;"><p>TIMESTAMP(6)</p></td>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>BOOLEAN</p></td>
<td style="text-align: left;"><p>BOOLEAN</p></td>
<td style="text-align: left;"><p>NUMBER(1)</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>TINYINT</p></td>
<td style="text-align: left;"><p>NUMERIC(1)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>BLOB</p></td>
<td style="text-align: left;"><p>BLOB</p></td>
<td style="text-align: left;"><p>LONG RAW</p></td>
<td style="text-align: left;"><p>BYTEA</p></td>
<td style="text-align: left;"><p>LONGBLOB</p></td>
<td style="text-align: left;"><p>BLOB</p></td>
</tr>
</tbody>
</table>
