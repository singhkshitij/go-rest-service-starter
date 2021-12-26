package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	v1 "github.com/singhkshitij/golang-rest-service-starter/src/api/v1"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateReq(user v1.User, ctx *gin.Context) bool {
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false
	}
	return true
}

func MakeHandler(r *gin.Engine) {
	RegisterHealthCheckRoutes(r)
	RegisterAppRoutes(r)
}

func RegisterAppRoutes(r *gin.Engine) {
	r.POST("/v1/user", func(ctx *gin.Context) {
		var user v1.User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if ValidateReq(user, ctx){
			result := v1.GetUser()
			ctx.JSON(200, gin.H{
				"message": result,
			})
		}

	})
}

func RegisterHealthCheckRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
