# Copyright (C) 2022 The go-sqlparser Authors. All rights reserved.
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

all: queries

%.sql : %.pl $(wildcard data/*.pict)
	perl $< > $@

QUERY_FILES = \
	create_table.sql \
	insert.sql \
	select.sql \
	update.sql \
	delete.sql \
	create_index.sql

queries: ${QUERY_FILES}
