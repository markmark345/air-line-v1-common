package responses

import (
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, msgKey string, datas interface{}) {
	// dataLookup, err := GetLookup(msgKey)
	// if err != nil {
	// 	ctx.JSON(500, err)
	// }

	// res := &model.Response{
	// 	Status: model.Status {
	// 		Code: dataLookup.Code,
	// 		Description: dataLookup.DescEN,
	// 	},
	// 	Data: datas,
	// }

	// ctx.JSON(dataLookup.HTTPCode, res)
	ctx.JSON(200, "sss")
	// return
}

// func Write(ctx *gin.Context, res interface{}) error {
// 	var httpCode int
// 	// var code int
// 	// var description string

// 	return ctx.JSON(httpCode, res)
// }

// , res interface{} ctx *gin.Context,
