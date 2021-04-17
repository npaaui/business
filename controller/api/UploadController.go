package api

import (
	"fmt"
	"github.com/gin-gonic/gin"

	. "business/common"
)

type UploadController struct{}

func NewUploadController() *UploadController {
	return &UploadController{}
}

type UploadFileArgs struct {
	ResType string `json:"res_type"`
}

func (c *UploadController) UploadFile(g *gin.Context) {
	args := &UploadFileArgs{}
	ValidatePostForm(g, map[string]string{
		"res_type": "string|required|enum:task,recharge||资源类型",
	}, args)

	name := LoadPostFile(g, "file", args.ResType)

	srcConf, err := Conf.GetSection("SRC")
	if err != nil {
		panic(NewSysErr(fmt.Errorf("get src conf error: %w", err)))
	}

	data := map[string]string{
		"file_path": srcConf["src"] + name,
		"file_name": name,
	}

	ReturnData(g, data)
	return
}
