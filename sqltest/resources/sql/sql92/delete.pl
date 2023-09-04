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
my $line_no = 0;
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
  for (my $pr_key_idx = 0; $pr_key_idx < scalar(@row); $pr_key_idx++) {
    my $type_name = lc($data_type_row[$pr_key_idx]);
    my $column_name = "c" . $type_name;
    print "DELETE * FROM ${tbl_name} WHERE $column_name = $row[$pr_key_idx];\n";  
  }
}
close(IN);
