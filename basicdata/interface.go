package basicdata

import (
	"errors"
	"gpk/http"
	"io"
)

func GetData(url string, param any) ([]byte, error) {
	resp, err := http.Get(url, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("response code error")
	}

	body_bytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body_bytes, nil
}

func CreatData() {

}

func DeleteData() {

}
