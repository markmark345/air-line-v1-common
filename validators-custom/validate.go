package validatorscustom

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/markmark345/air-line-v1-common/api/responses"
)

func ValidateCustomError(ctx *gin.Context, err error) {

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				responses.Failure(ctx, "invaild_parameters", errors.New("Missing required field "+"\""+err.Field()+"\""))
				return
			case "email":
				responses.Failure(ctx, "invaild_parameters", errors.New("Invalid email format for field "+"\""+err.Field()+"\""))
				return
			case "max":
				responses.Failure(ctx, "invaild_parameters", errors.New("Value for field \""+err.Field()+"\" exceeds the maximum allowed"))
				return
			case "min":
				responses.Failure(ctx, "invaild_parameters", errors.New("Value for field \""+err.Field()+"\" does not meet the minimum required"))
				return
			case "boolean":
				responses.Failure(ctx, "invaild_parameters", errors.New("Invalid value for boolean field \""+err.Field()+"\""))
				return
			case "len":
				responses.Failure(ctx, "invaild_parameters", errors.New("Invalid length for field \""+err.Field()+"\""))
				return
			case "oneof":
				responses.Failure(ctx, "invaild_parameters", errors.New("Invalid value for field \""+err.Field()+"\". Allowed values are M, N, or W"))
				return
			}
		}
	}
}
