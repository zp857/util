package client

import (
	"github.com/imroc/req/v3"
	"github.com/zp857/util/maputil"
	"github.com/zp857/util/sliceutil"
	"github.com/zp857/util/stack"
	"github.com/zp857/util/structutil"
	"go.uber.org/zap"
	"net/http"
)

type HTTPClient struct {
	logger             *zap.SugaredLogger
	reqClient          *req.Client
	baseUrl            string
	noStatusDebugPrint bool
	ignoreApis         []string
}

func NewHTTPClient(reqClient *req.Client, url string, statusDebugPrint bool, ignoreApiList []string) *HTTPClient {
	return &HTTPClient{
		logger:             zap.L().Named("[http-client]").Sugar(),
		reqClient:          reqClient,
		baseUrl:            url,
		noStatusDebugPrint: !statusDebugPrint,
		ignoreApis:         ignoreApiList,
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *HTTPClient) SendWithResponse(api string, code int, data interface{}, msg string) {
	r := c.reqClient.R()
	body := Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	r.SetBody(body)
	resp, err := r.Post(c.baseUrl + api)
	if err != nil {
		c.logger.Errorf("request err: %v", err)
		return
	}
	if c.noStatusDebugPrint {
		if sliceutil.Contain(c.ignoreApis, api) {
			return
		}
	}
	c.logger.Infof("request: [%v]\n%v", api, structutil.JsonMarshalIndent(data))
	c.logger.Infof("response: [%v]\n%v", resp.StatusCode, structutil.JsonMarshalIndent(maputil.BytesToMap(resp.Bytes())))
	return
}

func (c *HTTPClient) SendOk(api string, data interface{}, msg string) {
	c.SendWithResponse(api, http.StatusOK, data, msg)
}

func (c *HTTPClient) SendErrorWithMessage(api string, msg string, err error) {
	if err != nil {
		c.logger.Error(
			msg,
			zap.Error(err),
			zap.Any("stack", string(stack.GetStack(2))),
		)
	}
	c.SendWithResponse(api, http.StatusInternalServerError, struct{}{}, msg)
}

func (c *HTTPClient) SendJson(api string, data interface{}) {
	r := c.reqClient.R()
	r.SetBody(data)
	resp, err := r.Post(c.baseUrl + api)
	if err != nil {
		c.logger.Errorf("request err: %v", err)
		return
	}
	if c.noStatusDebugPrint {
		if sliceutil.Contain(c.ignoreApis, api) {
			return
		}
	}
	c.logger.Infof("request: [%v]\n%v", api, structutil.JsonMarshalIndent(data))
	c.logger.Infof("response: [%v]\n%v", resp.StatusCode, structutil.JsonMarshalIndent(maputil.BytesToMap(resp.Bytes())))
	return
}
