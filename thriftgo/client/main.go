package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/config"
	"thirft_client/hertz_client/student_service"
	"thirft_client/model/thriftgo_example"
)

func main() {
	generatedClient, err := student_service.NewStudentServiceClient(
		"http://127.0.0.1:8888", // 指定 psm 作为域名
	)

	req := thriftgo_example.AddStudentReq{
		ID:       1,
		Name:     "张铭铭",
		Favorite: "MajSoul",
	}

	// 在发起调用的时候可指定请求级别的配置
	resp, rawResp, err := generatedClient.AddStudentMethod(
		context.Background(),
		&req,
		config.WithSD(true),                   // 指定请求级别的设置，用来开启服务发现
		config.WithReadTimeout(1000000000000), // 指定请求读超时
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
	fmt.Println(rawResp)

	queryReq := thriftgo_example.QueryStudentReq{
		ID: 1,
	}

	Queryresp, QueryrawResp, Queryerr := generatedClient.QueryStudentMethod(
		context.Background(),
		&queryReq,
		config.WithSD(true),
		config.WithReadTimeout(1000000000000),
	)
	if Queryerr != nil {
		fmt.Println(Queryerr)
		return
	}

	fmt.Println(Queryresp)
	fmt.Println(QueryrawResp)
}
