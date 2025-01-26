# Copyright (C) 2024 The go-sqlparser Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

GO_SQLTEST_DIR=../../../../../go-sqltest
TEST_FILE_PREFIX=sqltest-

all: insert select update delete

.PHONY: insert select update delete

insert:
	find ${GO_SQLTEST_DIR} -name "*.qst" | xargs more | grep INSERT > ${TEST_FILE_PREFIX}insert.sql
	git add ${TEST_FILE_PREFIX}insert.sql
	git commit -m "Update ${TEST_FILE_PREFIX}insert.sql" ${TEST_FILE_PREFIX}insert.sql

select:
	find ${GO_SQLTEST_DIR} -name "*.qst" | xargs more | grep SELECT > ${TEST_FILE_PREFIX}select.sql
	git add ${TEST_FILE_PREFIX}select.sql
	git commit -m "Update ${TEST_FILE_PREFIX}select.sql" ${TEST_FILE_PREFIX}select.sql

update:
	find ${GO_SQLTEST_DIR} -name "*.qst" | xargs more | grep UPDATE | grep -v pgbench > ${TEST_FILE_PREFIX}update.sql
	git add ${TEST_FILE_PREFIX}update.sql
	git commit -m "Update ${TEST_FILE_PREFIX}update.sql" ${TEST_FILE_PREFIX}update.sql

delete:
	find ${GO_SQLTEST_DIR} -name "*.qst" | xargs more | grep DELETE > ${TEST_FILE_PREFIX}delete.sql
	git add ${TEST_FILE_PREFIX}delete.sql
	git commit -m "Update ${TEST_FILE_PREFIX}delete.sql" ${TEST_FILE_PREFIX}delete.sql