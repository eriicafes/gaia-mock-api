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
	r.GET("/:accountId", c.GetAll)
	r.GET("/:accountId/:name", c.Get)
	r.PUT("/:accountId/:name", c.Put)
}

func (c *bucketsController) GetAll(ctx *gin.Context) {
	res := response.New(ctx)
	user := ctx.MustGet("auth").(*models.User)

	// users can only access their bucket list
	accountId := ctx.Param("accountId")
	if accountId != user.AccountID {
		res.SetStatus(http.StatusUnauthorized).ErrJSON()
		return
	}

	buckets := c.bucketsService.GetAll(accountId)

	if buckets == nil {
		buckets = []models.Bucket{}
	}

	res.SetData(buckets).JSON()
}

func (c *bucketsController) Get(ctx *gin.Context) {
	res := response.New(ctx)

	// other authorized users can access a user's bucket
	name := ctx.Param("name")
	accountId := ctx.Param("accountId")

	bucket, err := c.bucketsService.Get(accountId, name)

	if err != nil {
		res.SetStatusMessage(http.StatusNotFound, err.Error()).ErrJSON()
		return
	}

	res.SetData(bucket).JSON()
}

func (c *bucketsController) Put(ctx *gin.Context) {
	res := response.New(ctx)
	user := ctx.MustGet("auth").(*models.User)

	// users can only modify their bucket
	name := ctx.Param("name")
	accountId := ctx.Param("accountId")
	if accountId != user.AccountID {
		res.SetStatus(http.StatusUnauthorized).ErrJSON()
		return
	}

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

	bucket := c.bucketsService.Put(user.AccountID, name, bucketData)

	res.SetData(bucket).JSON()
}
