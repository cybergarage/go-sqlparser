#!/usr/bin/perl
# Copyright (C) 2022 The go-sqlparser Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

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
    next;
  }
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
}

close(IN);
