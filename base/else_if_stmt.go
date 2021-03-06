package base

import (
	"gengine/context"
	"gengine/core/errors"
	"reflect"
)

type ElseIfStmt struct {
	Expression       *Expression
	StatementList    *Statements
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}


func (ef *ElseIfStmt) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	it ,err := ef.Expression.Evaluate(Vars)
	if err != nil {
		return nil, err
	}

	if reflect.ValueOf(it).Bool() {
		if ef.StatementList == nil{
			return nil,nil
		}else {
			return ef.StatementList.Evaluate(Vars)
		}
	}else {
		return nil,nil
	}
}


func (ef *ElseIfStmt) Initialize(kc *KnowledgeContext,  dc *context.DataContext) {
	ef.knowledgeContext = kc
	ef.dataCtx = dc

	if ef.Expression != nil {
		ef.Expression.Initialize(kc, dc)
	}

	if ef.StatementList != nil {
		ef.StatementList.Initialize(kc, dc)
	}
}


func (ef *ElseIfStmt)AcceptExpression(expr *Expression) error{

	if ef.Expression == nil {
		ef.Expression = expr
		return nil
	}
	return errors.New("ElseIfStmt's Expression set twice!")
}


func (ef *ElseIfStmt)AcceptStatements(stmts *Statements)error{
	if ef.StatementList == nil {
		ef.StatementList = stmts
		return nil
	}
	return errors.New("ElseIfStmt's statements set twice!")
}
