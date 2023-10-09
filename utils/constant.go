package util

const ErrCd0 = "0000"
const ErrMsg0 = "success"
const ErrCd1 = "1001"
const ErrMsg1 = "merchant not found"
const ErrCd2 = "1002"
const ErrMsg2 = "merchant user not found"
const ErrCd3 = "1003"
const ErrMsg3 = "product not found"
const ErrCd4 = "1004"
const ErrMsg4 = "category not found"
const ErrCd5 = "1005"
const ErrMsg5 = "bank not found"
const ErrCd6 = "1006"
const ErrMsg6 = "merchant product not found"
const ErrCd7 = "1007"
const ErrMsg7 = "trx date wrong format"
const ErrCd8 = "1008"
const ErrMsg8 = "wrong merchant token"
const ErrCd9 = "1009"
const ErrMsg9 = "ref id already use"
const ErrCd10 = "1010"
const ErrMsg10 = "biller error"

const ErrCd99 = "9999"
const ErrMsg99 = "General Error"

type PdamCd struct {
	PdamTb string
}

func LoadPdamCd() PdamCd {
	pdamCd := PdamCd{
		PdamTb: "PDAMTBK",
	}

	return pdamCd
}

type PdamAdmin struct {
	PdamTb   int64
	PdamTbVa int64
}

func LoadPdamAdmin() PdamAdmin {
	pdamAdmin := PdamAdmin{
		PdamTb:   3000,
		PdamTbVa: 5000,
	}

	return pdamAdmin
}
