// Code generated from SQLiteParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // SQLiteParser
import "github.com/antlr4-go/antlr/v4"

// SQLiteParserListener is a complete listener for a parse tree produced by SQLiteParser.
type SQLiteParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterSql_stmt_list is called when entering the sql_stmt_list production.
	EnterSql_stmt_list(c *Sql_stmt_listContext)

	// EnterSql_stmt is called when entering the sql_stmt production.
	EnterSql_stmt(c *Sql_stmtContext)

	// EnterAlter_database_stmt is called when entering the alter_database_stmt production.
	EnterAlter_database_stmt(c *Alter_database_stmtContext)

	// EnterRename_database_to is called when entering the rename_database_to production.
	EnterRename_database_to(c *Rename_database_toContext)

	// EnterAlter_table_stmt is called when entering the alter_table_stmt production.
	EnterAlter_table_stmt(c *Alter_table_stmtContext)

	// EnterRename_table_to is called when entering the rename_table_to production.
	EnterRename_table_to(c *Rename_table_toContext)

	// EnterRename_table_colum is called when entering the rename_table_colum production.
	EnterRename_table_colum(c *Rename_table_columContext)

	// EnterAdd_table_column is called when entering the add_table_column production.
	EnterAdd_table_column(c *Add_table_columnContext)

	// EnterAdd_table_index is called when entering the add_table_index production.
	EnterAdd_table_index(c *Add_table_indexContext)

	// EnterDrop_table_column is called when entering the drop_table_column production.
	EnterDrop_table_column(c *Drop_table_columnContext)

	// EnterAnalyze_stmt is called when entering the analyze_stmt production.
	EnterAnalyze_stmt(c *Analyze_stmtContext)

	// EnterAttach_stmt is called when entering the attach_stmt production.
	EnterAttach_stmt(c *Attach_stmtContext)

	// EnterBegin_stmt is called when entering the begin_stmt production.
	EnterBegin_stmt(c *Begin_stmtContext)

	// EnterCommit_stmt is called when entering the commit_stmt production.
	EnterCommit_stmt(c *Commit_stmtContext)

	// EnterRollback_stmt is called when entering the rollback_stmt production.
	EnterRollback_stmt(c *Rollback_stmtContext)

	// EnterSavepoint_stmt is called when entering the savepoint_stmt production.
	EnterSavepoint_stmt(c *Savepoint_stmtContext)

	// EnterRelease_stmt is called when entering the release_stmt production.
	EnterRelease_stmt(c *Release_stmtContext)

	// EnterCreate_index_stmt is called when entering the create_index_stmt production.
	EnterCreate_index_stmt(c *Create_index_stmtContext)

	// EnterIndexed_column is called when entering the indexed_column production.
	EnterIndexed_column(c *Indexed_columnContext)

	// EnterCreate_database_stmt is called when entering the create_database_stmt production.
	EnterCreate_database_stmt(c *Create_database_stmtContext)

	// EnterIf_not_exists is called when entering the if_not_exists production.
	EnterIf_not_exists(c *If_not_existsContext)

	// EnterCreate_table_stmt is called when entering the create_table_stmt production.
	EnterCreate_table_stmt(c *Create_table_stmtContext)

	// EnterTable_options is called when entering the table_options production.
	EnterTable_options(c *Table_optionsContext)

	// EnterMysql_table_options is called when entering the mysql_table_options production.
	EnterMysql_table_options(c *Mysql_table_optionsContext)

	// EnterColumn_def is called when entering the column_def production.
	EnterColumn_def(c *Column_defContext)

	// EnterWith_param is called when entering the with_param production.
	EnterWith_param(c *With_paramContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterColumn_constraint is called when entering the column_constraint production.
	EnterColumn_constraint(c *Column_constraintContext)

	// EnterPrimary_key_constraint is called when entering the primary_key_constraint production.
	EnterPrimary_key_constraint(c *Primary_key_constraintContext)

	// EnterIndex_constraint is called when entering the index_constraint production.
	EnterIndex_constraint(c *Index_constraintContext)

	// EnterSigned_number is called when entering the signed_number production.
	EnterSigned_number(c *Signed_numberContext)

	// EnterTable_constraint is called when entering the table_constraint production.
	EnterTable_constraint(c *Table_constraintContext)

	// EnterPrimary_key_def is called when entering the primary_key_def production.
	EnterPrimary_key_def(c *Primary_key_defContext)

	// EnterUnique_key_def is called when entering the unique_key_def production.
	EnterUnique_key_def(c *Unique_key_defContext)

	// EnterForeign_key_clause is called when entering the foreign_key_clause production.
	EnterForeign_key_clause(c *Foreign_key_clauseContext)

	// EnterConflict_clause is called when entering the conflict_clause production.
	EnterConflict_clause(c *Conflict_clauseContext)

	// EnterCreate_trigger_stmt is called when entering the create_trigger_stmt production.
	EnterCreate_trigger_stmt(c *Create_trigger_stmtContext)

	// EnterCreate_view_stmt is called when entering the create_view_stmt production.
	EnterCreate_view_stmt(c *Create_view_stmtContext)

	// EnterCreate_virtual_table_stmt is called when entering the create_virtual_table_stmt production.
	EnterCreate_virtual_table_stmt(c *Create_virtual_table_stmtContext)

	// EnterWith_clause is called when entering the with_clause production.
	EnterWith_clause(c *With_clauseContext)

	// EnterCte_table_name is called when entering the cte_table_name production.
	EnterCte_table_name(c *Cte_table_nameContext)

	// EnterRecursive_cte is called when entering the recursive_cte production.
	EnterRecursive_cte(c *Recursive_cteContext)

	// EnterCommon_table_expression is called when entering the common_table_expression production.
	EnterCommon_table_expression(c *Common_table_expressionContext)

	// EnterDelete_stmt is called when entering the delete_stmt production.
	EnterDelete_stmt(c *Delete_stmtContext)

	// EnterDelete_stmt_limited is called when entering the delete_stmt_limited production.
	EnterDelete_stmt_limited(c *Delete_stmt_limitedContext)

	// EnterDetach_stmt is called when entering the detach_stmt production.
	EnterDetach_stmt(c *Detach_stmtContext)

	// EnterDrop_stmt is called when entering the drop_stmt production.
	EnterDrop_stmt(c *Drop_stmtContext)

	// EnterDrop_database_stmt is called when entering the drop_database_stmt production.
	EnterDrop_database_stmt(c *Drop_database_stmtContext)

	// EnterDrop_index_stmt is called when entering the drop_index_stmt production.
	EnterDrop_index_stmt(c *Drop_index_stmtContext)

	// EnterDrop_table_stmt is called when entering the drop_table_stmt production.
	EnterDrop_table_stmt(c *Drop_table_stmtContext)

	// EnterDrop_trigger_stmt is called when entering the drop_trigger_stmt production.
	EnterDrop_trigger_stmt(c *Drop_trigger_stmtContext)

	// EnterDrop_view_stmt is called when entering the drop_view_stmt production.
	EnterDrop_view_stmt(c *Drop_view_stmtContext)

	// EnterIf_exists is called when entering the if_exists production.
	EnterIf_exists(c *If_existsContext)

	// EnterUse_stmt is called when entering the use_stmt production.
	EnterUse_stmt(c *Use_stmtContext)

	// EnterCopy_stmt is called when entering the copy_stmt production.
	EnterCopy_stmt(c *Copy_stmtContext)

	// EnterCopy_column_list is called when entering the copy_column_list production.
	EnterCopy_column_list(c *Copy_column_listContext)

	// EnterCopy_format is called when entering the copy_format production.
	EnterCopy_format(c *Copy_formatContext)

	// EnterSource_name is called when entering the source_name production.
	EnterSource_name(c *Source_nameContext)

	// EnterTruncate_stmt is called when entering the truncate_stmt production.
	EnterTruncate_stmt(c *Truncate_stmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterComparison_expr is called when entering the comparison_expr production.
	EnterComparison_expr(c *Comparison_exprContext)

	// EnterAnd_expr is called when entering the and_expr production.
	EnterAnd_expr(c *And_exprContext)

	// EnterOr_expr is called when entering the or_expr production.
	EnterOr_expr(c *Or_exprContext)

	// EnterArithmetic_expr is called when entering the arithmetic_expr production.
	EnterArithmetic_expr(c *Arithmetic_exprContext)

	// EnterRaise_function is called when entering the raise_function production.
	EnterRaise_function(c *Raise_functionContext)

	// EnterLiteral_value is called when entering the literal_value production.
	EnterLiteral_value(c *Literal_valueContext)

	// EnterBind_param is called when entering the bind_param production.
	EnterBind_param(c *Bind_paramContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterValue_row is called when entering the value_row production.
	EnterValue_row(c *Value_rowContext)

	// EnterValues_clause is called when entering the values_clause production.
	EnterValues_clause(c *Values_clauseContext)

	// EnterInsert_stmt is called when entering the insert_stmt production.
	EnterInsert_stmt(c *Insert_stmtContext)

	// EnterReturning_clause is called when entering the returning_clause production.
	EnterReturning_clause(c *Returning_clauseContext)

	// EnterUpsert_clause is called when entering the upsert_clause production.
	EnterUpsert_clause(c *Upsert_clauseContext)

	// EnterPragma_stmt is called when entering the pragma_stmt production.
	EnterPragma_stmt(c *Pragma_stmtContext)

	// EnterPragma_value is called when entering the pragma_value production.
	EnterPragma_value(c *Pragma_valueContext)

	// EnterReindex_stmt is called when entering the reindex_stmt production.
	EnterReindex_stmt(c *Reindex_stmtContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterJoin_clause is called when entering the join_clause production.
	EnterJoin_clause(c *Join_clauseContext)

	// EnterSelect_core is called when entering the select_core production.
	EnterSelect_core(c *Select_coreContext)

	// EnterFactored_select_stmt is called when entering the factored_select_stmt production.
	EnterFactored_select_stmt(c *Factored_select_stmtContext)

	// EnterSimple_select_stmt is called when entering the simple_select_stmt production.
	EnterSimple_select_stmt(c *Simple_select_stmtContext)

	// EnterCompound_select_stmt is called when entering the compound_select_stmt production.
	EnterCompound_select_stmt(c *Compound_select_stmtContext)

	// EnterFrom is called when entering the from production.
	EnterFrom(c *FromContext)

	// EnterFrom_table is called when entering the from_table production.
	EnterFrom_table(c *From_tableContext)

	// EnterFrom_function is called when entering the from_function production.
	EnterFrom_function(c *From_functionContext)

	// EnterFrom_table_or_subquery is called when entering the from_table_or_subquery production.
	EnterFrom_table_or_subquery(c *From_table_or_subqueryContext)

	// EnterFrom_select is called when entering the from_select production.
	EnterFrom_select(c *From_selectContext)

	// EnterTable_or_subquery is called when entering the table_or_subquery production.
	EnterTable_or_subquery(c *Table_or_subqueryContext)

	// EnterResult_column is called when entering the result_column production.
	EnterResult_column(c *Result_columnContext)

	// EnterJoin_operator is called when entering the join_operator production.
	EnterJoin_operator(c *Join_operatorContext)

	// EnterJoin_constraint is called when entering the join_constraint production.
	EnterJoin_constraint(c *Join_constraintContext)

	// EnterCompound_operator is called when entering the compound_operator production.
	EnterCompound_operator(c *Compound_operatorContext)

	// EnterUpdate_stmt is called when entering the update_stmt production.
	EnterUpdate_stmt(c *Update_stmtContext)

	// EnterUpdate_column_set is called when entering the update_column_set production.
	EnterUpdate_column_set(c *Update_column_setContext)

	// EnterColumn_name_list is called when entering the column_name_list production.
	EnterColumn_name_list(c *Column_name_listContext)

	// EnterUpdate_stmt_limited is called when entering the update_stmt_limited production.
	EnterUpdate_stmt_limited(c *Update_stmt_limitedContext)

	// EnterQualified_table_name is called when entering the qualified_table_name production.
	EnterQualified_table_name(c *Qualified_table_nameContext)

	// EnterVacuum_stmt is called when entering the vacuum_stmt production.
	EnterVacuum_stmt(c *Vacuum_stmtContext)

	// EnterFilter_clause is called when entering the filter_clause production.
	EnterFilter_clause(c *Filter_clauseContext)

	// EnterWindow_defn is called when entering the window_defn production.
	EnterWindow_defn(c *Window_defnContext)

	// EnterOver_clause is called when entering the over_clause production.
	EnterOver_clause(c *Over_clauseContext)

	// EnterFrame_spec is called when entering the frame_spec production.
	EnterFrame_spec(c *Frame_specContext)

	// EnterFrame_clause is called when entering the frame_clause production.
	EnterFrame_clause(c *Frame_clauseContext)

	// EnterSimple_function_invocation is called when entering the simple_function_invocation production.
	EnterSimple_function_invocation(c *Simple_function_invocationContext)

	// EnterAggregate_function_invocation is called when entering the aggregate_function_invocation production.
	EnterAggregate_function_invocation(c *Aggregate_function_invocationContext)

	// EnterWindow_function_invocation is called when entering the window_function_invocation production.
	EnterWindow_function_invocation(c *Window_function_invocationContext)

	// EnterCommon_table_stmt is called when entering the common_table_stmt production.
	EnterCommon_table_stmt(c *Common_table_stmtContext)

	// EnterOrder_by_stmt is called when entering the order_by_stmt production.
	EnterOrder_by_stmt(c *Order_by_stmtContext)

	// EnterLimit_stmt is called when entering the limit_stmt production.
	EnterLimit_stmt(c *Limit_stmtContext)

	// EnterOrdering_term is called when entering the ordering_term production.
	EnterOrdering_term(c *Ordering_termContext)

	// EnterAsc_desc is called when entering the asc_desc production.
	EnterAsc_desc(c *Asc_descContext)

	// EnterFrame_left is called when entering the frame_left production.
	EnterFrame_left(c *Frame_leftContext)

	// EnterFrame_right is called when entering the frame_right production.
	EnterFrame_right(c *Frame_rightContext)

	// EnterFrame_single is called when entering the frame_single production.
	EnterFrame_single(c *Frame_singleContext)

	// EnterWindow_function is called when entering the window_function production.
	EnterWindow_function(c *Window_functionContext)

	// EnterOffset is called when entering the offset production.
	EnterOffset(c *OffsetContext)

	// EnterDefault_value is called when entering the default_value production.
	EnterDefault_value(c *Default_valueContext)

	// EnterPartition_by is called when entering the partition_by production.
	EnterPartition_by(c *Partition_byContext)

	// EnterOrder_by_expr is called when entering the order_by_expr production.
	EnterOrder_by_expr(c *Order_by_exprContext)

	// EnterOrder_by_expr_asc_desc is called when entering the order_by_expr_asc_desc production.
	EnterOrder_by_expr_asc_desc(c *Order_by_expr_asc_descContext)

	// EnterExpr_asc_desc is called when entering the expr_asc_desc production.
	EnterExpr_asc_desc(c *Expr_asc_descContext)

	// EnterInitial_select is called when entering the initial_select production.
	EnterInitial_select(c *Initial_selectContext)

	// EnterRecursive_select is called when entering the recursive_select production.
	EnterRecursive_select(c *Recursive_selectContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterError_message is called when entering the error_message production.
	EnterError_message(c *Error_messageContext)

	// EnterModule_argument is called when entering the module_argument production.
	EnterModule_argument(c *Module_argumentContext)

	// EnterColumn_alias is called when entering the column_alias production.
	EnterColumn_alias(c *Column_aliasContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterDatabase_name is called when entering the database_name production.
	EnterDatabase_name(c *Database_nameContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterSchema_name is called when entering the schema_name production.
	EnterSchema_name(c *Schema_nameContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterTable_or_index_name is called when entering the table_or_index_name production.
	EnterTable_or_index_name(c *Table_or_index_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterCollation_name is called when entering the collation_name production.
	EnterCollation_name(c *Collation_nameContext)

	// EnterForeign_table is called when entering the foreign_table production.
	EnterForeign_table(c *Foreign_tableContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterTrigger_name is called when entering the trigger_name production.
	EnterTrigger_name(c *Trigger_nameContext)

	// EnterView_name is called when entering the view_name production.
	EnterView_name(c *View_nameContext)

	// EnterModule_name is called when entering the module_name production.
	EnterModule_name(c *Module_nameContext)

	// EnterPragma_name is called when entering the pragma_name production.
	EnterPragma_name(c *Pragma_nameContext)

	// EnterSavepoint_name is called when entering the savepoint_name production.
	EnterSavepoint_name(c *Savepoint_nameContext)

	// EnterTable_alias is called when entering the table_alias production.
	EnterTable_alias(c *Table_aliasContext)

	// EnterTransaction_name is called when entering the transaction_name production.
	EnterTransaction_name(c *Transaction_nameContext)

	// EnterWindow_name is called when entering the window_name production.
	EnterWindow_name(c *Window_nameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterBase_window_name is called when entering the base_window_name production.
	EnterBase_window_name(c *Base_window_nameContext)

	// EnterSimple_func is called when entering the simple_func production.
	EnterSimple_func(c *Simple_funcContext)

	// EnterAggregate_func is called when entering the aggregate_func production.
	EnterAggregate_func(c *Aggregate_funcContext)

	// EnterTable_function_name is called when entering the table_function_name production.
	EnterTable_function_name(c *Table_function_nameContext)

	// EnterAny_name is called when entering the any_name production.
	EnterAny_name(c *Any_nameContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitSql_stmt_list is called when exiting the sql_stmt_list production.
	ExitSql_stmt_list(c *Sql_stmt_listContext)

	// ExitSql_stmt is called when exiting the sql_stmt production.
	ExitSql_stmt(c *Sql_stmtContext)

	// ExitAlter_database_stmt is called when exiting the alter_database_stmt production.
	ExitAlter_database_stmt(c *Alter_database_stmtContext)

	// ExitRename_database_to is called when exiting the rename_database_to production.
	ExitRename_database_to(c *Rename_database_toContext)

	// ExitAlter_table_stmt is called when exiting the alter_table_stmt production.
	ExitAlter_table_stmt(c *Alter_table_stmtContext)

	// ExitRename_table_to is called when exiting the rename_table_to production.
	ExitRename_table_to(c *Rename_table_toContext)

	// ExitRename_table_colum is called when exiting the rename_table_colum production.
	ExitRename_table_colum(c *Rename_table_columContext)

	// ExitAdd_table_column is called when exiting the add_table_column production.
	ExitAdd_table_column(c *Add_table_columnContext)

	// ExitAdd_table_index is called when exiting the add_table_index production.
	ExitAdd_table_index(c *Add_table_indexContext)

	// ExitDrop_table_column is called when exiting the drop_table_column production.
	ExitDrop_table_column(c *Drop_table_columnContext)

	// ExitAnalyze_stmt is called when exiting the analyze_stmt production.
	ExitAnalyze_stmt(c *Analyze_stmtContext)

	// ExitAttach_stmt is called when exiting the attach_stmt production.
	ExitAttach_stmt(c *Attach_stmtContext)

	// ExitBegin_stmt is called when exiting the begin_stmt production.
	ExitBegin_stmt(c *Begin_stmtContext)

	// ExitCommit_stmt is called when exiting the commit_stmt production.
	ExitCommit_stmt(c *Commit_stmtContext)

	// ExitRollback_stmt is called when exiting the rollback_stmt production.
	ExitRollback_stmt(c *Rollback_stmtContext)

	// ExitSavepoint_stmt is called when exiting the savepoint_stmt production.
	ExitSavepoint_stmt(c *Savepoint_stmtContext)

	// ExitRelease_stmt is called when exiting the release_stmt production.
	ExitRelease_stmt(c *Release_stmtContext)

	// ExitCreate_index_stmt is called when exiting the create_index_stmt production.
	ExitCreate_index_stmt(c *Create_index_stmtContext)

	// ExitIndexed_column is called when exiting the indexed_column production.
	ExitIndexed_column(c *Indexed_columnContext)

	// ExitCreate_database_stmt is called when exiting the create_database_stmt production.
	ExitCreate_database_stmt(c *Create_database_stmtContext)

	// ExitIf_not_exists is called when exiting the if_not_exists production.
	ExitIf_not_exists(c *If_not_existsContext)

	// ExitCreate_table_stmt is called when exiting the create_table_stmt production.
	ExitCreate_table_stmt(c *Create_table_stmtContext)

	// ExitTable_options is called when exiting the table_options production.
	ExitTable_options(c *Table_optionsContext)

	// ExitMysql_table_options is called when exiting the mysql_table_options production.
	ExitMysql_table_options(c *Mysql_table_optionsContext)

	// ExitColumn_def is called when exiting the column_def production.
	ExitColumn_def(c *Column_defContext)

	// ExitWith_param is called when exiting the with_param production.
	ExitWith_param(c *With_paramContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitColumn_constraint is called when exiting the column_constraint production.
	ExitColumn_constraint(c *Column_constraintContext)

	// ExitPrimary_key_constraint is called when exiting the primary_key_constraint production.
	ExitPrimary_key_constraint(c *Primary_key_constraintContext)

	// ExitIndex_constraint is called when exiting the index_constraint production.
	ExitIndex_constraint(c *Index_constraintContext)

	// ExitSigned_number is called when exiting the signed_number production.
	ExitSigned_number(c *Signed_numberContext)

	// ExitTable_constraint is called when exiting the table_constraint production.
	ExitTable_constraint(c *Table_constraintContext)

	// ExitPrimary_key_def is called when exiting the primary_key_def production.
	ExitPrimary_key_def(c *Primary_key_defContext)

	// ExitUnique_key_def is called when exiting the unique_key_def production.
	ExitUnique_key_def(c *Unique_key_defContext)

	// ExitForeign_key_clause is called when exiting the foreign_key_clause production.
	ExitForeign_key_clause(c *Foreign_key_clauseContext)

	// ExitConflict_clause is called when exiting the conflict_clause production.
	ExitConflict_clause(c *Conflict_clauseContext)

	// ExitCreate_trigger_stmt is called when exiting the create_trigger_stmt production.
	ExitCreate_trigger_stmt(c *Create_trigger_stmtContext)

	// ExitCreate_view_stmt is called when exiting the create_view_stmt production.
	ExitCreate_view_stmt(c *Create_view_stmtContext)

	// ExitCreate_virtual_table_stmt is called when exiting the create_virtual_table_stmt production.
	ExitCreate_virtual_table_stmt(c *Create_virtual_table_stmtContext)

	// ExitWith_clause is called when exiting the with_clause production.
	ExitWith_clause(c *With_clauseContext)

	// ExitCte_table_name is called when exiting the cte_table_name production.
	ExitCte_table_name(c *Cte_table_nameContext)

	// ExitRecursive_cte is called when exiting the recursive_cte production.
	ExitRecursive_cte(c *Recursive_cteContext)

	// ExitCommon_table_expression is called when exiting the common_table_expression production.
	ExitCommon_table_expression(c *Common_table_expressionContext)

	// ExitDelete_stmt is called when exiting the delete_stmt production.
	ExitDelete_stmt(c *Delete_stmtContext)

	// ExitDelete_stmt_limited is called when exiting the delete_stmt_limited production.
	ExitDelete_stmt_limited(c *Delete_stmt_limitedContext)

	// ExitDetach_stmt is called when exiting the detach_stmt production.
	ExitDetach_stmt(c *Detach_stmtContext)

	// ExitDrop_stmt is called when exiting the drop_stmt production.
	ExitDrop_stmt(c *Drop_stmtContext)

	// ExitDrop_database_stmt is called when exiting the drop_database_stmt production.
	ExitDrop_database_stmt(c *Drop_database_stmtContext)

	// ExitDrop_index_stmt is called when exiting the drop_index_stmt production.
	ExitDrop_index_stmt(c *Drop_index_stmtContext)

	// ExitDrop_table_stmt is called when exiting the drop_table_stmt production.
	ExitDrop_table_stmt(c *Drop_table_stmtContext)

	// ExitDrop_trigger_stmt is called when exiting the drop_trigger_stmt production.
	ExitDrop_trigger_stmt(c *Drop_trigger_stmtContext)

	// ExitDrop_view_stmt is called when exiting the drop_view_stmt production.
	ExitDrop_view_stmt(c *Drop_view_stmtContext)

	// ExitIf_exists is called when exiting the if_exists production.
	ExitIf_exists(c *If_existsContext)

	// ExitUse_stmt is called when exiting the use_stmt production.
	ExitUse_stmt(c *Use_stmtContext)

	// ExitCopy_stmt is called when exiting the copy_stmt production.
	ExitCopy_stmt(c *Copy_stmtContext)

	// ExitCopy_column_list is called when exiting the copy_column_list production.
	ExitCopy_column_list(c *Copy_column_listContext)

	// ExitCopy_format is called when exiting the copy_format production.
	ExitCopy_format(c *Copy_formatContext)

	// ExitSource_name is called when exiting the source_name production.
	ExitSource_name(c *Source_nameContext)

	// ExitTruncate_stmt is called when exiting the truncate_stmt production.
	ExitTruncate_stmt(c *Truncate_stmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitComparison_expr is called when exiting the comparison_expr production.
	ExitComparison_expr(c *Comparison_exprContext)

	// ExitAnd_expr is called when exiting the and_expr production.
	ExitAnd_expr(c *And_exprContext)

	// ExitOr_expr is called when exiting the or_expr production.
	ExitOr_expr(c *Or_exprContext)

	// ExitArithmetic_expr is called when exiting the arithmetic_expr production.
	ExitArithmetic_expr(c *Arithmetic_exprContext)

	// ExitRaise_function is called when exiting the raise_function production.
	ExitRaise_function(c *Raise_functionContext)

	// ExitLiteral_value is called when exiting the literal_value production.
	ExitLiteral_value(c *Literal_valueContext)

	// ExitBind_param is called when exiting the bind_param production.
	ExitBind_param(c *Bind_paramContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitValue_row is called when exiting the value_row production.
	ExitValue_row(c *Value_rowContext)

	// ExitValues_clause is called when exiting the values_clause production.
	ExitValues_clause(c *Values_clauseContext)

	// ExitInsert_stmt is called when exiting the insert_stmt production.
	ExitInsert_stmt(c *Insert_stmtContext)

	// ExitReturning_clause is called when exiting the returning_clause production.
	ExitReturning_clause(c *Returning_clauseContext)

	// ExitUpsert_clause is called when exiting the upsert_clause production.
	ExitUpsert_clause(c *Upsert_clauseContext)

	// ExitPragma_stmt is called when exiting the pragma_stmt production.
	ExitPragma_stmt(c *Pragma_stmtContext)

	// ExitPragma_value is called when exiting the pragma_value production.
	ExitPragma_value(c *Pragma_valueContext)

	// ExitReindex_stmt is called when exiting the reindex_stmt production.
	ExitReindex_stmt(c *Reindex_stmtContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitJoin_clause is called when exiting the join_clause production.
	ExitJoin_clause(c *Join_clauseContext)

	// ExitSelect_core is called when exiting the select_core production.
	ExitSelect_core(c *Select_coreContext)

	// ExitFactored_select_stmt is called when exiting the factored_select_stmt production.
	ExitFactored_select_stmt(c *Factored_select_stmtContext)

	// ExitSimple_select_stmt is called when exiting the simple_select_stmt production.
	ExitSimple_select_stmt(c *Simple_select_stmtContext)

	// ExitCompound_select_stmt is called when exiting the compound_select_stmt production.
	ExitCompound_select_stmt(c *Compound_select_stmtContext)

	// ExitFrom is called when exiting the from production.
	ExitFrom(c *FromContext)

	// ExitFrom_table is called when exiting the from_table production.
	ExitFrom_table(c *From_tableContext)

	// ExitFrom_function is called when exiting the from_function production.
	ExitFrom_function(c *From_functionContext)

	// ExitFrom_table_or_subquery is called when exiting the from_table_or_subquery production.
	ExitFrom_table_or_subquery(c *From_table_or_subqueryContext)

	// ExitFrom_select is called when exiting the from_select production.
	ExitFrom_select(c *From_selectContext)

	// ExitTable_or_subquery is called when exiting the table_or_subquery production.
	ExitTable_or_subquery(c *Table_or_subqueryContext)

	// ExitResult_column is called when exiting the result_column production.
	ExitResult_column(c *Result_columnContext)

	// ExitJoin_operator is called when exiting the join_operator production.
	ExitJoin_operator(c *Join_operatorContext)

	// ExitJoin_constraint is called when exiting the join_constraint production.
	ExitJoin_constraint(c *Join_constraintContext)

	// ExitCompound_operator is called when exiting the compound_operator production.
	ExitCompound_operator(c *Compound_operatorContext)

	// ExitUpdate_stmt is called when exiting the update_stmt production.
	ExitUpdate_stmt(c *Update_stmtContext)

	// ExitUpdate_column_set is called when exiting the update_column_set production.
	ExitUpdate_column_set(c *Update_column_setContext)

	// ExitColumn_name_list is called when exiting the column_name_list production.
	ExitColumn_name_list(c *Column_name_listContext)

	// ExitUpdate_stmt_limited is called when exiting the update_stmt_limited production.
	ExitUpdate_stmt_limited(c *Update_stmt_limitedContext)

	// ExitQualified_table_name is called when exiting the qualified_table_name production.
	ExitQualified_table_name(c *Qualified_table_nameContext)

	// ExitVacuum_stmt is called when exiting the vacuum_stmt production.
	ExitVacuum_stmt(c *Vacuum_stmtContext)

	// ExitFilter_clause is called when exiting the filter_clause production.
	ExitFilter_clause(c *Filter_clauseContext)

	// ExitWindow_defn is called when exiting the window_defn production.
	ExitWindow_defn(c *Window_defnContext)

	// ExitOver_clause is called when exiting the over_clause production.
	ExitOver_clause(c *Over_clauseContext)

	// ExitFrame_spec is called when exiting the frame_spec production.
	ExitFrame_spec(c *Frame_specContext)

	// ExitFrame_clause is called when exiting the frame_clause production.
	ExitFrame_clause(c *Frame_clauseContext)

	// ExitSimple_function_invocation is called when exiting the simple_function_invocation production.
	ExitSimple_function_invocation(c *Simple_function_invocationContext)

	// ExitAggregate_function_invocation is called when exiting the aggregate_function_invocation production.
	ExitAggregate_function_invocation(c *Aggregate_function_invocationContext)

	// ExitWindow_function_invocation is called when exiting the window_function_invocation production.
	ExitWindow_function_invocation(c *Window_function_invocationContext)

	// ExitCommon_table_stmt is called when exiting the common_table_stmt production.
	ExitCommon_table_stmt(c *Common_table_stmtContext)

	// ExitOrder_by_stmt is called when exiting the order_by_stmt production.
	ExitOrder_by_stmt(c *Order_by_stmtContext)

	// ExitLimit_stmt is called when exiting the limit_stmt production.
	ExitLimit_stmt(c *Limit_stmtContext)

	// ExitOrdering_term is called when exiting the ordering_term production.
	ExitOrdering_term(c *Ordering_termContext)

	// ExitAsc_desc is called when exiting the asc_desc production.
	ExitAsc_desc(c *Asc_descContext)

	// ExitFrame_left is called when exiting the frame_left production.
	ExitFrame_left(c *Frame_leftContext)

	// ExitFrame_right is called when exiting the frame_right production.
	ExitFrame_right(c *Frame_rightContext)

	// ExitFrame_single is called when exiting the frame_single production.
	ExitFrame_single(c *Frame_singleContext)

	// ExitWindow_function is called when exiting the window_function production.
	ExitWindow_function(c *Window_functionContext)

	// ExitOffset is called when exiting the offset production.
	ExitOffset(c *OffsetContext)

	// ExitDefault_value is called when exiting the default_value production.
	ExitDefault_value(c *Default_valueContext)

	// ExitPartition_by is called when exiting the partition_by production.
	ExitPartition_by(c *Partition_byContext)

	// ExitOrder_by_expr is called when exiting the order_by_expr production.
	ExitOrder_by_expr(c *Order_by_exprContext)

	// ExitOrder_by_expr_asc_desc is called when exiting the order_by_expr_asc_desc production.
	ExitOrder_by_expr_asc_desc(c *Order_by_expr_asc_descContext)

	// ExitExpr_asc_desc is called when exiting the expr_asc_desc production.
	ExitExpr_asc_desc(c *Expr_asc_descContext)

	// ExitInitial_select is called when exiting the initial_select production.
	ExitInitial_select(c *Initial_selectContext)

	// ExitRecursive_select is called when exiting the recursive_select production.
	ExitRecursive_select(c *Recursive_selectContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitError_message is called when exiting the error_message production.
	ExitError_message(c *Error_messageContext)

	// ExitModule_argument is called when exiting the module_argument production.
	ExitModule_argument(c *Module_argumentContext)

	// ExitColumn_alias is called when exiting the column_alias production.
	ExitColumn_alias(c *Column_aliasContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitDatabase_name is called when exiting the database_name production.
	ExitDatabase_name(c *Database_nameContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitSchema_name is called when exiting the schema_name production.
	ExitSchema_name(c *Schema_nameContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitTable_or_index_name is called when exiting the table_or_index_name production.
	ExitTable_or_index_name(c *Table_or_index_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitCollation_name is called when exiting the collation_name production.
	ExitCollation_name(c *Collation_nameContext)

	// ExitForeign_table is called when exiting the foreign_table production.
	ExitForeign_table(c *Foreign_tableContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitTrigger_name is called when exiting the trigger_name production.
	ExitTrigger_name(c *Trigger_nameContext)

	// ExitView_name is called when exiting the view_name production.
	ExitView_name(c *View_nameContext)

	// ExitModule_name is called when exiting the module_name production.
	ExitModule_name(c *Module_nameContext)

	// ExitPragma_name is called when exiting the pragma_name production.
	ExitPragma_name(c *Pragma_nameContext)

	// ExitSavepoint_name is called when exiting the savepoint_name production.
	ExitSavepoint_name(c *Savepoint_nameContext)

	// ExitTable_alias is called when exiting the table_alias production.
	ExitTable_alias(c *Table_aliasContext)

	// ExitTransaction_name is called when exiting the transaction_name production.
	ExitTransaction_name(c *Transaction_nameContext)

	// ExitWindow_name is called when exiting the window_name production.
	ExitWindow_name(c *Window_nameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitBase_window_name is called when exiting the base_window_name production.
	ExitBase_window_name(c *Base_window_nameContext)

	// ExitSimple_func is called when exiting the simple_func production.
	ExitSimple_func(c *Simple_funcContext)

	// ExitAggregate_func is called when exiting the aggregate_func production.
	ExitAggregate_func(c *Aggregate_funcContext)

	// ExitTable_function_name is called when exiting the table_function_name production.
	ExitTable_function_name(c *Table_function_nameContext)

	// ExitAny_name is called when exiting the any_name production.
	ExitAny_name(c *Any_nameContext)
}
