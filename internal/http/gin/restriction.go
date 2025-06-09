package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdcg/access-control-service/restriction"
)

func CreateRestriction(ru restriction.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r restriction.Restriction
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := ru.CreateRestriction(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"received": r})
	}
}

func RestrictionHandlers(r *gin.RouterGroup, ru restriction.UseCase) {
	r.POST("/restrictions", CreateRestriction(ru))
}
