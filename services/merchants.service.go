package services

import (
	"context"
	"github.com/ariandi/kilat-be-go1/dto"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MerchantService struct {
	Store  *mongo.Database
	Config util.Config
	Client *resty.Client
}

var merchantService *MerchantService
var ctxDefault = func() context.Context {
	return context.Background()
}()

const tbMerchant = "merchants"

type MerchantInterface interface {
	// CreateMerchantService(ctx *gin.Context, in dto.CreateMerchantRequest) (db.Merchant, error)
	GetMerchantService(ctx *gin.Context, in string) (*dto.MerchantRes, error)
	GetMerchantUserService(ctx *gin.Context, in dto.MerchantReq) (*dto.MerchantRes, error)
	ListMerchantService(ctx *gin.Context, in dto.ListMerchantRequest) ([]dto.MerchantRes, error)
	// UpdateMerchantService(ctx *gin.Context, in dto.UpdateMerchantRequest) (db.Merchant, error)
}

// GetMerchantService is
func GetMerchantService(config util.Config, store *mongo.Database, client *resty.Client) MerchantInterface {

	if merchantService == nil {
		merchantService = &MerchantService{
			Store:  store,
			Config: config,
			Client: client,
		}
	}

	return merchantService
}

//func (o *MerchantService) CreateMerchantService(ctx *gin.Context, in dto.CreateMerchantRequest) (db.Merchant, error) {
//	// just for temporary not with dto mapping
//	logrus.Println("[MerchantService CreateMerchantService] start.")
//	var result db.Merchant
//
//	uId := uuid.New().String()
//	arg := db.CreateMerchantParams{
//		ID:   uId,
//		Name: in.Name,
//		Balance: sql.NullString{
//			String: util.FloatToString(in.Balance),
//			Valid:  true,
//		},
//		UpdatedAt: sql.NullTime{
//			Time:  time.Now(),
//			Valid: true,
//		},
//		CreatedBy: sql.NullString{String: "1", Valid: true},
//		UpdatedBy: sql.NullString{String: "1", Valid: true},
//	}
//
//	_, err := o.Store.CreateMerchant(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse(err))
//		return result, err
//	}
//
//	out, _ := o.GetMerchantService(ctx, uId)
//
//	return out, nil
//}

func (o *MerchantService) GetMerchantService(ctx *gin.Context, in string) (*dto.MerchantRes, error) {
	logrus.Println("[MerchantService GetMerchantService] start.")
	_id, _ := primitive.ObjectIDFromHex(in)
	out := new(dto.MerchantRes)

	err := o.Store.Collection(tbMerchant).FindOne(ctx, bson.M{"_id": _id}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (o *MerchantService) GetMerchantUserService(ctx *gin.Context, in dto.MerchantReq) (*dto.MerchantRes, error) {
	logrus.Println("[MerchantService GetMerchantService] start.")
	out := new(dto.MerchantRes)

	err := o.Store.Collection(tbMerchant).FindOne(ctx, bson.M{"kode": in.Kode, "user": in.User}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (o *MerchantService) ListMerchantService(ctx *gin.Context, in dto.ListMerchantRequest) ([]dto.MerchantRes, error) {
	logrus.Println("[MerchantService ListMerchantService] start.")

	Limit := in.PageSize
	Offset := (in.PageID - 1) * in.PageSize

	collection := o.Store.Collection(tbMerchant)
	cursor, err := collection.Find(ctxDefault, bson.M{"active": "1"}, options.Find().SetSkip(Offset).SetLimit(Limit).SetSort(bson.D{primitive.E{Key: "name", Value: 1}}))
	if err != nil {
		logrus.Info("[MerchantService ListMerchantService] error when select data")
		log.Fatal(err)
	}

	var out []dto.MerchantRes
	for cursor.Next(context.Background()) {
		var merchant dto.MerchantRes
		errMerchant := cursor.Decode(&merchant)
		if errMerchant != nil {
			logrus.Info("[MerchantService ListMerchantService] error when looping data")
			log.Fatal(errMerchant)
		}

		out = append(out, merchant)
	}

	return out, nil
}

//func (o *MerchantService) UpdateMerchantService(ctx *gin.Context, in dto.UpdateMerchantRequest) (db.Merchant, error) {
//	arg := db.UpdateMerchantParams{
//		SetBalance: sql.NullString{String: "1", Valid: true},
//		Balance: sql.NullString{
//			String: util.FloatToString(in.Balance),
//			Valid:  true,
//		},
//		ID: in.Id,
//	}
//
//	if in.Name != "" {
//		arg.SetName = "1"
//		arg.Name = in.Name
//	}
//
//	_, err := o.Store.UpdateMerchant(ctx, arg)
//	if err != nil {
//		return db.Merchant{}, err
//	}
//
//	out, _ := o.GetMerchantService(ctx, in.Id)
//
//	return out, nil
//}
