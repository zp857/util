package producer

import (
	"github.com/zp857/util/kafka"
	"github.com/zp857/util/stack"
	"go.uber.org/zap"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendWithResponse(topic string, code int, data interface{}, msg string, p *kafka.Producer) {
	response := Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	p.SendJSON(topic, response)
}

func SendOk(topic string, data interface{}, msg string, p *kafka.Producer) {
	SendWithResponse(topic, http.StatusOK, data, msg, p)
}

func SendOkWithMessage(topic string, msg string, p *kafka.Producer) {
	SendWithResponse(topic, http.StatusOK, struct{}{}, msg, p)
}

func SendErrorWithMessage(topic string, msg string, err error, p *kafka.Producer) {
	if err != nil {
		zap.L().Named("[kafka-producer]").Error(
			msg,
			zap.Error(err),
			zap.Any("stack", string(stack.GetStack(2))),
		)
	}
	SendWithResponse(topic, http.StatusInternalServerError, struct{}{}, msg, p)
}

func BadRequestWithMessage(topic string, msg string, p *kafka.Producer) {
	SendWithResponse(topic, http.StatusBadRequest, struct{}{}, msg, p)
}

func SendJson(topic string, data interface{}, p *kafka.Producer) {
	p.SendJSON(topic, data)
}
