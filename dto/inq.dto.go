package dto

type InqReq struct {
	TrxDate       string `json:"trx_date" bson:"trx_date" binding:"required"`
	ProductCode   string `json:"product_code" bson:"product_code" binding:"required"`
	BillId        string `json:"bill_id" bson:"bill_id" binding:"required"`
	MerchantCode  string `json:"merchant_code" bson:"merchant_code" binding:"required"`
	Category      string `json:"category" bson:"category" binding:"required"`
	User          string `json:"user" bson:"user" binding:"required"`
	Method        string `json:"method" bson:"method" binding:"required"`
	RefId         string `json:"ref_id" bson:"ref_id" binding:"required"`
	MerchantToken string `json:"merchant_token" bson:"merchant_token" binding:"required"`
}

type InqRes struct {
	ResultCd        string       `json:"result_cd" bson:"result_cd"`
	ResultMsg       string       `json:"result_msg" bson:"result_msg"`
	TxDate          string       `json:"tx_date"`
	TxId            string       `json:"tx_id"`
	RefId           string       `json:"ref_id" bson:"ref_id"`
	ProductCode     string       `json:"product_code" bson:"product_code"`
	BillId          string       `json:"bill_id" bson:"bill_id"`
	MerchantCode    string       `json:"merchant_code" bson:"merchant_code"`
	Category        string       `json:"category" bson:"category"`
	User            string       `json:"user" bson:"user"`
	Method          string       `json:"method" bson:"method"`
	Sign            string       `json:"sign" bson:"sign"`
	Amount          int          `json:"amount"`
	Admin           int          `json:"admin"`
	TotalAmount     int          `json:"total_amount"`
	DeductedBalance int          `json:"deducted_balance"`
	Detail          DetailInqRes `json:"detail"`
}

type DetailInqRes struct {
	ProductName   string                `json:"product_name"`
	Periode       string                `json:"periode"`
	CustName      string                `json:"cust_name"`
	CustAddress   string                `json:"cust_address"`
	Type          string                `json:"type"`
	RefProvider   string                `json:"ref_provider"`
	DetailBilling []BillingDetailInqRes `json:"detail_billing"`
}

type BillingDetailInqRes struct {
	Info          string `json:"info"`
	StandMeter    string `json:"stand_meter"`
	Period        string `json:"period"`
	Billamount    int    `json:"billamount"`
	Fine          int    `json:"fine"`
	AdminPpob     int    `json:"admin_ppob"`
	Total         int    `json:"total"`
	AdditionInfo  string `json:"addition_info"`
	AdditionInfo2 string `json:"addition_info_2"`
}
