# Function Types

`go-sqparser` supports a variety of function types that can be used in SQL queries in the `fn` package. Below is a list of the supported function types along with their descriptions.

<table>
<colgroup>
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
<col style="width: 5%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Function Name</th>
<th style="text-align: left;">go-sqlparser</th>
<th style="text-align: left;">SQL Standard</th>
<th style="text-align: left;">Oracle</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">SQL Server</th>
<th style="text-align: left;">Notes</th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>String Functions</p></td>
<td style="text-align: left;"><p>LENGTH / CHAR_LENGTH</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>LENGTH / OCTET_LENGTH / CHAR_LENGTH</p></td>
<td style="text-align: left;"><p>LENGTH</p></td>
<td style="text-align: left;"><p>LENGTH / CHAR_LENGTH</p></td>
<td style="text-align: left;"><p>LENGTH / CHAR_LENGTH</p></td>
<td style="text-align: left;"><p>LEN / DATALENGTH</p></td>
<td style="text-align: left;"><p>LEN for characters</p></td>
<td style="text-align: left;"><p>DATALENGTH for bytes. CHAR_LENGTH is generally for characters.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String Functions</p></td>
<td style="text-align: left;"><p>SUBSTRING</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>SUBSTRING(string FROM start FOR length)</p></td>
<td style="text-align: left;"><p>SUBSTR(string</p></td>
<td style="text-align: left;"><p>start</p></td>
<td style="text-align: left;"><p>length)</p></td>
<td style="text-align: left;"><p>SUBSTRING(string</p></td>
<td style="text-align: left;"><p>start</p></td>
<td style="text-align: left;"><p>length)</p></td>
<td style="text-align: left;"><p>SUBSTRING(string FROM start FOR length)</p></td>
<td style="text-align: left;"><p>SUBSTRING(string</p></td>
<td style="text-align: left;"><p>start</p></td>
<td style="text-align: left;"><p>length)</p></td>
<td style="text-align: left;"><p>Syntax varies slightly.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String Functions</p></td>
<td style="text-align: left;"><p>CONCATENATION</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>|| (double pipe)</p></td>
<td style="text-align: left;"><p>|| (double pipe) / CONCAT(string1</p></td>
<td style="text-align: left;"><p>string2)</p></td>
<td style="text-align: left;"><p>CONCAT(string1</p></td>
<td style="text-align: left;"><p>string2</p></td>
<td style="text-align: left;"><p>…​)</p></td>
<td style="text-align: left;"><p>|| (double pipe) / CONCAT(string1</p></td>
<td style="text-align: left;"><p>string2)</p></td>
<td style="text-align: left;"><p>+ (plus sign) / CONCAT(string1</p></td>
<td style="text-align: left;"><p>string2</p></td>
<td style="text-align: left;"><p>…​)</p></td>
<td style="text-align: left;"><p>CONCAT is an alternative</p></td>
<td style="text-align: left;"><p>+ is common.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String Functions</p></td>
<td style="text-align: left;"><p>LOWER</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>LOWER(string)</p></td>
<td style="text-align: left;"><p>LOWER(string)</p></td>
<td style="text-align: left;"><p>LOWER(string)</p></td>
<td style="text-align: left;"><p>LOWER(string)</p></td>
<td style="text-align: left;"><p>LOWER(string)</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String Functions</p></td>
<td style="text-align: left;"><p>UPPER</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>UPPER(string)</p></td>
<td style="text-align: left;"><p>UPPER(string)</p></td>
<td style="text-align: left;"><p>UPPER(string)</p></td>
<td style="text-align: left;"><p>UPPER(string)</p></td>
<td style="text-align: left;"><p>UPPER(string)</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String Functions</p></td>
<td style="text-align: left;"><p>TRIM</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>TRIM([LEADING | TRAILING | BOTH] [character FROM] string)</p></td>
<td style="text-align: left;"><p>TRIM([character FROM] string)</p></td>
<td style="text-align: left;"><p>TRIM([BOTH | LEADING | TRAILING] [character FROM] string)</p></td>
<td style="text-align: left;"><p>TRIM([LEADING | TRAILING | BOTH] [character FROM] string)</p></td>
<td style="text-align: left;"><p>TRIM([character FROM] string) / LTRIM / RTRIM</p></td>
<td style="text-align: left;"><p>Syntax variations for specifying character and position. LTRIM/RTRIM are common.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String Functions</p></td>
<td style="text-align: left;"><p>LPAD / RPAD</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>No direct standard</p></td>
<td style="text-align: left;"><p>LPAD(string</p></td>
<td style="text-align: left;"><p>length</p></td>
<td style="text-align: left;"><p>pad_string)</p></td>
<td style="text-align: left;"><p>LPAD(string</p></td>
<td style="text-align: left;"><p>length</p></td>
<td style="text-align: left;"><p>pad_string)</p></td>
<td style="text-align: left;"><p>LPAD(string</p></td>
<td style="text-align: left;"><p>length</p></td>
<td style="text-align: left;"><p>pad_string)</p></td>
<td style="text-align: left;"><p>No direct equivalent (can be simulated)</p></td>
<td style="text-align: left;"><p>Commonly used for padding strings.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String Functions</p></td>
<td style="text-align: left;"><p>REPLACE</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>No direct standard</p></td>
<td style="text-align: left;"><p>REPLACE(string</p></td>
<td style="text-align: left;"><p>search</p></td>
<td style="text-align: left;"><p>replace)</p></td>
<td style="text-align: left;"><p>REPLACE(string</p></td>
<td style="text-align: left;"><p>search</p></td>
<td style="text-align: left;"><p>replace)</p></td>
<td style="text-align: left;"><p>REPLACE(string</p></td>
<td style="text-align: left;"><p>search</p></td>
<td style="text-align: left;"><p>replace)</p></td>
<td style="text-align: left;"><p>REPLACE(string</p></td>
<td style="text-align: left;"><p>search</p></td>
<td style="text-align: left;"><p>replace)</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Numeric Functions</p></td>
<td style="text-align: left;"><p>ABS</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>ABS(number)</p></td>
<td style="text-align: left;"><p>ABS(number)</p></td>
<td style="text-align: left;"><p>ABS(number)</p></td>
<td style="text-align: left;"><p>ABS(number)</p></td>
<td style="text-align: left;"><p>ABS(number)</p></td>
<td style="text-align: left;"><p>Absolute value.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Numeric Functions</p></td>
<td style="text-align: left;"><p>CEIL / CEILING</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>CEILING(number)</p></td>
<td style="text-align: left;"><p>CEIL(number)</p></td>
<td style="text-align: left;"><p>CEIL(number)</p></td>
<td style="text-align: left;"><p>CEIL(number)</p></td>
<td style="text-align: left;"><p>CEILING(number)</p></td>
<td style="text-align: left;"><p>Rounds up to the nearest integer.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Numeric Functions</p></td>
<td style="text-align: left;"><p>FLOOR</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>FLOOR(number)</p></td>
<td style="text-align: left;"><p>FLOOR(number)</p></td>
<td style="text-align: left;"><p>FLOOR(number)</p></td>
<td style="text-align: left;"><p>FLOOR(number)</p></td>
<td style="text-align: left;"><p>FLOOR(number)</p></td>
<td style="text-align: left;"><p>Rounds down to the nearest integer.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Numeric Functions</p></td>
<td style="text-align: left;"><p>ROUND</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>ROUND(number</p></td>
<td style="text-align: left;"><p>decimal_places)</p></td>
<td style="text-align: left;"><p>ROUND(number</p></td>
<td style="text-align: left;"><p>decimal_places)</p></td>
<td style="text-align: left;"><p>ROUND(number</p></td>
<td style="text-align: left;"><p>decimal_places)</p></td>
<td style="text-align: left;"><p>ROUND(number</p></td>
<td style="text-align: left;"><p>decimal_places)</p></td>
<td style="text-align: left;"><p>ROUND(number</p></td>
<td style="text-align: left;"><p>decimal_places)</p></td>
<td style="text-align: left;"><p>Rounds to specified decimal places.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Numeric Functions</p></td>
<td style="text-align: left;"><p>MOD</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>MOD(dividend</p></td>
<td style="text-align: left;"><p>divisor)</p></td>
<td style="text-align: left;"><p>MOD(dividend</p></td>
<td style="text-align: left;"><p>divisor)</p></td>
<td style="text-align: left;"><p>MOD(dividend</p></td>
<td style="text-align: left;"><p>divisor)</p></td>
<td style="text-align: left;"><p>MOD(dividend</p></td>
<td style="text-align: left;"><p>divisor)</p></td>
<td style="text-align: left;"><p>% (modulo operator)</p></td>
<td style="text-align: left;"><p>Modulo operation (remainder).</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Numeric Functions</p></td>
<td style="text-align: left;"><p>POWER</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>POWER(base</p></td>
<td style="text-align: left;"><p>exponent)</p></td>
<td style="text-align: left;"><p>POWER(base</p></td>
<td style="text-align: left;"><p>exponent)</p></td>
<td style="text-align: left;"><p>POW(base</p></td>
<td style="text-align: left;"><p>exponent) / POWER(base</p></td>
<td style="text-align: left;"><p>exponent)</p></td>
<td style="text-align: left;"><p>POWER(base</p></td>
<td style="text-align: left;"><p>exponent)</p></td>
<td style="text-align: left;"><p>POWER(base</p></td>
<td style="text-align: left;"><p>exponent)</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Numeric Functions</p></td>
<td style="text-align: left;"><p>SQRT</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SQRT(number)</p></td>
<td style="text-align: left;"><p>SQRT(number)</p></td>
<td style="text-align: left;"><p>SQRT(number)</p></td>
<td style="text-align: left;"><p>SQRT(number)</p></td>
<td style="text-align: left;"><p>SQRT(number)</p></td>
<td style="text-align: left;"><p>Square root.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Date/Time Functions</p></td>
<td style="text-align: left;"><p>CURRENT_DATE</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CURRENT_DATE</p></td>
<td style="text-align: left;"><p>SYSDATE (date + time)</p></td>
<td style="text-align: left;"><p>CURDATE()</p></td>
<td style="text-align: left;"><p>CURRENT_DATE</p></td>
<td style="text-align: left;"><p>GETDATE() (date + time) / CURRENT_TIMESTAMP</p></td>
<td style="text-align: left;"><p>Returns current date. SYSDATE/GETDATE() in Oracle/SQL Server include time.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Date/Time Functions</p></td>
<td style="text-align: left;"><p>CURRENT_TIME</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CURRENT_TIME</p></td>
<td style="text-align: left;"><p>SYSDATE (date + time)</p></td>
<td style="text-align: left;"><p>CURTIME()</p></td>
<td style="text-align: left;"><p>CURRENT_TIME</p></td>
<td style="text-align: left;"><p>GETDATE() (date + time) / CURRENT_TIMESTAMP</p></td>
<td style="text-align: left;"><p>Returns current time. SYSDATE/GETDATE() in Oracle/SQL Server include date.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Date/Time Functions</p></td>
<td style="text-align: left;"><p>CURRENT_TIMESTAMP</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CURRENT_TIMESTAMP</p></td>
<td style="text-align: left;"><p>SYSTIMESTAMP / LOCALTIMESTAMP</p></td>
<td style="text-align: left;"><p>NOW() / CURRENT_TIMESTAMP</p></td>
<td style="text-align: left;"><p>NOW() / CURRENT_TIMESTAMP</p></td>
<td style="text-align: left;"><p>GETDATE() / SYSDATETIME() / CURRENT_TIMESTAMP</p></td>
<td style="text-align: left;"><p>Returns current date and time. SYSDATETIME() for higher precision.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Date/Time Functions</p></td>
<td style="text-align: left;"><p>EXTRACT</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>EXTRACT(part FROM datetime)</p></td>
<td style="text-align: left;"><p>EXTRACT(part FROM datetime)</p></td>
<td style="text-align: left;"><p>EXTRACT(part FROM datetime) / YEAR()</p></td>
<td style="text-align: left;"><p>MONTH()</p></td>
<td style="text-align: left;"><p>DAY()</p></td>
<td style="text-align: left;"><p>etc.</p></td>
<td style="text-align: left;"><p>EXTRACT(part FROM datetime)</p></td>
<td style="text-align: left;"><p>DATEPART(part</p></td>
<td style="text-align: left;"><p>date) / YEAR()</p></td>
<td style="text-align: left;"><p>MONTH()</p></td>
<td style="text-align: left;"><p>DAY()</p></td>
<td style="text-align: left;"><p>etc.</p></td>
<td style="text-align: left;"><p>Used to extract parts like year</p></td>
<td style="text-align: left;"><p>month</p></td>
<td style="text-align: left;"><p>day.</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Date/Time Functions</p></td>
<td style="text-align: left;"><p>DATE_ADD / INTERVAL</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>No direct standard</p></td>
<td style="text-align: left;"><p>DATE_ADD (using INTERVAL syntax)</p></td>
<td style="text-align: left;"><p>DATE_ADD(date</p></td>
<td style="text-align: left;"><p>INTERVAL value unit)</p></td>
<td style="text-align: left;"><p>date + INTERVAL value unit</p></td>
<td style="text-align: left;"><p>DATEADD(unit</p></td>
<td style="text-align: left;"><p>value</p></td>
<td style="text-align: left;"><p>date)</p></td>
<td style="text-align: left;"><p>Adds/subtracts time intervals.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Date/Time Functions</p></td>
<td style="text-align: left;"><p>DATE_DIFF</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>No direct standard</p></td>
<td style="text-align: left;"><p>DATE_DIFF (using specific functions like MONTHS_BETWEEN)</p></td>
<td style="text-align: left;"><p>DATEDIFF(expr1</p></td>
<td style="text-align: left;"><p>expr2)</p></td>
<td style="text-align: left;"><p>AGE(timestamp1</p></td>
<td style="text-align: left;"><p>timestamp2)</p></td>
<td style="text-align: left;"><p>DATEDIFF(unit</p></td>
<td style="text-align: left;"><p>startdate</p></td>
<td style="text-align: left;"><p>enddate)</p></td>
<td style="text-align: left;"><p>Calculates difference between dates.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Aggregate Functions</p></td>
<td style="text-align: left;"><p>COUNT</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>COUNT(*) / COUNT(column)</p></td>
<td style="text-align: left;"><p>COUNT(*) / COUNT(column)</p></td>
<td style="text-align: left;"><p>COUNT(*) / COUNT(column)</p></td>
<td style="text-align: left;"><p>COUNT(*) / COUNT(column)</p></td>
<td style="text-align: left;"><p>COUNT(*) / COUNT(column)</p></td>
<td style="text-align: left;"><p>Counts rows or non-NULL values.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Aggregate Functions</p></td>
<td style="text-align: left;"><p>SUM</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SUM(column)</p></td>
<td style="text-align: left;"><p>SUM(column)</p></td>
<td style="text-align: left;"><p>SUM(column)</p></td>
<td style="text-align: left;"><p>SUM(column)</p></td>
<td style="text-align: left;"><p>SUM(column)</p></td>
<td style="text-align: left;"><p>Calculates sum of values.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Aggregate Functions</p></td>
<td style="text-align: left;"><p>AVG</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>AVG(column)</p></td>
<td style="text-align: left;"><p>AVG(column)</p></td>
<td style="text-align: left;"><p>AVG(column)</p></td>
<td style="text-align: left;"><p>AVG(column)</p></td>
<td style="text-align: left;"><p>AVG(column)</p></td>
<td style="text-align: left;"><p>Calculates average of values.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Aggregate Functions</p></td>
<td style="text-align: left;"><p>MIN</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>MIN(column)</p></td>
<td style="text-align: left;"><p>MIN(column)</p></td>
<td style="text-align: left;"><p>MIN(column)</p></td>
<td style="text-align: left;"><p>MIN(column)</p></td>
<td style="text-align: left;"><p>MIN(column)</p></td>
<td style="text-align: left;"><p>Finds minimum value.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Aggregate Functions</p></td>
<td style="text-align: left;"><p>MAX</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>MAX(column)</p></td>
<td style="text-align: left;"><p>MAX(column)</p></td>
<td style="text-align: left;"><p>MAX(column)</p></td>
<td style="text-align: left;"><p>MAX(column)</p></td>
<td style="text-align: left;"><p>MAX(column)</p></td>
<td style="text-align: left;"><p>Finds maximum value.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Aggregate Functions</p></td>
<td style="text-align: left;"><p>GROUP_CONCAT</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>No direct standard</p></td>
<td style="text-align: left;"><p>LISTAGG (Oracle 11g+)</p></td>
<td style="text-align: left;"><p>GROUP_CONCAT(column [ORDER BY] [SEPARATOR])</p></td>
<td style="text-align: left;"><p>STRING_AGG(expression</p></td>
<td style="text-align: left;"><p>delimiter)</p></td>
<td style="text-align: left;"><p>STRING_AGG(expression</p></td>
<td style="text-align: left;"><p>delimiter) (SQL Server 2017+)</p></td>
<td style="text-align: left;"><p>Concatenates strings within a group.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Conditional Expressions</p></td>
<td style="text-align: left;"><p>CASE</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CASE WHEN condition THEN result [ELSE result] END</p></td>
<td style="text-align: left;"><p>CASE WHEN condition THEN result [ELSE result] END</p></td>
<td style="text-align: left;"><p>CASE WHEN condition THEN result [ELSE result] END</p></td>
<td style="text-align: left;"><p>CASE WHEN condition THEN result [ELSE result] END</p></td>
<td style="text-align: left;"><p>CASE WHEN condition THEN result [ELSE result] END</p></td>
<td style="text-align: left;"><p>Standard conditional logic.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Type Conversion</p></td>
<td style="text-align: left;"><p>CAST</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CAST(expression AS type)</p></td>
<td style="text-align: left;"><p>CAST(expression AS type) / TO_CHAR</p></td>
<td style="text-align: left;"><p>TO_DATE</p></td>
<td style="text-align: left;"><p>TO_NUMBER</p></td>
<td style="text-align: left;"><p>CAST(expression AS type) / CONVERT</p></td>
<td style="text-align: left;"><p>CAST(expression AS type) / to_char</p></td>
<td style="text-align: left;"><p>to_date</p></td>
<td style="text-align: left;"><p>to_number</p></td>
<td style="text-align: left;"><p>CAST(expression AS type) / CONVERT(type</p></td>
<td style="text-align: left;"><p>expression)</p></td>
<td style="text-align: left;"><p>Converts one data type to another.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Ranking Functions</p></td>
<td style="text-align: left;"><p>ROW_NUMBER</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>ROW_NUMBER() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>ROW_NUMBER() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>No direct equivalent (can be simulated)</p></td>
<td style="text-align: left;"><p>ROW_NUMBER() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>ROW_NUMBER() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>Assigns a unique row number within a partition.</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Ranking Functions</p></td>
<td style="text-align: left;"><p>RANK</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>RANK() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>RANK() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>No direct equivalent (can be simulated)</p></td>
<td style="text-align: left;"><p>RANK() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>RANK() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>Assigns a rank within a partition (with gaps for ties).</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Ranking Functions</p></td>
<td style="text-align: left;"><p>DENSE_RANK</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>DENSE_RANK() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>DENSE_RANK() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>No direct equivalent (can be simulated)</p></td>
<td style="text-align: left;"><p>DENSE_RANK() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>DENSE_RANK() OVER (PARTITION BY …​ ORDER BY …​)</p></td>
<td style="text-align: left;"><p>Assigns a rank within a partition (without gaps for ties).</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Window Functions</p></td>
<td style="text-align: left;"><p>LEAD / LAG</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>LEAD(column</p></td>
<td style="text-align: left;"><p>offset</p></td>
<td style="text-align: left;"><p>default) OVER (…​)</p></td>
<td style="text-align: left;"><p>LEAD(column</p></td>
<td style="text-align: left;"><p>offset</p></td>
<td style="text-align: left;"><p>default) OVER (…​)</p></td>
<td style="text-align: left;"><p>No direct equivalent (can be simulated)</p></td>
<td style="text-align: left;"><p>LEAD(column</p></td>
<td style="text-align: left;"><p>offset</p></td>
<td style="text-align: left;"><p>default) OVER (…​)</p></td>
<td style="text-align: left;"><p>LEAD(column</p></td>
<td style="text-align: left;"><p>offset</p></td>
<td style="text-align: left;"><p>default) OVER (…​)</p></td>
<td style="text-align: left;"><p>Accesses a row at a given physical offset after/before the current row.</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>
