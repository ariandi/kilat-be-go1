package api

import (
	"errors"
	"fmt"
	"github.com/ariandi/kilat-be-go1/dto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Inq(ctx *gin.Context) {
	logrus.Println("[digi DigiStatus] start.")
	var req dto.InqReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		errorValidator(ctx, err)
		return
	}
	out, errOut := inqService.InqPdamService(ctx, req)
	if errOut != nil {
		if out.ResultCd != "" {
			ctx.AbortWithStatusJSON(http.StatusOK, out)
		} else {
			errMsg := errors.New(errOut.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse(errMsg))
		}
		return
	}

	ctx.JSON(http.StatusOK, out)
	return
}

func errorValidator(ctx *gin.Context, err error) {
	errs, _ := err.(validator.ValidationErrors)
	logrus.Info("ok", errs)
	for _, v := range errs {
		field := v.Field()
		tag := v.Tag()

		errMsg := fmt.Sprintf("%v: %v", field, tag)
		errObj := errors.New(errMsg)
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse(errObj))
		break
	}

	if len(errs) == 0 {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse(err))
	}
	return
}
