package schemas

type CreateBudget struct {
	UserID      string `json:"user_id" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateBudget struct {
	UserID      string `json:"user_id" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	Description string `json:"description" binding:"required"`
}
