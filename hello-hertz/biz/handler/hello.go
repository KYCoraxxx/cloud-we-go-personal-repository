package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Hello(ctx context.Context, c *app.RequestContext) {
	name, _ := c.GetQuery("name")
	c.JSON(consts.StatusOK, utils.H{
		"message": "Hello, " + name,
	})
}
