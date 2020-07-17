package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"github.com/proximax-storage/xpx-catapult-faucet"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
	"net/http"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type errorResp struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func getAddressParam(ctx *gin.Context) (*string, error) {
	add := ctx.Param("address")
	if add != "" {
		if err := utils.IsAddressValid(add, Faucet.Config.NetworkType()); err != nil {
			return nil, Faucet.AddressInvalid
		}
	} else {
		return nil, Faucet.AddressMissing
	}
	return &add, nil
}

func getMosaicIdParam(ctx *gin.Context) (*string, error) {
	id := ctx.Param("mosaic")
	if id == "" {
		return nil, Faucet.MosaicMissing
	}
	return &id, nil
}

func respOk(ctx *gin.Context, res interface{}) {
	ctx.JSON(http.StatusOK, res)
}

func respError(ctx *gin.Context, err error) {
	if _, ok := err.(Faucet.BadRequestError); ok {
		respErrorWithCode(ctx, err, http.StatusBadRequest)
	} else {
		respErrorWithCode(ctx, err, http.StatusInternalServerError)
	}
}

func respErrorWithCode(ctx *gin.Context, err error, errorCode int) {
	res := &errorResp{
		Message: strings.TrimSpace(fmt.Sprint(err)),
	}

	ctx.AbortWithStatusJSON(errorCode, res)
	utils.Logger(3, "%d | %s", errorCode, res.Message)
}
