package buckets

import (
	"net/http"

	"github.com/eriicafes/go-api-starter/models"
	"github.com/eriicafes/go-api-starter/response"
	"github.com/gin-gonic/gin"
)

type bucketsController struct {
	bucketsService BucketsService
}

func NewBucketsController(bucketsService BucketsService) *bucketsController {
	return &bucketsController{
		bucketsService: bucketsService,
	}
}

func (c *bucketsController) Routes(r *gin.RouterGroup) {
	r.GET("", c.Get)
	r.PUT("", c.Put)
}

func (c *bucketsController) Get(ctx *gin.Context) {
	res := response.New(ctx)

	user := ctx.MustGet("auth").(*models.User)

	bucket, err := c.bucketsService.Get(user.AccountID)

	if err != nil {
		res.SetStatusMessage(http.StatusNotFound, err.Error()).ErrJSON()
		return
	}

	res.SetData(bucket).JSON()
}

func (c *bucketsController) Put(ctx *gin.Context) {
	res := response.New(ctx)

	user := ctx.MustGet("auth").(*models.User)

	var bucketDto struct {
		Data interface{} `json:"data" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&bucketDto); err != nil {
		res.SetStatusMessage(http.StatusUnprocessableEntity, "validation error").SetError(err.Error()).ErrJSON()
		return
	}

	bucketData := models.Bucket{
		Data: bucketDto.Data,
	}

	bucket := c.bucketsService.Put(user.AccountID, bucketData)

	res.SetData(bucket).JSON()
}
