package dto

type TransactionCreate struct {
	BillId       string `json:"bill_id" bson:"bill_id"`
	TxId         string `json:"tx_id" bson:"tx_id"`
	RefId        string `json:"ref_id" bson:"ref_id"`
	Type         string `json:"type" bson:"type"`
	Status       string `json:"status" bson:"status"`
	Sign         string `json:"sign" bson:"sign" copier:"MerchantToken"`
	MerchantId   string `json:"merchant_id" bson:"merchant_id" copier:"MerchantCode"`
	BankId       string `json:"bank_id" bson:"bank_id" copier:"Method"`
	CategoryId   string `json:"category_id" bson:"category_id" copier:"Category"`
	CategoryName string `json:"category_name" bson:"category_name"`
	ProductId    string `json:"product_id" bson:"product_id" copier:"ProductCode"`
	ProductName  string `json:"product_name" bson:"product_name"`
	ReqParam     InqReq `json:"req_param" bson:"req_param"`
	ResParam     InqRes `json:"res_param" bson:"res_param"`
	BatchTime    int    `json:"batch_time" bson:"batch_time"`
	CreatedBy    string `json:"created_by" bson:"created_by" copier:"User"`
	TxDate       string `json:"tx_date" bson:"tx_date" copier:"TrxDate"`
}
