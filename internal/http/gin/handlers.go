package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mdcg/access-control-service/restriction"
)

func Handlers(ru restriction.UseCase) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	RestrictionHandlers(v1, ru)

	return r
}
