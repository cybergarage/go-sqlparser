#!/usr/bin/perl
# Copyright (C) 2022 The go-sqlparser Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

my $pr_key_type = "";
if (1 <= @ARGV){
  $pr_key_type = lc($ARGV[0]);
}

print<<HEADER;
- Copyright (C) 2020 The go-sqlparser Authors. All rights reserved.
-
- Licensed under the Apache License, Version 2.0 (the "License");
- you may not use this file except in compliance with the License.
- You may obtain a copy of the License at
-
-  http:-www.apache.org/licenses/LICENSE-2.0
-
- Unless required by applicable law or agreed to in writing, software
- distributed under the License is distributed on an "AS IS" BASIS,
- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
- See the License for the specific language governing permissions and
- limitations under the License.

HEADER

my $data_type_file = "data/data_type.pict";
open(IN, $data_type_file) or die "Failed to open $data_type_file: $!";
my $line_no = 0;
my $pr_key_idx = -1;
my @data_type_row;
my $tbl_name = "test";
while(<IN>){
  $line_no++;
  chomp($_);
  my @row = split(/\t/, $_, -1);
  if ($line_no <= 1) {
    @data_type_row = @row;
    print "CREATE TABLE ${tbl_name} (\n";  
    for (my $i = 0; $i < scalar(@row); $i++) {
      my $type_name = lc($row[$i]);
      my $column_type = uc($type_name);
      my $column_name = "c" . $type_name;
      if ($type_name eq $pr_key_type) {
        print "\t$column_name $column_type PRIMARY KEY";  
        $pr_key_idx = $i;
      } else {
        print "\t$column_name $column_type";  
      }
      if ($i < ((@row) - 1)) {
        print ",";  
      }
      print "\n";  
    }
    print ");\n";  
    print "{\n";  
    print "}\n";  
    next;
  }
  if ($pr_key_idx < 0) {
    die "The primary key type ($pr_key_type) is not found in $data_type_file";
  }
  if ($line_no <= 2) {
    print "INSERT INTO ${tbl_name} (";  
    for (my $i = 0; $i < scalar(@row); $i++) {
      my $type_name = lc($data_type_row[$i]);
      my $column_name = "c" . $type_name;
      if (0 < $i) {
        print ", ";
      }
      print $column_name;
    }
    print ") VALUES (";  
    for (my $i = 0; $i < scalar(@row); $i++) {
      if (0 < $i) {
        print ", ";  
      }
      print $row[$i];
    }
    print ");\n";  
    print "{\n";  
    print "}\n";  
  } else {
    print "UPDATE ${tbl_name} SET ";  
    my $n_colums = 0;
    for (my $i = 0; $i < scalar(@row); $i++) {
      my $type_name = lc($data_type_row[$i]);
      my $column_name = "c" . $type_name;
      if ($i == $pr_key_idx) {
        next;
      }
      if (0 < $n_colums) {
        print ", ";  
      }
      print "$column_name = $row[$i]";
      $n_colums++;
    }
    my $type_name = lc($data_type_row[$pr_key_idx]);
    my $column_name = "c" . $type_name;
    print " WHERE $column_name = $row[$pr_key_idx];\n";    
    print "{\n";  
    print "}\n";  
  }
  my $type_name = lc($data_type_row[$pr_key_idx]);
  my $column_name = "c" . $type_name;
  print "SELECT * FROM ${tbl_name} WHERE $column_name = $row[$pr_key_idx];\n";  
  print "{\n";  
  print "\t\"rows\" :\n";  
  print "\t[\n";  
  print "\t\t{\n";  
  for (my $i = 0; $i < scalar(@row); $i++) {
    my $type_name = lc($data_type_row[$i]);
    my $column_name = "c" . $type_name;
    my $column_val = $row[$i];
    $column_val =~ s/'/"/g;
    print "\t\t\t\"$column_name\" : $column_val";
    if ($i < ((@row) - 1)) {
      print ",";  
    }
    print "\n";
  }
  print "\t\t}\n";  
  print "\t]\n";  
  print "}\n";  
}
close(IN);
