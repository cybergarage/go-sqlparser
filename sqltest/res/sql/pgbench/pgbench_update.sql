UPDATE pgbench_accounts SET abalance = abalance + 1071 WHERE aid = 30142;
UPDATE pgbench_tellers SET tbalance = tbalance + 1071 WHERE tid = 1;
UPDATE pgbench_branches SET bbalance = bbalance + 1071 WHERE bid = 1;
UPDATE pgbench_accounts SET abalance = abalance + -1065 WHERE aid = 15801;
UPDATE pgbench_tellers SET tbalance = tbalance + -1065 WHERE tid = 9;
UPDATE pgbench_branches SET bbalance = bbalance + -1065 WHERE bid = 1;
UPDATE pgbench_accounts SET abalance = abalance + -4317 WHERE aid = 93403;
UPDATE pgbench_tellers SET tbalance = tbalance + -4317 WHERE tid = 3;
UPDATE pgbench_branches SET bbalance = bbalance + -4317 WHERE bid = 1;
UPDATE pgbench_accounts SET abalance = abalance + 178 WHERE aid = 1209;
UPDATE pgbench_tellers SET tbalance = tbalance + 178 WHERE tid = 7;
UPDATE pgbench_branches SET bbalance = bbalance + 178 WHERE bid = 1;
UPDATE pgbench_accounts SET abalance = abalance + 350 WHERE aid = 89359;
UPDATE pgbench_tellers SET tbalance = tbalance + 350 WHERE tid = 8;
UPDATE pgbench_branches SET bbalance = bbalance + 350 WHERE bid = 1;

