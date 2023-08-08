select count(*) from pgbench_branches;
# select o.n, p.partstrat, pg_catalog.count(i.inhparent) from pg_catalog.pg_class as c join pg_catalog.pg_namespace as n on (n.oid = c.relnamespace) cross join lateral (select pg_catalog.array_position(pg_catalog.current_schemas(true), n.nspname)) as o(n) left join pg_catalog.pg_partitioned_table as p on (p.partrelid = c.oid) left join pg_catalog.pg_inherits as i on (c.oid = i.inhparent) where c.relname = 'pgbench_accounts' and o.n is not null group by 1, 2 order by 1 asc limit 1;
SELECT abalance FROM pgbench_accounts WHERE aid = 20855;
SELECT abalance FROM pgbench_accounts WHERE aid = 62497;
SELECT abalance FROM pgbench_accounts WHERE aid = 60047;
SELECT abalance FROM pgbench_accounts WHERE aid = 35003;
SELECT abalance FROM pgbench_accounts WHERE aid = 19368;
SELECT abalance FROM pgbench_accounts WHERE aid = 6235;
SELECT abalance FROM pgbench_accounts WHERE aid = 46600;
SELECT abalance FROM pgbench_accounts WHERE aid = 9640;
SELECT abalance FROM pgbench_accounts WHERE aid = 79406;
SELECT abalance FROM pgbench_accounts WHERE aid = 92129;
SELECT abalance FROM pgbench_accounts WHERE aid = 21537;
SELECT abalance FROM pgbench_accounts WHERE aid = 4732;
SELECT abalance FROM pgbench_accounts WHERE aid = 17855;