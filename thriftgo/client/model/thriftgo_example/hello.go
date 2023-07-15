// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package thriftgo_example

import (
	"context"
	"fmt"
)

type HelloReq struct {
	Name string `thrift:"Name,1" json:"Name" query:"name"`
}

func NewHelloReq() *HelloReq {
	return &HelloReq{}
}

func (p *HelloReq) GetName() (v string) {
	return p.Name
}

func (p *HelloReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloReq(%+v)", *p)
}

type HelloResp struct {
	RespBody string `thrift:"RespBody,1" form:"RespBody" json:"RespBody" query:"RespBody"`
}

func NewHelloResp() *HelloResp {
	return &HelloResp{}
}

func (p *HelloResp) GetRespBody() (v string) {
	return p.RespBody
}

func (p *HelloResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloResp(%+v)", *p)
}

type QueryStudentReq struct {
	ID int32 `thrift:"ID,1" json:"ID" query:"id"`
}

func NewQueryStudentReq() *QueryStudentReq {
	return &QueryStudentReq{}
}

func (p *QueryStudentReq) GetID() (v int32) {
	return p.ID
}

func (p *QueryStudentReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("QueryStudentReq(%+v)", *p)
}

type QueryStudentResp struct {
	ID       int32  `thrift:"ID,1" form:"ID" json:"ID" query:"ID"`
	Name     string `thrift:"Name,2" form:"Name" json:"Name" query:"Name"`
	Favorite string `thrift:"Favorite,3" form:"Favorite" json:"Favorite" query:"Favorite"`
}

func NewQueryStudentResp() *QueryStudentResp {
	return &QueryStudentResp{}
}

func (p *QueryStudentResp) GetID() (v int32) {
	return p.ID
}

func (p *QueryStudentResp) GetName() (v string) {
	return p.Name
}

func (p *QueryStudentResp) GetFavorite() (v string) {
	return p.Favorite
}

func (p *QueryStudentResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("QueryStudentResp(%+v)", *p)
}

type AddStudentReq struct {
	ID       int32  `thrift:"ID,1" form:"ID" json:"ID"`
	Name     string `thrift:"Name,2" form:"Name" json:"Name"`
	Favorite string `thrift:"Favorite,3" form:"Favorite" json:"Favorite"`
}

func NewAddStudentReq() *AddStudentReq {
	return &AddStudentReq{}
}

func (p *AddStudentReq) GetID() (v int32) {
	return p.ID
}

func (p *AddStudentReq) GetName() (v string) {
	return p.Name
}

func (p *AddStudentReq) GetFavorite() (v string) {
	return p.Favorite
}

func (p *AddStudentReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddStudentReq(%+v)", *p)
}

type AddStudentResp struct {
	RespBody string `thrift:"RespBody,1" form:"RespBody" json:"RespBody" query:"RespBody"`
}

func NewAddStudentResp() *AddStudentResp {
	return &AddStudentResp{}
}

func (p *AddStudentResp) GetRespBody() (v string) {
	return p.RespBody
}

func (p *AddStudentResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddStudentResp(%+v)", *p)
}

type OtherReq struct {
	Other string `thrift:"Other,1" form:"other" json:"other"`
}

func NewOtherReq() *OtherReq {
	return &OtherReq{}
}

func (p *OtherReq) GetOther() (v string) {
	return p.Other
}

func (p *OtherReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OtherReq(%+v)", *p)
}

type OtherResp struct {
	Resp string `thrift:"Resp,1" form:"Resp" json:"Resp" query:"Resp"`
}

func NewOtherResp() *OtherResp {
	return &OtherResp{}
}

func (p *OtherResp) GetResp() (v string) {
	return p.Resp
}

func (p *OtherResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OtherResp(%+v)", *p)
}

type HelloService interface {
	HelloMethod(ctx context.Context, request *HelloReq) (r *HelloResp, err error)

	OtherMethod(ctx context.Context, request *OtherReq) (r *OtherResp, err error)
}

type StudentService interface {
	QueryStudentMethod(ctx context.Context, request *QueryStudentReq) (r *QueryStudentResp, err error)

	AddStudentMethod(ctx context.Context, request *AddStudentReq) (r *AddStudentResp, err error)
}