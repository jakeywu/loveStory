package initial

import (
	"github.com/astaxie/beego/validation"
)

var msg = map[string]string{
	"Required":     "不能为空",
	"Min":          "最小为%d",
	"Max":          "最大为%d",
	"Range":        "Range is %d to %d",
	"MinSize":      "最短为%d",
	"MaxSize":      "最长为%d",
	"Length":       "Required length is %d",
	"Alpha":        "Must be valid alpha characters",
	"Numeric":      "Must be valid numeric characters",
	"AlphaNumeric": "Must be valid alpha or numeric characters",
	"Match":        "Must match %s",
	"NoMatch":      "Must not match %s",
	"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
	"Email":        "Must be a valid email address",
	"IP":           "Must be a valid ip address",
	"Base64":       "Must be valid base64 characters",
	"Mobile":       "Must be valid mobile number",
	"Tel":          "Must be valid telephone number",
	"Phone":        "Must be valid telephone or mobile phone number",
	"ZipCode":      "Must be valid zipcode",
}

func initValidMsg() {
	for name := range msg {
		validation.MessageTmpls[name] = msg[name]
	}
}
