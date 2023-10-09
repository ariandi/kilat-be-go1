package services

import (
	"github.com/ariandi/kilat-be-go1/dto"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BankService struct {
	Store  *mongo.Database
	Config util.Config
}

var bankService *BankService

const tbBanks = "banks"

type BankInterface interface {
	GetBank(ctx *gin.Context, in dto.GetBanks) (*dto.Banks, error)
}

// GetBankService is
func GetBankService(config util.Config, store *mongo.Database) BankInterface {

	if bankService == nil {
		bankService = &BankService{
			Store:  store,
			Config: config,
		}
	}

	return bankService
}

func (o *BankService) GetBank(ctx *gin.Context, in dto.GetBanks) (*dto.Banks, error) {
	logrus.Println("[BankService GetBank] start.")
	out := new(dto.Banks)

	err := o.Store.Collection(tbBanks).FindOne(ctx, bson.M{"kode": in.Kode}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
