package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type null struct{}

var Null *null

type httpResponse struct {
	ctx     *gin.Context
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func New(ctx *gin.Context) httpResponse {
	return httpResponse{
		ctx:    ctx,
		Status: http.StatusOK,
	}
}

func (r httpResponse) JSON() {
	r.Error = nil
	if r.Message == "" {
		r.Message = http.StatusText(r.Status)
	}
	r.ctx.JSON(r.Status, r)
}

func (r httpResponse) ErrJSON() {
	r.Data = nil
	if r.Status < 300 {
		r.Status = http.StatusInternalServerError
	}
	if r.Message == "" {
		r.Message = http.StatusText(r.Status)
	}
	r.ctx.AbortWithStatusJSON(r.Status, r)
}

func (r httpResponse) SetData(data interface{}) httpResponse {
	r.Data = data
	return r
}

func (r httpResponse) SetError(err interface{}) httpResponse {
	r.Error = err
	return r
}

func (r httpResponse) SetStatus(status int) httpResponse {
	r.Status = status
	return r
}

func (r httpResponse) SetMessage(message string) httpResponse {
	r.Message = message
	return r
}

func (r httpResponse) SetStatusMessage(status int, message string) httpResponse {
	r.Status, r.Message = status, message
	return r
}
