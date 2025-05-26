package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/novadm/backend/internal"
)

func checksHealthReady(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func checksHealthAlive(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func checksVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"version": internal.Version})
}
