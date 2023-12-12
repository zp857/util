package kafka

import (
	"encoding/json"
	"github.com/zp857/util/funcutil"
)

type HandlerBinding struct {
	TopicName   string
	HandlerName string
	HandlerFunc HandlerFunc
	Handler     string
}

func DirectBinding(topicName, handlerName string, handlerFunc HandlerFunc) HandlerBinding {
	return HandlerBinding{
		TopicName:   topicName,
		HandlerName: handlerName,
		HandlerFunc: handlerFunc,
		Handler:     funcutil.NameOfFunction(handlerFunc),
	}
}

func BindJSON(msg []byte, obj any) (err error) {
	err = json.Unmarshal(msg, obj)
	return
}
