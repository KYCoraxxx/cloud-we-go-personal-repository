// Code generated by hertz generator.

package student_service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	thriftgo_example "thirft_client/model/thriftgo_example"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	QueryStudentMethod(context context.Context, req *thriftgo_example.QueryStudentReq, reqOpt ...config.RequestOption) (resp *thriftgo_example.QueryStudentResp, rawResponse *protocol.Response, err error)

	AddStudentMethod(context context.Context, req *thriftgo_example.AddStudentReq, reqOpt ...config.RequestOption) (resp *thriftgo_example.AddStudentResp, rawResponse *protocol.Response, err error)
}

type StudentServiceClient struct {
	client *cli
}

func NewStudentServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &StudentServiceClient{
		client: cli,
	}, nil
}

func (s *StudentServiceClient) QueryStudentMethod(context context.Context, req *thriftgo_example.QueryStudentReq, reqOpt ...config.RequestOption) (resp *thriftgo_example.QueryStudentResp, rawResponse *protocol.Response, err error) {
	httpResp := &thriftgo_example.QueryStudentResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"id": req.GetID(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/query")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *StudentServiceClient) AddStudentMethod(context context.Context, req *thriftgo_example.AddStudentReq, reqOpt ...config.RequestOption) (resp *thriftgo_example.AddStudentResp, rawResponse *protocol.Response, err error) {
	httpResp := &thriftgo_example.AddStudentResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/add-student-info")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewStudentServiceClient("http://127.0.0.1:8888")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewStudentServiceClient("http://127.0.0.1:8888", ops...)
	return
}

func QueryStudentMethod(context context.Context, req *thriftgo_example.QueryStudentReq, reqOpt ...config.RequestOption) (resp *thriftgo_example.QueryStudentResp, rawResponse *protocol.Response, err error) {
	return defaultClient.QueryStudentMethod(context, req, reqOpt...)
}

func AddStudentMethod(context context.Context, req *thriftgo_example.AddStudentReq, reqOpt ...config.RequestOption) (resp *thriftgo_example.AddStudentResp, rawResponse *protocol.Response, err error) {
	return defaultClient.AddStudentMethod(context, req, reqOpt...)
}
