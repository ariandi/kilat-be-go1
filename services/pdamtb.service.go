package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ariandi/kilat-be-go1/dto"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type PdamTbService struct {
	Config util.Config
	Client *resty.Client
	Store  *mongo.Database
}

const productName = "PDAM KOTA TANGERANG"
const pdamTangerangKota = "pdamtangerang"
const bankTb = "8"
const locketTb = "KTM"
const sKey = "FXDD23411"
const redBoxUser = "redbox01"

var pdamTbService *PdamTbService
var proxySvc ProxyInterface
var bankSvc BankInterface

type PdamTbInterface interface {
	InqPdamTb(ctx *gin.Context, in dto.InqReq) (dto.InqRes, error)
}

func GetPdamTbService(config util.Config, store *mongo.Database, client *resty.Client) PdamTbInterface {

	if pdamTbService == nil {
		proxySvc = GetProxyService(config, client)
		bankSvc = GetBankService(config, store)
		pdamTbService = &PdamTbService{
			Config: config,
			Client: client,
		}
	}

	return pdamTbService
}

func (s *PdamTbService) InqPdamTb(ctx *gin.Context, in dto.InqReq) (dto.InqRes, error) {
	logrus.Println("[PdamTbService InqPdamTb] start.")
	tb := dto.PdamTb{}
	out := dto.InqRes{ResultCd: util.ErrCd10, ResultMsg: util.ErrMsg10}

	bankArg := dto.GetBanks{Kode: in.Method}
	bank, err := bankSvc.GetBank(ctx, bankArg)
	if err != nil {
		return out, errors.New("get bank data error")
	}

	urlBank := new(dto.UrlBanks)
	err = json.Unmarshal([]byte(bank.Url), &urlBank)
	if err != nil {
		return out, errors.New("bank url json unmarshall error")
	}

	requestData := map[string]string{
		"produk": pdamTangerangKota,
		"bank":   bankTb,
		"loket":  locketTb,
		"nopel":  in.BillId,
	}

	headerArg := &dto.HeaderProxy{
		Key:   "skey",
		Value: sKey,
	}

	resp, err := proxySvc.PostFormService(urlBank.Url+urlBank.UrlInq, requestData, headerArg)
	if err != nil {
		logrus.Println("[[PdamTbService InqPdamTb]] send request to tb error : ", err)
		return out, err
	}

	err = json.Unmarshal(resp.Body(), &tb)
	if err != nil {
		logrus.Println("[PdamTbService InqPdamTb] error unmarshal tb response : ", err)
		return out, err
	}
	logrus.Info("[PdamTbService InqPdamTb] success send to provider")
	arg := dto.TbSetResponse{
		InqRequest:  in,
		InqResponse: out,
		TbResponse:  tb,
	}
	out = s.mappingResKilatTb(ctx, arg)

	return out, nil
}

func (s *PdamTbService) mappingResKilatTb(ctx *gin.Context, in dto.TbSetResponse) dto.InqRes {
	var out dto.InqRes
	if in.TbResponse.ResultCode != 0 {
		logrus.Println("[PdamTbService mappingResKilatTb] transaction not success in biller")
		out.ResultCd = util.ErrCd10
		out.ResultMsg = util.ErrMsg10 + " : " + util.Int64ToString(in.TbResponse.ResultCode) + " " + s.responseMsg(in.TbResponse)
		return out
	}

	out.ResultCd = util.ErrCd0
	out.ResultMsg = util.ErrMsg0

	arg := dto.TbSetResponse{
		InqRequest:  in.InqRequest,
		InqResponse: out,
		TbResponse:  in.TbResponse,
	}
	out = s.setInqRes(ctx, arg)
	return out
}

func (s *PdamTbService) setInqRes(ctx *gin.Context, in dto.TbSetResponse) dto.InqRes {
	totalAdmin := in.TbResponse.BillQty * s.Config.PdamAdmin.PdamTb

	merchantProdArg := dto.GetMerchantProducts{MerchantCode: in.InqRequest.MerchantCode, ProductCode: in.InqRequest.ProductCode, BankCode: in.InqRequest.Method}
	merchantProd, err := getMerchantProductService.GetMerchantProduct(ctx, merchantProdArg)
	if err != nil {
		return dto.InqRes{ResultCd: util.ErrCd6, ResultMsg: util.ErrMsg6}
	}

	feeMerchant := s.Config.PdamAdmin.PdamTb - int64(merchantProd.FeeMerchant)
	feeMerchantRest := in.TbResponse.BillQty * feeMerchant

	if in.InqRequest.User == redBoxUser {
		totalAdmin = s.Config.PdamAdmin.PdamTbVa
		feeMerchantRest = 0
	}

	out := dto.InqRes{
		ResultCd:        in.InqResponse.ResultCd,
		ResultMsg:       in.InqResponse.ResultMsg,
		TxDate:          in.InqRequest.TrxDate,
		TxId:            util.SetTxID(),
		RefId:           in.InqRequest.RefId,
		ProductCode:     in.InqRequest.ProductCode,
		BillId:          in.InqRequest.BillId,
		MerchantCode:    in.InqRequest.MerchantCode,
		Category:        in.InqRequest.Category,
		User:            in.InqRequest.User,
		Method:          in.InqRequest.Method,
		Sign:            in.InqRequest.MerchantToken,
		Amount:          int(in.TbResponse.TotalAmount),
		Admin:           int(totalAdmin),
		TotalAmount:     int(in.TbResponse.TotalAmount) + int(totalAdmin),
		DeductedBalance: int(feeMerchantRest),
		Detail:          s.setInqResDetail(in),
	}

	return out
}

func (s *PdamTbService) setInqResDetail(in dto.TbSetResponse) dto.DetailInqRes {
	var periodTb []string

	var billingDetal []dto.BillingDetailInqRes

	for _, detTb := range in.TbResponse.Detail {
		fmt.Println(detTb)
		logrus.Info("[PdamTbService setInqResDetail] detTb : ", detTb)
		periodTb = append(periodTb, detTb.Period)
		billingDetal = append(billingDetal, s.setInqResDetailArr(in, detTb))
	}

	return dto.DetailInqRes{
		ProductName:   productName,
		Periode:       strings.Join(periodTb, ", "),
		CustName:      in.TbResponse.BillName,
		CustAddress:   in.TbResponse.Address,
		Type:          in.TbResponse.Type, // type is golongan tarif
		RefProvider:   "-",
		DetailBilling: billingDetal,
	}
}

func (s *PdamTbService) setInqResDetailArr(in dto.TbSetResponse, resDetail dto.DetailPdamTb) dto.BillingDetailInqRes {
	out := dto.BillingDetailInqRes{
		Info:          resDetail.Jenis,
		StandMeter:    resDetail.Usage,
		Period:        resDetail.Period,
		Billamount:    util.StringToInt(resDetail.Total) - util.StringToInt(resDetail.Fine),
		Fine:          util.StringToInt(resDetail.Fine),
		AdminPpob:     int(s.Config.PdamAdmin.PdamTb),
		Total:         0,
		AdditionInfo:  "",
		AdditionInfo2: "",
	}
	if in.InqRequest.User == redBoxUser {
		out.AdminPpob = 0
	}
	out.Total = out.Billamount + out.AdminPpob + out.Fine

	return out
}

func (s *PdamTbService) responseMsg(in dto.PdamTb) string {
	var resultMsgTb string
	switch in.ResultCode {
	case 0:
		resultMsgTb = "Success"
		break
	case 7000:
		resultMsgTb = "System maintenance"
		break
	case 7004:
		resultMsgTb = "Pembayaran Gagal"
		break
	case 7120:
		resultMsgTb = "Contact IT Support"
		break
	case 7012:
		resultMsgTb = "Syntax format wrong"
		break
	case 7020:
		resultMsgTb = "Invalid account / Account not Found"
		break
	case 7026:
		resultMsgTb = "Invalid format"
		break
	case 7030:
		resultMsgTb = "Tidak ada Tagihan"
		break
	case 7107:
		resultMsgTb = "Invalid product code"
		break
	case 7110:
		resultMsgTb = "hanya bisa dilunasi di loket pdam"
		break
	case 7200:
		resultMsgTb = "Forbidden Acces"
		break
	case 7201:
		resultMsgTb = "CLOSING AT END OF MONTH"
		break
	case 7400:
		resultMsgTb = "Ticket Expired"
		break
	case 7500:
		resultMsgTb = "Rekon melebihi waktu yang ditetapkan"
		break
	default:
		resultMsgTb = "Forbidden Acces"
		break
	}

	return resultMsgTb
}
