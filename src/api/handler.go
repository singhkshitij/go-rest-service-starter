package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	stats "github.com/semihalev/gin-stats"
	v1 "github.com/singhkshitij/golang-rest-service-starter/src/api/v1"
	"github.com/singhkshitij/golang-rest-service-starter/src/metrics"
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
		//This is sample app metric. Change this to something meaningful
		metrics.Increment(metrics.UserMetric.Name, metrics.UserMetric.Labels)
		var user v1.User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if ValidateReq(user, ctx) {
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

	//Display request stats on this endpoint
	r.GET("/request/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})
}
