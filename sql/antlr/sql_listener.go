// Code generated from SQL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package antlr // SQL
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// SQLListener is a complete listener for a parse tree produced by SQLParser.
type SQLListener interface {
	antlr.ParseTreeListener

	// EnterQueries is called when entering the queries production.
	EnterQueries(c *QueriesContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterShowStmt is called when entering the showStmt production.
	EnterShowStmt(c *ShowStmtContext)

	// EnterUseStmt is called when entering the useStmt production.
	EnterUseStmt(c *UseStmtContext)

	// EnterCreateStmt is called when entering the createStmt production.
	EnterCreateStmt(c *CreateStmtContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterCreateCollection is called when entering the createCollection production.
	EnterCreateCollection(c *CreateCollectionContext)

	// EnterDrop_collectionStmt is called when entering the drop_collectionStmt production.
	EnterDrop_collectionStmt(c *Drop_collectionStmtContext)

	// EnterCreate_indexStmt is called when entering the create_indexStmt production.
	EnterCreate_indexStmt(c *Create_indexStmtContext)

	// EnterDrop_indexStmt is called when entering the drop_indexStmt production.
	EnterDrop_indexStmt(c *Drop_indexStmtContext)

	// EnterSelectStmt is called when entering the selectStmt production.
	EnterSelectStmt(c *SelectStmtContext)

	// EnterColumns is called when entering the columns production.
	EnterColumns(c *ColumnsContext)

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

	// EnterInsertStmt is called when entering the insertStmt production.
	EnterInsertStmt(c *InsertStmtContext)

	// EnterInsert_columns_section is called when entering the insert_columns_section production.
	EnterInsert_columns_section(c *Insert_columns_sectionContext)

	// EnterInsert_values_section is called when entering the insert_values_section production.
	EnterInsert_values_section(c *Insert_values_sectionContext)

	// EnterUpdateStmt is called when entering the updateStmt production.
	EnterUpdateStmt(c *UpdateStmtContext)

	// EnterProperty_section is called when entering the property_section production.
	EnterProperty_section(c *Property_sectionContext)

	// EnterDeleteStmt is called when entering the deleteStmt production.
	EnterDeleteStmt(c *DeleteStmtContext)

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

	// EnterCollectionName is called when entering the collectionName production.
	EnterCollectionName(c *CollectionNameContext)

	// EnterDatabaseName is called when entering the databaseName production.
	EnterDatabaseName(c *DatabaseNameContext)

	// EnterColumn is called when entering the column production.
	EnterColumn(c *ColumnContext)

	// EnterIndex_section is called when entering the index_section production.
	EnterIndex_section(c *Index_sectionContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterWhere_section is called when entering the where_section production.
	EnterWhere_section(c *Where_sectionContext)

	// EnterUnreserved_keyword is called when entering the unreserved_keyword production.
	EnterUnreserved_keyword(c *Unreserved_keywordContext)

	// ExitQueries is called when exiting the queries production.
	ExitQueries(c *QueriesContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitShowStmt is called when exiting the showStmt production.
	ExitShowStmt(c *ShowStmtContext)

	// ExitUseStmt is called when exiting the useStmt production.
	ExitUseStmt(c *UseStmtContext)

	// ExitCreateStmt is called when exiting the createStmt production.
	ExitCreateStmt(c *CreateStmtContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitCreateCollection is called when exiting the createCollection production.
	ExitCreateCollection(c *CreateCollectionContext)

	// ExitDrop_collectionStmt is called when exiting the drop_collectionStmt production.
	ExitDrop_collectionStmt(c *Drop_collectionStmtContext)

	// ExitCreate_indexStmt is called when exiting the create_indexStmt production.
	ExitCreate_indexStmt(c *Create_indexStmtContext)

	// ExitDrop_indexStmt is called when exiting the drop_indexStmt production.
	ExitDrop_indexStmt(c *Drop_indexStmtContext)

	// ExitSelectStmt is called when exiting the selectStmt production.
	ExitSelectStmt(c *SelectStmtContext)

	// ExitColumns is called when exiting the columns production.
	ExitColumns(c *ColumnsContext)

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

	// ExitInsertStmt is called when exiting the insertStmt production.
	ExitInsertStmt(c *InsertStmtContext)

	// ExitInsert_columns_section is called when exiting the insert_columns_section production.
	ExitInsert_columns_section(c *Insert_columns_sectionContext)

	// ExitInsert_values_section is called when exiting the insert_values_section production.
	ExitInsert_values_section(c *Insert_values_sectionContext)

	// ExitUpdateStmt is called when exiting the updateStmt production.
	ExitUpdateStmt(c *UpdateStmtContext)

	// ExitProperty_section is called when exiting the property_section production.
	ExitProperty_section(c *Property_sectionContext)

	// ExitDeleteStmt is called when exiting the deleteStmt production.
	ExitDeleteStmt(c *DeleteStmtContext)

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

	// ExitCollectionName is called when exiting the collectionName production.
	ExitCollectionName(c *CollectionNameContext)

	// ExitDatabaseName is called when exiting the databaseName production.
	ExitDatabaseName(c *DatabaseNameContext)

	// ExitColumn is called when exiting the column production.
	ExitColumn(c *ColumnContext)

	// ExitIndex_section is called when exiting the index_section production.
	ExitIndex_section(c *Index_sectionContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitWhere_section is called when exiting the where_section production.
	ExitWhere_section(c *Where_sectionContext)

	// ExitUnreserved_keyword is called when exiting the unreserved_keyword production.
	ExitUnreserved_keyword(c *Unreserved_keywordContext)
}
