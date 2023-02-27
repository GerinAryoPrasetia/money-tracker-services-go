package controllers

import (
	"context"

	db "github.com/GerinAryoPrasetia/go-money-tracking-services/database/sqlc"
)

// package controllers

type BudgetController struct {
	db  *db.Queries
	ctx context.Context
}

func NewBudgetController(db *db.Queries, ctx context.Context) *BudgetController {
	return &BudgetController{db, ctx}
}
