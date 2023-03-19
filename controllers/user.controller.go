package controllers

import (
	"context"
	"database/sql"
	"net/http"

	db "github.com/GerinAryoPrasetia/go-money-tracking-services/database/sqlc"
	"github.com/GerinAryoPrasetia/go-money-tracking-services/pkg"
	"github.com/GerinAryoPrasetia/go-money-tracking-services/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// package controllers

type UserController struct {
	db  *db.Queries
	ctx context.Context
}

func NewUserController(db *db.Queries, ctx context.Context) *UserController {
	return &UserController{db, ctx}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var payload *schemas.CreateUser

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "invalid json payload",
			"message": err.Error(),
		})
		return
	}

	id := uuid.New().String()
	hashPassword, err := pkg.HashPassword(payload.Password)
	if err != nil {
		//TODO: LOGGER ERROR
		return
	}
	args := &db.CreateUserParams{
		ID:        id,
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  hashPassword,
		CreatedBy: payload.Email,
		UpdatedBy: sql.NullString{String: payload.Email, Valid: true},
	}

	user, err := uc.db.CreateUser(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	userToken, err := pkg.GenerateJWT(user.ID, user.Email, user.Name)
	if err != nil {
		//TODO : Logger
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user.ID,
		"token":  userToken,
	})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var payload *schemas.Login

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "invalid json payload",
			"message": err.Error(),
		})
		return
	}

	user, err := uc.db.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "invalid email or password",
				"success": "false",
			})
			return
		}
	}

	hashPass := pkg.CheckPasswordHash(payload.Password, user.Password)
	if !hashPass {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "invalid email or password",
			"success": "false",
		})
		return
	}

	userToken, err := pkg.GenerateJWT(user.ID, user.Email, user.Name)
	if err != nil {
		//TODO : Logger
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user.ID,
		"token":  userToken,
	})
}
