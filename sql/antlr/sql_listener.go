// Code generated from SQL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package antlr // SQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQLListener is a complete listener for a parse tree produced by SQLParser.
type SQLListener interface {
	antlr.ParseTreeListener

	// EnterQueries is called when entering the queries production.
	EnterQueries(c *QueriesContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterShow_stmt is called when entering the show_stmt production.
	EnterShow_stmt(c *Show_stmtContext)

	// EnterUse_stmt is called when entering the use_stmt production.
	EnterUse_stmt(c *Use_stmtContext)

	// EnterCreate_collection_stmt is called when entering the create_collection_stmt production.
	EnterCreate_collection_stmt(c *Create_collection_stmtContext)

	// EnterDrop_collection_stmt is called when entering the drop_collection_stmt production.
	EnterDrop_collection_stmt(c *Drop_collection_stmtContext)

	// EnterCreate_index_stmt is called when entering the create_index_stmt production.
	EnterCreate_index_stmt(c *Create_index_stmtContext)

	// EnterDrop_index_stmt is called when entering the drop_index_stmt production.
	EnterDrop_index_stmt(c *Drop_index_stmtContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterSelect_core is called when entering the select_core production.
	EnterSelect_core(c *Select_coreContext)

	// EnterResult_column_section is called when entering the result_column_section production.
	EnterResult_column_section(c *Result_column_sectionContext)

	// EnterFrom_section is called when entering the from_section production.
	EnterFrom_section(c *From_sectionContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterData_source is called when entering the data_source production.
	EnterData_source(c *Data_sourceContext)

	// EnterGrouping_section is called when entering the grouping_section production.
	EnterGrouping_section(c *Grouping_sectionContext)

	// EnterHaving_section is called when entering the having_section production.
	EnterHaving_section(c *Having_sectionContext)

	// EnterSorting_section is called when entering the sorting_section production.
	EnterSorting_section(c *Sorting_sectionContext)

	// EnterSorting_item is called when entering the sorting_item production.
	EnterSorting_item(c *Sorting_itemContext)

	// EnterSorting_specification is called when entering the sorting_specification production.
	EnterSorting_specification(c *Sorting_specificationContext)

	// EnterLimit_section is called when entering the limit_section production.
	EnterLimit_section(c *Limit_sectionContext)

	// EnterOffset_section is called when entering the offset_section production.
	EnterOffset_section(c *Offset_sectionContext)

	// EnterInsert_stmt is called when entering the insert_stmt production.
	EnterInsert_stmt(c *Insert_stmtContext)

	// EnterInsert_columns_section is called when entering the insert_columns_section production.
	EnterInsert_columns_section(c *Insert_columns_sectionContext)

	// EnterInsert_values_section is called when entering the insert_values_section production.
	EnterInsert_values_section(c *Insert_values_sectionContext)

	// EnterUpdate_stmt is called when entering the update_stmt production.
	EnterUpdate_stmt(c *Update_stmtContext)

	// EnterProperty_section is called when entering the property_section production.
	EnterProperty_section(c *Property_sectionContext)

	// EnterDelete_stmt is called when entering the delete_stmt production.
	EnterDelete_stmt(c *Delete_stmtContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterExpression_literal is called when entering the expression_literal production.
	EnterExpression_literal(c *Expression_literalContext)

	// EnterExpression_literal_value is called when entering the expression_literal_value production.
	EnterExpression_literal_value(c *Expression_literal_valueContext)

	// EnterExpression_dictionary is called when entering the expression_dictionary production.
	EnterExpression_dictionary(c *Expression_dictionaryContext)

	// EnterDictionary_literal is called when entering the dictionary_literal production.
	EnterDictionary_literal(c *Dictionary_literalContext)

	// EnterExpression_array is called when entering the expression_array production.
	EnterExpression_array(c *Expression_arrayContext)

	// EnterArray_literal is called when entering the array_literal production.
	EnterArray_literal(c *Array_literalContext)

	// EnterExpression_logic_operator is called when entering the expression_logic_operator production.
	EnterExpression_logic_operator(c *Expression_logic_operatorContext)

	// EnterExpression_binary_operator is called when entering the expression_binary_operator production.
	EnterExpression_binary_operator(c *Expression_binary_operatorContext)

	// EnterExpression_function is called when entering the expression_function production.
	EnterExpression_function(c *Expression_functionContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterFunction_value is called when entering the function_value production.
	EnterFunction_value(c *Function_valueContext)

	// EnterExpression_operator is called when entering the expression_operator production.
	EnterExpression_operator(c *Expression_operatorContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

	// EnterLogical_operator is called when entering the logical_operator production.
	EnterLogical_operator(c *Logical_operatorContext)

	// EnterProperty_literal is called when entering the property_literal production.
	EnterProperty_literal(c *Property_literalContext)

	// EnterInteger_literal is called when entering the integer_literal production.
	EnterInteger_literal(c *Integer_literalContext)

	// EnterReal_literal is called when entering the real_literal production.
	EnterReal_literal(c *Real_literalContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterTrue_literal is called when entering the true_literal production.
	EnterTrue_literal(c *True_literalContext)

	// EnterFalse_literal is called when entering the false_literal production.
	EnterFalse_literal(c *False_literalContext)

	// EnterSync_operator is called when entering the sync_operator production.
	EnterSync_operator(c *Sync_operatorContext)

	// EnterCompound_operator is called when entering the compound_operator production.
	EnterCompound_operator(c *Compound_operatorContext)

	// EnterCondition_operator is called when entering the condition_operator production.
	EnterCondition_operator(c *Condition_operatorContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterCollection_section is called when entering the collection_section production.
	EnterCollection_section(c *Collection_sectionContext)

	// EnterCollection_name is called when entering the collection_name production.
	EnterCollection_name(c *Collection_nameContext)

	// EnterColumn_section is called when entering the column_section production.
	EnterColumn_section(c *Column_sectionContext)

	// EnterIndex_section is called when entering the index_section production.
	EnterIndex_section(c *Index_sectionContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterWhere_section is called when entering the where_section production.
	EnterWhere_section(c *Where_sectionContext)

	// ExitQueries is called when exiting the queries production.
	ExitQueries(c *QueriesContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitShow_stmt is called when exiting the show_stmt production.
	ExitShow_stmt(c *Show_stmtContext)

	// ExitUse_stmt is called when exiting the use_stmt production.
	ExitUse_stmt(c *Use_stmtContext)

	// ExitCreate_collection_stmt is called when exiting the create_collection_stmt production.
	ExitCreate_collection_stmt(c *Create_collection_stmtContext)

	// ExitDrop_collection_stmt is called when exiting the drop_collection_stmt production.
	ExitDrop_collection_stmt(c *Drop_collection_stmtContext)

	// ExitCreate_index_stmt is called when exiting the create_index_stmt production.
	ExitCreate_index_stmt(c *Create_index_stmtContext)

	// ExitDrop_index_stmt is called when exiting the drop_index_stmt production.
	ExitDrop_index_stmt(c *Drop_index_stmtContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitSelect_core is called when exiting the select_core production.
	ExitSelect_core(c *Select_coreContext)

	// ExitResult_column_section is called when exiting the result_column_section production.
	ExitResult_column_section(c *Result_column_sectionContext)

	// ExitFrom_section is called when exiting the from_section production.
	ExitFrom_section(c *From_sectionContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitData_source is called when exiting the data_source production.
	ExitData_source(c *Data_sourceContext)

	// ExitGrouping_section is called when exiting the grouping_section production.
	ExitGrouping_section(c *Grouping_sectionContext)

	// ExitHaving_section is called when exiting the having_section production.
	ExitHaving_section(c *Having_sectionContext)

	// ExitSorting_section is called when exiting the sorting_section production.
	ExitSorting_section(c *Sorting_sectionContext)

	// ExitSorting_item is called when exiting the sorting_item production.
	ExitSorting_item(c *Sorting_itemContext)

	// ExitSorting_specification is called when exiting the sorting_specification production.
	ExitSorting_specification(c *Sorting_specificationContext)

	// ExitLimit_section is called when exiting the limit_section production.
	ExitLimit_section(c *Limit_sectionContext)

	// ExitOffset_section is called when exiting the offset_section production.
	ExitOffset_section(c *Offset_sectionContext)

	// ExitInsert_stmt is called when exiting the insert_stmt production.
	ExitInsert_stmt(c *Insert_stmtContext)

	// ExitInsert_columns_section is called when exiting the insert_columns_section production.
	ExitInsert_columns_section(c *Insert_columns_sectionContext)

	// ExitInsert_values_section is called when exiting the insert_values_section production.
	ExitInsert_values_section(c *Insert_values_sectionContext)

	// ExitUpdate_stmt is called when exiting the update_stmt production.
	ExitUpdate_stmt(c *Update_stmtContext)

	// ExitProperty_section is called when exiting the property_section production.
	ExitProperty_section(c *Property_sectionContext)

	// ExitDelete_stmt is called when exiting the delete_stmt production.
	ExitDelete_stmt(c *Delete_stmtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitExpression_literal is called when exiting the expression_literal production.
	ExitExpression_literal(c *Expression_literalContext)

	// ExitExpression_literal_value is called when exiting the expression_literal_value production.
	ExitExpression_literal_value(c *Expression_literal_valueContext)

	// ExitExpression_dictionary is called when exiting the expression_dictionary production.
	ExitExpression_dictionary(c *Expression_dictionaryContext)

	// ExitDictionary_literal is called when exiting the dictionary_literal production.
	ExitDictionary_literal(c *Dictionary_literalContext)

	// ExitExpression_array is called when exiting the expression_array production.
	ExitExpression_array(c *Expression_arrayContext)

	// ExitArray_literal is called when exiting the array_literal production.
	ExitArray_literal(c *Array_literalContext)

	// ExitExpression_logic_operator is called when exiting the expression_logic_operator production.
	ExitExpression_logic_operator(c *Expression_logic_operatorContext)

	// ExitExpression_binary_operator is called when exiting the expression_binary_operator production.
	ExitExpression_binary_operator(c *Expression_binary_operatorContext)

	// ExitExpression_function is called when exiting the expression_function production.
	ExitExpression_function(c *Expression_functionContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitFunction_value is called when exiting the function_value production.
	ExitFunction_value(c *Function_valueContext)

	// ExitExpression_operator is called when exiting the expression_operator production.
	ExitExpression_operator(c *Expression_operatorContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

	// ExitLogical_operator is called when exiting the logical_operator production.
	ExitLogical_operator(c *Logical_operatorContext)

	// ExitProperty_literal is called when exiting the property_literal production.
	ExitProperty_literal(c *Property_literalContext)

	// ExitInteger_literal is called when exiting the integer_literal production.
	ExitInteger_literal(c *Integer_literalContext)

	// ExitReal_literal is called when exiting the real_literal production.
	ExitReal_literal(c *Real_literalContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitTrue_literal is called when exiting the true_literal production.
	ExitTrue_literal(c *True_literalContext)

	// ExitFalse_literal is called when exiting the false_literal production.
	ExitFalse_literal(c *False_literalContext)

	// ExitSync_operator is called when exiting the sync_operator production.
	ExitSync_operator(c *Sync_operatorContext)

	// ExitCompound_operator is called when exiting the compound_operator production.
	ExitCompound_operator(c *Compound_operatorContext)

	// ExitCondition_operator is called when exiting the condition_operator production.
	ExitCondition_operator(c *Condition_operatorContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitCollection_section is called when exiting the collection_section production.
	ExitCollection_section(c *Collection_sectionContext)

	// ExitCollection_name is called when exiting the collection_name production.
	ExitCollection_name(c *Collection_nameContext)

	// ExitColumn_section is called when exiting the column_section production.
	ExitColumn_section(c *Column_sectionContext)

	// ExitIndex_section is called when exiting the index_section production.
	ExitIndex_section(c *Index_sectionContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitWhere_section is called when exiting the where_section production.
	ExitWhere_section(c *Where_sectionContext)
}