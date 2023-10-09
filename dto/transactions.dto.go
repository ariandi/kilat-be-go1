package dto

type TransactionCreate struct {
	BillId       string `json:"bill_id" bson:"bill_id"`
	TxId         string `json:"tx_id" bson:"tx_id"`
	RefId        string `json:"ref_id" bson:"ref_id"`
	Type         string `json:"type" bson:"type"`
	Status       string `json:"status" bson:"status"`
	Sign         string `json:"sign" bson:"sign"`
	MerchantId   string `json:"merchant_id" bson:"merchant_id"`
	BankId       string `json:"bank_id" bson:"bank_id"`
	CategoryId   string `json:"category_id" bson:"category_id"`
	CategoryName string `json:"category_name" bson:"category_name"`
	ProductId    string `json:"product_id" bson:"product_id"`
	ProductName  string `json:"product_name" bson:"product_name"`
	ReqParam     InqReq `json:"req_param" bson:"req_param"`
	BatchTime    int    `json:"batch_time" bson:"batch_time"`
	CreatedBy    string `json:"created_by" bson:"created_by"`
	TxDate       string `json:"tx_date" bson:"tx_date"`
}
