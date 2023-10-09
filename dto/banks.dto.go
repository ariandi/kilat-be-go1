package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GetBanks struct {
	Kode string `json:"kode" bson:"Kode,omitempty"`
}

type Banks struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Kode       string             `json:"kode" bson:"Kode,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	User       string             `json:"user" bson:"user,omitempty"`
	Password   string             `json:"password" bson:"password,omitempty"`
	SecretCode string             `json:"secret_code" bson:"secret_code,omitempty"`
	Url        string             `json:"url" bson:"url,omitempty"`
	Active     int                `json:"active" bson:"active,omitempty"`
	CreatedBy  string             `json:"created_by" bson:"created_by,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"`
}

type UrlBanks struct {
	Url    string `json:"url" bson:"url,omitempty"`
	Method string `json:"method" bson:"method,omitempty"`
	Param  string `json:"param" bson:"param,omitempty"`
	UrlInq string `json:"url_inq" bson:"url_inq,omitempty"`
	UrlPay string `json:"url_pay" bson:"url_pay,omitempty"`
	UrlAdv string `json:"url_adv" bson:"url_adv,omitempty"`
	UrlRev string `json:"url_rev" bson:"url_rev,omitempty"`
}
