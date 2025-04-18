# Query Types

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
<th style="text-align: left;">Query Type</th>
<th style="text-align: left;">Operation</th>
<th style="text-align: left;">SQL99</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">SQLite3</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>ALTER DATABSE</p></td>
<td style="text-align: left;"><p>TO</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>ALTER TABLE</p></td>
<td style="text-align: left;"><p>RENAME TABLE</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>RENAME COLUMN</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>ADD COLUMN</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>ADD CONSTRAINT</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>ADD CONSTRAINT INDEX</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>DESCRIBE</p></td>
<td style="text-align: left;"><p>TABLE</p></td>
<td style="text-align: left;"><p>*1)</p></td>
<td style="text-align: left;"><p>\d</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>*2)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>SHOW</p></td>
<td style="text-align: left;"><p>COLUMNS</p></td>
<td style="text-align: left;"><p>*1)</p></td>
<td style="text-align: left;"><p>\d</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>*2)</p></td>
</tr>
</tbody>
</table>

1.  The SQL99 standard does not define a DESCRIBE and SHOT statements. However, it defines the INFORMATION\_SCHEMA.TABLES and INFORMATION\_SCHEMA.COLUMNS views.

2.  [sqlite\_schema - The Schema Table](https://www.sqlite.org/schematab.html)
