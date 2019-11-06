// Code generated from SQL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package antlr // SQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSQLListener is a complete listener for a parse tree produced by SQLParser.
type BaseSQLListener struct{}

var _ SQLListener = &BaseSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQueries is called when production queries is entered.
func (s *BaseSQLListener) EnterQueries(ctx *QueriesContext) {}

// ExitQueries is called when production queries is exited.
func (s *BaseSQLListener) ExitQueries(ctx *QueriesContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSQLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSQLListener) ExitStatement(ctx *StatementContext) {}

// EnterShowStmt is called when production showStmt is entered.
func (s *BaseSQLListener) EnterShowStmt(ctx *ShowStmtContext) {}

// ExitShowStmt is called when production showStmt is exited.
func (s *BaseSQLListener) ExitShowStmt(ctx *ShowStmtContext) {}

// EnterUseStmt is called when production useStmt is entered.
func (s *BaseSQLListener) EnterUseStmt(ctx *UseStmtContext) {}

// ExitUseStmt is called when production useStmt is exited.
func (s *BaseSQLListener) ExitUseStmt(ctx *UseStmtContext) {}

// EnterCreate_collectionStmt is called when production create_collectionStmt is entered.
func (s *BaseSQLListener) EnterCreate_collectionStmt(ctx *Create_collectionStmtContext) {}

// ExitCreate_collectionStmt is called when production create_collectionStmt is exited.
func (s *BaseSQLListener) ExitCreate_collectionStmt(ctx *Create_collectionStmtContext) {}

// EnterDrop_collectionStmt is called when production drop_collectionStmt is entered.
func (s *BaseSQLListener) EnterDrop_collectionStmt(ctx *Drop_collectionStmtContext) {}

// ExitDrop_collectionStmt is called when production drop_collectionStmt is exited.
func (s *BaseSQLListener) ExitDrop_collectionStmt(ctx *Drop_collectionStmtContext) {}

// EnterCreate_indexStmt is called when production create_indexStmt is entered.
func (s *BaseSQLListener) EnterCreate_indexStmt(ctx *Create_indexStmtContext) {}

// ExitCreate_indexStmt is called when production create_indexStmt is exited.
func (s *BaseSQLListener) ExitCreate_indexStmt(ctx *Create_indexStmtContext) {}

// EnterDrop_indexStmt is called when production drop_indexStmt is entered.
func (s *BaseSQLListener) EnterDrop_indexStmt(ctx *Drop_indexStmtContext) {}

// ExitDrop_indexStmt is called when production drop_indexStmt is exited.
func (s *BaseSQLListener) ExitDrop_indexStmt(ctx *Drop_indexStmtContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseSQLListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseSQLListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterColumns is called when production columns is entered.
func (s *BaseSQLListener) EnterColumns(ctx *ColumnsContext) {}

// ExitColumns is called when production columns is exited.
func (s *BaseSQLListener) ExitColumns(ctx *ColumnsContext) {}

// EnterFrom_section is called when production from_section is entered.
func (s *BaseSQLListener) EnterFrom_section(ctx *From_sectionContext) {}

// ExitFrom_section is called when production from_section is exited.
func (s *BaseSQLListener) ExitFrom_section(ctx *From_sectionContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseSQLListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseSQLListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterData_source is called when production data_source is entered.
func (s *BaseSQLListener) EnterData_source(ctx *Data_sourceContext) {}

// ExitData_source is called when production data_source is exited.
func (s *BaseSQLListener) ExitData_source(ctx *Data_sourceContext) {}

// EnterGrouping_section is called when production grouping_section is entered.
func (s *BaseSQLListener) EnterGrouping_section(ctx *Grouping_sectionContext) {}

// ExitGrouping_section is called when production grouping_section is exited.
func (s *BaseSQLListener) ExitGrouping_section(ctx *Grouping_sectionContext) {}

// EnterHaving_section is called when production having_section is entered.
func (s *BaseSQLListener) EnterHaving_section(ctx *Having_sectionContext) {}

// ExitHaving_section is called when production having_section is exited.
func (s *BaseSQLListener) ExitHaving_section(ctx *Having_sectionContext) {}

// EnterSorting_section is called when production sorting_section is entered.
func (s *BaseSQLListener) EnterSorting_section(ctx *Sorting_sectionContext) {}

// ExitSorting_section is called when production sorting_section is exited.
func (s *BaseSQLListener) ExitSorting_section(ctx *Sorting_sectionContext) {}

// EnterSorting_item is called when production sorting_item is entered.
func (s *BaseSQLListener) EnterSorting_item(ctx *Sorting_itemContext) {}

// ExitSorting_item is called when production sorting_item is exited.
func (s *BaseSQLListener) ExitSorting_item(ctx *Sorting_itemContext) {}

// EnterSorting_specification is called when production sorting_specification is entered.
func (s *BaseSQLListener) EnterSorting_specification(ctx *Sorting_specificationContext) {}

// ExitSorting_specification is called when production sorting_specification is exited.
func (s *BaseSQLListener) ExitSorting_specification(ctx *Sorting_specificationContext) {}

// EnterLimit_section is called when production limit_section is entered.
func (s *BaseSQLListener) EnterLimit_section(ctx *Limit_sectionContext) {}

// ExitLimit_section is called when production limit_section is exited.
func (s *BaseSQLListener) ExitLimit_section(ctx *Limit_sectionContext) {}

// EnterOffset_section is called when production offset_section is entered.
func (s *BaseSQLListener) EnterOffset_section(ctx *Offset_sectionContext) {}

// ExitOffset_section is called when production offset_section is exited.
func (s *BaseSQLListener) ExitOffset_section(ctx *Offset_sectionContext) {}

// EnterInsertStmt is called when production insertStmt is entered.
func (s *BaseSQLListener) EnterInsertStmt(ctx *InsertStmtContext) {}

// ExitInsertStmt is called when production insertStmt is exited.
func (s *BaseSQLListener) ExitInsertStmt(ctx *InsertStmtContext) {}

// EnterInsert_columns_section is called when production insert_columns_section is entered.
func (s *BaseSQLListener) EnterInsert_columns_section(ctx *Insert_columns_sectionContext) {}

// ExitInsert_columns_section is called when production insert_columns_section is exited.
func (s *BaseSQLListener) ExitInsert_columns_section(ctx *Insert_columns_sectionContext) {}

// EnterInsert_values_section is called when production insert_values_section is entered.
func (s *BaseSQLListener) EnterInsert_values_section(ctx *Insert_values_sectionContext) {}

// ExitInsert_values_section is called when production insert_values_section is exited.
func (s *BaseSQLListener) ExitInsert_values_section(ctx *Insert_values_sectionContext) {}

// EnterUpdateStmt is called when production updateStmt is entered.
func (s *BaseSQLListener) EnterUpdateStmt(ctx *UpdateStmtContext) {}

// ExitUpdateStmt is called when production updateStmt is exited.
func (s *BaseSQLListener) ExitUpdateStmt(ctx *UpdateStmtContext) {}

// EnterProperty_section is called when production property_section is entered.
func (s *BaseSQLListener) EnterProperty_section(ctx *Property_sectionContext) {}

// ExitProperty_section is called when production property_section is exited.
func (s *BaseSQLListener) ExitProperty_section(ctx *Property_sectionContext) {}

// EnterDeleteStmt is called when production deleteStmt is entered.
func (s *BaseSQLListener) EnterDeleteStmt(ctx *DeleteStmtContext) {}

// ExitDeleteStmt is called when production deleteStmt is exited.
func (s *BaseSQLListener) ExitDeleteStmt(ctx *DeleteStmtContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSQLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSQLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaseSQLListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaseSQLListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterExpression_literal is called when production expression_literal is entered.
func (s *BaseSQLListener) EnterExpression_literal(ctx *Expression_literalContext) {}

// ExitExpression_literal is called when production expression_literal is exited.
func (s *BaseSQLListener) ExitExpression_literal(ctx *Expression_literalContext) {}

// EnterExpression_literal_value is called when production expression_literal_value is entered.
func (s *BaseSQLListener) EnterExpression_literal_value(ctx *Expression_literal_valueContext) {}

// ExitExpression_literal_value is called when production expression_literal_value is exited.
func (s *BaseSQLListener) ExitExpression_literal_value(ctx *Expression_literal_valueContext) {}

// EnterExpression_dictionary is called when production expression_dictionary is entered.
func (s *BaseSQLListener) EnterExpression_dictionary(ctx *Expression_dictionaryContext) {}

// ExitExpression_dictionary is called when production expression_dictionary is exited.
func (s *BaseSQLListener) ExitExpression_dictionary(ctx *Expression_dictionaryContext) {}

// EnterDictionary_literal is called when production dictionary_literal is entered.
func (s *BaseSQLListener) EnterDictionary_literal(ctx *Dictionary_literalContext) {}

// ExitDictionary_literal is called when production dictionary_literal is exited.
func (s *BaseSQLListener) ExitDictionary_literal(ctx *Dictionary_literalContext) {}

// EnterExpression_array is called when production expression_array is entered.
func (s *BaseSQLListener) EnterExpression_array(ctx *Expression_arrayContext) {}

// ExitExpression_array is called when production expression_array is exited.
func (s *BaseSQLListener) ExitExpression_array(ctx *Expression_arrayContext) {}

// EnterArray_literal is called when production array_literal is entered.
func (s *BaseSQLListener) EnterArray_literal(ctx *Array_literalContext) {}

// ExitArray_literal is called when production array_literal is exited.
func (s *BaseSQLListener) ExitArray_literal(ctx *Array_literalContext) {}

// EnterExpression_logic_operator is called when production expression_logic_operator is entered.
func (s *BaseSQLListener) EnterExpression_logic_operator(ctx *Expression_logic_operatorContext) {}

// ExitExpression_logic_operator is called when production expression_logic_operator is exited.
func (s *BaseSQLListener) ExitExpression_logic_operator(ctx *Expression_logic_operatorContext) {}

// EnterExpression_binary_operator is called when production expression_binary_operator is entered.
func (s *BaseSQLListener) EnterExpression_binary_operator(ctx *Expression_binary_operatorContext) {}

// ExitExpression_binary_operator is called when production expression_binary_operator is exited.
func (s *BaseSQLListener) ExitExpression_binary_operator(ctx *Expression_binary_operatorContext) {}

// EnterExpression_function is called when production expression_function is entered.
func (s *BaseSQLListener) EnterExpression_function(ctx *Expression_functionContext) {}

// ExitExpression_function is called when production expression_function is exited.
func (s *BaseSQLListener) ExitExpression_function(ctx *Expression_functionContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseSQLListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseSQLListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterFunction_value is called when production function_value is entered.
func (s *BaseSQLListener) EnterFunction_value(ctx *Function_valueContext) {}

// ExitFunction_value is called when production function_value is exited.
func (s *BaseSQLListener) ExitFunction_value(ctx *Function_valueContext) {}

// EnterExpression_operator is called when production expression_operator is entered.
func (s *BaseSQLListener) EnterExpression_operator(ctx *Expression_operatorContext) {}

// ExitExpression_operator is called when production expression_operator is exited.
func (s *BaseSQLListener) ExitExpression_operator(ctx *Expression_operatorContext) {}

// EnterBinary_operator is called when production binary_operator is entered.
func (s *BaseSQLListener) EnterBinary_operator(ctx *Binary_operatorContext) {}

// ExitBinary_operator is called when production binary_operator is exited.
func (s *BaseSQLListener) ExitBinary_operator(ctx *Binary_operatorContext) {}

// EnterLogical_operator is called when production logical_operator is entered.
func (s *BaseSQLListener) EnterLogical_operator(ctx *Logical_operatorContext) {}

// ExitLogical_operator is called when production logical_operator is exited.
func (s *BaseSQLListener) ExitLogical_operator(ctx *Logical_operatorContext) {}

// EnterProperty_literal is called when production property_literal is entered.
func (s *BaseSQLListener) EnterProperty_literal(ctx *Property_literalContext) {}

// ExitProperty_literal is called when production property_literal is exited.
func (s *BaseSQLListener) ExitProperty_literal(ctx *Property_literalContext) {}

// EnterInteger_literal is called when production integer_literal is entered.
func (s *BaseSQLListener) EnterInteger_literal(ctx *Integer_literalContext) {}

// ExitInteger_literal is called when production integer_literal is exited.
func (s *BaseSQLListener) ExitInteger_literal(ctx *Integer_literalContext) {}

// EnterReal_literal is called when production real_literal is entered.
func (s *BaseSQLListener) EnterReal_literal(ctx *Real_literalContext) {}

// ExitReal_literal is called when production real_literal is exited.
func (s *BaseSQLListener) ExitReal_literal(ctx *Real_literalContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseSQLListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseSQLListener) ExitString_literal(ctx *String_literalContext) {}

// EnterTrue_literal is called when production true_literal is entered.
func (s *BaseSQLListener) EnterTrue_literal(ctx *True_literalContext) {}

// ExitTrue_literal is called when production true_literal is exited.
func (s *BaseSQLListener) ExitTrue_literal(ctx *True_literalContext) {}

// EnterFalse_literal is called when production false_literal is entered.
func (s *BaseSQLListener) EnterFalse_literal(ctx *False_literalContext) {}

// ExitFalse_literal is called when production false_literal is exited.
func (s *BaseSQLListener) ExitFalse_literal(ctx *False_literalContext) {}

// EnterSync_operator is called when production sync_operator is entered.
func (s *BaseSQLListener) EnterSync_operator(ctx *Sync_operatorContext) {}

// ExitSync_operator is called when production sync_operator is exited.
func (s *BaseSQLListener) ExitSync_operator(ctx *Sync_operatorContext) {}

// EnterCompound_operator is called when production compound_operator is entered.
func (s *BaseSQLListener) EnterCompound_operator(ctx *Compound_operatorContext) {}

// ExitCompound_operator is called when production compound_operator is exited.
func (s *BaseSQLListener) ExitCompound_operator(ctx *Compound_operatorContext) {}

// EnterCondition_operator is called when production condition_operator is entered.
func (s *BaseSQLListener) EnterCondition_operator(ctx *Condition_operatorContext) {}

// ExitCondition_operator is called when production condition_operator is exited.
func (s *BaseSQLListener) ExitCondition_operator(ctx *Condition_operatorContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseSQLListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseSQLListener) ExitProperty(ctx *PropertyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSQLListener) ExitValue(ctx *ValueContext) {}

// EnterName is called when production name is entered.
func (s *BaseSQLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseSQLListener) ExitName(ctx *NameContext) {}

// EnterCollection_section is called when production collection_section is entered.
func (s *BaseSQLListener) EnterCollection_section(ctx *Collection_sectionContext) {}

// ExitCollection_section is called when production collection_section is exited.
func (s *BaseSQLListener) ExitCollection_section(ctx *Collection_sectionContext) {}

// EnterCollection_name is called when production collection_name is entered.
func (s *BaseSQLListener) EnterCollection_name(ctx *Collection_nameContext) {}

// ExitCollection_name is called when production collection_name is exited.
func (s *BaseSQLListener) ExitCollection_name(ctx *Collection_nameContext) {}

// EnterColumn is called when production column is entered.
func (s *BaseSQLListener) EnterColumn(ctx *ColumnContext) {}

// ExitColumn is called when production column is exited.
func (s *BaseSQLListener) ExitColumn(ctx *ColumnContext) {}

// EnterIndex_section is called when production index_section is entered.
func (s *BaseSQLListener) EnterIndex_section(ctx *Index_sectionContext) {}

// ExitIndex_section is called when production index_section is exited.
func (s *BaseSQLListener) ExitIndex_section(ctx *Index_sectionContext) {}

// EnterIndex_name is called when production index_name is entered.
func (s *BaseSQLListener) EnterIndex_name(ctx *Index_nameContext) {}

// ExitIndex_name is called when production index_name is exited.
func (s *BaseSQLListener) ExitIndex_name(ctx *Index_nameContext) {}

// EnterWhere_section is called when production where_section is entered.
func (s *BaseSQLListener) EnterWhere_section(ctx *Where_sectionContext) {}

// ExitWhere_section is called when production where_section is exited.
func (s *BaseSQLListener) ExitWhere_section(ctx *Where_sectionContext) {}
