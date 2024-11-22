package service

import (
	"encoding/base64"
	"strings"
)

func PicAnalyseRequest(image []byte) {

}

func CovertPicToBase64(image []byte) string {
	// var prefix = "data:image/jpg;base64,"
	var prefix = ""
	return prefix + base64.StdEncoding.EncodeToString(image)
}

func CovertBase64ToPic(base64str string) ([]byte, error) {
	str := strings.Split(base64str, ",")
	return base64.StdEncoding.DecodeString(str[len(str)-1])
}
