package services

import (
	"encoding/json"
	"errors"
	"github.com/ariandi/kilat-be-go1/cache"
	"github.com/ariandi/kilat-be-go1/dto"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type InqService struct {
	Store  *mongo.Database
	Config util.Config
	Client *resty.Client
	Redis  cache.RedisInterface
}

var inqService *InqService
var getPdamTbService PdamTbInterface
var getMerchantProductService MerchantProductInterface

type InqInterface interface {
	InqPdamService(ctx *gin.Context, in dto.InqReq) (dto.InqRes, error)
}

func GetInqService(config util.Config, store *mongo.Database, client *resty.Client, redis cache.RedisInterface) InqInterface {

	if inqService == nil {
		getPdamTbService = GetPdamTbService(config, store, client)
		getMerchantProductService = GetMerchantProductService(config, store)
		inqService = &InqService{
			Store:  store,
			Config: config,
			Client: client,
			Redis:  redis,
		}
	}

	return inqService
}

func (i *InqService) InqPdamService(ctx *gin.Context, in dto.InqReq) (dto.InqRes, error) {
	logrus.Println("[InqService InqPdamService] start.")
	out := dto.InqRes{
		ResultCd:  util.ErrCd99,
		ResultMsg: util.ErrMsg99,
	}
	out, err := i.inqValidation(ctx, in)
	if err != nil {
		return out, err
	}

	out, err = i.sendToPdamService(ctx, in, out)
	if err != nil {
		return out, err
	}

	if out.ResultCd != util.ErrCd0 {
		return out, nil
	}

	return out, nil
}

func (i *InqService) inqValidation(ctx *gin.Context, in dto.InqReq) (dto.InqRes, error) {
	logrus.Println("[InqService inqValidation] start.")
	out := dto.InqRes{
		ResultCd:  util.ErrCd0,
		ResultMsg: util.ErrMsg0,
	}
	checkProdReq := false
	merchantArg := dto.MerchantReq{User: in.User, Kode: in.MerchantCode}
	merchant, err := merchantService.GetMerchantUserService(ctx, merchantArg)
	if err != nil {
		out = dto.InqRes{ResultCd: util.ErrCd1, ResultMsg: util.ErrMsg1}
	}

	prodListStr, _ := i.Redis.GetRedisVal("PDAM_PROD_LIST")
	prodListData := new([]string)
	_ = json.Unmarshal([]byte(prodListStr), prodListData)

	for _, prodData := range *prodListData {
		if in.ProductCode == prodData {
			checkProdReq = true
			break
		}
	}

	if !checkProdReq {
		out = dto.InqRes{ResultCd: util.ErrCd3, ResultMsg: util.ErrMsg3}
	}

	bankListStr, _ := i.Redis.GetRedisVal("PDAM_BANK_LIST")
	bankListData := new([]string)
	_ = json.Unmarshal([]byte(bankListStr), bankListData)

	checkProdReq = false
	for _, bankData := range *bankListData {
		if in.Method == bankData {
			checkProdReq = true
			break
		}
	}

	if !checkProdReq {
		out = dto.InqRes{ResultCd: util.ErrCd5, ResultMsg: util.ErrMsg5}
	}

	cat, _ := i.Redis.GetRedisVal("CATEGORY")
	if in.Category != cat {
		out = dto.InqRes{ResultCd: util.ErrCd4, ResultMsg: util.ErrMsg4}
	}

	merchantProdArg := dto.GetMerchantProducts{MerchantCode: in.MerchantCode, ProductCode: in.ProductCode, BankCode: in.Method}
	_, err = getMerchantProductService.GetMerchantProduct(ctx, merchantProdArg)
	if err != nil {
		out = dto.InqRes{ResultCd: util.ErrCd6, ResultMsg: util.ErrMsg6}
	}

	_, err = time.Parse("20060102150405", in.TrxDate)
	if err != nil {
		out = dto.InqRes{ResultCd: util.ErrCd7, ResultMsg: util.ErrMsg7}
	}

	setLocalToken := in.BillId + in.ProductCode + in.User + in.RefId + merchant.SecretCode + in.TrxDate
	localToken := util.StoreSha256([]byte(setLocalToken))
	logrus.Println("[InqService inqValidation] localToken is : ", localToken)
	logrus.Println("[InqService inqValidation] merchantToken is ", in.MerchantToken)

	if localToken != in.MerchantToken {
		logrus.Println("[InqService inqValidation] start.")
		out = dto.InqRes{ResultCd: util.ErrCd8, ResultMsg: util.ErrMsg8}
	}

	// TODO check by redis is ref id already used or not

	if out.ResultCd != util.ErrCd0 {
		return out, errors.New("validation error")
	}

	return out, nil
}

func (i *InqService) sendToPdamService(ctx *gin.Context, in dto.InqReq, out dto.InqRes) (dto.InqRes, error) {
	switch in.ProductCode {
	case i.Config.PdamCd.PdamTb:
		out, err := getPdamTbService.InqPdamTb(ctx, in)
		if err != nil {
			return out, err
		}
		logrus.Println("[InqService InqPdamService] start.", out)

		if out.ResultCd != util.ErrCd0 {
			logrus.Println("[InqService InqPdamService] transaction not success in biller")
			return out, nil
		}

		return out, nil
	}

	return out, nil
}
