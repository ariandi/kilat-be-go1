package dto

type PdamTb struct {
	TotalAmount int64          `json:"totalamount" bson:"totalamount"`
	BillQty     int64          `json:"billqty" bson:"billqty"`
	Ts          string         `json:"ts" bson:"ts"`
	BillNumber  string         `json:"billnumber" bson:"billnumber"`
	BillName    string         `json:"billname" bson:"billname"`
	Address     string         `json:"address" bson:"address"`
	Type        string         `json:"type" bson:"type"`
	ResultCode  int64          `json:"resultcode" bson:"resultcode"`
	Detail      []DetailPdamTb `json:"detail" bson:"detail"`
}

type DetailPdamTb struct {
	Jenis        string `json:"jenis" bson:"jenis"`
	Usage        string `json:"usage" bson:"usage"`
	Period       string `json:"period" bson:"period"`
	Air          string `json:"air" bson:"air"`
	Adm          string `json:"adm" bson:"adm"`
	Pemeliharaan string `json:"pemeliharaan" bson:"pemeliharaan"`
	Meterai      string `json:"meterai" bson:"meterai"`
	Fine         string `json:"fine" bson:"fine"`
	Diskon       string `json:"diskon" bson:"diskon"`
	NonAir       int64  `json:"non_air" bson:"non_air"`
	Total        string `json:"total" bson:"total"`
}

type TbSetResponse struct {
	InqRequest  InqReq
	InqResponse InqRes
	TbResponse  PdamTb
}
