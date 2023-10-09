package services

import (
	"github.com/ariandi/kilat-be-go1/dto"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MerchantProductService struct {
	Store  *mongo.Database
	Config util.Config
}

var merchantProductService *MerchantProductService

const tbMerchantProducts = "merchantproducts"

type MerchantProductInterface interface {
	GetMerchantProduct(ctx *gin.Context, in dto.GetMerchantProducts) (*dto.MerchantProducts, error)
}

// GetMerchantProductService is
func GetMerchantProductService(config util.Config, store *mongo.Database) MerchantProductInterface {

	if merchantProductService == nil {
		merchantProductService = &MerchantProductService{
			Store:  store,
			Config: config,
		}
	}

	return merchantProductService
}

func (o *MerchantProductService) GetMerchantProduct(ctx *gin.Context, in dto.GetMerchantProducts) (*dto.MerchantProducts, error) {
	logrus.Println("[MerchantProductService GetMerchantProduct] start.")
	out := new(dto.MerchantProducts)

	err := o.Store.Collection(tbMerchantProducts).FindOne(ctx, bson.M{"merchant_code": in.MerchantCode, "product_code": in.ProductCode, "bank_code": in.BankCode}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
