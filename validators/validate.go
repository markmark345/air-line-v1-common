package validators

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/markmark345/air-line-v1-common/api/responses"
)

func validate(ctx *gin.Context, err error) {
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				responses.Error(ctx, errors.New("Missing required field "+"\""+err.Field()+"\""))
				return
			case "email":
				responses.Error(ctx, errors.New("Invalid email format for field "+"\""+err.Field()+"\""))
				return
			case "max":
				responses.Error(ctx, errors.New("Value for field \""+err.Field()+"\" exceeds the maximum allowed"))
				return
			case "min":
				responses.Error(ctx, errors.New("Value for field \""+err.Field()+"\" does not meet the minimum required"))
				return
			case "boolean":
				responses.Error(ctx, errors.New("Invalid value for boolean field \""+err.Field()+"\""))
				return
			case "len":
				responses.Error(ctx, errors.New("Invalid length for field \""+err.Field()+"\""))
				return
			case "oneof":
				responses.Error(ctx, errors.New("Invalid value for field \""+err.Field()+"\". Allowed values are M, N, or W"))
				return
			}
		}
	}
	return
}
