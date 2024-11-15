package service

import "encoding/base64"

func PicAnalyseRequest(image []byte) {

}

func CovertPicToBase64(image []byte) string {
	var prefix = "data:image/jpg;base64,"
	return prefix + base64.StdEncoding.EncodeToString(image)
}

func CovertBase64ToPic(base64str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64str)
}
