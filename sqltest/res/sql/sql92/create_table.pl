#!/usr/bin/perl
# Copyright (C) 2022 The go-sqlparser Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

my $data_type_file = "data/data_type.pict";
if (1 <= @ARGV){
  $data_type_file = $ARGV[0];
}

open(IN, $data_type_file) or die "Failed to open $data_type_file: $!";

my $first_line = <IN>;
unless ($first_line) {
  die "Failed to read $data_type_file: $!";
}

chomp($first_line);
my @data_types = split(/\t/, $first_line, -1);

my $tbl_name = "test";

for (my $n = 0; $n < scalar(@data_types); $n++) {
  print "CREATE TABLE ${tbl_name} (\n";  
  for (my $i = 0; $i < scalar(@data_types); $i++) {
      my $type_name = lc($data_types[$i]);
      my $column_type = uc($type_name);
      my $column_name = "c" . $type_name;
      if ($i == $n) {
        print "\t$column_name $column_type PRIMARY KEY";  
      } else {
        print "\t$column_name $column_type";  
      }
      if ($i < ((@data_types) - 1)) {
        print ",";  
      }
      print "\n";
    }
  print ");\n\n";  
}

close(IN);
