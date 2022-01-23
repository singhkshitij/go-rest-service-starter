package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	stats "github.com/semihalev/gin-stats"
	v1 "github.com/singhkshitij/golang-rest-service-starter/src/api/v1"
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
	r.GET("/api/v1/tweets/:category", func(ctx *gin.Context) {
		category := ctx.Param("category")
		pageNumber := ctx.Query("page")
		result, totalTweets, err := v1.GetTweetsForCategory(category, pageNumber)
		if err != nil {
			ctx.JSON(200, gin.H{
				"success": false,
				"error":   err,
			})
		} else {
			ctx.JSON(200, gin.H{
				"success": true,
				"data": &v1.TweetCategoryResponse{
					Results:     result,
					TotalTweets: totalTweets,
				},
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
