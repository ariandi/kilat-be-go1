package services

import (
	"github.com/ariandi/kilat-be-go1/dto"
	"github.com/ariandi/kilat-be-go1/models"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TransactionService struct {
	Store  *mongo.Database
	Config util.Config
	Client *resty.Client
}

var trxService *TransactionService

const tbTrx = "transactions"

// GetTransactionService is
func GetTransactionService(config util.Config, store *mongo.Database) TransactionInterface {

	if trxService == nil {
		trxService = &TransactionService{
			Store:  store,
			Config: config,
		}
	}

	return trxService
}

type TransactionInterface interface {
	CreateTrxService(ctx *gin.Context, in dto.TransactionCreate) (*mongo.InsertOneResult, error)
	GetTrxByIdService(ctx *gin.Context, in string) (*dto.MerchantRes, error)
	GetTrxByTxIdService(ctx *gin.Context, in dto.MerchantReq) (*dto.MerchantRes, error)
	// UpdateMerchantService(ctx *gin.Context, in dto.UpdateMerchantRequest) (db.Merchant, error)
}

func (t *TransactionService) CreateTrxService(ctx *gin.Context, in dto.TransactionCreate) (*mongo.InsertOneResult, error) {
	logrus.Println("[TransactionService CreateTrxService] start.")

	arg := models.Transaction{
		BillId:       in.BillId,
		TxId:         in.TxId,
		RefId:        in.RefId,
		Type:         in.Type,
		Status:       "3",
		Sign:         in.Sign,
		MerchantId:   in.MerchantId,
		BankId:       in.BankId,
		CategoryId:   in.CategoryId,
		CategoryName: in.CategoryName,
		ProductId:    in.ProductId,
		ProductName:  in.ProductName,
		ReqParam:     in.ReqParam,
		ResParam:     in.ResParam,
		BatchTime:    1,
		Active:       1,
		CreatedAt:    time.Time.Format(time.Now(), "2006-01-02 15:04:05"),
		CreatedBy:    in.CreatedBy,
		UpdatedAt:    time.Time.Format(time.Now(), "2006-01-02 15:04:05"),
		UpdatedBy:    in.CreatedBy,
		TxDate:       in.TxDate,
	}
	out, err := t.Store.Collection(tbTrx).InsertOne(ctx, arg)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (t *TransactionService) GetTrxByIdService(ctx *gin.Context, in string) (*dto.MerchantRes, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionService) GetTrxByTxIdService(ctx *gin.Context, in dto.MerchantReq) (*dto.MerchantRes, error) {
	//TODO implement me
	panic("implement me")
}
