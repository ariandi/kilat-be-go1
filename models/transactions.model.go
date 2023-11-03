package models

import (
	"github.com/ariandi/kilat-be-go1/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BillId          string             `json:"bill_id" bson:"bill_id"`
	TxId            string             `json:"tx_id" bson:"tx_id"`
	RefId           string             `json:"ref_id" bson:"ref_id"`
	Amount          int                `json:"amount" bson:"amount"`
	FeeMerchant     int                `json:"fee_merchant" bson:"fee_merchant"`
	FeeBiller       int                `json:"fee_biller" bson:"fee_biller"`
	FeeKilat        int                `json:"fee_kilat" bson:"fee_kilat"`
	FeeBenefit      int                `json:"fee_benefit" bson:"fee_benefit"`
	TotalAmount     int                `json:"total_amount" bson:"total_amount"`
	TotalAdmin      int                `json:"total_admin" bson:"total_admin"`
	Type            string             `json:"type" bson:"type"`
	Status          string             `json:"status" bson:"status"`
	Sign            string             `json:"sign" bson:"sign"`
	MerchantId      string             `json:"merchant_id" bson:"merchant_id"`
	BankId          string             `json:"bank_id" bson:"bank_id"`
	CategoryId      string             `json:"category_id" bson:"category_id"`
	CategoryName    string             `json:"category_name" bson:"category_name"`
	ProductId       string             `json:"product_id" bson:"product_id"`
	ProductName     string             `json:"product_name" bson:"product_name"`
	LastBalance     int                `json:"last_balance" bson:"last_balance"`
	FirstBalance    int                `json:"first_balance" bson:"first_balance"`
	DeductedBalance int                `json:"deducted_balance" bson:"deducted_balance"`
	ReqParam        dto.InqReq         `json:"req_param" bson:"req_param"`
	ResParam        dto.InqRes         `json:"res_param" bson:"res_param"`
	BatchTime       int                `json:"batch_time" bson:"batch_time"`
	TxDetail        string             `json:"tx_detail" bson:"tx_detail"`
	Active          int                `json:"active" bson:"active"`
	CreatedAt       string             `json:"created_at" bson:"created_at"`
	CreatedBy       string             `json:"created_by" bson:"created_by"`
	UpdatedAt       string             `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy       string             `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	TxDate          string             `json:"tx_date" bson:"tx_date"`
}
