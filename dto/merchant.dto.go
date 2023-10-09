package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ListMerchantRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=5,max=200"`
}

type MerchantReq struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Kode        string             `json:"kode" bson:"kode" validate:"required"`
	Name        string             `json:"name" bson:"name"`
	User        string             `json:"user" bson:"user"`
	Active      string             `json:"active" bson:"active"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	TypePayment string             `json:"type_payment,omitempty" bson:"type_payment,omitempty"`
}

type MerchantRes struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Kode           string             `json:"kode" bson:"kode" validate:"required"`
	Name           string             `json:"name" bson:"name"`
	User           string             `json:"user" bson:"user"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	SecretCode     string             `json:"secret_code" bson:"secret_code"`
	CurrentBalance int                `json:"current_balance" bson:"current_balance"`
	BankAccount    string             `json:"bank_account,omitempty" bson:"bank_account,omitempty"`
	BankName       string             `json:"bank_name,omitempty" bson:"bank_name,omitempty"`
	BankBranch     string             `json:"bank_branch,omitempty" bson:"bank_branch,omitempty"`
	Active         string             `json:"active" bson:"active"`
	CreatedBy      string             `json:"created_by" bson:"created_by"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy      *string            `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	CallbackUrl    string             `json:"callback_url,omitempty" bson:"callback_url,omitempty"`
	TypePayment    string             `json:"type_payment,omitempty" bson:"type_payment,omitempty"`
}
