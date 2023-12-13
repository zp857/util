package response

import (
	"github.com/gin-gonic/gin"
	"github.com/zp857/util/stack"
	"github.com/zp857/util/structutil"
	"go.uber.org/zap"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	DumpFormat = "response:\n%v"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	response := Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	zap.L().Named("[gin]").Sugar().Infof(DumpFormat, structutil.JsonMarshalIndent(response))
	c.JSON(http.StatusOK, response)
}

func Ok(data interface{}, msg string, c *gin.Context) {
	Result(http.StatusOK, data, msg, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(http.StatusOK, struct{}{}, msg, c)
}

func ErrorWithMessage(msg string, err error, c *gin.Context) {
	if err != nil {
		zap.L().Named("[gin]").Error(
			msg,
			zap.Error(err),
			zap.Any("stack", string(stack.GetStack(2))),
		)
	}
	Result(http.StatusInternalServerError, struct{}{}, msg, c)
}

func UnAuthWithMessage(msg string, c *gin.Context) {
	Result(http.StatusUnauthorized, struct{}{}, msg, c)
}

func BadRequestWithMessage(msg string, c *gin.Context) {
	Result(http.StatusBadRequest, struct{}{}, msg, c)
}
