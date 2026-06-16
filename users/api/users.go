package api

import (
	db "cloud-computing/users/db/sqlc"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listUsersRequest struct {
	Limit  int32 `uri:"limit" binding:"omitempty,min=1,max=100"`
	Offset int32 `uri:"offset" binding:"omitempty,min=1"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Offset == 0 {
		req.Offset = 1
	}

	params := db.ListUsersParams{
		Limit:  int32(req.Limit),
		Offset: int32((req.Offset - 1) * req.Limit),
	}

	users, err := server.store.ListUsers(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve users: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
		"pagination": gin.H{
			"page":  req.Offset,
			"limit": req.Limit,
		},
	})

}

type deleteUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err := server.store.DeleteUser(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
		"id":      req.ID,
	})
}
