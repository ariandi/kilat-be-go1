package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GetMerchantProducts struct {
	MerchantCode string `json:"merchant_code" bson:"merchant_code,omitempty"`
	ProductCode  string `json:"product_code" bson:"product_code,omitempty"`
	BankCode     string `json:"bank_code" bson:"bank_code,omitempty"`
}

type MerchantProducts struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FeeKilat     int                `json:"fee_kilat" bson:"fee_kilat,omitempty"`
	FeeBiller    int                `json:"fee_biller" bson:"fee_biller,omitempty"`
	FeeMerchant  int                `json:"fee_merchant" bson:"fee_merchant,omitempty"`
	PercentPrice int                `json:"percent_price" bson:"percent_price,omitempty"`
	BankCode     string             `json:"bank_code" bson:"bank_code,omitempty"`
	ProductCode  string             `json:"product_code" bson:"product_code,omitempty"`
	MerchantCode string             `json:"merchant_code" bson:"merchant_code,omitempty"`
	Active       int                `json:"active" bson:"active,omitempty"`
	CreatedBy    string             `json:"created_by" bson:"created_by,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at,omitempty"`
}
