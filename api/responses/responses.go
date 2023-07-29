package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/markmark345/air-line-v1-common/api/responses/model"
)

func Success(ctx *gin.Context, msgKey string, datas interface{}) {
	dataLookup, err := GetLookup(msgKey)
	if err != nil {
		ctx.JSON(500, err)
	}

	res := &model.Response{
		Status: model.Status{
			Code:        dataLookup.Code,
			Description: dataLookup.DescEN,
		},
		Data: datas,
	}

	ctx.JSON(dataLookup.HTTPCode, res)
}

func Failure(ctx *gin.Context, msgKey string, details interface{}) {
	dataLookup, err := GetLookup(msgKey)
	if err != nil {
		ctx.JSON(500, err)
	}

	res := &model.Response{
		Status: model.Status{
			Code:        dataLookup.Code,
			Description: dataLookup.DescEN,
			Details:     details,
		},
	}

	ctx.JSON(dataLookup.HTTPCode, res)
}

func Error(ctx *gin.Context, err error) {
	res := &model.Response{
		Status: model.Status{
			Code:        9999,
			Description: "Internal Server Error",
			Details: model.Details{
				Error: err.Error(),
			},
		},
	}

	ctx.JSON(500, res)
}
